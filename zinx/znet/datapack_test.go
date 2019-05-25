package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T)  {
	fmt.Println("testing datapack...")

	/*
	  模拟写一个server
	  收到二进制流 进行解包
	*/
	// 1 创建listenner
	listenner ,err := net.Listen("tcp","127.0.0.1:7777")
	if err != nil{
		fmt.Println("server listener ere",err)
		return
	}
	// 2 Accept TCP
	go func() {
		for  {
			conn,err := listenner.Accept()
			if err!= nil {
				fmt.Println("server accept err" ,err)
				return
			}
			//读写业务
			go func(conn *net.Conn) {
				//读取客户端的请求
				//拆包过程

				dp:= NewDataPack()
				for  {
					//进行第一次从conn读，吧head读出来
					headData := make([]byte,dp.GetHeadLen())
					_,err := io.ReadFull(*conn,headData)
					if err != nil {
						fmt.Println("read head err",err)
						break
					}
					msgHead ,err := dp.UnPack(headData)
					if err !=nil {
						fmt.Println("server unpack err",err)
						return
					}
					if msgHead.GetMsgLen()>0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte,msg.GetMsgLen())
						_,err := io.ReadFull(*conn,msg.Data)
						if err !=nil {
							fmt.Println("servet unpack data err",err)
							return
						}
						fmt.Println("--->Recv MsgId = ",msg.Id,"datalen= ",msg.Datalen,"data =",string(msg.Data))

					}
				}
			}(&conn)
		}
	}()
	/*
	 模拟写一个client  封包之后再发包
	*/
	//connection Dail
	conn,err := net.Dial("tcp","127.0.0.1:7777")
	if err!= nil {
		fmt.Println("client dail err",err)
		return
	}
	dp := NewDataPack()

	msg1 := &Message{
		Id:1,
		Datalen:5,
		Data:[]byte{'h','e','l','l','o'},
	}
	sendData1,err := dp.Pack(msg1)
	if err!= nil {
		fmt.Println("client send data1 error")
		return
	}
	msg2 := &Message{
		Id:2,
		Datalen:4,
		Data:[]byte{'z','i','n','x'},
	}
	sendData2,err := dp.Pack(msg2)
	if err!= nil {
		fmt.Println("client send data1 error")
		return
	}

	sendData1 = append(sendData1,sendData2...)
	conn.Write(sendData1)

	//让test不结束
	select {}
}
