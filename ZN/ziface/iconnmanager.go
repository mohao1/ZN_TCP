package ziface

type IConnManager interface {
	// Add 添加连接
	Add(IConnection)
	// Remove 删除连接
	Remove(IConnection)
	// Get 获取连接
	Get(uint32) (IConnection, error)
	// Len 得到连接总数
	Len() int
	// ClearConn 清除所有连接
	ClearConn()
}
