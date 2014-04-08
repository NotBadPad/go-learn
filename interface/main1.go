package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) say() {
	fmt.Println("Human,", h.name)
}

func (s *Student) study() {
	fmt.Println("Student,", s.name, " is studying")
}

func (e *Employee) work() {
	fmt.Println("Employee,", e.name, " is working")
}

type Men interface {
	say()
}

func main() {
	gj := Student{Human{"gj", 25, "13423141241"}, "FAWFAW", 0.00}
	lm := Employee{Human{"lm", 25, "14134231231"}, "HERTGR", 5000}
	hm := Human{"gj1", 25, "12412312321"}
	var i Men

	i = &hm
	i.say()

	i = &gj
	i.say()

	i = &lm
	i.say()
}
