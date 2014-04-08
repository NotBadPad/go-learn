package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":8899"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		fmt.Printf("new connect:", conn.RemoteAddr().String())
		go handleCient(conn)
	}
}

func handleCient(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()

	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			break
		}

		fmt.Printf(string(read_len))

		// if read_len == 0 {
		// 	break
		// } else if string(request) == "timestamp" {
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		// }
	}

	request = make([]byte, 128)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
		os.Exit(1)
	}
}
