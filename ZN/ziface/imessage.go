package ziface

type IMessage interface {

	// GetMsgID 获取消息ID
	GetMsgID() uint32
	// SetMsgID 修改消息ID
	SetMsgID(uint32)
	// GetDataLen 获取消息长度
	GetDataLen() uint32
	// SetDataLen 修改消息长度
	SetDataLen(uint32)
	// GetData 获取消息内容
	GetData() []byte
	// SetData 修改消息内容
	SetData([]byte)
}
