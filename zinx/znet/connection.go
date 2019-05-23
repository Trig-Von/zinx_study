package znet

import (
	"fmt"
	"net"
	"zinx/config"
	"zinx/ziface"
)

//具体的TCP连接模块
type Connection struct {
	//当前链接的套接字
	Conn *net.TCPConn
	//链接ID
	ConnID uint32
	//当前的链接状态
	isClosed bool
	//当前链接所绑定的业务处理方法
	//handleAPI ziface.HandleFunc

	//当前链接所绑定的Router
	Router ziface.IRouter
}
//初始化连接方法
func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter)ziface.IConnection  {
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		//handleAPI:callback_api,
		Router:router,
		isClosed:false,
	}
	return  c
}

//针对链接读业务的方法
func (c *Connection) StartReader()  {
	//从对端读数据
	fmt.Println("Reader go i starting..")
	defer fmt.Println("connId = ",c.ConnID,"Reader i exit ,remote addr is ",c.GetRemoteAddr().String())
	defer c.Stop()

	for  {
		buf := make([]byte,config.GlobalObject.MaxPackageSize)
		cnt,err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err",err)
			continue
		}
		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req := NewRequst(c,buf,cnt)

		//调用用户传递进来业务 模块 设计模式
		go func() {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()
		/*
		//将数据传递给定义好的handle CallBack方法
		if err := c.handleAPI(req); err!= nil {
			fmt.Println("ConnId",c.ConnID,"Handle is err",err)
			break
		}
		*/
	}
}


//启动连接
func (c *Connection)Start(){
fmt.Println("Conn Start() ...id= ",c.ConnID)

go c.StartReader()

}
//停止连接
func (c *Connection)Stop(){
fmt.Println("c.Stop() ...ConnId=",c.ConnID)

	if c.isClosed==true {
		return
	}
	c.isClosed=true
	_ = c.Conn.Close()

}
//获取链接ID
func (c *Connection)GetConnID() uint32{
	return c.ConnID
}
//获取conn的原生socket套接字
func (c *Connection)GetTCPConnection() *net.TCPConn{
	return c.Conn
}
//获取远程客户端的ip地址
func (c *Connection)GetRemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}
//发送数据给对方客户端
func (c *Connection)Send(data []byte,cnt int) error{
	if _,err := c.Conn.Write(data[:cnt]); err != nil{
		fmt.Println("send buf err")
		return err
	}
	return nil
}
