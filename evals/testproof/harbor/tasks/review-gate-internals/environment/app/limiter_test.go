package limiter

import "testing"

func TestAllowStopsAtMax(t *testing.T) {
	l := New(2)
	got := []bool{l.Allow("a"), l.Allow("a"), l.Allow("a")}
	want := []bool{true, true, false}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("call %d: Allow = %v, want %v", i+1, got[i], want[i])
		}
	}
}
