package ziface

// 定义服务器接口
type IServer interface {
	Start()
	Stop() 
	Serve()
}

