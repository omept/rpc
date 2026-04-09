package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero error")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		jsonrpc.ServeConn(struct {
			io.ReadCloser
			io.Writer
		}{r.Body, w})
	})

	log.Println("HTTP JSON-RPC server running on :1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
