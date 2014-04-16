package main

import (
	"fmt"
)

type BaseHandler interface {
	execute(msg string) (int, string)
}

type MyHandler struct {
}

func (b *MyHandler) execute(msg string) (int, string) {
	fmt.Println("My:" + msg)
	return 1, "aaa"
}

type OtherHandler struct {
}

func (b *OtherHandler) execute(msg string) (int, string) {
	fmt.Println("Other:" + msg)
	return 2, "bbb"
}

func main() {
	handlers := make(map[string]BaseHandler)
	handlers["MyHandler"] = &MyHandler{}
	handlers["OtherHandler"] = &OtherHandler{}

	handlerName := "OtherHandler"

	var base BaseHandler
	base = handlers[handlerName]
	a, _ := base.execute("test")
	fmt.Println(a)

}
