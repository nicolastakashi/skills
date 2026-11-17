// Package durations parses short duration strings such as "30s" or "5m".
package durations

import (
	"errors"
	"strconv"
	"time"
)

var ErrNegative = errors.New("durations: negative duration")

var units = map[byte]time.Duration{'s': time.Second, 'm': time.Minute, 'h': time.Hour}

// Parse reads a number followed by one unit: s, m or h.
func Parse(s string) (time.Duration, error) {
	if len(s) < 2 {
		return 0, errors.New("durations: too short")
	}
	unit, ok := units[s[len(s)-1]]
	if !ok {
		return 0, errors.New("durations: unknown unit")
	}
	n, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, ErrNegative
	}
	return time.Duration(n) * unit, nil
}
