/**
 zinx v0.3 应用
*/
package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

//PreHandle方法  ---  用户可以在处理业务之前  自定义一些业务， 实现这个方法
//Handler方法  ---- 用户可以定义一个 业务处理的 核心方法
//PostHandle方法  --- 用户可以在处理业务之后 定义一些业务，实现这个方法
type PingRouter struct {
	znet.BaseRouter
}

//处理业务之前的方法
func (this *PingRouter)PreHandle(request ziface.IRequest){
	fmt.Println("Call Router PreHandle...")
	//回写客户端
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("Call back before ping error")
	}
}
//处理业务的方法
func (this *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	//回写客户端
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n"))
	if err != nil {
		fmt.Println("Call back  ping error")
	}
}
//处理业务之后的方法
func (this *PingRouter)PostHandle(request ziface.IRequest){
	fmt.Println("Call Router afterHandle...")
	//回写客户端
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("Call back after ping error")
	}
}

func main() {
	//创建一个zinx server对象
	s := znet.NewServer("zinx v0.1")

	//注册一些自定义的业务
	s.AddRouter(&PingRouter{})

	//让server对象 启动服务
	s.Serve()

	return
}
