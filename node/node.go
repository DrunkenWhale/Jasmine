package node

import (
	"Jasmine/cache"
	"fmt"
	"time"
)

type Node struct {
	name     string
	cache    *cache.Cache
	callback Callback
}

type Callback func(key string) ([]byte, error)

type OOMError struct {
	name string
}

func (e *OOMError) Error() string {
	return fmt.Sprintf("Node %v: Cache Out of Memory", e.name)
}

const defaultRespiration = int64(time.Hour * 3)

func NewNode(name string, maxMemory int, callback Callback) *Node {
	return &Node{
		name:     name,
		cache:    cache.NewCache(maxMemory),
		callback: callback,
	}
}

func (node *Node) autoClearExpireCache(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		go func() {
			node.cache.CleanExpireCache()
		}()
	}
}

func (node *Node) Name() string {
	return node.name
}

func (node *Node) Get(key string) ([]byte, error) {
	res, mark := node.cache.Get(key)
	if mark {
		return res, nil
	} else {
		r, err := node.callback(key)
		if err != nil {
			return nil, err
		} else {
			err := node.Put(key, r, defaultRespiration)
			if err != nil {
				return r, nil
			} else {
				return nil, &OOMError{node.name}
			}
		}
	}
	//return node.cache.Get(key)
}

func (node *Node) Put(key string, value []byte, respiration int64) error {
	f := node.cache.Put(key, value, respiration)
	if f {
		return nil
	} else {
		return &OOMError{node.name}
	}
}
