package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestBattleLogMainLevel(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqMainLevelSettle{
			LevelMonsterId: int32(1101003),
		}
		Push(pbreq.GrpId16, pbreq.CmdReqMainLevelSettleId, roleId, p)
	})
}
