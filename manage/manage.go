package manage

import "Jasmine/consistent"

type Manager struct {

	// key 	 : node name
	// value : node host
	nodes map[string]string

	hashMap *consistent.Map
}

func NewManger() *Manager {
	return &Manager{
		nodes:   make(map[string]string),
		hashMap: consistent.NewMap(nil),
	}
}

func (m *Manager) AddNode(nodeName string, nodeAddress string) {
	m.nodes[nodeName] = nodeAddress
}

func (m *Manager) Register() {
	var t []string
	for k, _ := range m.nodes {
		t = append(t, k)
	}
	m.hashMap.Add(t...)
}

func (m *Manager) FindNode(key string) string {
	return m.hashMap.Get(key)
}
