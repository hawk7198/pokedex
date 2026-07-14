package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return &c
}
