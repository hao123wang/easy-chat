package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(req ziface.IRequest) {
	conn, ok := req.GetConnection().(*znet.Connection)
	if !ok {
		fmt.Println("assert req.GetConnection err")
		return
	}
	_, err := conn.GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (pr *PingRouter) Handle(req ziface.IRequest) {
	conn, _ := req.GetConnection().(*znet.Connection)
	_, err := conn.GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (pr *PingRouter) PostHandle(req ziface.IRequest) {
	conn, _ := req.GetConnection().(*znet.Connection)
	_, err := conn.GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {

	server := znet.NewServer("zinx-copy")
	server.AddRouter(&PingRouter{})
	server.Serve()
}
