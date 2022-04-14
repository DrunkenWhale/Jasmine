package manage

import (
	"Jasmine/consistent"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	host := m.nodes[m.FindNode(key)]
	bytes, err := m.getValueFromRemoteNode(host, key)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func (m *Manager) FindNode(key string) string {
	// node name
	nodeName := m.hashMap.Get(key)
	log.Printf("[Manager] find key in [Node: %v]", nodeName)
	return nodeName
}

func (m *Manager) getValueFromRemoteNode(nodeAddress string, key string) ([]byte, error) {
	url := fmt.Sprintf("%v/%v/?key=%v", nodeAddress, "__jasmine__", key)
	get, err := http.Get(url)
	if err != nil {
		log.Printf("[Manager] %v", err)
		return nil, err
	}
	defer get.Body.Close()
	if get.StatusCode != 200 {
		log.Printf("[node %v] return %v", nodeAddress, get.StatusCode)
		return nil, &NodeNoResponse{}
	} else {
		bytes, err := ioutil.ReadAll(get.Body)
		if err != nil {
			return nil, nil
		}
		return bytes, nil
	}
}

func (m *Manager) StartManageServer(host string) {
	http.HandleFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		log.Printf("[Manager] search key: %v", key)
		bytes, err := m.Query(key)
		if err != nil {
			log.Printf("[Manager] %v", err)
			http.Error(writer, err.Error(), 404)
			return
		}
		writer.Header().Set("Content-Type", "application/octet-stream")
		_, err = writer.Write(bytes)
		if err != nil {
			log.Printf("[Manager] %v", err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	if strings.HasPrefix(host, "localhost") {
		log.Printf("[Manager] start listen [ http://%v ]", host)
	} else if strings.HasPrefix(host, "http://") {
		log.Printf("[Manager] start listen [ %v ]", host)
	} else if strings.HasPrefix(host, ":") {
		log.Printf("[Manager] start listen [ http://localhost%v ]", host)
	} else {
		log.Printf("[Manager] start listen [ %v ]", host)
	}
	http.ListenAndServe(host, nil)
}

type NodeNoResponse struct {
}

func (e *NodeNoResponse) Error() string {
	return "Node has no response"
}
