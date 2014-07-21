package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIp   string `xml:"serverIP"`
}

func test1() {
	/******* 写文件 ******/
	// v := &Servers{Version: "1"}
	// v.Svs = append(v.Svs, server{"ShangHai", "127.0.0.1"})
	// v.Svs = append(v.Svs, server{"BieJing", "127.0.0.2"})
	// output, err := xml.MarshalIndent(v, "\t", "\t")
	// if err != nil {
	// 	fmt.Printf("error:%v", err)
	// }

	// fout, err := os.Create("Server.xml")
	// defer fout.Close()
	// if err != nil {
	// 	fmt.Println("Server.xml create fail:", err)
	// 	return
	// }

	// fout.Write([]byte(xml.Header))
	// fout.Write(output)

	/******* 读文件 ******/

	fin, err := os.Open("Server.xml")

	if err != nil {
		fmt.Println("Server.xml read fail:", err)
		return
	}
	defer fin.Close()

	data, err := ioutil.ReadAll(fin)
	if err != nil {
		fmt.Println("Server.xml read fail:", err)
		return
	}
	v := Servers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("Server.xml read fail:", err)
		return
	}

	fmt.Println(v)
}

func test2() {
	data := `<serviceResponse>
			    <authenticationSuccess>
			    <user>username</user>
			     <proxyGrantingTicket>PGTIOU-84678-8a9d...</proxyGrantingTicket>
			    </authenticationSuccess>
			</serviceResponse>`
	type AuthenticationSuccess struct {
		User string `xml:"user"`
	}
	type Resp struct {
		XMLName xml.Name              `xml:"serviceResponse"`
		Auth    AuthenticationSuccess `xml:"authenticationSuccess"`
	}

	value := &Resp{}
	err := xml.Unmarshal([]byte(data), &value)
	if err != nil {
		fmt.Println("Server.xml read fail:", err)
		return
	}

	fmt.Println(value)
}

func main() {
	test2()
}
