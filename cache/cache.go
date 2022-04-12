package cache

import "time"

type Cache struct {
	core      map[string]*Value
	memory    int
	maxMemory int
}

func NewCache(maxMemory int) *Cache {
	return &Cache{
		core:      make(map[string]*Value),
		memory:    0,
		maxMemory: maxMemory,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if res, ok := c.core[key]; ok {
		if time.Now().Unix() > res.expiration {
			delete(c.core, key)
			c.memory -= len(res.value)
			return nil, false
		}
		return res.value, true
	} else {
		return nil, false
	}
}

func (c *Cache) Put(key string, value []byte, expiration int64) {
	c.core[key] = NewValue(value, expiration)
	c.memory += len(value)
}

func (c *Cache) CacheMemory() int {
	return c.memory
}

func (c *Cache) MaxCacheMemory() int {
	return c.maxMemory
}
