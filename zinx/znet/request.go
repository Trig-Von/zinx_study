package znet

import (
	"zinx/ziface"
)

type Request struct {
	//链接信息
	conn ziface.IConnection
	//客户端发送的消息
	msg ziface.IMessage
}

func NewRequst(conn ziface.IConnection,msg ziface.IMessage)ziface.IRequest {
	req := &Request{
		conn:conn,
		msg:msg,
	}
	return req
}

//得到请求连接
func (r *Request)GetConnection() ziface.IConnection{
	return r.conn
}
//得到连接数据
func (r *Request)GetMsg() ziface.IMessage{
	return r.msg
}

