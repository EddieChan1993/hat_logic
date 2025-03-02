package goRuntime

import (
	"context"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/util"
	"sync"
)

type tGoRuntime struct {
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

var goRuntime *tGoRuntime

func InitGoRuntime() {
	rootCtx, cancel := context.WithCancel(context.Background())
	goRuntime = &tGoRuntime{
		wg:     &sync.WaitGroup{},
		ctx:    rootCtx,
		cancel: cancel,
	}
}

func GoRun(fn func(ctx context.Context)) {
	goRuntime.wg.Add(1)
	go func() {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				klog.Error(panicErr)
				util.PanicStack()
			}
			goRuntime.wg.Done()
		}()
		fn(goRuntime.ctx)
	}()
}

func CloseGoRuntime() {
	goRuntime.cancel()
	goRuntime.wg.Wait()
}
