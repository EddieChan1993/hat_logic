package timewheel

import (
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"github.com/RussellLuo/timingwheel"
	"hat_logic/tool/timewheel/register"
	"time"
)

type CallBack func()

var tw TimeWheeler //离线推送时间轮

func InitTimeWheel() {
	tw = registerTimerWheel()
}

// StartTick 开一个周期循环调用的timer 关闭的话，直接用返回的timer调用Stop即可
func StartTick(d time.Duration, f CallBack) *timingwheel.Timer {
	return tw.ScheduleFunc(&register.DefaultScheduler{Interval: d}, f)
}

// StartTimer 开一个一次性的timer 关闭的话，直接用返回的timer调用Stop即可
func StartTimer(d time.Duration, f CallBack) *timingwheel.Timer {
	return tw.AfterFunc(d, f)
}

func CloseTw() {
	tw.CloseTw()
	klog.Info("timeWheel 关闭")
}
