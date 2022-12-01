package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestUpTalent(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUpTalent{
			Id: 1,
		}
		Push(pbreq.GrpId17, pbreq.CmdReqUpTalentId, roleId, p)
	})
}

func TestResetTalent(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqResetTalent{}
		Push(pbreq.GrpId17, pbreq.CmdReqResetTalentId, roleId, p)
	})
}
