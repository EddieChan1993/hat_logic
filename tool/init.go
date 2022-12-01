package tool

import (
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/tool/bi"
	"hat_logic/tool/consul"
	"hat_logic/tool/redisdb"
	resty "hat_logic/tool/restyHttp"
	"hat_logic/tool/timewheel"
)

func InitTool() {
	consul.InitConsul()
	redisdb.InitRedisDb()
	bi.InitBi()
	timewheel.InitTimeWheel()
	resty.InitResty()
}

func ReleaseTool() {
	timewheel.CloseTw()
	klog.Info("tool close complete")
}
