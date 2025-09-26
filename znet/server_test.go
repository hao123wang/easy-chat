package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test start...")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Printf("client dial err: %v", err)
		return
	}

	_, err = conn.Write([]byte("heelo"))
	if err != nil {
		fmt.Printf("write conn err: %v", err)
		return
	}
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn read err: %v", err)
		return
	}
	fmt.Println("Server:", string(buf[:n]))
}

func TestServer(t *testing.T) {
	server := NewServer("[TEST] server")
	go ClientTest()
	server.Serve()
}
