package components

import (
	"bytes"
	"github.com/astaxie/beego"
	"html/template"
	"time"
)

/**
 * 获取模板内容，模板默认放在template下
 */
func GetTemplateContent(tempName string, data map[string]interface{}) (content string, err error) {
	t := template.New(tempName)
	t, err = t.ParseFiles("template/" + tempName)
	if err != nil {
		beego.Error("Error template :", err.Error())
	}

	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		beego.Error("Error template :", err.Error())
	}
	content = buff.String()
	return
}

/**
 * 转化时间
 */
func ConverUnixToTime(value int64) string {
	return time.Unix(value, 0).Format("2006-01-02 15:04:05")
}

func IsShow(value int64) string {
	if value == 0 {
		return ""
	}
	return "show"
}

func init() {
	beego.AddFuncMap("ConverUnixToTime", ConverUnixToTime)
	beego.AddFuncMap("IsShow", IsShow)
}
