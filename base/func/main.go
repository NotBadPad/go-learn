package main

import (
	"fmt"
	"os"
)

func main() {
	encryptStr := os.Args[1]
	if encryptStr == "" {
		fmt.Errorf("Error s%", "参数为空")
		return
	}
	fmt.Println(encryptStr)
}
