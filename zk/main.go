package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	zk "github.com/samuel/go-zookeeper/zk"
	"time"
)

func getConnect(zkList []string) (conn *zk.Conn) {
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
	conn.Create("/go_servers", nil, 0, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 测试临时节点
 * @return {[type]}
 */
func test2() {
	zkList := []string{"localhost:2183"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/testadaadsasdsaw", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 获取所有节点
 */
func test3() {
	zkList := []string{"localhost:2183"}
	conn := getConnect(zkList)

	defer conn.Close()

	children, _, err := conn.Children("/go_servers")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", children)

}

func main() {
	test3()
}
