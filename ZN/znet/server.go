package znet

import (
	"ZN/ZN/utils"
	"ZN/ZN/ziface"
	"fmt"
	"net"
)

type Server struct {
	//服务名称
	SName string
	//服务绑定版本
	IPVersion string
	//服务监听IP
	IP string
	//服务监听端口
	Port int
	//路由管理模块
	msgHandler ziface.IMsgHandler
	//连接管理模块
	connManger ziface.IConnManager
	//连接创建之后调用Hook
	OnConnStartHook ziface.OnHookFunc
	//连接断开之前调用Hook
	OnConnStopHook ziface.OnHookFunc
}

// Start 启动服务
func (s *Server) Start() {
	fmt.Printf("[ZN]:Server Run Name:%s IP: %d Port:%s \n",
		utils.GlobalOBJ.Name,
		utils.GlobalOBJ.TcpPort,
		utils.GlobalOBJ.Host,
	)
	go func() {
		//开启业务的线程池
		s.msgHandler.StartWorkerPool()
		//获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resoleTCError:", err)
			return
		}

		//监听服务器的地址
		tcp, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen:", s.IPVersion, "err:", err)
			return
		}

		//定义ConnID
		var ConnID uint32 = 0
		//阻塞等待用户的连接,处理客户业务(for进行循环的监听阻塞等待简历连接)
		for {
			conn, err := tcp.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			//进行连接限制
			if s.connManger.Len() >= utils.GlobalOBJ.MaxConn {
				//TODO 返回连接过多
				//fmt.Println("返回连接过多")
				//进行断开连接
				conn.Close()
				continue
			}

			//使用connection来代替原理的操作的方法
			Connection := NewConnection(s, conn, ConnID, s.msgHandler)
			Connection.Start()
			ConnID++

		}
	}()
}

// Server 运行服务
func (s *Server) Server() {

	//启动服务
	s.Start()

	//进行阻塞防止主线程的结束
	select {}

}

// Stop 停止服务
func (s *Server) Stop() {
	//回收资源
	s.connManger.ClearConn()

}

// AddRouter 用户添加Router方法
func (s *Server) AddRouter(id uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(id, router)
}

// GetConnManger 获取连接模块
func (s *Server) GetConnManger() ziface.IConnManager {
	return s.connManger
}

// SetOnConnStartHook 注册OnConnStartHook方法
func (s *Server) SetOnConnStartHook(hook ziface.OnHookFunc) {
	s.OnConnStartHook = hook
}

// SetOnConnStopHook 注册OnConnStopHook方法
func (s *Server) SetOnConnStopHook(hook ziface.OnHookFunc) {
	s.OnConnStopHook = hook
}

// CallOnConnStartHook 调用OnConnStartHook方法
func (s *Server) CallOnConnStartHook(conn ziface.IConnection) {
	//判断是否存在Hook
	if s.OnConnStartHook != nil {
		//fmt.Println("OnConnStartHook Run....")
		s.OnConnStartHook(conn)
	}
}

// CallOnConnStopHook 调用OnConnStopHook方法
func (s *Server) CallOnConnStopHook(conn ziface.IConnection) {
	//判断是否存在Hook
	if s.OnConnStopHook != nil {
		//fmt.Println("OnConnStopHook Run....")
		s.OnConnStopHook(conn)
	}
}

func NewZServer() ziface.IServer {
	s := Server{
		SName:      utils.GlobalOBJ.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalOBJ.Host,
		Port:       utils.GlobalOBJ.TcpPort,
		msgHandler: NewMsgHandler(),  //路由管理模块
		connManger: NewConnManager(), //连接管理模块
	}

	return &s
}
