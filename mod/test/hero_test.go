package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestUpLvHero(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUpLvHero{
			Id:  10121,
			Num: 149,
		}
		Push(pbreq.GrpId14, pbreq.CmdReqUpLvHeroId, roleId, p)
	})
}

func TestUpStageHero(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUpStageHero{
			Id: 10021,
		}
		for i := 0; i < 8; i++ {
			Push(pbreq.GrpId14, pbreq.CmdReqUpStageHeroId, roleId, p)
		}
	})
}
