// Command verify scores /app/review.json against /tests/task.json and writes
// /logs/verifier/reward.json. The README in evals/reviewproof/harbor explains each metric.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

const (
	appDir    = "/app"
	testsDir  = "/tests"
	rewardOut = "/logs/verifier/reward.json"

	maxTitle = 100
	maxWords = 60 // the rule in the fixture's AGENTS.md
)

var (
	evidenceRe = regexp.MustCompile(`\w+\.go:\d+|\bTest[A-Z]\w*|\bgo (test|vet|run|build)\b`)
	labelRe    = regexp.MustCompile(`^(praise|nitpick|suggestion|issue|todo|question|thought|chore|note|typo|polish|quibble)(\s?\([a-z ,-]+\))?: \S`)
)

type finding struct {
	ID       string `json:"id"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Verdict  string `json:"verdict"`
	Evidence string `json:"evidence"`
	Comment  string `json:"comment"`
}

type match struct {
	Name     string   `json:"name"`
	ID       string   `json:"id"`
	File     string   `json:"file"`
	From     int      `json:"from"`
	To       int      `json:"to"`
	Verdicts []string `json:"verdicts"` // expect only
}

type task struct {
	Expect  []match  `json:"expect"`
	Forbid  []match  `json:"forbid"`
	Require []string `json:"require"`
}

func main() {
	var spec task
	if err := readJSON(filepath.Join(testsDir, "task.json"), &spec); err != nil {
		write(map[string]any{"score": 0, "error": err.Error()})
		return
	}
	var got struct {
		Findings []finding `json:"findings"`
	}
	if err := readJSON(filepath.Join(appDir, "review.json"), &got); err != nil {
		write(map[string]any{"score": 0, "error": "review.json: " + err.Error()})
		return
	}
	write(score(spec, got.Findings))
}

func score(spec task, findings []finding) map[string]any {
	reward := map[string]any{}
	for _, want := range spec.Expect {
		ok := false
		for _, got := range findings {
			if want.matches(got) && slices.Contains(want.Verdicts, verdict(got)) {
				ok = true
			}
		}
		reward[want.Name] = b2i(ok)
	}
	for _, want := range spec.Forbid {
		ok := true
		for _, got := range findings {
			if want.matches(got) && posted(got) {
				ok = false
			}
		}
		reward[want.Name] = b2i(ok)
	}

	reward["evidence"] = 1
	for _, got := range findings {
		if slices.Contains([]string{"fix", "decide", "drop"}, verdict(got)) && !evidenceRe.MatchString(got.Evidence) {
			reward["evidence"] = 0
			reward["evidence_fail"] = fmt.Sprintf("%s:%d has no file:line, test or command", got.File, got.Line)
		}
	}

	nPosted := 0
	reward["format"] = 1
	for _, got := range findings {
		if !posted(got) {
			continue
		}
		nPosted++
		if why := formatProblem(got); why != "" {
			reward["format"] = 0
			reward["format_fail"] = fmt.Sprintf("%s:%d %s", got.File, got.Line, why)
		}
	}
	if nPosted == 0 {
		reward["format"] = 0
		reward["format_fail"] = "no posted comments"
	}
	reward["posted"] = nPosted

	all := 1
	for _, name := range spec.Require {
		if value, _ := reward[name].(int); value != 1 {
			all = 0
		}
	}
	reward["score"] = all
	return reward
}

func formatProblem(got finding) string {
	title, _, _ := strings.Cut(strings.TrimSpace(got.Comment), "\n")
	switch {
	case !labelRe.MatchString(title):
		return "no conventional-comments label"
	case utf8.RuneCountInString(title) > maxTitle:
		return "title over one line"
	case strings.Contains(got.Comment, "—"):
		return "em dash"
	case strings.Contains(got.Comment, "**"):
		return "bold"
	case len(strings.Fields(got.Comment)) > maxWords:
		return fmt.Sprintf("over %d words", maxWords)
	}
	data, err := os.ReadFile(filepath.Join(appDir, cleanPath(got.File)))
	if err != nil {
		return "file not found"
	}
	if got.Line < 1 || got.Line > strings.Count(string(data), "\n")+1 {
		return "line out of range"
	}
	return ""
}

func (want match) matches(got finding) bool {
	if want.ID != "" && strings.TrimSpace(got.ID) == want.ID {
		return true
	}
	return want.File != "" && cleanPath(got.File) == want.File && got.Line >= want.From && got.Line <= want.To
}

func posted(got finding) bool {
	return slices.Contains([]string{"fix", "decide"}, verdict(got)) && strings.TrimSpace(got.Comment) != ""
}

func verdict(got finding) string { return strings.ToLower(strings.TrimSpace(got.Verdict)) }

// cleanPath turns "/app/a.go", "./a.go" and "b/a.go" (diff prefix) into "a.go".
func cleanPath(path string) string {
	path = strings.TrimPrefix(strings.TrimSpace(path), appDir+"/")
	path = strings.TrimPrefix(path, "./")
	if _, err := os.Stat(filepath.Join(appDir, path)); err != nil {
		path = strings.TrimPrefix(strings.TrimPrefix(path, "a/"), "b/")
	}
	return path
}

// write keeps only numbers in reward.json, because Harbor rejects any other value.
// The full result, with the reasons, goes to stdout and details.json.
func write(result map[string]any) {
	numbers := map[string]any{}
	for name, value := range result {
		if _, isInt := value.(int); isInt {
			numbers[name] = value
		}
	}
	_ = os.MkdirAll(filepath.Dir(rewardOut), 0o755)
	data, _ := json.Marshal(numbers)
	if err := os.WriteFile(rewardOut, data, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "write reward:", err)
	}
	details, _ := json.Marshal(result)
	_ = os.WriteFile(filepath.Join(filepath.Dir(rewardOut), "details.json"), details, 0o644)
	fmt.Println(string(details))
}

func b2i(ok bool) int {
	if ok {
		return 1
	}
	return 0
}

func readJSON(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}
