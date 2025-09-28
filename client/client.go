package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Printf("client start err: %v\n", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 512)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from terminal err: %v\n", err)
			return
		}
		conn.Write([]byte(str))

		go func() {
			for {
				n, err := conn.Read(buf)
				if err != nil && err != io.EOF {
					fmt.Printf("conn read err: %v\n", err)
					return
				}
				fmt.Println(string(buf[:n]))
			}
		}()
	}
}
