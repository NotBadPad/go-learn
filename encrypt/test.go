package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
)

/**
 * 获取微信验签
 * 算法：参数值按ascii顺序排序拼接，做sha1
 */
func GenerateWeiXinSign(kwargs map[string]string) string {
	if kwargs == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(kwargs))
	for _, v := range kwargs {
		keys = append(keys, v)
	}
	sort.Strings(keys) //按key ascii 顺序排序
	for _, v := range keys {
		buf.WriteString(url.QueryEscape(v))
	}
	// beego.Debug(buf.String())
	//做sha1处理，取出16进制全小写字符串
	msg := buf.String()
	h := sha1.New()
	h.Write([]byte(msg))
	bs := h.Sum(nil)
	mySign := fmt.Sprintf("%x", bs)
	return mySign
}

func main() {
	params := make(map[string]string, 0)
	params["timestamp"] = "1448439054"
	params["token"] = "CQN029CBHHhJOQc1OX8VOyaEqFSk9RinGQ"
	params["app"] = "5"

	sign := GenerateWeiXinSign(params)
	fmt.Println(sign)
}
