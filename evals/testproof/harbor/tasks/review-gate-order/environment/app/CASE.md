New test for `Lookup` in `registry.go`, to be added at the end of `registry_test.go`:

```go
func TestLookupReturnsRegisteredEncoder(t *testing.T) {
	e, ok := Lookup("json")
	if !ok {
		t.Fatal("json is not registered")
	}
	if got := e("x"); got != `"x"` {
		t.Fatalf("encoder(x) = %s, want \"x\"", got)
	}
}
```

`go test ./...` passes with it.
