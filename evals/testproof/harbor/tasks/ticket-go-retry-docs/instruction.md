RET-16: Fix typos in the retry policy docs

The Go package in /app has the retry policy in `policy.go`.
A reviewer found three typos in the doc comments:

- "retires" should be "retries".
- "defualts" should be "defaults".
- "then that delay" should be "than that delay".

Fix only the comments. The code does not change.

Make sure the package builds in /app before you finish.
