package retry

import (
	"testing"
	"time"
)


func TestNextDelay_DefaultSequence(t *testing.T) {
	want := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 400 * time.Millisecond, 800 * time.Millisecond}
	for i, w := range want {
		if got, _ := NextDelay(Policy{}, i+1, 0); got != w {
			t.Errorf("attempt %d: got %v, want %v", i+1, got, w)
		}
	}
}




func TestNextDelay_StopsAfterDefaultMaxAttempts(t *testing.T) {
	if _, ok := NextDelay(Policy{}, 4, 0); !ok {
		t.Fatal("attempt 4 should still retry")
	}
	if _, ok := NextDelay(Policy{}, 5, 0); ok {
		t.Fatal("attempt 5 is the last one")
	}
}

func TestNextDelay_CapsAtMax(t *testing.T) {
	p := Policy{Base: time.Second, Max: 5 * time.Second, MaxAttempts: 10}
	if got, _ := NextDelay(p, 4, 0); got != 5*time.Second {
		t.Fatalf("got %v, want 5s", got)
	}
}

func TestNextDelay_CapsAtDefaultMax(t *testing.T) {
	p := Policy{Base: 3 * time.Second, MaxAttempts: 10}
	if got, _ := NextDelay(p, 4, 0); got != 10*time.Second {
		t.Fatalf("got %v, want 10s", got)
	}
}


func TestNextDelay_RetryAfterWinsWhenLonger(t *testing.T) {
	p := Policy{Base: time.Second, MaxAttempts: 10}
	if got, _ := NextDelay(p, 1, 3*time.Second); got != 3*time.Second {
		t.Fatalf("got %v, want 3s", got)
	}
}


func TestNextDelay_ShortRetryAfterDoesNotShortenDelay(t *testing.T) {
	p := Policy{Base: time.Second, MaxAttempts: 10}
	if got, _ := NextDelay(p, 3, 2*time.Second); got != 4*time.Second {
		t.Fatalf("got %v, want 4s", got)
	}
}

func TestNextDelay_RetryAfterIsCapped(t *testing.T) {
	if got, _ := NextDelay(Policy{}, 1, 5*time.Minute); got != 60*time.Second {
		t.Fatalf("got %v, want 60s", got)
	}
}
