package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("test")
	t, _ = t.Parse("Hello,{{.UserName}}")
	user := Person{"guojing"}
	t.Execute(os.Stdout, user)
}
