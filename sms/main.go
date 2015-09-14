package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strings"
	"time"
)

//curl -d 'phone=15501000792&msg=预定测试&caller=predetermine&night=1&quota=1' http://sms.uboxol.com/send_sms

const (
	CONNECT_TIME_OUT    = 10 * time.Second //连接超时时间
	READ_WRITE_TIME_OUT = 10 * time.Second //读写超时时间
)

var (
	SmsUrl  string
	PushUrl string
	WxUrl   string
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
	Desc string `json:"desc"`
}

// 短信响应
type SmsResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

/**
 * 短信接口-新
 */
func SendSmsNew(phone, msg string) (err error) {
	defer func() {
		if x := recover(); x != nil {
			errStr := fmt.Sprintf("Send sms fail: phone-%s, msg-%s, err-%v", phone, msg, x)
			err = errors.New(errStr)
		}
	}()

	phone = strings.TrimSpace(phone)
	msg = strings.TrimSpace(msg)
	if phone == "" || msg == "" {
		return errors.New("无效参数")
	}

	req := httplib.Post("http://sms.uboxol.com/send_sms").SetTimeout(CONNECT_TIME_OUT, READ_WRITE_TIME_OUT).Header("Content-Type", "application/x-www-form-urlencoded")
	req.Param("phone", phone)
	req.Param("msg", msg)
	req.Param("caller", "predetermine")

	resp := &SmsResp{}
	err = req.ToJson(resp)

	fmt.Println(resp, err.Error())
	return nil
}

func main() {
	SendSmsNew("15501000792", "预定测试")
	return
}
