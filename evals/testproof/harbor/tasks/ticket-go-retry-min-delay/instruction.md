RET-12: Add a minimum delay to the retry policy

The Go package in /app has the retry policy in `policy.go`.
Some upstream servers reject retries that come too fast, so we need a floor on the delay.

- `Policy` gets a `Min time.Duration` field. The default is 0, which means no floor.
- `NextDelay` never returns a delay below `Min`. This includes the case where a Retry-After hint is shorter than `Min`.
- If `Min` is greater than `Max`, use `Max` as the floor.
- Update the doc comments.

Make sure the package builds in /app before you finish.
