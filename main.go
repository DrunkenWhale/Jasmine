package main

import (
	"fmt"
	"time"
)

func test() {
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C{
			fmt.Println("114514")
		}
	}()
}

func main() {
	test()
	time.Sleep(time.Second*7)
	fmt.Println(114)
	time.Sleep(time.Second*7)
	//node := node2.NewNode("114514", 1, nil)

	// seconds , not milliseconds
	fmt.Println(time.Now().Unix())

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C {
			fmt.Println("114514")
		}
	}()

}
