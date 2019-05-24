/**
 zinx v0.2 应用
*/
package main

import (
	"zinx/znet"
)

func main() {
	//创建一个zinx server对象
	s := znet.NewServer("zinx v0.2")

	//注册一些自定义的业务
	//s.AddRouter(1, &dsad)

	//让server对象 启动服务
	s.Serve()

	return
}
