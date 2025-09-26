package znet

import (
	"errors"
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// 回显处理：将客户端发送的内容原样返回
func CallBackHandler(conn *net.TCPConn, data []byte, n int) error {
	fmt.Println("Callback to client")
	if _, err := conn.Write(data[:n]); err != nil {
		fmt.Printf("Write back to client err: %v", err)
		return errors.New("callback error")
	}
	return nil
}

// 开启服务器
func (s *Server) Start() {

	// 使用协程做服务端Listener业务
	go func() {
		// 获取一个tcp的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("resolve tcp addr err: %v\n", err)
			return
		}

		// 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listen %s err: %v\n", s.IPVersion, err)
			return
		}

		fmt.Printf("[START] Server %s listener at address: %s:%d\n", s.Name, s.IP, s.Port)

		var cid uint32

		// 接受客户端连接
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("accept err: %v\n", err)
				continue
			}

			// 每连接一个客户端，为其创建一个Connection
			connection := NewConnection(conn, cid, CallBackHandler)
			cid++
			
			go connection.Start()

		}

	}()
}

func (s *Server) Stop() {
	fmt.Printf("[STOP] Server %s\n", s.Name)
}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      3333,
	}
}
