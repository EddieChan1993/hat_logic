package util

import (
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"runtime"
)

//PanicStack 捕获recover的panic堆栈
func PanicStack() {
	buf := make([]byte, 1<<10)
	runtime.Stack(buf, true)
	klog.Error(string(buf))
}
