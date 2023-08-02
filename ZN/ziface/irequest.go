package ziface

type IRequest interface {
	// GetConnection 获取连接对象Connection
	GetConnection() IConnection
	// GetMsgData 获取请求数据
	GetMsgData() []byte
	// GetMsgID 获取请求数据ID
	GetMsgID() uint32
	//// GetReqID 获取请求数据ID
	//GetReqID() uint
}
