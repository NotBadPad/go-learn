package components

import (
	"encoding/json"
	"errors"
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
func GetHttpRequest(url string, data interface{}) (code int, err error) {
	req := httplib.Get(url)
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
