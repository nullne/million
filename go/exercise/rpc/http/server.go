package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Arith int
type Args struct {
	A, B int
}

func (t * Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
