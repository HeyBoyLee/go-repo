package main

import (
	"net/rpc/jsonrpc"
	"log"
	"fmt"
)
type Args struct {
	A, B int
}
type Reply struct {
	C int
}
func main() {

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	reply := new(Reply)
	err1 := client.Call("Arith.Add", args, reply)
	if err1 == nil {
		fmt.Println(reply)
	}
}