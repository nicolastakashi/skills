package cache

import "testing"

func TestGetMissingKey(t *testing.T) {
	c := New()
	if _, ok := c.Get("nope"); ok {
		t.Fatal("Get on an empty cache should miss")
	}
}

func TestPurgeCallsOnEvict(t *testing.T) {
	c := New()
	c.items["a"] = entry{value: "1"}
	var evicted []string
	c.Purge(func(k string) { evicted = append(evicted, k) })
	if len(evicted) != 1 || evicted[0] != "a" {
		t.Fatalf("evicted = %v", evicted)
	}
}
