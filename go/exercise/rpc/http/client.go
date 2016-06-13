package main

import (
	"fmt"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var reply int
	err = client.Call("Arith.Multiply", Args{3, 2}, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("res", reply)
}
