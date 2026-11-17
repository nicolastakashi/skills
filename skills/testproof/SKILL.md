---
name: testproof
description: Use when implementing any feature, fix or ticket that changes code, before the first edit, even when the change looks like a one-liner. Plans which tests the change needs and checks that each one earns its place. Also use when writing, reviewing, cleaning up, pruning, deduplicating, merging, or deleting tests. Not for docs-only or formatting changes. Any language.
---

# testproof

A test earns its place when it can fail for a real bug.
This skill has two modes:

- **Gate** (rules `G1`, `G2`, …): a new or changed test. Decide `keep` or `cut`.
- **Audit** (rules `A1`, `A2`, …): a proposal to delete or merge tests. Decide `allow` or `block`.

Each rule has a code and a name, for example `G1 checks-result`.
Always cite both. The code is short. The name says what it means.

## Gate mode

Before you write tests, read the code, its docs and the tests next to it. List every default value, cap, limit and error path.
Note the test helpers, fixtures and table tests the package already uses.
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
| **G7** | `public-behavior` | The test sets up and checks through the public API or observable output. | It reads or writes private fields, or checks how the code works instead of what it returns. A refactor that keeps the behavior would break it. |
| **G8** | `repo-conventions` | The test fits the tests next to it and adds something they miss. | It repeats a case an existing test already guards, or it adds a new block where the package has a table, helper or fixture for that case. |
| **G9** | `isolated` | The test passes alone and in any order. | It needs another test to run first, or it leaves shared state behind. |

### Exceptions

- **Nil or empty safety.** A "does not panic" test is fine when safety is the contract and the name says so.
- **Golden and table tests.** A helper that compares against expected data is a real assertion.
- **Mocks.** Many mocks are fine when the test asserts an observable outcome. Mock count is not a rule.
- **No public way in.** A test may touch private state when the public API cannot reach the case. Say why in a comment, and keep the touch in one place.

### How to check G5 fails-before-fix

Read the fix. Find the input that triggered the bug.
Run the test on the code without the fix. Use [Run a mutation safely](#run-a-mutation-safely).

### How to check G6 covers-edges

List every default value, cap and error path in the code and its docs.
For each one, name the test that would fail if it changed. If none would fail, the tests are not done. Add one.
G6 judges the set of tests, not one test. A gap never cuts a test that passes G1 to G5 and G7 to G9. Keep that test and report the gap next to it.

### How to check G7 public-behavior

Ask: if someone rewrote the function with the same inputs and outputs, would this test still pass?
If not, the test checks structure. Set up and assert through the public API instead.

### How to check G8 repo-conventions

A new test must kill a mutant that no existing test kills.
Break the code the test guards. If an existing test fails too, the new test adds nothing. Cut it, or move its input into the existing test.
When the package has a table test for the function, add a row. Do not add a new test function next to it.
Use the helpers the other tests use, for example a flag parser or a fixture builder.

### How to check G9 isolated

Run the test alone, for example `go test -run '^TestName$'`, `vitest -t "name"` or `pytest path::name`.
Run the package in shuffled order, for example `go test -shuffle=on -count=3`.
If it fails either way, find the shared state and set it up inside the test.

## Audit mode

A deleted test often looks redundant, but is only partly covered elsewhere. The gap is invisible until a bug ships.
So default to `block`. Allow a deletion only when all checks pass.

| Rule | Name | Check |
|---|---|---|
| **A1** | `names-keeper` | The proposal names a **keeper**: `file::test name`. The keeper checks **the same case** with the same or a stronger assertion. "Covered by the integration suite" is not a keeper. |
| **A2** | `keeps-negatives` | Negative cases and defaults keep a test. Check rejections, error paths, "does not" cases and default values first. Keepers usually cover the happy path only. |
| **A3** | `mutation-proof` | Prove it with a mutation. Break the production code in the way the deleted test guards. The keeper must fail. Use [Run a mutation safely](#run-a-mutation-safely). |
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

## Run a mutation safely

A mutation is proof only if it applied, the code built, and the tree is the same after.
Uncommitted work is the user's. Never lose it.

1. **Use a tool when one exists**, scoped to the diff: `cargo mutants --in-diff`, `gremlins unleash --diff`, `stryker run --mutate <file>:<lines>`, `mutmut run`, PIT with history. They work in a copy.
2. **By hand, save first.** Copy each file you will change: `cp file /tmp/file.orig`. Never restore with `git checkout`, `git restore` or `git stash`. They drop uncommitted edits.
3. **Check that it applied.** `diff /tmp/file.orig file` shows exactly your change. An empty diff or an empty file list means the run proves nothing.
4. **Check that it built.** A compile error is not a failing test. Read the output and find the test name in the failure.
5. **Restore and compare.** `cp /tmp/file.orig file`, then `cmp` the two. Run the tests once more. They must pass.

When a mutant survives, decide if it is equivalent: no input can tell the two versions apart. Say why. Only a mutant that is not equivalent is a gap.

## Output

One row per test. Every verdict names the mutation or run that proves it.

| Test | Verdict | Rule | Proof |
|---|---|---|---|
| `file::name` | keep / cut / allow / block | code and name, or none | the mutation or run, and its result |

When you cut, add the fix in one line. For an audit, add the evidence records, then a summary: allowed, blocked, and why.
