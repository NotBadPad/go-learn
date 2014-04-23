package main

import (
	"fmt"
)

func main() {
	pbytes := []byte("12345")
	bytes := make([]byte, len(pbytes))
	for i, b := range pbytes {
		bytes[i] = b ^ 3
	}
	pwd := string(bytes)

	fmt.Println(pwd)
}
