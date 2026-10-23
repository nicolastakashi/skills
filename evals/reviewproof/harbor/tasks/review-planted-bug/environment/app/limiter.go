// Package limiter caps how many requests each user can make.
package limiter

import "time"

// Limiter allows at most Max requests per user in each fixed window.
// A new window starts when the current one has run for Window.
type Limiter struct {
	Max    int
	Window time.Duration

	windowStart time.Time
	counts      map[string]int
}

// New returns a Limiter with the given cap and window.
func New(max int, window time.Duration) *Limiter {
	return &Limiter{Max: max, Window: window, counts: map[string]int{}}
}

// Allow reports whether user may make another request at now.
func (l *Limiter) Allow(user string, now time.Time) bool {
	if l.windowStart.IsZero() {
		l.windowStart = now
	}
	if now.Sub(l.windowStart) >= l.Window {
		l.windowStart = now
	}
	if l.counts[user] >= l.Max {
		return false
	}
	l.counts[user]++
	return true
}
