package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	fmt.Fprintf(w, "Hello")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Fatal("Start fail!")
	}
}
