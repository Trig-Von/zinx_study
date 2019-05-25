package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
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
		/*buf := make([]byte,config.GlobalObject.MaxPackageSize)
		cnt,err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err",err)
			continue
		}*/
		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		//创建拆包封包的对象
		dp := NewDataPack()

		//读取客户端消息的头部
		headData := make([]byte,dp.GetHeadLen())
		if _,err := io.ReadFull(c.Conn,headData);err!= nil {
			fmt.Println("read msg head error",err)
			break
		}
		//根据头部,获取数据的长度，进行第二次读取
		msg,err := dp.UnPack(headData)
		if err !=nil {
			fmt.Println("unpack error",err)
			return
		}
		//根据长度再次读取
		var data []byte
		if msg.GetMsgLen()>0 {
			//有内容
			data = make([]byte,msg.GetMsgLen())
			if _,err := io.ReadFull(c.Conn,data);err!= nil {
				fmt.Println("read msg data error",err)
				break
			}
		}
		msg.SetData(data)
		//将读出来的msg 组装一个request
		//将当前一次性得到的对端客户端请求的数据 封装成一个request
		req := NewRequst(c,msg)

		//调用用户传递进来业务 模块 设计模式
		go func() {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()

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
func (c *Connection)Send(msgId uint32,msgData []byte) error{
	if c.isClosed == true {
		return errors.New("Connection closed ..send Msg")
	}
	//封装成msg
	dp := NewDataPack()
	binaryMsg,err:= dp.Pack(NewMsgPackage(msgId,msgData))
	if err != nil {
		fmt.Println("Pack error msg id = ",msgId)
		return err
	}
	//将binaryMsg发送给客户端

	if _,err := c.Conn.Write(binaryMsg); err != nil{
		fmt.Println("send buf err")
		return err
	}
	return nil
}
