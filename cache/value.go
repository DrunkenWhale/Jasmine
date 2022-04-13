package cache

import "time"

type Value struct {
	value      []byte
	expiration int64
}

func NewValue(value []byte, expiration int64) *Value {
	return &Value{
		value:      value,
		expiration: expiration + time.Now().Unix(),
	}
}
