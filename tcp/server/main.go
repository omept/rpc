package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"

	"github.com/omept/rpc/utils"
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

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	utils.CheckFatalError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckFatalError(err)
	log.Println("Listening on ", tcpAddr.String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error: ", err.Error())
			continue
		}
		rpc.ServeConn(conn)
	}

}
