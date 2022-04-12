package node

import "Jasmine/cache"

type Node struct {
	cache *cache.Cache
}

func (node *Node) Get(key string) ([]byte, bool) {
	return node.cache.Get(key)
}

func (node *Node) Put(key string, value []byte, respiration int64) {
	node.cache.Put(key, value, respiration)
}

