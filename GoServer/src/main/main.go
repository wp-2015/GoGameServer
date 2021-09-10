package main

import (
	"GameServer/src/Managers"
	"GameServer/src/MessageHandles"
	"GameServer/utils"
	"GameServer/ziface"
	"GameServer/znet"
	"fmt"
)

type TestRouter struct {
	znet.BaseRouter
}

func (this *TestRouter) Handle(request ziface.IRequest)  {
	fmt.Println("TestRouter Handle")
}

func main() {
	fmt.Println("asdasd")
	s := znet.NewServer()
	//Demo.AddHandle(s)
	MessageHandles.AddHandles(s)
	//3 开启服务
	s.Serve()

	//计时器
	var interval int64 = 1000 / 30
	var timeRecord int64 = 0
	timestamp := utils.GetUnixMillis()     //时间戳
	fmt.Println("1", timestamp)
	for true {
		timestamp = utils.GetUnixMillis()
		//mapCreater.Update(timestamp)
		Managers.GMInstance.Update(timestamp)
		if timestamp - timeRecord > interval {
			MainLoop(timestamp)
			timeRecord = timestamp
		}
	}
}

func MainLoop(millisecond int64){
	//fmt.Println(millisecond)
}