package znet

import (
	"fmt"
	"zinx/config"
	"zinx/ziface"
)

type MsgHandler struct {
	//存放路由集合的 map
	Apis map[uint32] ziface.IRouter
	//负责worker取任务的消息队列
	TaskQueue []chan ziface.IRequest

	//worker工作池的worker数量'
	WorkerPoolSize uint32
}

func NewMsgHandler()  ziface.IMsgHandler  {
	return &MsgHandler{
		Apis:make(map[uint32]ziface.IRouter),
		WorkerPoolSize:config.GlobalObject.WorkerPoolSize,
		TaskQueue:make([]chan ziface.IRequest,config.GlobalObject.WorkerPoolSize),
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

//一个worker处理业务的goroutine函数
func (mh *MsgHandler) startOneWorker(workerID int,taskQueue chan ziface.IRequest)  {
	fmt.Println("workerId = ",workerID,"is starting...")

	for  {
		select {
		case req:= <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

//启动Worker工作池
func (mh *MsgHandler) StartWorkerPool()  {
	fmt.Println("WorkerPool is starting...")

	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest,config.GlobalObject.MaxWorkerTaskLen)

		go mh.startOneWorker(i,mh.TaskQueue[i])
	}

}

//将消息添加到工作池中
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest)  {
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize

	mh.TaskQueue[workerID] <- request
}