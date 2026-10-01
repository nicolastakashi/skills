---
name: testproof
description: Use when adding, writing, changing, or reviewing tests, including requests like "add tests for", "write tests", "cover this with tests", or fixing a bug. Also use when cleaning up, pruning, deduplicating, merging, or deleting tests, like "clean up this test file" or "remove redundant tests". Decides if each test earns its place. Any language.
---

# testproof

A test earns its place when it can fail for a real bug.
This skill has two modes:

- **Gate** (rules `G1`, `G2`, …): a new or changed test. Decide `keep` or `cut`.
- **Audit** (rules `A1`, `A2`, …): a proposal to delete or merge tests. Decide `allow` or `block`.

Each rule has a code and a name, for example `G1 checks-result`.
Always cite both. The code is short. The name says what it means.

## Gate mode

Before you write tests, read the code and its docs. List every default value, cap, limit and error path.
Write one test for each item on the list. Then check each test against the rules below.

Check the test against each rule. One failed rule means `cut`, unless an exception applies.
When you cut, say how to fix the test. Fixing is better than deleting.

| Rule | Name | Check | Fails when |
|---|---|---|---|
| **G1** | `checks-result` | The test checks the result. | It only checks "no error", "no panic", "not nil", "truthy" or a length. |
| **G2** | `name-matches` | The name matches the assertions. | The name promises more than the assertions check, or the opposite. |
| **G3** | `tests-real-code` | The test runs production code. | It only exercises a test helper, fake, or fixture. |
| **G4** | `independent-expected` | The expected value is independent. | The expected value comes from the code under test, or a mock produces the asserted result. |
| **G5** | `fails-before-fix` | A regression test fails on the old code. | The test does not reach the input that triggered the bug. |
| **G6** | `covers-edges` | Each documented default, limit and rejection has a test. | A default, cap or error path in the code or docs has no test. |

### Exceptions

- **Nil or empty safety.** A "does not panic" test is fine when safety is the contract and the name says so.
- **Golden and table tests.** A helper that compares against expected data is a real assertion.
- **Mocks.** Many mocks are fine when the test asserts an observable outcome. Mock count is not a rule.

### How to check G5 fails-before-fix

Read the fix. Find the input that triggered the bug.
Ask: on the pre-fix code, would this test fail? If you can run it, run it on the old code.

### How to check G6 covers-edges

List every default value, cap and error path in the code and its docs.
For each one, name the test that would fail if it changed. If none would fail, the tests are not done. Add one.

## Audit mode

A deleted test often looks redundant, but is only partly covered elsewhere. The gap is invisible until a bug ships.
So default to `block`. Allow a deletion only when all checks pass.

| Rule | Name | Check |
|---|---|---|
| **A1** | `names-keeper` | The proposal names a **keeper**: `file::test name`. The keeper checks **the same case** with the same or a stronger assertion. "Covered by the integration suite" is not a keeper. |
| **A2** | `keeps-negatives` | Negative cases and defaults keep a test. Check rejections, error paths, "does not" cases and default values first. Keepers usually cover the happy path only. |
| **A3** | `mutation-proof` | Prove it with a mutation. Break the production code in the way the deleted test guards. The keeper must fail. Then restore the code. |
| **A4** | `extra-care` | Take extra care with security, lifecycle and ordering, public contracts, and bug regressions. A gap here is costly and hard to notice. |
| **A5** | `hints-not-proof` | Heuristics are hints, not proof. "Many mocks", "snapshot", "string grep" or "only one expect" do not justify a deletion by themselves. |

### Allowed without a keeper

- **Dead code.** The production code is deleted in the same change, and it has no other callers.
- **Exact duplicate.** Another test, or a table row, has the same input and the same expected output.

### Evidence record

Write this for every deletion before you edit. A missing field means `block`.

```
test:      <file>::<name>
guards:    <input> -> <observable result>
keeper:    <file>::<name>, or "dead code", or "exact duplicate"
same case: <yes/no + one line>
mutation:  <what you broke> -> <keeper failed: yes/no>
category:  <security | lifecycle | contract | regression | default | negative | other>
```

## Output

End with one verdict line:

```
VERDICT: {"verdict": "keep|cut|allow|block", "rule": "<code and name, e.g. G1 checks-result, or none>"}
```

For an audit of many tests, write one evidence record per test, then a summary: allowed, blocked, and why.
