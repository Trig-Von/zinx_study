package main

import (
	"fmt"
	"net"
	"time"
)

//模拟客户端
func main ()  {
	fmt.Println("client online...")
	time.Sleep(1*time.Second)

	conn,err := net.Dial("tcp","127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err",err)
		return
	}
	for  {
		//写
		_,err := conn.Write([]byte("Hello Zinx..."))
		if err != nil {
			fmt.Println("write conn err",err)
			return
		}
		//读
		buf := make([]byte,512)
		cnt,err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err",err)
			return
		}
		fmt.Printf("server call back:%s,cnt:%d\n",buf,cnt)

		time.Sleep(1* time.Second)
	}
}