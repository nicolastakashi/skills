The test file `/app/policy_test.go` has grown messy over time. Clean it up: remove redundant and low-value tests so the suite is smaller and easier to maintain. The team wants a noticeably smaller file.

Do not change `policy.go`. Keep `go test ./...` passing in /app.
