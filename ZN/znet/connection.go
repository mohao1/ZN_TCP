package znet

import (
	"ZN/ZN/utils"
	"ZN/ZN/ziface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	//Conn对应关联到的Server
	TcpServer ziface.IServer
	//当前连接conn对象（套接字-服务连接对象）
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前连接是否已经关闭
	isClosed bool
	//等待连接被退出的channel
	ExitChan chan bool
	//读写分离channel
	msgChan chan []byte
	//路由管理模块
	msgHandler ziface.IMsgHandler
	//连接属性集合
	property map[string]interface{}
	//设置连接属性集合的锁
	propertyLock sync.RWMutex
}

func (c *Connection) StartReader() {
	//fmt.Println("Conn StartReader ConnID:", c.ConnID)
	//defer fmt.Println("Conn StartReader Stop ConnID:", c.ConnID)
	defer c.Stop()

	//读取数据进行处理操作
	for {
		//进行拆包读取
		dp := NewDataPack()

		//读取Msg的数据
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			break
		}

		//获取头部数据进行拆包获取头部数据
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("read msg head error ", err)
			break
		}

		//根据头部长度数据获取Data数据
		var data []byte
		if msg.GetDataLen() > 0 {
			//开辟了一个长度为数据长度的空间
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		//加入数据到Msg中
		msg.SetData(data)
		//封装Request对象
		req := Request{
			conn: c,
			Msg:  msg,
		}

		if utils.GlobalOBJ.WorkerPoolSize > 0 {
			c.msgHandler.SendMsgToTakeQueue(&req)
		} else {
			//调用路由管理模块
			go c.msgHandler.DoMsgHandler(&req)
		}

	}
}

// StartWrite 写出数据操作
func (c *Connection) StartWrite() {
	//defer fmt.Println("Conn StartWrite Stop ConnID:", c.ConnID)
	//循环监听数据
	for {
		select {
		//取得数据
		case data := <-c.msgChan:
			{
				//发送数据
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("发送失败 Write err:", err)
					return
				}

			}
			//连接退出
		case <-c.ExitChan:
			return

		}

	}
}

// Start 启动连接
func (c *Connection) Start() {
	fmt.Println("Conn Start .... ConnID:", c.ConnID)
	//进行读取数据编写
	go c.StartReader()
	//实现返回数据操作
	go c.StartWrite()
	//调用建立了连接之后的Hook方法
	c.TcpServer.CallOnConnStartHook(c)
}

// Stop 关闭连接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop .... ConnID:", c.ConnID)

	//判断是否已经进行断开操作
	if c.isClosed == true {
		return
	}
	//进行断开操作
	c.isClosed = true

	//调用关闭了连接之后的Hook方法
	c.TcpServer.CallOnConnStopHook(c)

	//关闭连接
	c.Conn.Close()
	c.ExitChan <- true
	//退出连接管理模块
	c.TcpServer.GetConnManger().Remove(c)
	//关闭管道监听回收资源
	close(c.ExitChan)
}

// GetTCPConnection 获取连接的conn对象(套接字-服务连接对象)
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获得客户端的连接地址和端口
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 发送数据操作方法
func (c *Connection) SendMsg(msgId uint32, msg []byte) error {
	//判断是否关闭连接
	if c.isClosed == true {
		return errors.New("已经断开连接")
	}

	dp := NewDataPack()

	DataMsg := NewMessage(msgId, msg)

	bytes, err := dp.Pack(DataMsg)
	if err != nil {
		fmt.Println("封包失败 Pack err:", err)
		return errors.New("封包失败 Pack err")
	}
	//调用异步发送数据
	c.msgChan <- bytes
	return nil
}

// SetProperty 添加连接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	//上锁
	c.propertyLock.Lock()
	c.propertyLock.Unlock()
	//添加一个连接属性
	c.property[key] = value
}

// GetProperty 获取连接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	//上锁
	c.propertyLock.RLock()
	c.propertyLock.RUnlock()
	//获取一个连接属性
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New(fmt.Sprint("Key:", key, "Is Null"))
	}

}

// RemoveProperty 删除连接属性
func (c *Connection) RemoveProperty(key string) {
	//上锁
	c.propertyLock.Lock()
	c.propertyLock.Unlock()
	//删除一个连接属性
	delete(c.property, key)
}

// NewConnection 创建连接方法
func NewConnection(tcpserver ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,      //true==关闭
		msgHandler: msgHandler, //路由管理模块
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
		TcpServer:  tcpserver,
		property:   make(map[string]interface{}),
	}
	//将连接自身加入到连接管理模块之中
	c.TcpServer.GetConnManger().Add(c)

	return c
}
