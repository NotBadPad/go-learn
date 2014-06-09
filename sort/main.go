package main

import (
	"fmt"
)

func quicksort(a []int) {
	i := 0
	j := len(a) - 1
	if i >= j {
		return
	}

	temp := a[0]
	for {
		if temp <= a[j] {
			j--
		}else{
			a[i] = a[j]
			a[j] = temp
		}
	}
}


4 7 10 1 23 6 
1 7 10 4 23 6
1 4 10 7 23 6


1 4 6 7 23 10
1 4 6 7 10 23