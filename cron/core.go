package cron

import (
	"context"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/util"
	"sync"
	"time"
)

type ICronTaskFn interface {
	CronTime() time.Duration
	CronTask()
}

type Cron struct {
	ctxCron   context.Context
	ctxCancel context.CancelFunc
	wg        sync.WaitGroup
}

var cron *Cron

func newCron() {
	ctxCron, ctxCancel := context.WithCancel(context.Background())
	cron = &Cron{
		ctxCron:   ctxCron,
		ctxCancel: ctxCancel,
		wg:        sync.WaitGroup{},
	}
}

func (this_ *Cron) registerCronTask(cronTask ICronTaskFn) {
	tick := time.NewTicker(cronTask.CronTime())
	this_.wg.Add(1)
	go func() {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				klog.Error(panicErr)
				util.PanicStack()
			}
			this_.wg.Done()
			tick.Stop()
		}()
		for {
			select {
			case <-this_.ctxCron.Done():
				return
			case <-tick.C:
				cronTask.CronTask()
			}
		}
	}()
}
