package retry

import (
	"testing"
	"time"
)

func TestNextDelay_ZeroPolicyUsesDefaults(t *testing.T) {
	cases := []struct {
		attempt int
		want    time.Duration
		ok      bool
	}{
		{1, 100 * time.Millisecond, true},
		{2, 200 * time.Millisecond, true},
		{3, 400 * time.Millisecond, true},
		{4, 800 * time.Millisecond, true},
		{5, 0, false},
	}
	for _, c := range cases {
		got, ok := NextDelay(Policy{}, c.attempt, 0)
		if got != c.want || ok != c.ok {
			t.Errorf("attempt %d: got (%v, %v), want (%v, %v)", c.attempt, got, ok, c.want, c.ok)
		}
	}
}

func TestNextDelay_AttemptBelowOneIsFirst(t *testing.T) {
	for _, a := range []int{0, -3} {
		if got, ok := NextDelay(Policy{Base: time.Second, MaxAttempts: 3}, a, 0); got != time.Second || !ok {
			t.Errorf("attempt %d: got (%v, %v)", a, got, ok)
		}
	}
}

func TestNextDelay_DefaultMaxCapsGrowth(t *testing.T) {
	got, ok := NextDelay(Policy{Base: 3 * time.Second, MaxAttempts: 10}, 4, 0)
	if got != 10*time.Second || !ok {
		t.Fatalf("got (%v, %v), want 10s", got, ok)
	}
}

func TestNextDelay_CustomMaxCaps(t *testing.T) {
	got, _ := NextDelay(Policy{Base: time.Second, Max: 5 * time.Second, MaxAttempts: 10}, 4, 0)
	if got != 5*time.Second {
		t.Fatalf("got %v, want 5s", got)
	}
}

func TestNextDelay_RetryAfter(t *testing.T) {
	p := Policy{Base: time.Second, MaxAttempts: 10}
	if got, _ := NextDelay(p, 1, 3*time.Second); got != 3*time.Second {
		t.Errorf("longer hint: got %v, want 3s", got)
	}
	if got, _ := NextDelay(p, 3, 2*time.Second); got != 4*time.Second {
		t.Errorf("shorter hint: got %v, want 4s", got)
	}
	if got, _ := NextDelay(p, 1, -5*time.Second); got != time.Second {
		t.Errorf("negative hint: got %v, want 1s", got)
	}
	if got, _ := NextDelay(p, 1, 2*time.Minute); got != 60*time.Second {
		t.Errorf("default hint cap: got %v, want 60s", got)
	}
	if got, _ := NextDelay(Policy{MaxRetryAfter: 30 * time.Second}, 1, 2*time.Minute); got != 30*time.Second {
		t.Errorf("custom hint cap: got %v, want 30s", got)
	}
}
