package cache

// Purge removes every entry and calls onEvict for each removed key.
func (c *Cache) Purge(onEvict func(key string)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.items {
		delete(c.items, k)
		onEvict(k)
	}
}
