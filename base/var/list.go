package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := list.New()
	for i := 0; i < 10; i++ {
		a.PushBack(i)
		fmt.Println(i)
	}
	fmt.Println(list)

	a.Remove(2)
	a.Remove(7)

	for e := a.Front(); e != nil; e.Next() {
		fmt.Println(e)
	}
}
