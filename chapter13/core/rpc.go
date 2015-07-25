package core

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (s *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func rpcServer() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("error in net.Listen()", err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func rpcClient() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("error in net.Dial()", err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println("Error in rpc.Client.Call()", err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func RpcDemo() {
	go rpcServer()
	go rpcClient()

	var input string
	fmt.Scanln(&input)
}
