// Package quota keeps usage numbers inside a quota.
package quota

// Clamp limits n to the range 0..max.
func Clamp(n, max int) int {
	if n < 0 {
		return 0
	}
	if n > max {
		return max
	}
	return n
}
