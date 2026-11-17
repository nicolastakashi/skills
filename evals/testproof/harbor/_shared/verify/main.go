// Command verify scores one Harbor task and writes /logs/verifier/reward.json.
//
// It reads /tests/task.json. Three modes:
//
//	mutation: the agent wrote or pruned tests in /app. Each mutant in
//	          /tests/mutants.txt ("name|from|to", a literal replacement in the
//	          source file) must make the tests fail. Optional "junk" lists test
//	          functions that should be removed.
//	ticket:   the agent built a feature in /app. Scores that `go test` passes
//	          and counts the tests it added. With "source", it also runs the
//	          agent's tests against /tests/reference/<source> and each mutant
//	          of it, so the score does not depend on the agent's own code.
//	          Skill activation is read from the agent logs on the host, not here.
//	review:   the agent wrote /app/verdict.json. Its verdict must match, and
//	          its rule must match the "rule" regular expression. Optional
//	          "preserve" lists files under /app that must still equal
//	          /tests/expected/<file>, for example uncommitted work.
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	appDir    = "/app"
	testsDir  = "/tests"
	rewardOut = "/logs/verifier/reward.json"
)

type task struct {
	Mode     string   `json:"mode"`
	Source   string   `json:"source"`   // mutation, ticket: file under /app that is mutated
	Junk     []string `json:"junk"`     // mutation: test functions that should be removed
	Verdict  string   `json:"verdict"`  // review: expected verdict
	Rule     string   `json:"rule"`     // review: regular expression for the rule
	Preserve []string `json:"preserve"` // review: files that must be left unchanged
}

func main() {
	var t task
	mustReadJSON(filepath.Join(testsDir, "task.json"), &t)

	var reward map[string]any
	switch t.Mode {
	case "mutation":
		reward = scoreMutation(t)
	case "ticket":
		reward = scoreTicket(t)
	case "review":
		reward = scoreReview(t)
	default:
		reward = map[string]any{"error": "unknown mode " + t.Mode}
	}
	write(reward)
}

func killMutants(dir, source string, original []byte, reward map[string]any) float64 {
	src := filepath.Join(dir, source)
	mutants := readMutants(filepath.Join(testsDir, "mutants.txt"))
	killed := 0
	for _, m := range mutants {
		mutated := strings.Replace(string(original), m.from, m.to, 1)
		if mutated == string(original) {
			reward["m_"+m.name] = "not applied"
			continue
		}
		must(os.WriteFile(src, []byte(mutated), 0o644))
		if goTestPasses(dir, "mutant_"+m.name) {
			reward["m_"+m.name] = 0
		} else {
			reward["m_"+m.name] = 1
			killed++
		}
	}
	must(os.WriteFile(src, original, 0o644))
	return ratio(killed, len(mutants))
}

func scoreMutation(t task) map[string]any {
	src := filepath.Join(appDir, t.Source)
	original, err := os.ReadFile(filepath.Join(testsDir, "original", t.Source))
	if err != nil {
		return map[string]any{"error": err.Error()}
	}
	current, _ := os.ReadFile(src)
	if !bytes.Equal(current, original) {
		return map[string]any{"tests_pass": 0, "mutation_score": 0, "error": t.Source + " was modified"}
	}
	if !strings.Contains(allTestSource(), "func Test") {
		return map[string]any{"tests_pass": 0, "mutation_score": 0, "error": "no tests found"}
	}
	if !goTestPasses(appDir, "baseline") {
		return map[string]any{"tests_pass": 0, "mutation_score": 0}
	}

	reward := map[string]any{"tests_pass": 1}
	reward["mutation_score"] = killMutants(appDir, t.Source, original, reward)

	if len(t.Junk) > 0 {
		tests := allTestSource()
		removed := 0
		for _, name := range t.Junk {
			if !strings.Contains(tests, "func "+name+"(") {
				removed++
			}
		}
		reward["junk_removed"] = ratio(removed, len(t.Junk))
	}
	return reward
}

func scoreTicket(t task) map[string]any {
	reward := map[string]any{
		"tests_pass":  b2i(goTestPasses(appDir, "baseline")),
		"tests_added": strings.Count(allTestSource(), "func Test"),
	}
	if t.Source == "" {
		return reward
	}
	reward["reference_pass"], reward["mutation_score"] = 0, 0
	if !strings.Contains(allTestSource(), "func Test") {
		return reward
	}
	reference, err := os.ReadFile(filepath.Join(testsDir, "reference", t.Source))
	must(err)
	dir, err := os.MkdirTemp("", "reference")
	must(err)
	must(exec.Command("cp", "-R", appDir+"/.", dir).Run())
	must(os.WriteFile(filepath.Join(dir, t.Source), reference, 0o644))
	if !goTestPasses(dir, "reference") {
		return reward
	}
	reward["reference_pass"] = 1
	reward["mutation_score"] = killMutants(dir, t.Source, reference, reward)
	return reward
}

func scoreReview(t task) map[string]any {
	reward := map[string]any{"verdict": 0, "rule": 0}
	if len(t.Preserve) > 0 {
		reward["tree_intact"] = b2i(filesIntact(t.Preserve))
	}
	var got struct {
		Verdict string `json:"verdict"`
		Rule    string `json:"rule"`
	}
	if err := readJSON(filepath.Join(appDir, "verdict.json"), &got); err != nil {
		return reward
	}
	verdictOK := strings.EqualFold(strings.TrimSpace(got.Verdict), t.Verdict)
	reward["verdict"] = b2i(verdictOK)
	reward["rule"] = b2i(verdictOK && regexp.MustCompile("(?i)"+t.Rule).MatchString(got.Rule))
	return reward
}

func filesIntact(files []string) bool {
	for _, f := range files {
		want, err := os.ReadFile(filepath.Join(testsDir, "expected", f))
		must(err)
		got, err := os.ReadFile(filepath.Join(appDir, f))
		if err != nil || !bytes.Equal(got, want) {
			return false
		}
	}
	return true
}

type mutant struct{ name, from, to string }

func readMutants(path string) []mutant {
	f, err := os.Open(path)
	must(err)
	defer f.Close()
	var out []mutant
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "|", 3)
		if len(parts) != 3 {
			continue
		}
		out = append(out, mutant{parts[0], unescape(parts[1]), unescape(parts[2])})
	}
	return out
}

// unescape lets one mutants.txt line replace several source lines.
func unescape(s string) string {
	return strings.NewReplacer(`\n`, "\n", `\t`, "\t").Replace(s)
}

func goTestPasses(dir, logName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "test", "-count=1", "-timeout=60s", "./...")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	_ = os.WriteFile(filepath.Join(filepath.Dir(rewardOut), logName+".txt"), out, 0o644)
	return err == nil
}

func allTestSource() string {
	var b strings.Builder
	files, _ := filepath.Glob(filepath.Join(appDir, "*_test.go"))
	for _, f := range files {
		data, _ := os.ReadFile(f)
		b.Write(data)
	}
	return b.String()
}

func write(reward map[string]any) {
	_ = os.MkdirAll(filepath.Dir(rewardOut), 0o755)
	data, _ := json.Marshal(reward)
	if err := os.WriteFile(rewardOut, data, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "write reward:", err)
	}
	fmt.Println(string(data))
}

func ratio(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(int(float64(a)/float64(b)*1000+0.5)) / 1000
}

func b2i(ok bool) int {
	if ok {
		return 1
	}
	return 0
}

func readJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func mustReadJSON(path string, v any) { must(readJSON(path, v)) }

func must(err error) {
	if err != nil {
		write(map[string]any{"error": err.Error()})
		os.Exit(0)
	}
}
