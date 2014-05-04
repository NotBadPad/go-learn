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

func test3(){
	var a bool = false 
	timer := time.NewTicker(time.Second*3)
	for{
		if a {
			break
		}
		select{
		case <-timer.C:
			fmt.Println(time.Now())
		}
	}
}

func main() {
	test3()
	var a chan bool
	<-a
}