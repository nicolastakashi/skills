This is my pull request. The code in `/app` is my branch and `/app/PR.diff` is the change. A reviewer left the findings in `/app/findings.json`. Tell me what to do with each one. Do not change the code.

Write `/app/review.json` with one entry per finding, using the finding's `id`:

```json
{"findings": [{"id": "r1", "file": "example.go", "line": 12, "verdict": "fix", "evidence": "...", "comment": "..."}]}
```

- `verdict`: `fix` (confirmed bug that this PR should fix), `decide` (needs a human or product decision), `keep` (leave as is in this PR), or `drop` (checked, not a problem).
- `evidence`: what proves the verdict. It must cite `file:line`, a test name, or a command and its output.
- `comment`: the reply to post on the reviewer's thread. Leave it empty when no reply is needed.
