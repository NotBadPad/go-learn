package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case <-time.After(time.Second * 3):
			fmt.Println(time.Now().String())
		}
	}
}
