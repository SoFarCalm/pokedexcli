package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	for k, v := range c.data {
		if k == key {
			return v.val, true
		}
	}
	return nil, false
}
