package consistent

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	keys []int
	// key  :hash value
	// value:node name
	hashMap map[int]string

	replicas int

	hash Hash
}

const defaultReplicas = 30

func NewMap(hash Hash) *Map {
	m := &Map{
		hashMap:  make(map[int]string),
		replicas: defaultReplicas,
	}
	if hash == nil {
		m.hash = crc32.ChecksumIEEE
	} else {
		m.hash = hash
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(key + strconv.Itoa(i))))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	index := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[index%len(m.keys)]]
}
