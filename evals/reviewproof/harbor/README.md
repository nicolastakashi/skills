# reviewproof evals (Harbor)

Agent-agnostic evals for the `reviewproof` skill, run with [Harbor](https://github.com/harbor-framework/harbor).
Each task runs in Docker. A Go verifier scores the result. No LLM judge.

## Tasks

Each task is a small Go repo on a PR branch, with `PR.diff`, `comments.json` (comments already on the PR) and an `AGENTS.md`/`CLAUDE.md` with the user's comment rules.
The agent writes `/app/review.json`: one entry per finding with `file`, `line`, `verdict` (`fix`, `decide`, `keep`, `drop`), `evidence` and `comment`.

| Task | What it checks | `score` is 1 when |
|---|---|---|
| `review-planted-bug` | Recall: finds a bug the tests miss. | `fix` at the bug's lines. |
| `review-false-alarm` | Precision: Go 1.22+ loop variables look like a capture bug but are correct. | No posted comment on the trap, and the real bug is found. |
| `review-already-raised` | Dedup: one bug is already in `comments.json`. | No posted comment on it, and the other bug is found. |
| `receive-mixed-findings` | Four incoming findings on your own PR: 2 real, 1 false, 1 out of scope. | `fix`, `fix`, `drop`, and `keep` or `decide`. |
| `comment-format` | Posted comments are ready to post. | Label and one-line title, no em dash, no bold, at most 60 words, a real `file:line`. |

Every review task also has a real bug, so an empty review never scores.
Every task requires `evidence`: each `fix`, `decide` and `drop` cites `file:line`, a test name, or a `go` command.
Posted means verdict `fix` or `decide` with a non-empty comment.

`reward.json` also has the single metrics (for example `no_repeat`, `format`). The reasons are in `details.json` next to it.

## Layout

Same as [`evals/testproof/harbor`](../../testproof/harbor/README.md#layout). `tests/task.json` lists:

| Key | Meaning |
|---|---|
| `expect` | A finding must match, by `id` or by `file` and line range, with one of `verdicts`. |
| `forbid` | No posted finding may match. |
| `require` | The metrics that must all be 1 for `score` to be 1. |

Edit shared files in `_shared/`, then run `./sync.sh`. Run `./sync.sh --check` to confirm every task is up to date.

## Run

Needs Docker and an API key for the agent you test.

```bash
uv tool install harbor

# without the skill
harbor run -p tasks -a claude-code -m sonnet -k 5

# with the skill
harbor run -p tasks -a claude-code -m sonnet -k 5 --skill ../../../skills/reviewproof
```

## Check the harness

```bash
harbor run -p tasks -a oracle
harbor run -p tasks -a nop
```

Expect `oracle` to score 1 and `nop` to score 0 on every task.
