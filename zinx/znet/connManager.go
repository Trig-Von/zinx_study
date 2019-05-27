package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/ziface"
)

type ConnManager struct {
	//管理的全部链接
	connections map[uint32] ziface.IConnection
	connLock sync.RWMutex
}

func NewConnManager() ziface.IConnManager  {
	return  &ConnManager{
		connections:make(map[uint32] ziface.IConnection),
	}
}
//添加链接
func (connMgr *ConnManager) Add(conn ziface.IConnection){
	//枷锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	connMgr.connections[conn.GetConnID()] = conn

	fmt.Println("Add connId = ",conn.GetConnID(),"to manage succ!!")
}
//删除链接
func (connMgr *ConnManager) Remove(connId uint32){
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.connections,connId)
	fmt.Println("REmove connId = ",connId,"from manager succ!!")
}
//根据链接ID得到连接
func (connMgr *ConnManager) Get(connId uint32)(ziface.IConnection,error){
	//加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn,ok:= connMgr.connections[connId];ok {
		//找到了
		return conn,nil
	}else {
		//没找到
		return nil,errors.New("connnection not FOUND!!")
	}
}
//得到目前服务器的连接总个数
func (connMgr *ConnManager) Len() uint32{
	return uint32(len(connMgr.connections))
}
//清空全部链接
func (connMgr *ConnManager) ClearConn(){
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	//遍历删除
	for connID,conn := range connMgr.connections{
		//将全部的conn关闭
		conn.Stop()
		//删除链接
		delete(connMgr.connections,connID)
	}
	fmt.Println("Clear all Connection SUCC!!conn num = ",connMgr.Len())
}