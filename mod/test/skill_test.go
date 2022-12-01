package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqUseLeaderSkill(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUseLeaderSkill{
			SkillId: 999002,
		}
		Push(pbreq.GrpId21, pbreq.CmdReqUseLeaderSkillId, roleId, p)
		p = &pbreq.ReqUseLeaderSkill{
			SkillId: 999003,
		}
		Push(pbreq.GrpId21, pbreq.CmdReqUseLeaderSkillId, roleId, p)
	})
}
