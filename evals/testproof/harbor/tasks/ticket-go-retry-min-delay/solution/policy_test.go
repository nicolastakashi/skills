package retry

import (
	"testing"
	"time"
)

func TestNextDelayMin(t *testing.T) {
	ms := time.Millisecond
	cases := []struct {
		name    string
		policy  Policy
		attempt int
		hint    time.Duration
		want    time.Duration
	}{
		{"zero Min means no floor", Policy{Base: 100 * ms}, 1, 0, 100 * ms},
		{"Min raises the delay", Policy{Base: 100 * ms, Min: 250 * ms}, 1, 0, 250 * ms},
		{"Min below the delay does nothing", Policy{Base: 100 * ms, Min: 50 * ms}, 2, 0, 200 * ms},
		{"short hint still gets the floor", Policy{Base: 100 * ms, Min: 500 * ms}, 1, 300 * ms, 500 * ms},
		{"Min above Max uses Max", Policy{Base: 100 * ms, Max: 300 * ms, Min: time.Second}, 1, 0, 300 * ms},
		{"floor beats the hint cap", Policy{Base: 100 * ms, Min: 2 * time.Second, MaxRetryAfter: time.Second}, 1, 5 * time.Second, 2 * time.Second},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := NextDelay(c.policy, c.attempt, c.hint)
			if !ok || got != c.want {
				t.Fatalf("NextDelay = %v, %v; want %v, true", got, ok, c.want)
			}
		})
	}
}
