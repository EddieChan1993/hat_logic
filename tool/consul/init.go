package consul

import (
	"git.dhgames.cn/svr_comm/gcore/consul"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/tool/consul/static"
	"hat_logic/util"
)

var serviceInfo *consul.ServiceInfo

const MasterIndex = 1

func InitConsul() {
	var err error
	if util.IsLocalRun() {
		serviceInfo = &consul.ServiceInfo{
			Cluster: util.GetLocalClusterArgs(),
			Service: util.GetLocalServerArgs(),
			Index:   util.GetLocalArgsIndex(),
		}
		if err = consul.WatchConfigByService(serviceInfo, static.NewStatic()); err != nil {
			klog.Panic(err)
		}
	} else {
		if err = consul.WatchConfig(static.NewStatic()); err != nil {
			klog.Panic(err)
		}
		if serviceInfo, err = consul.GetServiceInfoByPath(); err != nil {
			klog.Panic(err)
		}
	}
}

func WhoAmI() *consul.ServiceInfo {
	return serviceInfo
}

func IsMasterIndex() bool {
	return serviceInfo.Index == MasterIndex
}
