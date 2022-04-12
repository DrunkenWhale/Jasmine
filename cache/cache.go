package cache

import "time"

type Cache struct {
	core map[string]*Value
}

func NewCache() *Cache {
	return &Cache{
		core: make(map[string]*Value),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
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

func (c *Cache) Put(key string, value []byte, expiration int64) {
	c.core[key] = NewValue(value, expiration)
}
