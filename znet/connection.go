package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn

	// 当前连接的id，全局唯一
	ConnID uint32

	// 当前连接的关闭状态
	isClosed bool

	// 该连接的处理方法
	handleAPI ziface.HandlerFunc

	// 用于监听连接退出的通道
	ExitChan chan struct{}
}

func NewConnection(conn *net.TCPConn, connID uint32, handleApi ziface.HandlerFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: handleApi,
		ExitChan:  make(chan struct{}),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("recv buf err: %v", err)
			c.ExitChan <- struct{}{}
			continue
		}
		if err := c.handleAPI(c.Conn, buf, n); err != nil {
			fmt.Printf("connID %d handle err: %v", c.ConnID, err)
			c.ExitChan <- struct{}{}
			return
		}
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
