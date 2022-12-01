package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestUpEquip(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUpEquip{
			EquipId: 0,
		}
		Push(pbreq.GrpId15, pbreq.CmdReqUpEquipId, roleId, p)
	})
}

func TestRecoveryItem(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqRecoveryItem{
			ItemId: 22,
			Num:    10000000,
		}
		for i := 0; i < 1000; i++ {
			Push(pbreq.GrpId15, pbreq.CmdReqRecoveryItemId, roleId, p)
		}
	})
}

func TestUseItem(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUseItem{
			ItemId: 1001,
			Num:    0,
		}
		Push(pbreq.GrpId15, pbreq.CmdReqUseItemId, roleId, p)
	})
}

func TestReqRecoveryEquip(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqRecoveryEquip{
			EquipId: []int32{1, 2},
		}
		Push(pbreq.GrpId15, pbreq.CmdReqRecoveryEquipId, roleId, p)
	})
}
