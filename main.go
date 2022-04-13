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
	go node.StartNodeServer(":9999")
	manager := manage.NewManger()
	manager.StartManageServer(":7777")
}
