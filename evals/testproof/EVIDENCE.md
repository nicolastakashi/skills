# Evidence

Where each rule comes from. Collected September 2026.

## Sources

| Source | What we looked at |
|---|---|
| [openclaw/openclaw](https://github.com/openclaw/openclaw) (MIT, TypeScript) | 131 test-pruning commits, Aug–Sep 2026: 31,908 deleted test declarations (moves excluded). Plus 12,512 kept tests as a control. |
| A private Go + TypeScript monorepo | Tests added over six months. Blind sample of 120 new tests, half from agent co-authored commits. |
| Four open-source Go repos (Apache-2.0) | ~4,200 tests added in 2026. Blind sample of 60. |

"Blind" means the labeller did not know if a test was deleted, kept, or who wrote it.

## Findings

### Writing tests (gate)

| Finding | Number | Rule |
|---|---|---|
| Most new tests are useful. | 7.5% (private) and 13% (open source) were low value. | — |
| Agent and human tests look the same. | 5 of 60 vs 4 of 60 low value. | — |
| The top problem is "only checks no error / no panic / not nil". | 10 of 17 low-value tests, counting names that over-promise. | G1 checks-result, G2 name-matches |
| Tests of test helpers. | 3 of 17. | G3 tests-real-code |
| Expected value from the mock or the code under test. | About 2% of labels. | G4 independent-expected (weak data) |
| Low-value tests do not go away by themselves. | 7 of 9 still existed months later. | — |
| Missing default values are the top gap when an agent writes tests. | Haiku without the skill: `default-max` survived in 5 of 5 retry-policy runs, `default-retryafter-cap` in 4 of 5. | G6 covers-edges |
| Listing edges before writing works better than a check after. | Haiku, skill read, retry-policy task, 5 runs each: G6 as a check only 75% of mutants caught; "list defaults, caps and errors first" 89%. No skill: about 80%. | G6 covers-edges |

### Deleting tests (audit)

| Finding | Number | Rule |
|---|---|---|
| Deleted tests mostly looked valuable on their own. | Only 7% of deleted tests looked low value in isolation (1% of kept tests). | A1 names-keeper |
| Coverage after deletion, sample of 60 valuable-looking deletions. | 38% covered, 43% partly covered, 18% lost. | A1 names-keeper, A3 mutation-proof |
| What partial coverage loses. | Mostly the negative case or the default. 6 of 11 losses. | A2 keeps-negatives |
| Themes restored inside the pruning PRs themselves. | 117 restore commits. Top themes: contracts 33, lifecycle/ordering 21, security 15, regressions 15. | A4 extra-care |
| Surface signals barely separate deleted from kept tests. | Heavy mocks 3.1% deleted vs 4.8% kept. String grep 1.3% vs 2.3%. Snapshot 2.9% vs 3.2%. | A5 hints-not-proof |
| Regex "no assertion" detection is unreliable. | Most Go hits were golden or helper-based tests. | A5 hints-not-proof |

### Does the skill help? (Harbor evals, `evals/testproof/harbor`)

| Model | Skill opened by itself | Without skill | With skill |
|---|---|---|---|
| Opus (default) | not measured | 100% on all tasks | 100% on all tasks |
| Sonnet | 15 of 15 runs | retry-policy: `default-max` missed in 2 of 5 runs | 0 misses in 15 runs |
| Haiku | 0 to 4 of 15 runs, any description | defaults missed in most runs | helps only when read |

When the skill is read on purpose, Haiku catches 89% of retry-policy mutants instead of about 80%.

## Limits

- One public TypeScript repo for deletions. Go deletion data is missing.
- Labels come from LLM reviewers. Two losses were checked by hand.
- Sample sizes are small. Expect about ±10 points on percentages.
- Agent authorship is only known from commit trailers. Many agent commits have none.
- G5 fails-before-fix has no data of its own. It comes from openclaw's practice.
