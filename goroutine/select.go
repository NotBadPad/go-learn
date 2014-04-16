package main

import (
	"fmt"
	"time"
)

func dataIn(c chan int, d chan int) {
	timer := time.NewTicker(3 * time.Second)
	i := 0
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		// case <-time.After(time.Second * 5):
		// 	fmt.Println(time.Now().String())
		// 	d <- 10
		case <-timer.C:
			fmt.Println(time.Now().String())
			i++
			d <- i
		}
	}
}

func test1() {
	c := make(chan int)
	d := make(chan int)
	go dataIn(c, d)
	for {
		if j := <-d; j > 10 {
			break
		}
	}
}

func dataIn1(c, d chan int) {
	select {
	case <-c:
		fmt.Println("cccc")
		d <- 10
	}
}

//死锁
func test2() {
	c := make(chan int)
	d := make(chan int)
	go dataIn1(c, d)
	<-d
	close(c)

}

func main() {
	test1()
}
