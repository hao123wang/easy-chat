package main

import "zinx/znet"

func main() {
	server := znet.NewServer("zinx-copy")
	server.Serve()
}