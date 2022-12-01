package main

import (
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	gchat "git.dhgames.cn/svr_comm/gclient/gchat/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/configs"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/api"
	"hat_logic/core"
	"hat_logic/cron"
	"hat_logic/pbgo/pbapi"
	"hat_logic/tool"
	"hat_logic/tool/bi"
	"hat_logic/util"
)

func init() {
	util.InitUtil()
	tool.InitTool()
	core.InitCore()
	klog.Info("server init ok")
}

func main() {
	klog.Info("start kite")

	if configs.Upgrade() {
		// 热更新这里不处理
	} else {
		cron.InitCron()
	}
	startKite()
}

func startKite() {
	if util.IsLocalRun() {
		configs.GlobalCommonObject.MetricsReportURL = ""
		kite.WhoAmI(util.GetLocalClusterArgs(), util.GetLocalServerArgs(), util.GetLocalArgsIndex())
	}
	pbapi.RegGmpApiServer(&api.Gm{})
	gchat.RegGameServerChatServer(&api.GChat{})
	pbgo.RegLogicServer(&api.GoGateApi{})
	kite.Serve(&Process{})
}

type Process struct {
}

func (this_ *Process) OldBefore(upGradeLevel int) {
	cron.CloseCron() // 避免重复发奖，热更前，停止任务
	if upGradeLevel == 1 {
		klog.Infof("========== upGradeLevel=%d 踢人热更开始 ==========", upGradeLevel)
		core.KickBeforeHotUpdate()
		klog.Infof("========== upGradeLevel=%d 踢人热更完成 ==========", upGradeLevel)
	}
}

func (this_ *Process) NewBefore(upGradeLevel int) {
}

func (this_ *Process) OldAfter(upGradeLevel int) {
}

func (this_ *Process) NewAfter(upGradeLevel int) {
	cron.InitCron() //热更完后启动定时任务
}

func (this_ *Process) SendData(send func(data []byte)) error {
	klog.Info("========== SendData 开始 ==========")
	core.HotSendData(send)
	klog.Info("========== SendData 完成 ==========")
	return nil
}

func (this_ *Process) ReceiveData(data []byte) error {
	core.HotReceiveData(data)
	klog.Info("ReceiveData 完成")
	return nil
}

func (this_ *Process) Stop() {
	AfterStopServe()
}

func AfterStopServe() {
	tool.ReleaseTool()
	core.CloseCore()
	cron.CloseCron()
	bi.Close()
}
