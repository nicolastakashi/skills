Review this pull request. The code in `/app` is the PR branch and `/app/PR.diff` is the change. Comments already on the PR are in `/app/comments.json`.

Nothing can be posted from here. Write your review to `/app/review.json`:

```json
{"findings": [{"id": "", "file": "example.go", "line": 12, "verdict": "fix", "evidence": "...", "comment": "..."}]}
```

- `verdict`: `fix` (confirmed bug), `decide` (needs a human or product decision), `keep` (leave as is), or `drop` (checked, not a problem).
- `evidence`: what proves the verdict, for example `example.go:12`, a test name, or a command and its output.
- `comment`: the text to post on the PR. Leave it empty when nothing should be posted.
