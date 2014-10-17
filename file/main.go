package main

import (
	"fmt"
	"path"
	"runtime"
)

func test1() {
	fmt.Println(path.Split(`test/aa/file.css`))
}

func test2() {
	_, file, line, ok := runtime.Caller(3)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
}

func main() {
	test2()
}
