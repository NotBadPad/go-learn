package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("aaaa ", s)
	}
}

func test1() {
	go say("bbbb")
	say("tttt")
	time.Sleep(5 * 10000)
}

func test2() {
	a := []string{"a", "b", "c"}
	fmt.Println(a)
	for _, b := range a {
		go func() {
			fmt.Println(b)
		}()
	}
}

func main() {
	test2()
	var c chan bool
	<-c
}
