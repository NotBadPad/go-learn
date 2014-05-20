package main

import (
	"fmt"
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
	fmt.Println(Heap)
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

func main() {
	a := []int{14, 2, 55, 3, 63, 43, 5}
	Heap = a
	// fmt.Println(Heap)
	build()
	// fmt.Println(Heap)
	sort()
	// fmt.Println(Heap)
}
