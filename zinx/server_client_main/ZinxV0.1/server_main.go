
package main

import "zinx/znet"

func main() {
	//创建一个zinx server对象
	s := znet.NewServer("zinx v0.1")

	//让server对象 启动服务
	s.Serve()
	select {}

}
