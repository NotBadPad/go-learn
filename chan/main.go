package main

import (
	"fmt"
	"runtime"
	"time"
)

var ChanMap map[string]chan bool

func ReciveMsg() {
	time.Sleep(time.Second * 10)
	a := ChanMap["0001"]
	a <- true
}

func test1() {
	ChanMap = make(map[string]chan bool)
	a := make(chan bool, 1)
	ChanMap["0001"] = a
	go ReciveMsg()
	b := <-a
	fmt.Println("recive:", b)
}

func test2(b chan bool) {
	a := 0
	for {
		if a > 3 {
			break
		}
		select {
		case <-time.After(time.Second * 3):
			fmt.Println(a)
			a++
		}
	}
	fmt.Println("aaaa")
	b <- true
}

func test3() {
	b := make(chan int)
	// b<-"testing"
	close(b)
	fmt.Println(b)
	// b<-"testing2"
	// fmt.Println("aaaa")
}

func test4() {
	b := make(chan string, 1)
	b <- "testing1"
	close(b)
	fmt.Println(<-b)
	b = make(chan string, 1)
	b <- "testing2"
	fmt.Println(<-b)
}

// func test5() {
// 	b := make(chan int)
// 	b <- "testing"
// 	close(b)
// 	if c, ok := <-b; ok {
// 		fmt.Println(b)
// 	}
// }

func test6() {
	runtime.GOMAXPROCS(4)
	ch := make(chan int, 3)
	go func() {
		fmt.Print("3")
		time.Sleep(20 * time.Second)
	}()
	fmt.Print("1")
	<-ch
	fmt.Print("2")
}

func main() {
	test6()
}
