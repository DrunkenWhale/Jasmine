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

type Callback func(key string) (interface{}, error)

type OOMError struct {
	name string
}

func (e *OOMError) Error() string {
	return fmt.Sprintf("Node %v: Cache Out of Memory", e.name)
}

const defaultAutoClearTime = time.Hour * 1
const defaultRespiration = time.Hour * 3

func NewNode(name string, maxMemory int, callback Callback) *Node {
	node := &Node{
		name:     name,
		cache:    cache.NewCache(maxMemory),
		callback: callback,
	}
	go node.autoClearExpireCache(defaultAutoClearTime)
	return node
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

func (node *Node) Get(key string) (interface{}, error) {
	res, mark := node.cache.Get(key)
	if mark {
		return res, nil
	} else {
		r, err := node.callback(key)
		if err != nil {
			return nil, err
		} else {
			node.Put(key, r, defaultRespiration)
			return r, nil
		}
	}
	//return node.cache.Get(key)
}

func (node *Node) Put(key string, value interface{}, respiration time.Duration) {
	node.cache.Put(key, value, int64(respiration))
}
