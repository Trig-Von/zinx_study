package ziface

type IConnManager interface {
	//添加链接
	Add(conn IConnection)
	//删除链接
	Remove(connId uint32)
	//根据链接ID得到连接
	Get(connId uint32)(IConnection,error)
	//得到目前服务器的连接总个数
	Len() uint32
	//清空全部链接的方法
	ClearConn()
}