package stats

import (
	"slices"
	"testing"
)

func TestPercentiles(t *testing.T) {
	values := []float64{5, 1, 4, 2, 3}
	got := Percentiles(values, 0.2, 1, 0.6)
	if want := []float64{1, 5, 3}; !slices.Equal(got, want) {
		t.Fatalf("Percentiles = %v, want %v", got, want)
	}
	if want := []float64{5, 1, 4, 2, 3}; !slices.Equal(values, want) {
		t.Fatalf("input changed to %v", values)
	}
}
