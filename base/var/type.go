package main

import (
	"fmt"
)

/*
	128:   	  1000 0000
	1<<10:100 0000 0000
*/

func test1() {
	fmt.Println(128 & (1 << uint(10)))
}

func main() {
	test1()
}
