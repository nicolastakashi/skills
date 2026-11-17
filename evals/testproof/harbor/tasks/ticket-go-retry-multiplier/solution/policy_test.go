package retry

import (
	"math"
	"testing"
	"time"
)

func TestNextDelayMultiplier(t *testing.T) {
	ms := time.Millisecond
	cases := []struct {
		name    string
		policy  Policy
		attempt int
		hint    time.Duration
		want    time.Duration
	}{
		{"zero uses default 2", Policy{Base: 100 * ms}, 3, 0, 400 * ms},
		{"below 1 uses default", Policy{Base: 100 * ms, Multiplier: 0.5}, 3, 0, 400 * ms},
		{"negative uses default", Policy{Base: 100 * ms, Multiplier: -3}, 3, 0, 400 * ms},
		{"NaN uses default", Policy{Base: 100 * ms, Multiplier: math.NaN()}, 3, 0, 400 * ms},
		{"exactly 1 keeps the delay", Policy{Base: 100 * ms, Multiplier: 1}, 3, 0, 100 * ms},
		{"fraction grows", Policy{Base: 100 * ms, Multiplier: 1.5}, 3, 0, 225 * ms},
		{"three triples", Policy{Base: 100 * ms, Multiplier: 3}, 3, 0, 900 * ms},
		{"capped at Max", Policy{Base: 100 * ms, Multiplier: 3, Max: 500 * ms}, 4, 0, 500 * ms},
		{"longer hint wins", Policy{Base: 100 * ms, Multiplier: 3}, 2, time.Second, time.Second},
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
