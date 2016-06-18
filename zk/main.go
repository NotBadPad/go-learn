package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	zk "github.com/samuel/go-zookeeper/zk"
	"time"
)

func test1() {
	conn, _, err := zk.Connect([]string{"localhost:2183"}, 50*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	conn.Create("/testadaw", nil, zk.FlagSequence, zk.WorldACL(zk.PermAll))
}

func main() {
	test1()
}
