package cache

import (
	"sync"
	"time"
)

type Cache struct {
	core      map[string]*Value
	memory    int
	maxMemory int
	mutex     sync.Mutex
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

func (c *Cache) Put(key string, value []byte, expiration int64) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	size := len(value)
	v, t := c.core[key]
	// key exist
	if t {
		if c.maxMemory-c.memory < len(value)-len(v.value) {
			return false
		} else {
			c.memory = c.memory + len(v.value) - len(value)
			c.core[key] = NewValue(value, expiration)
			return true
		}

	}
	if c.maxMemory-c.memory < size {
		return false
	}
	c.core[key] = NewValue(value, expiration)
	c.memory += size
	return true
}

func (c *Cache) CacheMemory() int {
	return c.memory
}

func (c *Cache) MaxCacheMemory() int {
	return c.maxMemory
}

func (c *Cache) CleanExpireCache() {
	for k, v := range c.core {
		if v.expiration < time.Now().Unix() {
			delete(c.core, k)
		}
	}
}
