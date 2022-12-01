package core

import (
	"git.dhgames.cn/svr_comm/kite/utils/klog"
)

func InitCore() {
	InitPlayerMgr()
}

func CloseCore() {
	playerMgr.stopServer()
	klog.Info("core close complete")
}
