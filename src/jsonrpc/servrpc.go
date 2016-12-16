package main

import (
	"net"
	"net/rpc"
	"log"
	"net/rpc/jsonrpc"
	"os"
)
type Args struct {
	A, B int
}
type Reply struct {
	C int
}
type Arith int

func (t *Arith) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		os.Exit(0)
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("Arith", new(Arith)); err != nil {
		//return err
		os.Exit(0)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
