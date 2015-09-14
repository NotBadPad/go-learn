package main

import (
	"fmt"
)

/*
	128:   	  1000 0000
	1<<10:100 0000 0000
*/

type ApiHead struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Desc string `json:"desc,omitempty"`
}

func test1() {
	fmt.Println(128 & (1 << uint(10)))
}

func test2() {
	datas := make([]map[string]interface{}, 10)
	fmt.Println(datas)
}

func test4(msg interface{}) {

}
func test5() {
	var flt float64
	var i int64
	flt = 0.23544645 * 100
	i = int64(flt)
	fmt.Println(i)
}

func main() {
	test5()
}
