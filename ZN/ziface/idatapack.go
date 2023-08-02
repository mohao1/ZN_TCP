package ziface

type IDataPack interface {
	// GetHeadLen 获取包头长度
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(IMessage) ([]byte, error)
	// UnPack 拆包方法
	UnPack([]byte) (IMessage, error)
}
