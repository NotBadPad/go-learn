package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func task() {
	fmt.Println("task begin")
	time.Sleep(10 * time.Second)
	fmt.Println("task end")
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	go task()
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/test", sayhelloName)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
