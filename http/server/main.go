package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Values struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Values, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Values, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	log.Println("Starting server on port 1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
