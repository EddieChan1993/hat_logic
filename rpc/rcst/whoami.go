package rcst

import (
	"git.dhgames.cn/svr_comm/kite"
	"hat_logic/tool/consul/static"
)

const (
	GateSvrName   = "gate"
	GDBSvrName    = "gdb"
	GPASvrName    = "gpay"
	BattleSvrName = "re_battle"
)

func RpcWhoAmIdGate() *kite.Destination {
	return whoAmISvrNameIndex(GateSvrName)
}

func RpcWhoAmIdGDB() *kite.Destination {
	return whoAmISvrNameIndex(GDBSvrName)
}

func RpcWhoAmIdGPay() kite.Destination {
	cluster, dc := static.StaticGPay()
	d := kite.Service(dc, cluster, GPASvrName)
	return d
}

func RpcWhoAmIBattle() *kite.Destination {
	return whoAmISvrNameIndex(BattleSvrName)
}

func whoAmISvrNameIndex(service string, index ...int) *kite.Destination {
	dc, cl, _, _ := kite.GetWhoAmI()
	des := &kite.Destination{
		DC:      dc,
		Cluster: cl,
		Service: service,
	}
	if len(index) == 0 {
		return des
	}
	des.Index = index[0]
	return des
}
