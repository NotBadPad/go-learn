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

func main() {
	test2()
}
