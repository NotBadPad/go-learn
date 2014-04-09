package main

import (
	"go-learn/socket/server/session"
)

func main() {
	session.Listen("127.0.0.1:8778")
}
