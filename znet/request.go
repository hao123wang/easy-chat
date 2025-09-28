package znet

import "zinx/ziface"

type Request struct {
	// 已经和客户端成功建立的连接
	conn ziface.IConnection
	// 客户端请求数据
	data []byte
}

// 获取已建立的连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取请求的数据
func (r *Request) GetData() []byte {
	return r.data
}
