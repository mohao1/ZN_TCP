package ziface

import "net"

type IConnection interface {
	// Start 启动连接
	Start()
	// Stop 关闭连接
	Stop()
	// GetTCPConnection 获取连接的conn对象(套接字-服务连接对象)
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取连接ID
	GetConnID() uint32
	// RemoteAddr 获得客户端的连接地址和端口
	RemoteAddr() net.Addr
	// SendMsg 发送数据操作方法
	SendMsg(uint32, []byte) error

	// SetProperty 添加连接属性
	SetProperty(string, interface{})
	// GetProperty 获取连接属性
	GetProperty(string) (interface{}, error)
	// RemoveProperty 删除连接属性
	RemoveProperty(string)
}

// HandleFunc 处理连接的回调函数的类型
type HandleFunc func(*net.TCPConn, []byte, int) error
