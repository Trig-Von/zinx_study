package znet

import (
	"bytes"
	"encoding/binary"
	"zinx/ziface"
)

type DataPack struct {

}

//初始化一个DataPack对象
func NewDataPack() *DataPack  {
	return  &DataPack{}
}

//获取二进制包的头部长度 固定返回8
func (dp *DataPack)GetHeadLen() uint32{
	//4+4
	return 8
}
//封包方法--将Message打包成|data
func (dp *DataPack)Pack(msg ziface.IMessage)([]byte,error){
	//创建一个存放二进制的字节缓冲
	dataBuffer := bytes.NewBuffer([]byte{})
	//将datalen写进buffer
	if err := binary.Write(dataBuffer,binary.LittleEndian,msg.GetMsgLen());err != nil {
		return nil ,err
	}
	//将dataId写进buffer
	if err := binary.Write(dataBuffer,binary.LittleEndian,msg.GetMsgId());err != nil {
		return nil ,err
	}
	//将data写进buffer
	if err := binary.Write(dataBuffer,binary.LittleEndian,msg.GetMsgData());err != nil {
		return nil ,err
	}
	//返回这个缓冲
	return dataBuffer.Bytes(),nil
}
//拆包方法--将|datalen|dataId|data
func (dp *DataPack)UnPack(binaryData []byte)(ziface.IMessage,error){
	//解包的时候 分两次  1：固定的长度8字节 2：根据len再次进行read
	msgHead := &Message{}

	//创建一个 读取二进制数据流的io.reader
	dataBuff := bytes.NewReader(binaryData)
	//将二进制流的dataLen放进Msg的dataLen属性中
	if err := binary.Read(dataBuff,binary.LittleEndian,&msgHead.Datalen);err !=nil{
		return nil,err
	}
	//将二进制流的dataId放进Msg的dataLen属性中
	if err := binary.Read(dataBuff,binary.LittleEndian,&msgHead.Id);err !=nil{
		return nil,err
	}
	return msgHead,nil
}