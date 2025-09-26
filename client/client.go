package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Printf("client start err: %v\n", err)
		return
	}

	if _, err := conn.Write([]byte("你好，Zinx服务器")); err != nil {
		fmt.Printf("conn write err: %v\n", err)
		return
	}

	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn read err: %v\n", err)
		return
	}
	fmt.Println(string(buf[:n]))

}
