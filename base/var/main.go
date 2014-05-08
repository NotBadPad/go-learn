package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
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

func test6() {
	a := "abcdefg"
	fmt.Println(a[0:3])
}

func test7() {
	a := "## \r"
	// b := []byte{13}
	fmt.Println([]byte(a))
}

func test8() {
	a := map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println(len(a))
	fmt.Println(a)
}
func test9() {
	a := []int32{1, 2, 3, 4, 5, 6}
	for i, b := range a {
		fmt.Println(i, b)
	}
}

func test10() {
	a := map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4}
	for key, value := range a {
		fmt.Println(key, value)
		if key == 3 {
			delete(a, key)
		}
	}

	fmt.Println(a)
}

func test11() {
	a := 10
	for a > 0 {
		a--
		fmt.Println(a)
	}
}

func test12() {
	a := "abcdef"
	fmt.Println(strings.Index(a, "c"))

	fmt.Println(strings.Repeat(a, 10))
}

func main() {
	test12()
}
