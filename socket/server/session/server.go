package session

import (
	"fmt"
	"net"
)

var connMap map[int]*IoSession

func Listen(connString string) (err error) {
	connMap = make(map[int]*IoSession)
	ln, err := net.Listen("tcp", connString)
	if err != nil {
		fmt.Errorf("Error listen to %s", connString)
		return err
	}
	fmt.Println("Server started, listen on %s \n", connString)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("Error accept conn: %s", err.Error())
			continue
		} else {
			fmt.Println("Accept: %s", conn.RemoteAddr().String())
		}
		session := NewSession(conn)
		go Dispatcher(session)
	}
}
