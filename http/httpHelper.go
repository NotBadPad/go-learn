package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
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
	Msg  string `json:"msg"`
}

/**
 * 获取http请求结果
 * 	url:请求url
 * 	data：body数据对象接口
 */
func GetHttpRequest(url string, data interface{}) (code int, err error) {
	begin := time.Now().UnixNano()
	beego.Info("Request url:", url)
	req := httplib.Get(url)
	resp := &ApiResp{}
	err = req.ToJson(resp)
	if err != nil {
		return
	}
	end := time.Now().UnixNano()
	beego.Info("Http response:", resp, ", cost:", end-begin)

	code = resp.Head.Code
	if resp.Head.Code != 200 {
		err = errors.New(resp.Head.Msg)
		return
	}

	//获取body数据
	err = GetRespBodyData(resp, data)
	resp.Body = &data
	return
}

/**
 * 获取http请求结果，只支持post
 * url:请求地址
 * params：post参数
 * data：返回数据
 */
func GetPostHttpRequest(url string, params map[string]string, data interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			errStr := fmt.Sprintf("get post fail: params-%v , err-%v", params, x)
			err = errors.New(errStr)
		}
	}()

	begin := time.Now().UnixNano()
	req := httplib.Post(url).SetTimeout(CONNECT_TIME_OUT, READ_WRITE_TIME_OUT).Header("Content-Type", "application/x-www-form-urlencoded")
	beego.Info("Request url:", url, params)
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
	end := time.Now().UnixNano()
	beego.Info("Http response:", resp, ", cost:", end-begin)

	if resp.Head.Code != 200 {
		err = errors.New(resp.Head.Msg)
		return
	}

	if data != nil {
		beego.Debug(resp.Body)
		//获取body数据
		err = GetRespBodyData(resp, data)
		resp.Body = &data
	}

	return
}

/**
 * 获取http请求结果，只支持post,返回结构自定义
 * url:请求地址
 * params：post参数
 * data：返回数据
 */
func GetPostHttpRequestData(url string, params map[string]string, respData interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			errStr := fmt.Sprintf("get post fail: params-%v , err-%v", params, x)
			err = errors.New(errStr)
		}
	}()
	begin := time.Now().UnixNano()
	req := httplib.Post(url).SetTimeout(CONNECT_TIME_OUT, READ_WRITE_TIME_OUT).Header("Content-Type", "application/x-www-form-urlencoded")
	beego.Info("Request url:", url, params)
	if params != nil && len(params) > 0 {
		for key, value := range params {
			req.Param(key, value)
		}
	}

	err = req.ToJson(respData)
	if err != nil {
		return
	}
	end := time.Now().UnixNano()
	beego.Info("Http response:", respData, ", cost:", end-begin)

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
