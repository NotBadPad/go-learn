package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("timer begin")
			go test2()
			fmt.Println("timer end")
		}
	}
}

func test2() {
	fmt.Println("test1 start")
	time.Sleep(5 * time.Second)
	fmt.Println("test1 end")
}
