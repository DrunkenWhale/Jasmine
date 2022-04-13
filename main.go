package main

import (
	"Jasmine/manage"
	node2 "Jasmine/node"
)

var db = map[string]interface{}{
	"114":    "114",
	"514":    1919810,
	"114514": "==",
}

func main() {
	node := node2.NewNode("pigeon", 114, func(key string) ([]byte, error) {
		return []byte("ssss"), nil
	})
	manager := manage.NewManger()
	manager.AddNode("pigeon", "http://localhost:9999")
	manager.Register()
	go manager.StartManageServer(":7777")

	node.StartNodeServer(":9999")
}
