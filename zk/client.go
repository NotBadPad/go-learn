package main

import (
	"errors"
	"fmt"
	"go-learn/zk/example"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		startClient()

		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	// service := "127.0.0.1:8899"
	//获取地址
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}

	fmt.Println("connect host: " + serverHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))

	return
}

func getServerHost() (host string, err error) {
	conn, err := example.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err := example.GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
