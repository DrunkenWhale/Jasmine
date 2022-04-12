package main

import (
	"fmt"
	"time"
)

func main() {
	// seconds , not milliseconds
	fmt.Println(time.Now().Unix())

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C {
			fmt.Println("114514")
		}
	}()

}
