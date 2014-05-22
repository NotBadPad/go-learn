package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func test2() {
	req := httplib.Get("http://www.baidu.com/")
	fmt.Println(req.Response().StatusCode)
}

func main() {
	test2()
}
