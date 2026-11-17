// Package limiter allows each key a fixed number of calls until it is reset.
package limiter

import "sync"

type Limiter struct {
	mu     sync.Mutex
	max    int
	counts map[string]int
}

func New(max int) *Limiter {
	return &Limiter{max: max, counts: map[string]int{}}
}

// Allow records one call for key and reports whether it is within the limit.
func (l *Limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.counts[key] >= l.max {
		return false
	}
	l.counts[key]++
	return true
}

// Reset gives key its full limit again.
func (l *Limiter) Reset(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.counts, key)
}
