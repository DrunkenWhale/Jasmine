package main

import (
	node2 "Jasmine/node"
	"fmt"
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
	fmt.Println(node.Get("114"))
}