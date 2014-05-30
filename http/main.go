package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

// 响应数据结构
type ApiResp struct {
	Head     ApiHead     `json:"head"`
	Body     interface{} `json:"body,omitempty"`
	byteBody []byte
}

// 响应头结构
type ApiHead struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

/**
 * 获取http请求结果
 * 	url:请求url
 * 	data：body数据对象接口
 */
func GetHttpRequest(url string, data interface{}) (resp *ApiResp, err error) {
	req := httplib.Get(url)
	resp = &ApiResp{}
	err = req.ToJson(resp)
	//获取body数据
	err = GetRespBodyData(resp, data)
	resp.Body = &data
	return
}

func GetRespBodyData(msg *ApiResp, data interface{}) (err error) {
	//将body转成bytes
	if msg.Body != nil {
		//判断body是否为空：body:[]或body:{}
		bodyLen := 0
		switch msg.Body.(type) {
		case map[string]interface{}:
			bodyLen = len(msg.Body.(map[string]interface{}))
		case []interface{}:
			bodyLen = len(msg.Body.([]interface{}))
		}
		if bodyLen != 0 {
			msg.byteBody, _ = json.Marshal(msg.Body)
			err = json.Unmarshal(msg.byteBody, data)
		} else {
			data = nil
			err = errors.New("Response body is nil")
		}
	}
	return
}

type RoleParm struct {
	Name string `json:"LDAP_NAME"`
}

func test1() {
	url := fmt.Sprintf("%s/interface/getEmpsByRoleId.ajax?roleId=%s", "http://ubox-acl.dev.uboxol.com", "11")
	roleParms := make([]*RoleParm, 0)
	resp, err := GetHttpRequest(url, roleParms)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(roleParms)
	result := *resp.Body.(*[]*RoleParm)
	fmt.Println(result)
}

func test2() {
	url := "http://m.5read.com/"
	req := httplib.Get(url)
	b, _ := req.Bytes()
	fmt.Println(string(b))
}

func main() {
	test2()
}
