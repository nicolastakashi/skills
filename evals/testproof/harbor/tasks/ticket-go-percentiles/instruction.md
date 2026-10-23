STAT-7: Compute several percentiles in one call

The Go package in /app has `Percentile` in `percentile.go`.
Dashboards ask for p50, p90 and p99 together, and they call `Percentile` three times.

- Add `Percentiles(values []float64, ps ...float64) []float64`.
- It returns one result for each `p`, in the same order.
- Each result follows the same rules as `Percentile`, including clamping and empty input.
- The input slice is never modified.

Make sure the package builds in /app before you finish.
