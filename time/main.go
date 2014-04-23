package main

import (
	"fmt"
	"time"
)

func test1() {
	// for {
	// 	select {
	// 	case <-time.After(time.Second * 3):
	// 		fmt.Println(time.Now().String())
	// 	}
	// }
}

func test2() {
	st := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(st)
}

func main() {
	test2()
}
