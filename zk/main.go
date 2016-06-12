package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	zk "github.com/samuel/go-zookeeper/zk"
)

func test1() {
	conn, session, err := zk.Connect([]string{"localhost:2181"}, 50000)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	event := <-session
	if event.State == zk.StateConnected {
		conn.Create("/testadaw", []byte{1, 2, 3, 4}, 0, zk.WorldACL(zk.PermAll))
	}
}

func main() {
	test1()
}
