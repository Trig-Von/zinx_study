package ziface

/*
抽象IRequest 一次性请求的数据封装
*/

type IRequest interface {
	//得到请求连接
	GetConnection() IConnection
	//得到连接数据
	GetData() []byte
	//得到连接长度
	GetDataLen() int
}