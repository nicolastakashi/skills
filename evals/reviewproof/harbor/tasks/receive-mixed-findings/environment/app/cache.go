// Package cache is an in-memory key-value cache.
package cache

import (
	"sync"
	"time"
)

type entry struct {
	value   string
	expires time.Time
}

// Cache is safe for concurrent use.
type Cache struct {
	mu    sync.Mutex
	items map[string]entry
	now   func() time.Time
}

// New returns an empty Cache.
func New() *Cache {
	return &Cache{items: map[string]entry{}, now: time.Now}
}

// Set stores value under key for ttl.
func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = entry{value: value, expires: c.now().Add(ttl)}
}

// Get returns the value for key. Expired entries are removed and reported
// as missing.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[key]
	if !ok {
		return "", false
	}
	if c.now().Before(e.expires) {
		delete(c.items, key)
		return "", false
	}
	return e.value, true
}

// Len returns the number of entries that have not expired.
func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}
