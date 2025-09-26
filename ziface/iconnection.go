package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConnID() uint32
}

// 定义一个统一处理连接业务的接口
type HandlerFunc func(*net.TCPConn, []byte, int) error
