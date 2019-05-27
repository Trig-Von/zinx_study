package znet

import (
	"fmt"
	"net"
	"zinx/config"
	"zinx/ziface"
)

type Server struct {
	IPVersoin string
	IP string
	Port int
	Name string

	//多路由管理
	MsgHandler ziface.IMsgHandler
	//连接管理
	connMgr ziface.IConnManager
	//server创建连接之后自动调用Hook函数
	OnConnStart func(conn ziface.IConnection)
	//server销毁链接之前自动调用的Hook函数
	OnConnStop func(conn ziface.IConnection)
}


//初始化new方法
func NewServer (name string) ziface.IServer  {
	s := &Server{
		Name:config.GlobalObject.Name,
		IPVersoin:"tcp4",
		IP:config.GlobalObject.Host,
		Port:config.GlobalObject.Port,
		MsgHandler:NewMsgHandler(),
		connMgr:NewConnManager(),
	}
	return  s
}

func (s *Server)Start()  {
	fmt.Printf("[start]Server listener at IP:%s,Port:%d,is starting",s.IP,s.Port)

	s.MsgHandler.StartWorkerPool()
	//创建套接字
	addr,err := net.ResolveTCPAddr(s.IPVersoin,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err !=nil {
		fmt.Println("resolve tcp addr error:",err)
		return
	}
	//监听服务器地址
	listener,err := net.ListenTCP(s.IPVersoin,addr)
	if err != nil {
		fmt.Println("listen",s.IPVersoin,"err,",err)
		return
	}

	var cid uint32
	cid = 0
	//阻塞等待客户端发送请求
	go func() {
		for  {
			//阻塞等待客户端请求
			conn,err := listener.AcceptTCP()
			if err !=nil{
				fmt.Println("Accept err :",err)
				continue
			}
			//判断当前server链接数量是否已经最大值
			if s.connMgr.Len()>=config.GlobalObject.MaxConn{
				fmt.Println("-->Too Many Connection MaxConn = ",config.GlobalObject.MaxConn)
				conn.Close()
				continue
			}
			//创建一个Connection对象
			dealConn := NewConnection(conn,cid, s.MsgHandler)
			cid++

			//此时conn就已经和对端客户端连接
			go dealConn.Start()
		}
	}()
}

func (s *Server)Stop()  {
	s.connMgr.ClearConn()
}
//运行服务器
func (s *Server)Serve()  {
	//启动server的监听功能
	s.Start()

	//阻塞
	select {}
}

func (s *Server)AddRouter(msgID uint32,router ziface.IRouter)  {
	s.MsgHandler.AddRouter(msgID,router)
	fmt.Println("Add Router Succ!!!msgID = ",msgID)
}

func (s *Server)GetConnMgr() ziface.IConnManager {
	return s.connMgr
}
//注册 创建连接之后 调用的Hook函数的方法
func (s *Server)AddOnConnStart(hookFunc func(conn ziface.IConnection)){
	s.OnConnStart = hookFunc
}
//注册 创建连接之前 调用的Hook函数的方法
func (s *Server)AddOnConnStop(hookFunc func(conn ziface.IConnection)){
	s.OnConnStop = hookFunc
}
//注册 创建连接之后的Hook函数的方法
func (s *Server)CallOnConnStart(conn ziface.IConnection){
	if s.OnConnStart != nil {
		fmt.Println("--> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}
//注册 销毁连接之后的Hook函数的方法
func (s *Server)CallOnConnStop(conn ziface.IConnection){
	if s.OnConnStop != nil {
		fmt.Println("-->Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}