package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	for k,v := range r.Form {
		fmt.Println("key",k)
		fmt.Println("val",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello")
}

func main() {
	http.HandleFunc("/",sayhelloName)
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
