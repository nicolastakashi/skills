package registry

import (
	"errors"
	"testing"
)

func TestRegisterRejectsDuplicate(t *testing.T) {
	if err := Register("json", Quote); err != nil {
		t.Fatalf("first Register: %v", err)
	}
	if err := Register("json", Quote); !errors.Is(err, ErrDuplicate) {
		t.Fatalf("second Register error = %v, want ErrDuplicate", err)
	}
}
