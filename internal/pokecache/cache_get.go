package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.entries[key]
	if ok {
		return value.val, true
	} else {
		return nil, false
	}
}
