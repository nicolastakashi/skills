---
name: reviewproof
description: Use when reviewing a pull request or diff, or when review findings arrive for your own PR from a reviewer, a bot, or another agent, including requests like "review this PR", "review my diff", "address the review comments", or "triage these findings". Verifies every finding with evidence, drops what does not reproduce, skips what is already raised, and drafts ready-to-post comments. Any language.
---

# reviewproof

A review comment earns its place when it is true, new, and worth acting on.
This skill has two modes:

- **Review**: someone else's PR or diff. You find the problems.
- **Receive**: findings that arrive for your own PR, from a reviewer, a bot, or another agent. You decide what to do with each one.

Both modes run the same five steps: find, verify, sort, draft, post.

## 1. Find

**Review mode.** Start from the host agent's review tool. If its findings are already in the conversation, for example from a `/code-review` the user ran, use them. Otherwise run it now. In Claude Code that is `/code-review`. Never pass `--comment`, because this skill posts only in step 5. Run it even when the diff is small, because it finds bugs that one read of the diff misses. A small diff or saving time is not a reason to skip it. Review the diff on your own only when the host has no review tool:

1. Read the diff, then read each changed function in full.
2. Write down the contract of each one: its doc comment, its name, and the units, limits and defaults it promises.
3. Check the code against that contract: boundaries, units, conditions, error paths, shared state, and resource cleanup.

Keep every candidate, including the ones that only look wrong. Step 2 decides.

**Receive mode.** The findings are your candidates. Do not add new ones unless the user asks.

In both modes, read the comments already on the PR, from people and bots, before step 3.

## 2. Verify

Every candidate needs evidence. A candidate with no evidence is a `drop`.

| Strength | Evidence |
|---|---|
| Best | A test or command that fails because of the problem. Paste the command and the failing line. |
| Good | `file:line` of the cause and `file:line` of the contract it breaks, for example a doc comment. |
| Not enough | "Looks wrong", "might race", "probably". |

- Write throwaway tests outside the change, run them, then delete them. Do not leave them behind unless the user asks.
- Check the language and tool version before you call something a bug. Semantics change between versions. For example, Python dicts keep insertion order only from 3.7.
- Prove or disprove concurrency claims with the language's race detector, or by citing the synchronization at `file:line`.
- When there are more than 3 candidates and the host supports subagents, verify each one in its own subagent. Otherwise verify them in order, in this session.

## 3. Sort

Give each candidate one verdict.

| Verdict | Meaning | Needs |
|---|---|---|
| `fix` | A change this PR should make to code the diff touches: a bug, a missing test for behavior the PR changes, or a wrong doc. | Evidence that shows it. The comment label says how serious it is. |
| `decide` | Real, but the right answer is a human or product choice. | The options, and who should choose. |
| `keep` | Leave the code as is. | The reason, for example "not in this diff, follow-up". |
| `drop` | Checked, and it is not a problem. | What you checked, with `file:line` or a command. |

- **Already raised.** If an existing comment covers the same root cause, do not raise it again. Match on file and cause, not on wording. Leave it out of the output, or mention it only in the summary.
- **Out of scope.** If the problem is in code the diff does not touch, it is `keep`, with a suggested follow-up. It is never `fix` for this PR.
- **Receive mode.** Never `fix` a finding you could not reproduce. Disagree with evidence, not with opinion.

## 4. Draft

Draft a comment for each `fix` and `decide`. In receive mode, draft a reply for every finding.

First read the user's instructions files: `AGENTS.md`, `CLAUDE.md`, or the host's equivalent, in the repo and in the user's home. Their rules on length, tone and words win over the defaults below.

Review comments:

| Rule | Check |
|---|---|
| **D1** `label` | Starts with a [conventional comments](https://conventionalcomments.org/) label: `issue:`, `suggestion:`, `question:`, `nitpick:`, `todo:`, `thought:`, `chore:`, `note:` or `praise:`. |
| **D2** `one-line-title` | The first line is the label and a title of 70 characters or less. A blank line, then the body. |
| **D3** `one-impact` | The body names one impact. Cut the follow-on consequences. |
| **D4** `cause-and-effect` | Cause and effect are in the same sentence: "X does Y, so Z". |
| **D5** `cites-line` | The body cites the `file:line` of the cause. |
| **D6** `plain` | No em dashes, no bold, no headings, no hedges. If you cannot show the impact, drop the comment. |

A comment that passes D1 to D6:

```
issue: Retry ignores the max delay

backoff.go:41 doubles the delay without applying MaxDelay, so after 8 retries a client waits 4 minutes instead of the documented 30 seconds.
```

Replies in receive mode: one or two sentences. Say what you will do and why, and cite `file:line` when you disagree. No label needed.

If a humanizer skill is installed, run it on each draft. Then check D1 to D6 again.

## 5. Post

Post only after the user says yes. Show the drafts first.

- Post `fix` and `decide` comments on their `file:line`.
- Do not post `keep` or `drop`. Report them to the user.
- In receive mode, resolve a thread only after the fix is on the branch.

## Output

Show a table to the user: `file:line`, verdict, evidence, and the draft comment. Under it, add this legend: `fix` = change it in this PR, `decide` = the author chooses, `keep` = leave it for now, `drop` = checked, not a problem.

Say whether the review tool ran. If it did not, say why.

When a machine reads the result (an eval, CI, another agent), write JSON:

```json
{"findings": [{"id": "", "file": "a.go", "line": 12, "verdict": "fix", "evidence": "a.go:12 ...", "comment": "issue: ...\n\n..."}]}
```

- `id`: the incoming finding's id in receive mode, otherwise empty.
- `comment`: the text to post, or empty when nothing should be posted. Keep the D2 shape: the label and title, a blank line (`\n\n`), then the body. The table may shorten it, the JSON may not.
