package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,MyMux")
}

// func main() {
// 	mux := &MyMux{}
// 	http.ListenAndServe(":8899", mux)
// }
