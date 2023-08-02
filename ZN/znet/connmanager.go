package znet

import (
	"ZN/ZN/ziface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	//存储连接
	connectionMap map[uint32]ziface.IConnection
	connLock      sync.RWMutex
}

// Add 添加连接
func (m *ConnManager) Add(conn ziface.IConnection) {
	//加锁
	m.connLock.Lock()
	defer m.connLock.Unlock()

	m.connectionMap[conn.GetConnID()] = conn

	//fmt.Println("ConnManager Add OK ConnID :", conn.GetConnID())
}

// Remove 删除连接
func (m *ConnManager) Remove(conn ziface.IConnection) {
	//加锁
	m.connLock.Lock()
	defer m.connLock.Unlock()

	delete(m.connectionMap, conn.GetConnID())

	//fmt.Println("ConnManager Delete OK ConnID :", conn.GetConnID())
}

// Get 获取连接
func (m *ConnManager) Get(id uint32) (ziface.IConnection, error) {
	//加锁
	m.connLock.RLock()
	defer m.connLock.RUnlock()

	if conn, ok := m.connectionMap[id]; ok {
		fmt.Println("ConnManager GetConn OK ConnID :", conn.GetConnID())
		return conn, nil
	} else {
		return nil, errors.New("get Conn Error: Conn Is Null")
	}

}

// Len 得到连接总数
func (m *ConnManager) Len() int {
	return len(m.connectionMap)
}

// ClearConn 清除所有连接
func (m *ConnManager) ClearConn() {
	//加锁
	m.connLock.Lock()
	defer m.connLock.Unlock()
	//停止连接并且删除连接
	for connId, conn := range m.connectionMap {
		//停止连接
		conn.Stop()
		//删除连接
		delete(m.connectionMap, connId)
	}
	//fmt.Println("ConnManager Delete All OK")
}

// NewConnManager 创建方法
func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		connectionMap: make(map[uint32]ziface.IConnection),
	}
}
