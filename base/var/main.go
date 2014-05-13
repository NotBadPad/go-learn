package main

import (
	"container/list"
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

func test1() {
	s1 := "12"
	strconv.ParseInt(s1, 10, 32)
	fmt.Println(s1)
}

func test2() {
	alphanum := "0123456789"
	var bytes = make([]byte, 3)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
		fmt.Println(b % byte(len(alphanum)))
	}

	fmt.Println(string(bytes))
}

func test3() {
	var array [5]rune
	array[0] = 'a'
	array[1] = 'b'
	array[2] = 'c'
	array[3] = 'd'
	array[4] = 'e'
	fmt.Println(string(array[0]))
}

func test4() {
	var arr2 *[]int = new([]int)
	fmt.Println((*arr2))
}

func test5() {
	var p []int
	fmt.Println(p)
}

func test6() {
	a := "abcdefg"
	fmt.Println(a[0:3])
}

func test7() {
	a := "## \r"
	// b := []byte{13}
	fmt.Println([]byte(a))
}

func test8() {
	a := map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println(len(a))
	fmt.Println(a)
}
func test9() {
	a := []int32{1, 2, 3, 4, 5, 6}
	for i, b := range a {
		fmt.Println(i, b)
	}
}

func test10() {
	a := map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4}
	for key, value := range a {
		fmt.Println(key, value)
		if key == 3 {
			delete(a, key)
		}
	}

	fmt.Println(a)
}

func test11() {
	a := 10
	for a > 0 {
		a--
		fmt.Println(a)
	}
}

func test12() {
	a := "abcdef"
	fmt.Println(strings.Index(a, "c"))

	fmt.Println(strings.Repeat(a, 10))
}

func test13() {
	a := make(map[int]int)
	for i := 1; i < 10; i++ {
		a[i] = i
	}

	for key, value := range a {
		fmt.Println(key, value)
	}
}

func test14() {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	a = append(a, 6)

	// j := len(a)
	// for i := 0; i < j; i++ {
	// 	a = a[1:]
	// 	fmt.Println(a)
	// }

	for _, value := range a {
		fmt.Println(value)
		a = a[1:]
	}

	fmt.Println(a)
}

type Order struct {
	OrderId int64 //订单号
}

func test15() {
	orderlist := list.New()
	order1 := &Order{OrderId: 1}
	order2 := &Order{OrderId: 2}
	order3 := &Order{OrderId: 3}
	order4 := &Order{OrderId: 4}
	order5 := &Order{OrderId: 5}
	orderlist.PushBack(*order1)
	orderlist.PushBack(*order2)
	orderlist.PushBack(*order3)
	orderlist.PushBack(*order4)
	orderlist.PushBack(*order5)
	eachList(orderlist)

	for elem := orderlist.Front(); elem != nil; {
		fmt.Println("-------------------")
		order := elem.Value.(Order)
		if order.OrderId == 3 {
			elem = deleteOrder(orderlist, elem)
		} else {
			elem = elem.Next()
		}
		eachList(orderlist)
	}
}

func deleteOrder(orderList *list.List, elem *list.Element) (nextElem *list.Element) {
	nextElem = elem.Next()
	orderList.Remove(elem)
	return
}

func eachList(orderList *list.List) {
	for elem := orderList.Front(); elem != nil; elem = elem.Next() {
		order := elem.Value.(Order)
		fmt.Println(order.OrderId)
	}
}

func main() {
	test15()
}
