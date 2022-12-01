package cron

import "git.dhgames.cn/svr_comm/kite/utils/klog"

/**
全局定时任务
*/

func InitCron() {
	newCron()
	registerCronTask()
	klog.Info("全局 cron ok")
}

func CloseCron() {
	cron.ctxCancel()
	cron.wg.Wait()
	klog.Info("全局 cron close complete")
}
