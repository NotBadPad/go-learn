package main

import (
	"container/list"
	"crypto/rand"
	"fmt"
	"sort"
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

func test16() {
	a := []string{"1", "2", "3"}
	b := []string{"3", "4", "5"}
	copy(a, b)
	fmt.Println(a)
}

func test17() {
	filePath := "/home/wwwlogs/productlogs/ubox002/service/exception.2014-05-11.log"
	index := strings.LastIndex(filePath, "/")
	if index > 1 {
		prePath := filePath[:index-1]
		index = strings.LastIndex(prePath, "/")
		prePath = prePath[index+1:]
		fmt.Println(prePath)
	}
}

type OrderList struct {
	Orders map[int64][]*Order
}

type Test1 struct {
	Name string
}

func test19() {
	a := []string{"aaa", "bbb"}
	test1 := &Test1{
		Name: a[1],
	}
	a = nil
	fmt.Println(test1.Name)
}

type OrderList1 struct {
	Orders map[string][3]float32
}

func test20() {
	a := make(map[string]*[3]float32, 0)
	v1 := [3]float32{0.05, 0.06, 13}
	v2 := [3]float32{0.05, 0.06, 13}
	v3 := [3]float32{0.05, 0.06, 13}
	a["v1"] = &v1
	a["v2"] = &v2
	a["v3"] = &v3

	p := a["v1"]
	p[1] = 0.09
	fmt.Println(a["v1"][1])
}

func test21(a map[string]string) {
	a["aaa"] = "bbbb"
}

func test22() {
	a := make(map[string]string, 0)
	a["aaa"] = "tttt"
	fmt.Println(a)
	test21(a)
	fmt.Println(a)
}

func test23() {
	var a float32
	a = 0.649
	a = a * 1000
	value := strconv.FormatFloat(float64(a), 'f', 0, 32)
	fmt.Println(strconv.Atoi(value))
}

func test24() {
	str := `'%%,%s,%%'`
	fmt.Println(fmt.Sprintf(str, "aaa"))
}

func test25() {
	str := `%X`
	fmt.Println(fmt.Sprintf(str, "aaa"))
}

func test26(format string, v ...interface{}) string {
	return fmt.Sprintf("[I] "+format, v...)

}

func test27() {
	fmt.Println(test26("logger initialized.%v", "aaaa"))
}

func test28(a []int) {
	a[2] = 10
}

func test29() {
	a := []int{1, 3, 4}
	test28(a)
	fmt.Println(a)
}

func test30() {
	url := "http://localhost:8081/test?token=gAGLi003OC3qKpRYCOq2clzYCuh5tt5hkXJppkwVZeE="
	index := strings.Index(url, "token=")
	index = index + 6
	fmt.Println(url[:index])
}

func test31() {
	a := make([]int, 3)
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(a), cap(a))
	fmt.Printf("%p\n", &a)
	a = append(a, b...)
	fmt.Println(len(a), cap(a))
	fmt.Printf("%p\n", &a)

}

func test32(a []int) {
	a[1] = 10
}
func test33() {
	a := []int{1, 2, 3, 4, 5}
	test32(a[:2])
	fmt.Println(a)
}

func test34() {
	a := "4.9E-324"
	b, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
}

func test35() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a[1:3]))
}

func test36() {
	var a float64
	a = 116.4807858
	b := strconv.FormatFloat(a, 'f', 6, 64)
	fmt.Println(b)
}

func test37() {
	a := []float64{116.480604, 116.481032, 116.480623, 116.480581, 116.481066, 116.481131, 116.480604}
	sort.Float64s(a)
	fmt.Println(a)
}

func test39() {
	a := 10
	for a > 0 {
		fmt.Println(a)
		a--
	}
}

func test40() {
	a := float64(5E-324)
	fmt.Println(a == 4.9E-324)
}
func test41() {
	a := []string{"a"}
	fmt.Println(a, strings.Join(a, ","))
}

func test42() {
	var a float64
	a = 0.0209615616
	val := strconv.FormatFloat(a, 'f', 4, 64)
	a, _ = strconv.ParseFloat(val, 64)
	fmt.Println(a)
}

func test43() {
	a := map[string]string{"aa": "ddd", "bb": "ccc"}
	b := a["rr"]
	fmt.Println(b == "")
}

func test44() {
	a := "1,2"
	b := strings.Split(a, ",")
	fmt.Println(b)
}

func test45() {
	a := 200
	b := float64(a) / 100
	c := strconv.FormatFloat(b, 'f', 2, 64)
	fmt.Println(c)
}

func test45() {
	a := make([][]string, 0)
	b := []string{"a", "b", "c", "d"}
	c := []string{"a", "b", "c", "d"}
	d := []string{"a", "b", "c", "d"}
	a = append(a, b)
	a = append(a, b)
	a = append(a, b)
}

func main() {
	test45()
}
