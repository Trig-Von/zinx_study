/**
 zinx v0.4 应用
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
func (this *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	//回写客户端
	err := request.GetConnection().Send(1,[]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter  struct {
	znet.BaseRouter
}

//处理业务之前的方法
func (this *HelloRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	//回写客户端
	err := request.GetConnection().Send(2,[]byte("hello Zinx!!\n"))
	if err != nil {
		fmt.Println(err)
	}
}


func main() {
	//创建一个zinx server对象
	s := znet.NewServer("zinx v0.6")

	//注册一些自定义的业务
	s.AddRouter(1,&PingRouter{})
	s.AddRouter(2,&HelloRouter{})
	//让server对象 启动服务
	s.Serve()

	return
}
