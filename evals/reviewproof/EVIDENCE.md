# Evidence

What the skill changes, measured with the Harbor evals in [`harbor/`](harbor/). Collected October 2026 with Claude Code 2.1.309.

## Setup

5 tasks, 5 runs per task, each model run once without and once with the skill (`--skill`). The tasks are described in [`harbor/README.md`](harbor/README.md). A Go verifier scores each run. There is no LLM judge. Each task is a git repository with a `main` and a `pr` branch, so review tools can diff the PR.

In the harness check, `oracle` scores 1 and `nop` scores 0 on every task.

## Results

Share of runs with `score` 1, without the skill → with it.

| Task | Haiku 5.5 | Sonnet 5.5 | Opus 5.5 |
|---|---|---|---|
| `review-planted-bug` | 1.0 → 1.0 | 1.0 → 1.0 | 1.0 → 1.0 |
| `review-false-alarm` | 1.0 → 1.0 | 0.6 → 0.8 | 1.0 → 1.0 |
| `review-already-raised` | 0.4 → 0.8 | 0.0 → 1.0 | 1.0 → 1.0 |
| `receive-mixed-findings` | 1.0 → 1.0 | 0.6 → 0.8 | 1.0 → 1.0 |
| `comment-format` | 0.0 → 1.0 | 0.0 → 1.0 | 0.0 → 1.0 |
| Mean | 0.68 → 0.96 | 0.44 → 0.92 | 0.80 → 1.00 |

The agent opened the skill by itself in 75 of 75 runs.

## What changed

| Finding | Without the skill | With it |
|---|---|---|
| Comments use a conventional-comments label and a title line. | 0 of 15 `comment-format` runs. | 15 of 15. |
| A bug that a bot already raised is not posted again. | Posted again in 8 of 15 runs. Sonnet did it in 5 of 5, often starting with "Agree with lint-bot". | 0 of 15. |
| Every `fix`, `decide` and `drop` cites `file:line`, a test, or a command. | Missed in 2 of 75 runs. | Missed in 1 of 75. |
| `/code-review` runs in review mode. | 0 of 60 runs. | Opus 10 of 20, Haiku 2 of 20, Sonnet 0 of 20. |
| Out-of-scope finding in receive mode is `keep` or `decide`. | 13 of 15. | 14 of 15. |
| Trap: Go 1.22+ loop variables that look like a capture bug. | No model posted it as a bug. | Same. |

Extra cost per run, from Claude Code's `total_cost_usd` (API prices): Haiku +32%, Sonnet +23%, Opus +70%. Most of the Opus increase is `/code-review`.

## Changes from real use

The first version was used on 11 real PRs. Four problems came up. This round addresses two of them:

| Problem in real use | Change |
|---|---|
| The agent skipped `/code-review` "to save time" or because "the diff is small", and missed a bug that the tool found in another session. | Run the review tool even on a small diff. Say why when it did not run. |
| `fix` was used for test gaps and wrong docs, and then changed after the user asked "are you sure?". The user also asked what the verdicts mean. | `fix` covers bugs, missing tests and wrong docs in code the diff touches. A legend goes under the table. |

In the evals, Haiku often wrote the whole comment on one line. With the old skill text it kept the title line in 4 of 5 runs, then in 0 of 5 on a second try. With the new text it did so in 1 of 5, twice. A full example comment under the draft rules fixed it: 5 of 5, then 5 of 5 again in the final run.

## Limits

- `/code-review` still runs in only some review runs. With Sonnet it did not run at all, and the final message mentioned the review tool in 20 of 20 runs, as the skill asks.
- The first round ran without git history, and `/code-review` did not run there. Its scores were close to the ones above. This round reran the baseline in the new setup.
- Earlier rounds found unfair checks: a valid finding on the trap's lines counted as a false alarm, an out-of-scope bug matched the instruction's own definition of "confirmed bug", and a numeric `id` crashed the verifier. All three are fixed.
- 4 of the 25 Opus runs with the skill started after the last edit to the `fix` row. The Sonnet `receive-mixed-findings` cell was rerun with that edit.
- The fixtures are small Go packages with one or two planted bugs each. Real PRs are larger and noisier.
- The verifier cannot check "one impact per comment" or "cause and effect in one sentence". It checks only the label, title length, em dashes, bold, word count and `file:line`.
- 5 runs per cell. One run is 0.2 of a cell.
