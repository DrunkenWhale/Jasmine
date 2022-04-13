package cache

import "time"

type Cache struct {
	core      map[string]*Value
}

func NewCache(maxMemory int) *Cache {
	return &Cache{
		core:      make(map[string]*Value),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if res, ok := c.core[key]; ok {
		if time.Now().Unix() > res.expiration {
			delete(c.core, key)
			return nil, false
		}
		return res.value, true
	} else {
		return nil, false
	}
}

func (c *Cache) Put(key string, value interface{}, expiration int64)  {
	c.core[key] = NewValue(value, expiration)
}

func (c *Cache) CleanExpireCache() {
	for k, v := range c.core {
		if v.expiration < time.Now().Unix() {
			delete(c.core, k)
		}
	}
}
