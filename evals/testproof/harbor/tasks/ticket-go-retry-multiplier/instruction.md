RET-15: Make the backoff growth configurable

The Go package in /app has the retry policy in `policy.go`.
The delay always doubles on each attempt. Some clients need slower growth.

- `Policy` gets a `Multiplier float64` field. The delay is multiplied by it on each attempt. The default is 2, so current behavior stays the same.
- A `Multiplier` below 1 uses the default.
- The `Max` cap and the Retry-After rules do not change.
- Update the doc comments.

Make sure the package builds in /app before you finish.
