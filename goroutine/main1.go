package main

import (
	"fmt"
	"time"
)

func main() {

	go test1()

	time.Sleep(10 * time.Second)
	fmt.Println("main end")
}

func test1() {
	fmt.Println("test1 start")
	go test2()
	fmt.Println("test1 end")
}

func test2() {
	fmt.Println("test2 start")
	time.Sleep(5 * time.Second)
	fmt.Println("test2 end")
}
