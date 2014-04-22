package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

/**
 * 校验签名，对拼装字符串进行sha1计算
 */
func CheckSign(kwargs map[string]string, priviteKey string) (asign string, ok bool) {
	if sign, ok := kwargs["sign"]; ok {
		//拼装字符串
		msg := GenerateMsg(kwargs, priviteKey)
		fmt.Println(msg)
		//做sha1处理，取出16进制全小写字符串
		h := sha1.New()
		h.Write([]byte(msg))
		bs := h.Sum(nil)
		mySign := fmt.Sprintf("%x", bs)
		asign = mySign
		if sign == mySign {
			ok = true
		}
	}
	return
}

/**
 * 功能：将安全校验码和参数排序
 * 签名原始串按以下方式组装成字符串：
 *  1、除sign字段外，所有参数按照字段名的ascii码从小到大排序后使用格式（即key1=value1key2=value2…）
 *  	拼接而成，空值不传递，不参与签名组串。
 *  2、签名原始串中，字段名和字段值都进行URL Encode。
 *  3、sign = sign + '_' + priviteKey
 */
func GenerateMsg(kwargs map[string]string, priviteKey string) string {
	if kwargs == nil || priviteKey == "" {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(kwargs))
	for k := range kwargs {
		keys = append(keys, k)
	}
	sort.Strings(keys) //按key ascii 顺序排序
	for _, k := range keys {
		v := kwargs[k]
		if strings.ToLower(strings.TrimSpace(k)) == "sign" || strings.TrimSpace(v) == "" {
			continue
		}
		//if buf.Len() > 0 {
		//	buf.WriteByte('&')
		//}
		buf.WriteString(url.QueryEscape(k) + "=")
		buf.WriteString(url.QueryEscape(v))
	}
	return buf.String() + "_" + priviteKey
}

func main() {
	kwargs := make(map[string]string, 0)
	kwargs["uid"] = "1000"
	kwargs["cts"] = "1234323"
	kwargs["sign"] = "aaa"
	sign, _ := CheckSign(kwargs, "we7H_E8Xbf_ZH7kT6PQxtsFsIxl-wbVXk4hsw9kRrbo=")
	fmt.Println(sign)
}
