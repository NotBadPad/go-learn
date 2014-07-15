package main

import (
	"code.google.com/p/go.text/encoding"
	"code.google.com/p/go.text/encoding/charmap"
	"code.google.com/p/go.text/transform"
	"encoding/csv"
	"os"
)

func test1() {
	f, _ := os.Create("test1.xls")
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.Write([]string{"姓名", "年龄"})
	w.Write([]string{"aaa", "12"})
	w.Write([]string{"bbb", "13"})
	w.Write([]string{"ccc", "14"})
	w.Flush()
}

func test2() {
	f, _ := os.Create("test1.xls")
	defer f.Close()

	f.WriteString("姓名,年龄,植物\n\r")
	f.WriteString("a,b,c\n\r")
	f.WriteString("a,b,c\n\r")
}

func test3() {

}

func main() {
	test2()
}
