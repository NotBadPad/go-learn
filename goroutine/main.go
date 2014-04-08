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

func main() {
	go say("bbbb")
	say("tttt")
	time.Sleep(5 * 10000)
}
