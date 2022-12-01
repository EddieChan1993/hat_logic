package static

import (
	"git.dhgames.cn/svr_comm/gcore/consul"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
)

var static = &StaticCfg{}

type StaticCfg struct {
	DebugMode             bool
	WaitReconnectWaitTime uint32 //离线后，等待重连时间，单位s
	Redis                 string //redis地址
	GPayCluster           string //gpay集群（废弃）
	GPayDC                string //gpay数据中心（废弃）
	GPayUrl               string //gpay http请求地址
	BiReport              bool   // bi日志发送开关
	BiDc                  string // bi日志数据中心
	BiCluster             string // bi日志集群
}

func NewStatic() *StaticCfg {
	return static
}

func (this_ *StaticCfg) Reload() {
	static = this_
	klog.Info("reload consul 静态配置完成")
}

func (this_ *StaticCfg) New() consul.IConfig {
	return &StaticCfg{}
}

//==================== 调用函数 ====================//

//StaticGPay gpay配置
func StaticGPay() (cluster, dc string) {
	return static.GPayCluster, static.GPayDC
}

//StaticGPayUrl gpay http请求地址
func StaticGPayUrl() string {
	return static.GPayUrl
}

//StaticRedisUrl redis地址
func StaticRedisUrl() string {
	return static.Redis
}

//StaticWaitReconnectWaitTime 等待重连时间
func StaticWaitReconnectWaitTime() uint32 {
	return static.WaitReconnectWaitTime
}

// StaticBi Bi上报配置
func StaticBi() (bool, string, string) {
	return static.BiReport, static.BiDc, static.BiCluster
}

//StaticDebugMode 调试模式
//gm 模块
func StaticDebugMode() bool {
	return static.DebugMode
}
