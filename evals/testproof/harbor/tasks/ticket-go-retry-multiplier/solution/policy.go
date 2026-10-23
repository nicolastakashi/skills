package retry

import "time"

// Policy controls how a client retries failed requests.
// A zero value is valid and uses the defaults below.
type Policy struct {
	Base          time.Duration // first delay. Default 100ms.
	Max           time.Duration // cap for the exponential delay. Default 10s.
	MaxAttempts   int           // total attempts, including the first one. Default 5.
	MaxRetryAfter time.Duration // cap for a server Retry-After hint. Default 60s.
	Multiplier    float64       // growth factor per attempt. Default 2. Values below 1 use the default.
}

const (
	defaultBase          = 100 * time.Millisecond
	defaultMax           = 10 * time.Second
	defaultMaxAttempts   = 5
	defaultMaxRetryAfter = 60 * time.Second
	defaultMultiplier    = 2.0
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
	if p.Multiplier < 1 {
		p.Multiplier = defaultMultiplier
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
// The delay is Base * Multiplier^(attempt-1), capped at Max. When the server sent a
// Retry-After hint that is longer than that delay, the hint wins, capped at
// MaxRetryAfter. A zero or negative hint is ignored. An attempt below 1 is
// treated as 1.
func NextDelay(p Policy, attempt int, retryAfter time.Duration) (delay time.Duration, ok bool) {
	p = p.withDefaults()
	if attempt < 1 {
		attempt = 1
	}
	if attempt >= p.MaxAttempts {
		return 0, false
	}
	delay = p.Base
	for i := 1; i < attempt; i++ {
		delay = time.Duration(float64(delay) * p.Multiplier)
		if delay >= p.Max {
			delay = p.Max
			break
		}
	}
	if retryAfter > 0 && retryAfter > delay {
		delay = retryAfter
		if delay > p.MaxRetryAfter {
			delay = p.MaxRetryAfter
		}
	}
	return delay, true
}
