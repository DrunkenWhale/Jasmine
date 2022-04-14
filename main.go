package main

import (
	node2 "Jasmine/node"
	"flag"
	"net/http"
	"strconv"
)

var db = map[string][]byte{
	"1":   []byte("14514"),
	"11":  []byte("4514"),
	"114": []byte("514"),
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8001, "Node's port")
	flag.Parse()
	node := node2.NewNode("pigeon", 377777, func(key string) ([]byte, error) {
		r, b := db[key]
		if b {
			return r, nil
		} else {
			return nil, http.ErrServerClosed
		}
	})
	node.StartNodeServer(":" + strconv.Itoa(port))
}

//
//func main() {
//	file, err := os.Open("nodes.json")
//	defer file.Close()
//	if err != nil {
//		log.Println("Config file : nodes.json not exists")
//		return
//	}
//	bytes, err := ioutil.ReadAll(file)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	var nodes []config.NodeDescribe
//	err = json.Unmarshal(bytes, &nodes)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	var port int
//	flag.IntVar(&port, "port", 7777, "Manager server port")
//	flag.Parse()
//	manager := manage.NewManger()
//	for _, nodes := range nodes {
//		manager.AddNode(nodes.Name, nodes.Address)
//	}
//	manager.Register()
//	manager.StartManageServer(":" + strconv.Itoa(port))
//
//}
