package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Emails   map[string]string
	Names    []string
}

func main() {
	t := template.New("test")
	t, _ = t.Parse(`

	 {{range $k,$v := .Emails}}

	 {{$k}}:{{$v}},

	 {{end}}
		{{with .Names}} {{}} {{end}}


		`)
	user := Person{"guojing", make(map[string]string, 0), make([]string, 3)}
	user.Emails["v1"] = "test1"
	user.Emails["v2"] = "test2"
	user.Names[0] = "aaaaa"
	t.Execute(os.Stdout, user)
}
