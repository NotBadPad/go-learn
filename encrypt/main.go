package main

import (
	"encoding/base64"
	"fmt"
	// "os"
)

func test1(str string) {
	pbytes := []byte(str)
	bytes := make([]byte, len(pbytes))
	for i, b := range pbytes {
		bytes[i] = b ^ 3
	}
	pwd := string(bytes)

	fmt.Println(pwd)
}

func test2() {
	var a [2]int
	a[0] = 2c
	fmt.Println(a[0] ^ a[1])
}

func test3() {
	str := "awayyqwe123"
	enStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(enStr)
	deStr, _ := base64.StdEncoding.DecodeString(enStr)
	fmt.Println(string(deStr))
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Error 参数为空")
	// 	return
	// }
	// encryptStr := os.Args[1]
	test3()
}
