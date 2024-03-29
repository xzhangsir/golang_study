package ziface

// 定义服务器接口
type IServer interface {
	// 启动服务器方法
	Start()
	// 停止服务器方法
	Stop()
	// 启动业务服务方法
	Serve()
	// 路由功能
	AddRouter(router IRouter)
}
