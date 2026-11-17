New test for the hours unit in `parse.go`, to be added to `parse_test.go`:

```go
func TestParseHours(t *testing.T) {
	d, err := Parse("2h")
	if err != nil {
		t.Fatal(err)
	}
	if d != 2*time.Hour {
		t.Fatalf("got %v, want 2h", d)
	}
}
```
