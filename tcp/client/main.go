package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/omept/rpc/utils"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s server:port \n ", os.Args[0])
	}
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	utils.CheckFatalError(err)
	client, err := rpc.Dial("tcp", addr.String())
	utils.CheckFatalError(err)

	// synchronous call
	args := Args{A: 17, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	utils.CheckFatalError(err)
	log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	utils.CheckFatalError(err)
	log.Printf("Arith: %d/%d=%d remainder %d", args.A, args.B, quo.Quo, quo.Rem)

}
