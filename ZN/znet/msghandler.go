package znet

import (
	"ZN/ZN/utils"
	"ZN/ZN/ziface"
	"fmt"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
	//定义一个消息队列
	TaskQueue []chan ziface.IRequest
	//定义一个Worker的工作池(线程)中的线程数量
	WorkerPoolSize uint32
}

// DoMsgHandler 定义调度路由方法
func (mh *MsgHandler) DoMsgHandler(req ziface.IRequest) {
	//判断id是否存在这个路由
	r, ok := mh.Apis[req.GetMsgID()]
	if !ok {
		fmt.Println("Router ID: ", req.GetMsgID(), " is Null")
	}
	//进行操作
	r.PreHandle(req)
	r.Handle(req)
	r.PostHandle(req)
}

// AddRouter 定义添加路由方法
func (mh *MsgHandler) AddRouter(id uint32, router ziface.IRouter) {
	//判断Apis中的路由是否存在
	if _, ok := mh.Apis[id]; ok {
		panic("router Add err ID:" + fmt.Sprint(id))
	}

	//进行绑定
	mh.Apis[id] = router
	fmt.Println("Router ID: ", id, " Add OK")

}

// StartWorkerPool 建立线程池的方法
func (mh *MsgHandler) StartWorkerPool() {
	//开启创建线程
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//开启线程通道
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalOBJ.MaxWorkerTaskSize)
		//开启线程
		go mh.startWorker(i, mh.TaskQueue[i])
	}
}

// StartWorker 开启一个工作线程方法
func (mh *MsgHandler) startWorker(workerId int, takeQueue chan ziface.IRequest) {
	//进行循环监听通道数据
	for {
		select {
		case req := <-takeQueue:
			{
				mh.DoMsgHandler(req)
			}
		}
	}

}

// SendMsgToTakeQueue 将消息给到TakeWorker
func (mh *MsgHandler) SendMsgToTakeQueue(req ziface.IRequest) {
	//编写平均分配Worker规则
	workerID := req.GetConnection().GetConnID() % mh.WorkerPoolSize
	//发送消息给到Take
	mh.TaskQueue[workerID] <- req
}

// NewMsgHandler 创建路由调度
func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalOBJ.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalOBJ.WorkerPoolSize,
	}
}
