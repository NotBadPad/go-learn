package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

// 响应数据结构
type ApiResp struct {
	Head ApiHead     `json:"head"`
	Body interface{} `json:"body,omitempty"`
}

// 响应头结构
type ApiHead struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func test2() {
	req := httplib.Get("http://localhost:8081/test")
	httpRes, _ := req.Response()
	fmt.Println(httpRes.StatusCode)

	result := &ApiResp{}
	err := req.ToJson(result)
	fmt.Println(err)
	fmt.Println(result)
}

func main() {
	test2()
}
