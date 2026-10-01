package retry

import "time"

// Policy controls how a client retries failed requests.
// A zero value is valid and uses the defaults below.
type Policy struct {
	Base          time.Duration // first delay. Default 100ms.
	Max           time.Duration // cap for the exponential delay. Default 10s.
	MaxAttempts   int           // total attempts, including the first one. Default 5.
	MaxRetryAfter time.Duration // cap for a server Retry-After hint. Default 60s.
}

const (
	defaultBase          = 100 * time.Millisecond
	defaultMax           = 10 * time.Second
	defaultMaxAttempts   = 5
	defaultMaxRetryAfter = 60 * time.Second
)

func (p Policy) withDefaults() Policy {
	if p.Base <= 0 {
		p.Base = defaultBase
	}
	if p.Max <= 0 {
		p.Max = defaultMax
	}
	if p.MaxAttempts <= 0 {
		p.MaxAttempts = defaultMaxAttempts
	}
	if p.MaxRetryAfter <= 0 {
		p.MaxRetryAfter = defaultMaxRetryAfter
	}
	return p
}

// NextDelay returns how long to wait before the next attempt, after attempt
// number `attempt` failed (attempts start at 1). It returns ok=false when no
// attempts are left.
//
// The delay is Base * 2^(attempt-1), capped at Max. A positive Retry-After
// hint, capped at MaxRetryAfter, replaces the delay when it is longer.
func NextDelay(p Policy, attempt int, retryAfter time.Duration) (delay time.Duration, ok bool) {
	p = p.withDefaults()
	if attempt < 1 {
		attempt = 1
	}
	if attempt >= p.MaxAttempts {
		return 0, false
	}
	delay = p.Base
	for i := 1; i < attempt && delay < p.Max; i++ {
		delay *= 2
	}
	if delay > p.Max {
		delay = p.Max
	}
	if retryAfter > 0 {
		hint := retryAfter
		if hint > p.MaxRetryAfter {
			hint = p.MaxRetryAfter
		}
		if hint > delay {
			delay = hint
		}
	}
	return delay, true
}
