package cache

import "time"

type Value struct {
	value      interface{}
	expiration int64
}

func NewValue(value interface{}, expiration int64) *Value {
	return &Value{
		value:      value,
		expiration: expiration + time.Now().Unix(),
	}
}
