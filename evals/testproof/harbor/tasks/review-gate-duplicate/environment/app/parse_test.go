package durations

import (
	"errors"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in      string
		want    time.Duration
		wantErr error
	}{
		{in: "30s", want: 30 * time.Second},
		{in: "5m", want: 5 * time.Minute},
		{in: "-1s", wantErr: ErrNegative},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Parse(tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Parse(%q) error = %v, want %v", tt.in, err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("Parse(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
