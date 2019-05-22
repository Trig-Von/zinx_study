package net

import (
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
	handleAPI ziface.HandleFunc
}
//初始化连接方法
func NewConnection(conn *net.TCPConn,connID uint32,callback_api ziface.HandleFunc)ziface.IConnection  {
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		handleAPI:callback_api,
		isClosed:false,
	}
	return  c
}

//启动连接
func (c *Connection)Start(){

}
//停止连接
func (c *Connection)Stop(){

}
//获取链接ID
func (c *Connection)GetConnID() uint32{
	return 0
}
//获取conn的原生socket套接字
func (c *Connection)GetTCPConnection() *net.TCPConn{
	return nil
}
//获取远程客户端的ip地址
func (c *Connection)GetRemoteAddr() net.Addr{
	return nil
}
//发送数据给对方客户端
func (c *Connection)Send(data []byte) error{
	return nil
}
