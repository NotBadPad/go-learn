package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func ReadFile(pathStr string) {
	file, err := os.Open(pathStr)
	if err != nil {
		fmt.Errorf("Error Read file %s: %s\n", pathStr, err.Error())
	}
	defer file.Close()
	fb := bufio.NewReader(file)
	buff := bytes.NewBuffer(make([]byte, 2048))
	count := 0
	for {
		line, isPrefix, err := fb.ReadLine()
		if err != nil { //结束
			fmt.Println(count)
			break
		}
		buff.Write(line)
		if !isPrefix { //如果前缀标记是false，则说明读取的是整行
			buff.String()
			buff.Reset() //读取后重置缓冲区

			//处理数据
			count++
		}
	}
}

func main() {
	begin := time.Now().UnixNano()
	strStr := "E:\\workspace\\go\\src\\logstatistics\\temp\\v.ubox.cn\\v.ubox.cn.2014-05-11.log"
	ReadFile(strStr)
	end := time.Now().UnixNano()
	fmt.Println(end - begin)
}

/**
 * 16.9M
 * 62035行
 * 48ms
 *
 * 794M
 * 2257557行
 * 9269ms
 *
 * 1842M
 * 6337175行
 * 14181ms
 */
