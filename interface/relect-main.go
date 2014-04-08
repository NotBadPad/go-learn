package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) say() {
	fmt.Println("Human,", h.name)
}

func main() {
	hm := Human{"gj", 25, "12412312312"}
	t := reflect.TypeOf(hm)

	fmt.Println(t.NumField())

	v := reflect.ValueOf(hm)

	fmt.Println(v.FieldByName("name"))
}
