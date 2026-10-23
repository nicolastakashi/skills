# testproof evals (Harbor)

Agent-agnostic evals for the `testproof` skill, run with [Harbor](https://github.com/harbor-framework/harbor).
Each task runs in Docker. A Go verifier scores the result. No LLM judge.

## Tasks

| Task | What the agent does | Score (`reward.json`) |
|---|---|---|
| `write-go-percentile` | Writes Go tests for a function. | `mutation_score`: share of mutants the tests catch. |
| `write-go-retry-policy` | Writes Go tests for a retry policy with defaults and caps. | `mutation_score` |
| `prune-go-retry-policy` | Cleans up a messy test file. | `mutation_score`: coverage kept. `junk_removed`: share of low-value tests removed. |
| `ticket-go-retry-min-delay` | Builds a small feature from a ticket that never mentions tests. | `tests_pass`, `tests_added`. Activation is read from the logs, see below. |
| `ticket-go-retry-multiplier` | Same, other feature. | same |
| `ticket-go-percentiles` | Same, other package. | same |
| `review-gate-noerror-notnil` | Reviews a new test. | `verdict`: 1 if it says `cut`. `rule`: 1 if the reason names the right rule. |
| `review-audit-lost-negative` | Reviews a test deletion. | `verdict`: 1 if it says `block`. `rule`: 1 if the reason names the right rule. |

## Layout

Every task has the same shape:

| Path | What it is |
|---|---|
| `instruction.md` | The request the agent gets. |
| `environment/Dockerfile` | The container the agent works in. Copied from `_shared/`. |
| `environment/app/` | The files the agent sees. |
| `solution/solve.sh` | The reference answer. Used by the `oracle` agent to prove the task is solvable. |
| `tests/test.sh` | Runs the verifier. Copied from `_shared/`. |
| `tests/verify/` | The Go verifier. Copied from `_shared/`. |
| `tests/task.json` | What to score: `mutation` or `review` mode, and its settings. |
| `tests/mutants.txt` | Mutation mode: `name\|text to find\|replacement`, one per line. |
| `tests/original/` | Mutation mode: the untouched source file. |

Edit shared files in `_shared/`, then run `./sync.sh`. Run `./sync.sh --check` to confirm every task is up to date.

## Run

Needs Docker and an API key for the agent you test.

```bash
uv tool install harbor

# without the skill
harbor run -p tasks -a claude-code -m sonnet -k 5

# with the skill
harbor run -p tasks -a claude-code -m sonnet -k 5 --skill ../../../skills/testproof

# another agent, same tasks
harbor run -p tasks -a codex -k 5 --skill ../../../skills/testproof
```

`-k 5` runs each task 5 times. Compare the two runs per task.

## Check the harness

`oracle` runs the reference solution. `nop` does nothing.
Expect `oracle` to score 1 and `nop` to score 0. For `prune`, `nop` keeps full coverage but removes no junk.

```bash
harbor run -p tasks -a oracle
harbor run -p tasks -a nop
```

## Check activation

Ticket tasks measure if the agent opens the skill on its own. The verifier cannot see that, so read the agent logs on the host:

```bash
./activation.sh jobs/<job>
```

Do not pass `-o` with a path outside this folder. Docker cannot mount it and every trial fails with `RewardFileNotFoundError`.
