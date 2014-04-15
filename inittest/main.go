package main

import (
	"fmt"
)

var Config map[string]string = make(map[string]string)

func Say() {
	Config["a"] = "b"
	Config["c"] = "d"
}

func init() {
	Say()
}

func main() {
	fmt.Println("aa")
}
