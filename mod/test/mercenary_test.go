package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqMercenarySelectTask(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqMercenarySelectTask{
			TaskId: 3,
		}
		Push(pbreq.GrpId20, pbreq.CmdReqMercenarySelectTaskId, roleId, p)
	})
}

func TestReqMercenaryRefresh(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqMercenaryRefresh{}
		Push(pbreq.GrpId20, pbreq.CmdReqMercenaryRefreshId, roleId, p)
	})
}

func TestReqMercenaryDrawTask(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqMercenaryDrawTask{}
		Push(pbreq.GrpId20, pbreq.CmdReqMercenaryDrawTaskId, roleId, p)
	})
}

func TestReqBuyMercenarySlot(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqBuyMercenarySlot{}
		Push(pbreq.GrpId20, pbreq.CmdReqBuyMercenarySlotId, roleId, p)
	})
}
