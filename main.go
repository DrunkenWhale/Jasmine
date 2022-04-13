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
	manager := manage.NewManger()

	node1 := node2.NewNode("thyme", 7777, func(key string) ([]byte, error) {
		r, b := db[key]
		if b {
			return r, nil
		} else {
			return nil, nil
		}
	})
	manager.AddNode("pigeon", "http://localhost:9999")
	manager.AddNode("thyme", "http://localhost:8888")
	manager.Register()
	go node1.StartNodeServer(":8888")
	go node.StartNodeServer(":9999")
	manager.StartManageServer(":7777")
}
