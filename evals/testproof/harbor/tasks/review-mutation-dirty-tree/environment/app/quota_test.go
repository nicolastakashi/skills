package quota

import "testing"

func TestClampNegativeReturnsZero(t *testing.T) {
	if got := Clamp(-3, 10); got != 0 {
		t.Fatalf("Clamp(-3, 10) = %d, want 0", got)
	}
}
