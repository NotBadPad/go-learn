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

func testHead() {
	a := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		a = append(a, rand.Intn(100000000))
	}
	b := a[:10]
	fmt.Println(a)
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

func test() {
	a := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		a = append(a, rand.Intn(100000000))
	}
	b := a[:10]

	for i := 10; i < 10000000; i++ {
		setRoot(a[i])
	}
}

func main() {
	testHead()
}
