package core

import "hat_logic/mod"

//newModMgr 模块管理
func newModMgr() *TModMgr {
	modMgr := initModMgr()
	modMgr.RegModBase(mod.NewGoods())
	modMgr.RegCompleteMod(mod.NewRole())
	modMgr.RegModBase(mod.NewCronTasks())
	modMgr.RegReqApiMod(mod.NewGm())
	modMgr.RegReqApiMod(mod.NewMail())
	modMgr.RegReqApiMod(mod.NewPay())
	modMgr.RegModBase(mod.NewBi())

	modMgr.RegCompleteMod(mod.NewBag())
	return modMgr
}
