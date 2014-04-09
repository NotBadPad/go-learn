package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

type Quotient struct {
	Quo int
	Rem int
}

type Arith int

func (a *Arith) Mulriply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error {
	if a {

	}
}
