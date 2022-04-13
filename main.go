package main

import (
	"Jasmine/manage"
	node2 "Jasmine/node"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var db = map[string][]byte{
	"114":    []byte("114"),
	"514":    []byte("1919810"),
	"114514": []byte("=="),
}

func main() {
	for i := 0; i < 1145414; i++ {
		k := strconv.Itoa(rand.Int())
		db[strconv.Itoa(i)] = []byte(k)
	}
	node := node2.NewNode("pigeon", 377777, func(key string) ([]byte, error) {
		r, b := db[key]
		if b {
			return r, nil
		} else {
			return nil, nil
		}
	})
	err := node.Put("114", []byte("514"), time.Second*7)
	err = node.Put("114", []byte("514"), time.Second*7)
	err = node.Put("114", []byte("514"), time.Second*7)
	err = node.Put("114", []byte("514"), time.Second*7)
	err = node.Put("114", []byte("514"), time.Second*7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(node.Memory())
	manager := manage.NewManger()
	manager.AddNode("pigeon", "http://localhost:9999")
	manager.Register()
	go manager.StartManageServer(":7777")

	node.StartNodeServer(":9999")
}
