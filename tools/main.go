package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	params := os.Args
	if len(params) < 2 {
		fmt.Printf("Error:no enough arguments.\n")
	}

	switch params[1] {
	case "encrypt":
		if len(params) == 3 {
			encrypt(params[2])
		} else {
			fmt.Printf("Error arguments: Need 3 args.\n")
		}
	case "sign":
	case "conv-time":
		if len(params) == 3 {
			unixToStanded(params[2])
		} else {
			fmt.Printf("Error arguments: Need 3 args.\n")
		}
	case "conv-unix-time":
	default:
	}
}

func encrypt(str string) {
	pbytes := []byte(str)
	bytes := make([]byte, len(pbytes))
	for i, b := range pbytes {
		bytes[i] = b ^ 3
	}
	pwd := string(bytes)

	fmt.Println(pwd)
}

func unixToStanded(str string) {
	t, _ := time.Parse("2006-01-02 15:04:05", str)
	fmt.Println(t)
}
