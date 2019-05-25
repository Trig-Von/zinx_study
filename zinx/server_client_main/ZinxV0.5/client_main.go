
package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

/*
	模拟客户端
 */
func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	//直接connect 服务器得到一个 已经建立好的conn句柄
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start errr", err)
		return
	}

	for {
		dp := znet.NewDataPack()

		binaryMsg,err:=dp.Pack(znet.NewMsgPackage(0,[]byte("Zinx 0.5 client Test Message..")))
		if err != nil {
			fmt.Println("Pack error ",err)
			return
		}
		if _,err := conn.Write(binaryMsg);err != nil {
			fmt.Println("write error ",err)
			return
		}
		//服务器会返回给我们一个消息Id为1的pingping TLV格式的二进制数据
		binaryHead := make([]byte,dp.GetHeadLen())
		if _,err := io.ReadFull(conn,binaryHead);err!= nil {
			fmt.Println("client unpack msgHead error",err)
			return
		}
		//根据头的长度进行第二次读取
		msgHead,err := dp.UnPack(binaryHead)
		if msgHead.GetMsgLen()>0 {
			//读取包体
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte,msg.GetMsgLen())
			if _,err := io.ReadFull(conn,msg.Data);err !=nil {
				fmt.Println("read msg data error",err)
				return
			}
			fmt.Println("-->Recv Server Msg : id =",msg.Id,"len = ",msg.Datalen,"data = ",string(msg.Data))
		}

		time.Sleep(1 *time.Second)
	}


}
