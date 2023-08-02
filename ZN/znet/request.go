package znet

import "ZN/ZN/ziface"

type Request struct {
	//连接对象Connection
	conn ziface.IConnection
	//请求数据
	Msg ziface.IMessage
	////请求数据ID
	//ID uint
}

// GetConnection 获取连接对象Connection
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetMsgData 获取请求数据
func (r *Request) GetMsgData() []byte {
	return r.Msg.GetData()
}

// GetMsgID 获取请求数据ID
func (r *Request) GetMsgID() uint32 {
	return r.Msg.GetMsgID()
}

//// GetReqID 获取请求数据ID
//func (r *Request) GetReqID() uint {
//	return r.ID
//}
