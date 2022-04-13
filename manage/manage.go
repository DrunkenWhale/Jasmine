package manage

type Manager struct {

	// key 	 : node name
	// value : node host
	nodes map[string]string
}

func NewManger() *Manager {
	return &Manager{
		nodes: make(map[string]string),
	}
}

func Add(nodeName string, nodeAddress string) {

}
