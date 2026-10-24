package limiter

import (
	"testing"
	"time"
)

func TestAllowCapsRequestsInOneWindow(t *testing.T) {
	l := New(2, time.Minute)
	now := time.Unix(0, 0)
	for i, want := range []bool{true, true, false} {
		if got := l.Allow("ana", now); got != want {
			t.Fatalf("request %d: Allow = %v, want %v", i+1, got, want)
		}
	}
}

func TestAllowCountsUsersSeparately(t *testing.T) {
	l := New(1, time.Minute)
	now := time.Unix(0, 0)
	if !l.Allow("ana", now) || !l.Allow("ben", now) {
		t.Fatal("each user should get their own first request")
	}
}
