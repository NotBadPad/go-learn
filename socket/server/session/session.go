package session

import (
	"bufio"
	"fmt"
	"github.com/tyrchen/goutil/uniq"
	"net"
)

type Message chan []byte

type IoSession struct {
	Id       int64
	Conn     net.Conn
	In       Message
	Out      Message
	reader   *bufio.Reader
	writer   *bufio.Writer
	Quit     chan bool
	Ready    chan string
	closed   bool
	mutx     chan bool
	ClientIp string
	Attrs    map[string]interface{}
}

func NewSession(conn net.Conn) *IoSession {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	session := &IoSession{
		Id:       int64(uniq.GetUniq()),
		Conn:     conn,
		In:       make(Message),
		Out:      make(Message),
		Quit:     make(chan bool),
		Ready:    make(chan string),
		reader:   reader,
		writer:   writer,
		closed:   false,
		mutx:     make(chan bool, 1),
		ClientIp: conn.RemoteAddr().String(),
		Attrs:    make(map[string]interface{}, 0),
	}
	return session
}

func Dispatcher(session *IoSession) {
	fmt.Println("Session Dispatcher", session.Id)
}
