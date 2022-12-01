package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestGetAttr(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqStartBattleEvent{
			Init: &pbreq.ReqInitBattle{
				Character: nil,
				Heroes:    nil,
				Level: &pbreq.BattleLevel{
					Type: pbreq.BattleType_BT_Tower,
					Id:   0,
				},
				Seed:    0,
				Inherit: nil,
			},
		}
		Push(pbreq.GrpId31, pbreq.CmdReqStartBattleEventId, roleId, p)
	})
}
