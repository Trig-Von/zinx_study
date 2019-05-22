package net

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	IPVersoin string
	IP string
	Port int
	Name string
}
//初始化new方法
func NewServer (name string) ziface.IServer  {
	s := &Server{
		Name:name,
		IPVersoin:"tcp4",
		IP:"0.0.0.0",
		Port:8999,
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
	//阻塞等待客户端发送请求
	go func() {
		for  {
			//阻塞等待客户端请求
			conn,err := listener.Accept()
			if err !=nil{
				fmt.Println("Accept err :",err)
				continue
			}
			//此时conn已经与对端客户端连接
			go func() {
				//客户端有数据请求，处理客户端业务（读、写）
				for  {
					buf := make([]byte,512)
					cnt,err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ",err)
						break
					}
					fmt.Printf("recv cllient buf %s,cnt :%d\n",buf,cnt)
					//回显功能
					if _,err := conn.Write(buf[:cnt]);err != nil {
						fmt.Println("write back buf err :",err)
						continue
					}
				}
			}()
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


