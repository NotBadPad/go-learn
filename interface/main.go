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

func (s *Student) say() {
	fmt.Println("Student,", s.name)
	s.Human.say()
}

func main() {
	var s Student
	s.name = "guojing"
	s.say()
}
