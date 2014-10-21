package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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
	fmt.Println(a[0] ^ a[1])
}

func test3() {
	str := "awayyqwe123"
	enStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(enStr)
	deStr, _ := base64.StdEncoding.DecodeString(enStr)
	fmt.Println(string(deStr))
}

func test4() {
	str := "daaaf542add66503eb8ae271a4d4eb32"
	h := md5.New()
	h.Write([]byte(str))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Error 参数为空")
	// 	return
	// }
	// encryptStr := os.Args[1]
	test1("96000")
}
