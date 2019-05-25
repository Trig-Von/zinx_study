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

	MsgHandler ziface.IMsgHandler
}


//初始化new方法
func NewServer (name string) ziface.IServer  {
	s := &Server{
		Name:config.GlobalObject.Name,
		IPVersoin:"tcp4",
		IP:config.GlobalObject.Host,
		Port:config.GlobalObject.Port,
		MsgHandler:NewMsgHandler(),
	}
	return  s
}

func (s *Server)Start()  {
	fmt.Printf("[start]Server listener at IP:%s,Port:%d,is starting",s.IP,s.Port)

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

			//创建一个Connection对象
			dealConn := NewConnection(conn,cid, s.MsgHandler)
			cid++

			//此时conn就已经和对端客户端连接
			go dealConn.Start()
		}
	}()
}

func (s *Server)Stop()  {
	
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

