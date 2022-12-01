package timewheel

import (
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"github.com/RussellLuo/timingwheel"
	"time"
)

const (
	twTick = 5 * time.Second
	twSize = 60
)

type timeWheel struct {
	tw *timingwheel.TimingWheel
}

func (this_ *timeWheel) AfterFunc(t time.Duration, f func()) *timingwheel.Timer {
	return this_.tw.AfterFunc(t, f)
}

func (this_ *timeWheel) DelTimer(timer *timingwheel.Timer) {
	if timer == nil {
		klog.Info("timer=nil,退出先前时间轮注册")
		return
	}
	if timer.Stop() {
		klog.Info("退出时间轮")
	}
}

func (this_ *timeWheel) ScheduleFunc(s timingwheel.Scheduler, f func()) *timingwheel.Timer {
	return this_.tw.ScheduleFunc(s, f)
}

func (this_ *timeWheel) CloseTw() {
	this_.tw.Stop()
}

type TimeWheeler interface {
	ScheduleFunc(s timingwheel.Scheduler, f func()) *timingwheel.Timer
	AfterFunc(t time.Duration, f func()) *timingwheel.Timer
	DelTimer(timer *timingwheel.Timer) //删除timer事件
	CloseTw()                          //关掉时间轮
}

func registerTimerWheel() TimeWheeler {
	tw := timingwheel.NewTimingWheel(twTick, twSize)
	tw.Start()
	twIns := &timeWheel{
		tw: tw,
	}
	return twIns
}
