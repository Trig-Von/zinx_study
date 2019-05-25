package ziface

type DataPack interface {
	//获取二进制包的头部长度
	GetHeadLen() uint32
	//封包方法--将Message打包成|data
	Pack(msg IMessage)([]byte,error)
	//拆包方法--将|datalen|dataId|data
	UnPack([]byte)(IMessage,error)
}
