package main

import (
	"bytes"
	"errors"
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"io/ioutil"
)

const (
	OUT_ENCODING = "gbk" //输出编码
)

func ExportCsv(head []string, data [][]string) (out []byte, err error) {
	fmt.Println(len(head), len(data))
	if len(head) == 0 {
		err = errors.New("ExportCsv Head is nil")
		return
	}

	columnCount := len(head)

	dataStr := bytes.NewBufferString("")
	//添加头
	for index, headElem := range head {
		separate := ","
		if index == columnCount-1 {
			separate = "\n"
		}
		dataStr.WriteString(headElem + separate)
	}

	//添加数据行
	for _, dataArray := range data {
		if len(dataArray) != columnCount { //数据项数小于列数
			err = errors.New("ExportCsv data format is error.")
		}
		for index, dataElem := range dataArray {
			separate := ","
			if index == columnCount-1 {
				separate = "\n"
			}
			dataStr.WriteString(dataElem + separate)
		}
	}

	//处理编码
	out = make([]byte, len(dataStr.Bytes()))
	iconv.Convert(dataStr.Bytes(), out, "utf-8", OUT_ENCODING)
	return
}

func main() {
	a := []string{"姓名", "年龄"}

	data := make([][]string, 0)
	b := []string{"郭靖", "21"}
	c := []string{"郭靖", "21"}
	d := []string{"郭靖", "21"}
	data = append(data, b)
	data = append(data, c)
	data = append(data, d)

	out, _ := ExportCsv(a, data)

	ioutil.WriteFile("test.csv", out, 0644)
}
