package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func test2() {
	req := httplib.Get("http://www.baidu.com/")
	fmt.Println(req.ToJson(v))
}

func main() {
	test2()
}
