package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
)

const (
	CONNECT_TIME_OUT    = 100 * time.Second //连接超时时间
	READ_WRITE_TIME_OUT = 30 * time.Second  //读写超时时间
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
	Msg  string `json:"msg,omitempty"`
	Desc string `json:"desc,,omitempty"`
}

/**
 * 获取http请求结果
 * 	url:请求url
 * 	data：body数据对象接口
 */
func GetHttpRequest(url string, data interface{}) (code int, err error) {
	req := httplib.Get(url).SetTimeout(CONNECT_TIME_OUT, READ_WRITE_TIME_OUT)
	resp := &ApiResp{}
	err = req.ToJson(resp)
	if err != nil {
		return
	}
	//获取body数据
	err = GetRespBodyData(resp, data)
	resp.Body = &data
	code = resp.Head.Code
	return
}

/**
 * 获取http请求结果，只支持post
 * url:请求地址
 * params：post参数
 * data：返回数据
 */
func GetPostHttpRequest(url string, params map[string]string, data interface{}) (err error) {
	req := httplib.Post(url).SetTimeout(CONNECT_TIME_OUT, READ_WRITE_TIME_OUT)
	if params != nil && len(params) > 0 {
		for key, value := range params {
			req.Param(key, value)
		}
	}

	resp := &ApiResp{}
	err = req.ToJson(resp)
	if err != nil {
		return
	}

	//获取body数据
	err = GetRespBodyData(resp, data)
	resp.Body = &data
	return
}

/**
 * 在此对body进行判断，解决body为nil、对象、列表的情况
 */
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

/**
 * 出货量接口数据
 */
type ProSellNum struct {
	ProId int64 `json:"proId"`
	Num   int64 `json:"num"`
}

type SellResult struct {
	VmCode  string        `json:"vmCode"`
	Sn      string        `json:"sn"`
	SellNum []*ProSellNum `json:"sellNum"`
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	req := httplib.Post(`http://192.168.8.30:8084/rest/release/soldout/`)
	req.Header("Content-type", "text/plain;charset=UTF-8")
	data := `{"vmCodes":[{"vmCode":"0001407","sn":"223389047060430848","ctime":"2014-09-08 00:00:01"},{"vmCode":"0001407","sn":"256","ctime":"2014-09-08 00:00:01"}]}`
	fmt.Println(data)
	req.Param("data", data)
	// resp := &ApiResp{}
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)
	// 	// err := req.ToJson(resp)
	// if err != nil {
	// 	return
	// }

	// var sellResults []*SellResult
	// //获取body数据
	// err = GetRespBodyData(resp, &sellResults)
	// fmt.Println(resp.Head, sellResults)
	return
}
