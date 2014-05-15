package main

import (
	"fmt"
	"strings"
	"time"
)

func test1() {
	// for {
	// 	select {
	// 	case <-time.After(time.Second * 3):
	// 		fmt.Println(time.Now().String())
	// 	}
	// }
}

func test2() {
	st := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(st)
}

func test3() {

}

func test4() {
	var a bool = false
	timer := time.NewTicker(time.Second * 3)
	for {
		if a {
			break
		}
		select {
		case <-timer.C:
			fmt.Println(time.Now())
		}
	}
}

func test5() {
	str := "11/May/2014:00:00:05 +0800"
	index := strings.Index(str, " ")
	str = str[:index]
	// strings.Split(str, sep)
	fmt.Println(str)

	// t1, _ := time.Parse("11/Jan/2006:15:04:05 -0700", str)
	// fmt.Println(t1)
}

func test6() {
	str := "12/May/2014:00:00:05 +0800"
	t1, _ := time.Parse("2/Jan/2006:15:04:05 -0700", str)
	fmt.Println(t1)
}

func test7() {
	ti := time.Now()
	fmt.Println(ti.Format("2006-01-02"))
}

func main() {
	test7()
}
