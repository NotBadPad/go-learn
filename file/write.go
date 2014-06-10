package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	TEST_STR   = `Stat returns the FileInfo structure describing file. If there is an error, it will be of type *PathError.\n\r`
	LINE_COUNT = 10000000
)

func WriteFile(pathStr string) {
	file, err := os.Create(pathStr)
	if err != nil {
		fmt.Errorf("Error create file:%s", err.Error())
	}
	defer file.Close()
	for i := 0; i < LINE_COUNT; i++ {
		file.WriteString(TEST_STR)
	}
}

func WriteFile1(pathStr string) {
	file, err := os.Create(pathStr)
	if err != nil {
		fmt.Errorf("Error Read file %s: %s\n", pathStr, err.Error())
	}
	defer file.Close()
	fb := bufio.NewWriter(file)
	for i := 0; i < LINE_COUNT; i++ {
		fb.WriteString(TEST_STR)
	}
}

func main() {
	begin := time.Now().UnixNano()
	pathStr := "E:/workspace/go/src/go-learn/temp/test2.txt"
	WriteFile(pathStr)
	end := time.Now().UnixNano()
	fmt.Println(end - begin)

	begin = time.Now().UnixNano()
	pathStr = "E:/workspace/go/src/go-learn/temp/test3.txt"
	WriteFile1(pathStr)
	end = time.Now().UnixNano()
	fmt.Println(end - begin)
}
