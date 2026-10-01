package stats

import (
	"math"
	"sort"
)

// Percentile returns the p-th percentile of values, using the nearest-rank method.
// p is a fraction between 0 and 1; values outside that range are clamped.
// It returns NaN for empty input. The input slice is never modified.
func Percentile(values []float64, p float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	if p > 1 {
		p = 1
	}
	s := make([]float64, len(values))
	copy(s, values)
	sort.Float64s(s)
	idx := int(math.Ceil(p*float64(len(s)))) - 1
	if idx < 0 {
		idx = 0
	}
	return s[idx]
}
