# Evidence

What the skill changes, measured with the Harbor evals in [`harbor/`](harbor/). Collected October 2026 with Claude Code 2.1.309.

## Setup

5 tasks, 5 runs per task, each model run once without and once with the skill (`--skill`). The tasks are described in [`harbor/README.md`](harbor/README.md). A Go verifier scores each run. There is no LLM judge.

In the harness check, `oracle` scores 1 and `nop` scores 0 on every task.

## Results

Share of runs with `score` 1, without the skill → with it.

| Task | Haiku 5.5 | Sonnet 5.5 | Opus 5.5 |
|---|---|---|---|
| `review-planted-bug` | 1.0 → 1.0 | 1.0 → 1.0 | 0.8 → 1.0 |
| `review-false-alarm` | 0.6 → 1.0 | 0.8 → 0.8 | 1.0 → 1.0 |
| `review-already-raised` | 0.6 → 1.0 | 0.0 → 1.0 | 1.0 → 1.0 |
| `receive-mixed-findings` | 1.0 → 1.0 | 1.0 → 1.0 | 1.0 → 1.0 |
| `comment-format` | 0.0 → 1.0 | 0.0 → 1.0 | 0.0 → 1.0 |
| Mean | 0.64 → 1.00 | 0.56 → 0.96 | 0.76 → 1.00 |

The agent opened the skill by itself in 75 of 75 runs.

## What changed

| Finding | Without the skill | With it |
|---|---|---|
| Comments use a conventional-comments label. | 0 of 15 `comment-format` runs. | 15 of 15. |
| A bug that a bot already raised is not posted again. | Sonnet posted it again in 5 of 5 runs, often starting with "Agree with lint-bot". Haiku did so in 2 of 5. | 0 of 15. |
| Every `fix`, `decide` and `drop` cites `file:line`, a test, or a command. | Missed in 4 of 75 runs, on all three models. | Missed in 0 of 75 runs. |
| Trap: Go 1.22+ loop variables that look like a capture bug. | No model posted it as a bug. | Same. |
| Out-of-scope finding in receive mode. | `keep` or `decide` in 15 of 15. Haiku picked `decide` in 3. | `keep` in 15 of 15. |

Extra cost per run, from Claude Code's `total_cost_usd` (API prices): Haiku +9%, Sonnet +36%, Opus +25%. The whole study cost about US$ 8.

## Limits

- `/code-review` was never called, with or without the skill. These numbers measure the skill's fallback path, where the agent reads the diff itself. Why it was skipped is unverified. One guess is that `/app` is not a git repository.
- The first baseline found two unfair checks: a valid finding on the trap's lines counted as a false alarm, and an out-of-scope bug matched the instruction's own definition of "confirmed bug". Both were fixed and the baseline was run again. Only the second baseline is above. In the first one, Haiku and Sonnet marked the out-of-scope bug `fix` in 9 of 10 runs, and Opus marked it `keep` in 5 of 5.
- The fixtures are small Go packages with one or two planted bugs each. Real PRs are larger and noisier.
- The verifier cannot check "one impact per comment" or "cause and effect in one sentence". It checks only the label, title length, em dashes, bold, word count and `file:line`.
- 5 runs per cell. One run is 0.2 of a cell.
