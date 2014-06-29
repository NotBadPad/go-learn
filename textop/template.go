package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Emails   map[string]interface{}
	Names    []string
}

func test1() {
	// 	t := template.New("test")
	// 	t, _ = t.Parse(`
	// {{.UserName}}
	// {{.Emails.v1}}
	// 		`)
	// 	user := Person{"guojing", make(map[string]string, 0), make([]string, 3)}
	// 	user.Emails["v1"] = "test1"
	// 	user.Emails["v2"] = "test2"
	// 	user.Names[0] = "aaaaa"
	// 	t.Execute(os.Stdout, user)
}

func test2() {
	t := template.New("test")
	t, _ = t.Parse(`
{{.UserName}}
{{.Emails.v2.b}}
{{range .Emails.v3}}
	{{.}}
{{end}}
		`)
	user := Person{"guojing", make(map[string]interface{}, 0), make([]string, 3)}
	user.Emails["v1"] = "test1"
	user.Emails["v2"] = map[string]string{"a": "4321", "b": "123"}
	user.Emails["v3"] = []int{1, 3, 4, 2, 3}
	user.Names[0] = "aaaaa"
	t.Execute(os.Stdout, user)
}

func main() {
	test2()
}
