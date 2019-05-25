package ziface

/*
抽象IRequest 一次性请求的数据封装
*/

type IRequest interface {
	//得到请求连接
	GetConnection() IConnection
	//得到请求的消息
	GetMsg() IMessage
}