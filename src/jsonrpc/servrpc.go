/**
socket则是对TCP/IP协议的封装和应用(程序员层面上)。
也可以说，TPC/IP协议是传输层协议，主要解决数据如何在网络中传输，
而HTTP是应用层协议，主要解决如何包装数据。

WEB使用HTTP协议作应用层协议，以封装HTTP文本信息，然后使用TCP/IP做传输层协议将它发到网络上。”
而我们平时说的最多的socket是什么呢，实际上socket是对TCP/IP协议的封装，Socket本身并不是协议，而是一个调用接口(API)。

网络协议：

应用层(会话层，表示层)　传输层　网络层　数据传输层　物理层

 */

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
