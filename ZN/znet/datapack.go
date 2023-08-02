package znet

import (
	"ZN/ZN/utils"
	"ZN/ZN/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct {
}

// GetHeadLen 获取包头长度
func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

// Pack 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	//定义一个转换的缓存区
	dataBuff := bytes.NewBuffer([]byte{})

	//进行插入数据
	//插入长度
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	//插入ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	//插入数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// UnPack 拆包方法
func (dp *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	//接收数据
	red := bytes.NewReader(data)

	//解析数据
	msg := &Message{}

	//解析长度
	if err := binary.Read(red, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//解析ID
	if err := binary.Read(red, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断长度是否超过最大读取数据长度
	if utils.GlobalOBJ.MaxPackageSize > 0 && msg.DataLen > utils.GlobalOBJ.MaxPackageSize {
		return nil, errors.New("too large msg data received : 数据长度过长大于可用长度")
	}

	return msg, nil
}

// NewDataPack 初始化拆装包对象
func NewDataPack() ziface.IDataPack {
	return &DataPack{}
}
