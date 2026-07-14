package pokecache

import (
	"time"
)

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.entries {
			if value.createdAt.Before(time.Now().Add(-c.interval)) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
