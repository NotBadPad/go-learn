package main

import (
	"bytes"
	"errors"
	iconv "github.com/djimenez/iconv-go"
)

/**
 * 导出处理
 */

const (
	OUT_ENCODING = "gbk" //输出编码
)

/**
 * 导出csv格式文件，输出byte数组
 * 输出编码通过OUT_ENCODING指定
 */
func ExportCsv(head []string, data [][]string) (out []byte, err error) {
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
