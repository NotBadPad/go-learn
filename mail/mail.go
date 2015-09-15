package components

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"net/smtp"
	"strings"
)

/**
 * 向指定用户发送邮件
 */
func SendMailToUsers(title, msg, users string) {
	passwd, _ := base64.StdEncoding.DecodeString(beego.AppConfig.String("mail.password"))

	auth := smtp.PlainAuth(
		"",
		beego.AppConfig.String("mail.username"),
		string(passwd),
		beego.AppConfig.String("mail.host"),
	)
	to := strings.Split(users, ";")

	content_type := "Content-Type: text/html" + "; charset=UTF-8"
	bmsg := []byte("To: " + users + "\r\nFrom: " + beego.AppConfig.String("mail.username") + "<" + beego.AppConfig.String("mail.username") + ">\r\nSubject: " + title + "\r\n" + content_type + "\r\n\r\n" + msg)

	err := SendMailTLS(beego.AppConfig.String("mail.host")+":"+beego.AppConfig.String("mail.port"), auth, beego.AppConfig.String("mail.username"), to, bmsg)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}

/**
 * 向指定用户发送邮件，用户邮箱在配置文件中配置（mail.receiver），多个使用;分隔
 */
func SendMail(title, msg string) {
	passwd, _ := base64.StdEncoding.DecodeString(beego.AppConfig.String("mail.password"))

	auth := smtp.PlainAuth(
		"",
		beego.AppConfig.String("mail.username"),
		string(passwd),
		beego.AppConfig.String("mail.host"),
	)
	to := strings.Split(beego.AppConfig.String("mail.receiver"), ";")

	content_type := "Content-Type: text/html" + "; charset=UTF-8"
	bmsg := []byte("To: " + beego.AppConfig.String("mail.receiver") + "\r\nFrom: " + beego.AppConfig.String("mail.username") + "<" + beego.AppConfig.String("mail.username") + ">\r\nSubject: " + title + "\r\n" + content_type + "\r\n\r\n" + msg)

	err := SendMailTLS(beego.AppConfig.String("mail.host")+":"+beego.AppConfig.String("mail.port"), auth, beego.AppConfig.String("mail.username"), to, bmsg)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil) //smtp包中net.Dial()当使用ssl连接时会卡住
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailTLS(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := Dial(addr)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer c.Quit()

	if ok, _ := c.Extension("Auth"); ok {
		if err = c.Auth(a); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	if err = c.Mail(from); err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}
