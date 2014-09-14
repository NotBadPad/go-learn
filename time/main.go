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

func test8() {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 24, 0, 0, 0, t.Location())
	fmt.Println(t1.Unix() - t.Unix())
}

func test9() {
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().UnixNano())
	}
}
func test10() {
	str := "2014-06-12 16:32:00"
	t1, _ := time.Parse("2006-01-02 15:04:05", str)
	fmt.Println(t1)
}

func test11() {
	t := time.Now()
	t = t.AddDate(0, -1, -1)
	fmt.Println(t.String())
}

func test12() {
	nowDate, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	begin := nowDate.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	end := nowDate.Format("2006-01-02 15:04:05")
	fmt.Println(nowDate.String(), begin, end)
}

func test13() {
	str := "2014-06-12 16:32:00"
	t1, _ := time.Parse("20060102150405", str)
	fmt.Println(t1)
}

func main() {
	test12()
}
