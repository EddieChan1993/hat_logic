package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqGachaShop(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqGachaShop{
			GachaId: 101,
			Nums:    10,
		}
		Push(pbreq.GrpId28, pbreq.CmdReqGachaShopId, roleId, p)
	})
}

func TestReqGachaBox(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqGachaShop{
			GachaId: 301,
		}
		Push(pbreq.GrpId28, pbreq.CmdReqGachaBoxId, roleId, p)
	})
}
