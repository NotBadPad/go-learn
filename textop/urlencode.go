package main

import (
	"fmt"
	"net/url"
	"strings"
)

func test1() {
	var str string = `http://workflow.uboxol.com/download/1441678136510089755_7%E5%8F%B7-%E4%B8%AD%E5%AD%A6%E7%89%88%E7%A7%9F%E8%B5%81%E5%8D%8F%E8%AE%AE%E9%95%BF%E9%83%A1%E5%8F%8C%E8%AF%AD.docx,http://workflow.uboxol.com/download/1441678136510089755_7%E5%8F%B7-%E4%B8%AD%E5%AD%A6%E7%89%88%E7%A7%9F%E8%B5%81%E5%8D%8F%E8%AE%AE%E9%95%BF%E9%83%A1%E5%8F%8C%E8%AF%AD.docx`
	strEn, _ := url.QueryUnescape(strings.Replace(str, `http://workflow.uboxol.com/download/`, "", -1))
	fmt.Println(strEn)

}

func ConvertAttachs(attachs string) string {
	if len(attachs) == 0 {
		return ""
	}

	if strings.Index(attachs, "/download/") >= 0 {
		return attachs
	}

	result := ""
	atts := strings.Split(attachs, ",")
	for _, att := range atts {
		result = result + "http://aaaa/download/" + url.QueryEscape(att) + ","
	}

	if len(result) > 0 {
		result = result[0 : len(result)-1]
	}
	return result
}

func main() {
	str := `1441678136510089755_7号-中学版租赁协议长郡双语.docx,1441678136510089755_7号-中学版租赁协议长郡双语.docx`
	fmt.Println(ConvertAttachs(str))
}
