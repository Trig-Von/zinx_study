package znet

import (
	"fmt"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32] ziface.IRouter
}

func NewMsgHandler()  ziface.IMsgHandler  {
	return &MsgHandler{
		Apis:make(map[uint32]ziface.IRouter),
	}
}

//添加路由到map集合中
func (mh *MsgHandler)AddRouter(msgID uint32,router ziface.IRouter){
	if _,ok := mh.Apis[msgID];ok {
		fmt.Println("repeat Api msgId= ",msgID)
		return
	}
	mh.Apis[msgID] = router
	fmt.Println("Apd APi MsgID = ",msgID,"succ!")
}
//调度路由 根据MsgID
func (mh *MsgHandler)DoMsgHandler(request ziface.IRequest){
	router,ok := mh.Apis[request.GetMsg().GetMsgId()]
	if !ok {
		fmt.Println("api MsgID = ",request.GetMsg().GetMsgId(),"Not Found!Need Add!")
		return
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}