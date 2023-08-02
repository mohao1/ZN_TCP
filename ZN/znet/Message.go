package znet

type Message struct {
	//消息ID
	Id uint32
	//消息长度
	DataLen uint32
	//消息内容
	Data []byte
}

// GetMsgID 获取消息ID
func (m *Message) GetMsgID() uint32 {
	return m.Id
}

// SetMsgID 修改消息ID
func (m *Message) SetMsgID(id uint32) {
	m.Id = id
}

// GetDataLen 获取消息长度
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

// SetDataLen 修改消息长度
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

// GetData 获取消息内容
func (m *Message) GetData() []byte {
	return m.Data
}

// SetData 修改消息内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}

// NewMessage 创建Message的方法
func NewMessage(msgId uint32, data []byte) *Message {
	return &Message{
		Id:      msgId,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}
