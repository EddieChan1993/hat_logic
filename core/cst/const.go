package cst

//常用常量

import "time"

const (
	PlayerStatusLoginOut PlayerStatus = 0 //玩家离线状态
	PlayerStatusLogin    PlayerStatus = 1 //玩家登陆状态
	PlayerStatusClose    PlayerStatus = 2 //玩家关闭
)

const ReqMsgTimeOut = 5 * time.Second //请求消息默认超时时间

const (
	PlayerOnlineDbSaveTick  = 300             //在线savedb 间隔单位s
	PlayerOfflineDbSaveTick = 3 * time.Minute //离线savedb
)

const (
	LogicErr = 1
)
