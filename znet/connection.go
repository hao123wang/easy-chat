package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn

	// 当前连接的id，全局唯一
	ConnID uint32

	// 当前连接的关闭状态
	isClosed bool

	// 处理业务的路由
	Router ziface.IRouter

	// 用于监听连接退出的通道
	ExitChan chan struct{}
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan struct{}),
	}
}

func (c *Connection) StartReader() {
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端%d断开连接\n", c.ConnID)
			} else {
				fmt.Printf("recv buf err: %v\n", err)
			}
			c.ExitChan <- struct{}{}
			return
		}

		// 得到当前客户端的请求
		req := Request{
			conn: c,
			data: buf[:n],
		}

		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)
	}
}

func (c *Connection) Start() {
	go c.StartReader()
	<-c.ExitChan
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()

	c.ExitChan <- struct{}{}
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
