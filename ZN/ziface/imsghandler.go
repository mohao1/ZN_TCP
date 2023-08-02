package ziface

// IMsgHandler 路由管理接口
type IMsgHandler interface {
	// DoMsgHandler 定义调度路由方法
	DoMsgHandler(IRequest)
	// AddRouter 定义添加路由方法
	AddRouter(uint32, IRouter)
	// StartWorkerPool 建立线程池的方法
	StartWorkerPool()
	// SendMsgToTakeQueue 将消息给到TakeWorker
	SendMsgToTakeQueue(req IRequest)
}
