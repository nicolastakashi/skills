/app is a git repo. My uncommitted change fixes a bug in `Clamp` (see `/app/CASE.md`) and adds the regression test `TestClampNegativeReturnsZero` in `quota_test.go`.
Review the test before I commit. Prove that it fails on the code without my fix, and keep my uncommitted work as it is.

Write your decision to `/app/verdict.json` as `{"verdict": "keep" or "cut", "rule": "<short reason or rule id>"}`, then explain it briefly.
