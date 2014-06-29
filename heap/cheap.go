package main

import (
	"container/heap"
	"fmt"
)

type HeapItem struct {
	Key string
	Value interface{}
	SortValue float64 //排序值
}

type HeapQueue []*HeapItem

func (hq HeapQueue) Len() int { 
	return len(hq) 
}

func (hq HeapQueue) Less(i, j int) bool {
	return hq[i].SortValue > hq[j].SortValue
}

func (hq HeapQueue) Swap(i, j int) {
	hq[i], hq[j] = hq[j], hq[i]
}

func (hq *HeapQueue) Push(x interface{}) {
	*hq = append(*hq, x.(*HeapItem))
}

func (hq *HeapQueue) Pop() interface{} {
	old := *hq
	n := len(old)
	item := old[n-1]
	*hq = old[0 : n-1]
	return item
}

func main() {
	hq := &HeapQueue{}
	a1:=&HeapItem{Key:"a1",Value:4,SortValue:10}
	a2:=&HeapItem{Key:"a2",Value:4,SortValue:2}
	a3:=&HeapItem{Key:"a3",Value:4,SortValue:35}
	a4:=&HeapItem{Key:"a4",Value:4,SortValue:100}
	hq.Push(a1)
	hq.Push(a2)
	hq.Push(a3)
	hq.Push(a4)
	heap.Init(hq)
	for _,item := range *hq{
		fmt.Println(item.Key)
	}
}