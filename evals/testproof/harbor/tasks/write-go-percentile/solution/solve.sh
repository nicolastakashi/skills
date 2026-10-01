#!/bin/bash
# Reference tests: kill every mutant.
cat > /app/percentile_test.go <<'GO'
package stats

import (
	"math"
	"testing"
)

func TestPercentile_EmptyReturnsNaN(t *testing.T) {
	if got := Percentile(nil, 0.5); !math.IsNaN(got) {
		t.Fatalf("got %v, want NaN", got)
	}
}

func TestPercentile_NearestRankOnUnsortedInput(t *testing.T) {
	values := []float64{5, 1, 4, 2, 3}
	cases := []struct {
		p    float64
		want float64
	}{{0, 1}, {0.2, 1}, {0.5, 3}, {0.9, 5}, {1, 5}, {1.5, 5}}
	for _, c := range cases {
		if got := Percentile(values, c.p); got != c.want {
			t.Errorf("Percentile(%v, %v) = %v, want %v", values, c.p, got, c.want)
		}
	}
}

func TestPercentile_DoesNotModifyInput(t *testing.T) {
	values := []float64{3, 1, 2}
	Percentile(values, 0.5)
	if values[0] != 3 || values[1] != 1 || values[2] != 2 {
		t.Fatalf("input was modified: %v", values)
	}
}
GO
cd /app && go test ./...
