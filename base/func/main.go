package main

import (
	"fmt"
	"os"
)

func test1() {
	encryptStr := os.Args[1]
	if encryptStr == "" {
		fmt.Errorf("Error s%", "参数为空")
		return
	}
	fmt.Println(encryptStr)
}

func test2() {
	type Param struct {
		Id   int64
		Name string
	}
	a := Param{Id: 12, Name: "guojing"}
	fmt.Printf("%s\n", a)
	fmt.Printf("%+s\n", a)
	fmt.Printf("%#s\n", a)
}

func test3() {
	a := []int{1, 2, 3}
	b := []int{7, 8, 9}
	c := make([]int, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	fmt.Println(c)
}

func test4() {
	a := []int{1, 2, 3}
	// b := []int{7, 8, 9}
	c := append(a, 3, 4, 5)
	fmt.Println(c)
}

func main() {
	test4()
}
