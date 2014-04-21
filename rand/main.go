package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	// "time"
)

func main() {
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

	// a := [5]int{3, 1, 3, 6, 7}
	// fmt.Println(a)
}
