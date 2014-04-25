package main

import (
	"fmt"
)

func test1() {
	pbytes := []byte("guojing")
	bytes := make([]byte, len(pbytes))
	for i, b := range pbytes {
		bytes[i] = b ^ 3
	}
	pwd := string(bytes)

	fmt.Println(pwd)
}

func test2() {
	var a [2]int
	a[0] = 2
	fmt.Println(a[0] ^ a[1])
}
func main() {
	test1()
}
