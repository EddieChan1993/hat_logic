package bi

import (
	"context"
	"git.dhgames.cn/svr_comm/gclient/utlog"
	"git.dhgames.cn/svr_comm/gclient/utlog/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/tool/consul/static"
	"hat_logic/util"
	"sync"
	"time"
)

var worker *logManger

type logManger struct {
	batchRecv  chan *pbgo.PushLogReq // 批量日志接收
	singleRecv chan *pbgo.PushLogReq // 单条日志接收
	logs       []*pbgo.PushLogReq    // 等待上报的批量日志
	ticker     *time.Ticker          // 上报批量日志的Ticker
	ctx        context.Context
	ctxCancel  context.CancelFunc
	wg         sync.WaitGroup
}

func InitBi() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	worker = &logManger{
		batchRecv:  make(chan *pbgo.PushLogReq, 100),
		singleRecv: make(chan *pbgo.PushLogReq, 3),
		logs:       make([]*pbgo.PushLogReq, 0, 100),
		ticker:     time.NewTicker(10 * time.Second),
		ctx:        ctx,
		ctxCancel:  ctxCancel,
		wg:         sync.WaitGroup{},
	}
	worker.wg.Add(1)
	go worker.run()
}

func Close() {
	worker.ctxCancel()
	worker.wg.Wait()
	klog.Info("bi close complete")
}

func (lm *logManger) run() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			klog.Error(panicErr)
			util.PanicStack()
		}
		lm.wg.Done()
		lm.ticker.Stop()
	}()
	for {
		select {
		case <-lm.ctx.Done():
			lm.pushAllLog()
			return
		case <-lm.ticker.C:
			//定时处理
			lm.pushAllLog()
		case log := <-lm.batchRecv:
			lm.addLog(log)
		case log := <-lm.singleRecv:
			lm.addLog(log)
		}
	}
}

// 上报批量日志
func (lm *logManger) pushAllLog() {
	if len(lm.logs) > 0 {
		_, dc, cluster := static.StaticBi()
		batchLog := &pbgo.PushBatchLogReq{BatchLog: lm.logs}
		rsp, err := utlog.PushBatchLog(dc, cluster, batchLog, kite.Cast)
		if err != nil {
			klog.Error("utlog.PushBatchLog err code:", err, rsp.ErrorCode)
		}
		lm.logs = lm.logs[:0]
	}
}

// 上报单条日志
func (lm *logManger) pushSingleLog(log *pbgo.PushLogReq) {
	if len(lm.logs) > 0 {
		_, dc, cluster := static.StaticBi()
		rsp, err := utlog.PushLog(dc, cluster, log, kite.Cast)
		if err != nil {
			klog.Error("utlog.PushLog err code:", err, rsp.ErrorCode)
		}
		lm.logs = lm.logs[:0]
	}
}

// 添加进批量日志
func (lm *logManger) addLog(log *pbgo.PushLogReq) {
	lm.logs = append(lm.logs, log)
	if len(lm.logs) >= 100 {
		lm.pushAllLog()
	}
}
