package ziface

type IRequest interface {
	// 获取当前请求已建立的连接的信息
	GetConnection() IConnection
	// 获取请求消息的数据
	GetData() []byte
}
