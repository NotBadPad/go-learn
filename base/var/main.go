package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "12"
	strconv.ParseInt(s1, 10, 32)
	fmt.Println(s1)
}
