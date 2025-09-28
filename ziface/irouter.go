package ziface

type IRouter interface {
	// 在处理请求前的钩子方法
	PreHandle(request IRequest)
	// 处理请求的方法
	Handle(request IRequest)
	// 处理请求之后的钩子方法
	PostHandle(request IRequest)
}
