package main

import (
	"github.com/astaxie/beego/toolbox"
)

func main() {
	task := toolbox.NewTask("stock_stat", "0 30 7 * * *", StockTaskNew)
	toolbox.AddTask("stock_stat", task)
	toolbox.StartTask()
}
