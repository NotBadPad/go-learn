package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	// "strconv"
	// "strings"
	"time"
	// "os"
)

func test1(str string) {
	pbytes := []byte(str)
	bytes := make([]byte, len(pbytes))
	for i, b := range pbytes {
		bytes[i] = b ^ 3
	}
	pwd := string(bytes)

	fmt.Println(pwd)
}

func test2() {
	var a [2]int
	fmt.Println(a[0] ^ a[1])
}

func test3() {
	str := "awayyqwe123"
	enStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(enStr)
	deStr, _ := base64.StdEncoding.DecodeString(enStr)
	fmt.Println(string(deStr))
}

func test4() {
	str := "daaaf542add66503eb8ae271a4d4eb32"
	h := md5.New()
	h.Write([]byte(str))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

func GenerateMsg(kwargs map[string]string) string {
	if kwargs == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(kwargs))
	for _, v := range kwargs {
		fmt.Println(v)
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

func signCheck() {
	params := make(map[string]string, 0)
	nowTime := time.Now().Unix()
	fmt.Println(nowTime)
	params["timestamp"] = "1419739624"
	params["token"] = "CQN029CBHHhJOQc1OX8VOyaEqFSk9RinGQ"
	params["app"] = "5"
	val := GenerateMsg(params)

	fmt.Println(val)
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Error 参数为空")
	// 	return
	// }
	// encryptStr := os.Args[1]
	signCheck()
}
