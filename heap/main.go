package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Heap []int

func parant(i int) int {
	return (i - 1) >> 1
}

func left(i int) int {
	return ((i + 1) << 1) - 1
}
func right(i int) int {
	return (i + 1) << 1
}

func swap(i, j int) {
	Heap[i], Heap[j] = Heap[j], Heap[i]
}

func headify(i, size int) {
	l := left(i)
	r := right(i)
	next := i
	if l < size && Heap[l]-Heap[i] > 0 {
		next = l
	}
	if r < size && Heap[r]-Heap[next] > 0 {
		next = r
	}
	if i == next {
		return
	}
	swap(i, next)
	headify(next, size)
}

func sort() {
	for i := len(Heap) - 1; i > 0; i-- {
		swap(0, i)
		headify(0, i)
	}
}

func setRoot(value int) {
	Heap[0] = value
	headify(0, len(Heap))
}

func build() {
	for i := len(Heap)/2 - 1; i >= 0; i-- {
		headify(i, len(Heap))
	}
}

/**
 * 测试堆，取最小的10个元素
 */
func testHead() {
	a := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		a = append(a, rand.Intn(100000000))
	}
	b := a[:10]
	begin := time.Now().UnixNano()
	Heap = a
	build()
	for i := 10; i < 10000000; i++ {
		setRoot(a[i])
	}
	sort()
	fmt.Println(b)
	end := time.Now().UnixNano()
	fmt.Println(end - begin)
}

/**
 * 数组排序，冒泡
 */
func sortArray(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

/**
 * 插入新元素
 */
func insertArray(a []int, v int) {
	if a[0] < v {
		return
	}
	i := len(a) - 1
	for {
		if a[i] < v {
			i--
		} else {
			break
		}
	}
	// 30 20 10 8 5
	//移动数组
	if i == 0 {
		a[0] = v
		return
	}

	temp := make([]int, len(a))
	copy(temp, a[1:i+1])
	copy(temp[i:], a[i+1:])
	a = temp
}

/**
 * 测试数组，取最小的10个元素
 */
func testArray() {
	a := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		a = append(a, rand.Intn(100000000))
	}
	b := a[:10]
	begin := time.Now().UnixNano()
	sortArray(a)
	for i := 10; i < 10000000; i++ {
		insertArray(b, a[i])
	}
	fmt.Println(b)
	end := time.Now().UnixNano()
	fmt.Println(end - begin)
}

func main() {
	// testHead()
	testArray()
}
