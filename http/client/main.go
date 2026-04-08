package main

import (
	"fmt"
	"log"
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
		log.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%s", serverAddress, "1234"))
	utils.CheckFatalError(err)

	// synchronous call
	args := Args{17, 8}
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	utils.CheckFatalError(err)

	log.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	utils.CheckFatalError(err)
	log.Printf("Arith: %d/%d=%d remainder %d \n", args.A, args.B, quot.Quo, quot.Rem)
}
