package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func test1() {
	s1 := "12"
	strconv.ParseInt(s1, 10, 32)
	fmt.Println(s1)
}

func test2() {
	alphanum := "0123456789"
	var bytes = make([]byte, 3)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
		fmt.Println(b % byte(len(alphanum)))
	}

	fmt.Println(string(bytes))
}

func main() {
	test2()
}
