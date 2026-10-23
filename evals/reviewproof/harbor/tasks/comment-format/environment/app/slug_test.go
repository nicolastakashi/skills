package slug

import "testing"

func TestMakeTwoWords(t *testing.T) {
	if got := Make("Hello World"); got != "hello-world" {
		t.Fatalf("Make = %q", got)
	}
}
