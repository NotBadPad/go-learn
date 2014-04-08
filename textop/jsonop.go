package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	/******* 写文件 ******/
	// var s Serverslice
	// s.Servers = append(s.Servers, Server{"ShangHai", "127.0.0.1"})
	// s.Servers = append(s.Servers, Server{"BeiJing", "127.0.0.2"})

	// b, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println("Server.json read fail:", err)
	// 	return
	// }

	// fout, err := os.Create("Server.json")
	// defer fout.Close()
	// if err != nil {
	// 	fmt.Println("Server.json create fail:", err)
	// 	return
	// }

	// fmt.Print(string(b))

	// fout.Write(b)

	/******* 读文件 ******/
	fin, err := os.Open("Server.json")

	if err != nil {
		fmt.Println("Server.json read fail:", err)
		return
	}
	defer fin.Close()

	data, err := ioutil.ReadAll(fin)
	if err != nil {
		fmt.Println("Server.json read fail:", err)
		return
	}

	var s Serverslice
	json.Unmarshal(data, &s)
	fmt.Println(s.Servers[0].ServerName)
}
