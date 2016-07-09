package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	zk "github.com/samuel/go-zookeeper/zk"
	"time"
)

func getConnect(string []zkList) (conn zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

/**
 * 测试连接
 * @return
 */
func test1() {
	zkList := []string{"localhost:2183"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/testadaw", nil, zk.FlagSequence, zk.WorldACL(zk.PermAll))
}

/**
 * 测试临时节点
 * @return {[type]}
 */
func test2() {

}

func main() {
	test1()
}
