package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func test1() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "000101"
	cell = row.AddCell()
	cell.Value = "中文"
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func test2() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file, _ = xlsx.OpenFile("MyXLSXFile.xlsx")
	sheet = file.Sheet["Sheet1"]
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "000101"
	cell = row.AddCell()
	cell.Value = "中文1"
	err = file.Save("MyXLSXFile1.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func main() {
	test1()
	test2()
}
