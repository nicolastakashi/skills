New test for negative input in `parse.go`, to be added to `parse_test.go`:

```go
func TestParseRejectsNegativeMinutes(t *testing.T) {
	_, err := Parse("-5m")
	if !errors.Is(err, ErrNegative) {
		t.Fatalf("Parse(-5m) error = %v, want ErrNegative", err)
	}
}
```
