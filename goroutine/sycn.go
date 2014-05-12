package main

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

type Test struct {
	items map[int]string
	lock  sync.Locker
	mutex chan bool
}

func getRandValue(n int) string {
	const alphanum = "0123456789"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func put(t *Test, i int) {
	value := getRandValue(5)
	t.items[i] = value
}

func putBylock(t *Test, i int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	value := getRandValue(5)
	t.items[i] = value
}

func putByChan(t *Test, i int) {
	t.mutex <- true
	value := getRandValue(5)
	t.items[i] = value
	<-t.mutex
}

func test2(t *Test) {
	for i := 0; i < 10000; i++ {
		go putBylock(t, j*10+i)
	}
}

func test1() {
	t := &Test{items: make(map[int]string), mutex: make(chan bool, 1)}

	go test2(t)
	time.Sleep(5 * time.Second)
	fmt.Println(len(t.items))
	for key, value := range t.items {
		fmt.Println(key, "-", value)
	}

	// a := make(chan bool)
	// <-a
}

func main() {
	test1()
}
