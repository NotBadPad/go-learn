package main

import (
	"fmt"
	"time"
)

func dataIn(c chan int, d chan int) {
	timer := time.NewTicker(3 * time.Second)
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		// case <-time.After(time.Second * 5):
		// 	fmt.Println(time.Now().String())
		// 	d <- 10
		// }
		case <-timer.C:
			fmt.Println(time.Now().String())
			d <- 10
		}
	}
}

func main() {
	c := make(chan int)
	d := make(chan int)
	go dataIn(c, d)
	for {
		<-d
	}
}
