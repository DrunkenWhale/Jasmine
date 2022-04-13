package manage

import (
	"Jasmine/consistent"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// manager is not a node
// it not include any data
// but user will use key find data from manager

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

func (m *Manager) Query(key string) ([]byte, error) {
	bytes, err := m.getValueFromRemoteNode(m.FindNode(key), key)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func (m *Manager) FindNode(key string) string {
	// node name
	return m.hashMap.Get(key)
}

func (m *Manager) getValueFromRemoteNode(nodeAddress string, key string) ([]byte, error) {
	url := fmt.Sprintf("%v/%v/?key=%v", nodeAddress, "__jasmine__", key)
	get, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()
	if get.StatusCode != 200 {
		log.Printf("[node %v] return %v", nodeAddress, get.StatusCode)
		return nil, nil
	} else {
		return ioutil.ReadAll(get.Body)
	}
}

func (m *Manager) StartManageServer(host string) {
	http.HandleFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		bytes, err := m.Query(key)
		if err != nil {
			http.Error(writer, err.Error(), 500)
			return
		}
		writer.Header().Set("Content-Type", "application/octet-stream")
		_, err = writer.Write(bytes)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(host, nil)
}
