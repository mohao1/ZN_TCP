package ziface

type IServer interface {

	// Start 启动服务
	Start()
	// Server 运行服务
	Server()
	// Stop 停止服务
	Stop()
	// AddRouter 用户添加Router方法
	AddRouter(uint32, IRouter)
	// GetConnManger 获取连接模块
	GetConnManger() IConnManager

	// SetOnConnStartHook 注册OnConnStartHook方法
	SetOnConnStartHook(OnHookFunc)
	// SetOnConnStopHook 注册OnConnStopHook方法
	SetOnConnStopHook(OnHookFunc)
	// CallOnConnStartHook 调用OnConnStartHook方法
	CallOnConnStartHook(IConnection)
	// CallOnConnStopHook 调用OnConnStopHook方法
	CallOnConnStopHook(IConnection)
}

type OnHookFunc func(IConnection)
