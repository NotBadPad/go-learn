package main

import (
	"crypto/rand"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math/big"
	"strconv"
	// "time"
)

func test1() {
	// ra := math.rand.New(math.rand.NewSource(time.Now().UnixNano()))
	// var a string
	// for i := 0; i < 5; i++ {
	// 	a = a + strconv.Itoa(ra.Intn(10))
	// }

	var a string
	for i := 0; i < 5; i++ {
		value, _ := rand.Int(rand.Reader, big.NewInt(10))
		a = a + strconv.Itoa(int(value.Int64()))
	}
	fmt.Println(a)
}

func test2() {
	fmt.Println(randomdata.Number(10000, 99999))
}

func test3(n int) string {
	// const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const alphanum = "0123456789"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func main() {
	fmt.Println(test3(5))
}
