package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{6, 7, 3, 2, 4, 1, -2, 10}
	c := make(chan int)
	go sum(a, c)

	result := <-c

	fmt.Println(result)
}
