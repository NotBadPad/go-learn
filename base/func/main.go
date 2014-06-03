package main

import (
	"fmt"
	"os"
	"strings"
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

func test5() {
	a := []int{1, 2, 3}
	fmt.Println(a[1:2])
}

func test6() {
	str := "aaaa?stageId={stageId}&status={status}"
	str = strings.Replace(str, "{stageId}", "11", 1)
	str = strings.Replace(str, "{status}", "2", 1)
	fmt.Println(str)
}

func test7() {
	str := "asdqwezxctyughjbmnvbn"
	a := "c"
	index := strings.Index(str, a)
	fmt.Print(index)
}

func main() {
	test7()
}
