package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqInsBossBattle(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqInsBossBattle{
			Damage: 1119,
		}
		Push(pbreq.GrpId22, pbreq.CmdReqInsBossBattleId, roleId, p)
	})
}

func TestReqInsBossRank(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqInsBossRank{}
		Push(pbreq.GrpId22, pbreq.CmdReqInsBossRankId, roleId, p)
	})
}
