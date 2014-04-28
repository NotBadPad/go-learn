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

func test3() {
	var array [5]rune
	array[0] = 'a'
	array[1] = 'b'
	array[2] = 'c'
	array[3] = 'd'
	array[4] = 'e'
	fmt.Println(string(array[0]))
}

func test4() {
	var arr2 *[]int = new([]int)
	fmt.Println((*arr2))
}

func test5() {
	var p []int
	fmt.Println(p)
}

func main() {
	test5()
}
