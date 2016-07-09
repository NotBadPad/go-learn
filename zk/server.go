package main

import (
	"fmt"
	"go-learn/zk/example"
	"net"
	"os"
	"time"
)

func main() {
	go starServer("127.0.0.1:8897")
	go starServer("127.0.0.1:8898")
	go starServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<-a
}

func starServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println(tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//注册zk节点q
	conn, err := example.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	err = example.RegistServer(conn, port)
	if err != nil {
		fmt.Printf(" regist node error: %s ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleCient(conn, port)
	}

	fmt.Println("aaaaaa")
}

func handleCient(conn net.Conn, port string) {
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(port + ": " + daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
		os.Exit(1)
	}
}
