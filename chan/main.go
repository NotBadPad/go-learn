package main

import (
	"fmt"
	"time"
)

var ChanMap map[string]chan bool

func ReciveMsg() {
	time.Sleep(time.Second * 10)
	a := ChanMap["0001"]
	a <- true
}

func main() {
	ChanMap = make(map[string]chan bool)
	a := make(chan bool, 1)
	ChanMap["0001"] = a
	go ReciveMsg()
	b := <-a
	fmt.Println("recive:", b)
}
