New test for `Reset` in `limiter.go`, to be added to `limiter_test.go`:

```go
func TestResetClearsKey(t *testing.T) {
	l := New(2)
	l.counts["a"] = 2
	l.Reset("a")
	if _, ok := l.counts["a"]; ok {
		t.Fatal("key a still has a count after Reset")
	}
}
```
