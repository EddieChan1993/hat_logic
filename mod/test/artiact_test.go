package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqRefreshArtifactSummoning(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqRefreshArtifactSummoning{}
		for i := 0; i < 10000; i++ {
			Push(pbreq.GrpId19, pbreq.CmdReqRefreshArtifactSummoningId, roleId, p)
		}
	})
}

func TestReqBuyArtifact(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqBuyArtifact{
			Index: 1,
		}
		for i := 0; i < 1000; i++ {
			Push(pbreq.GrpId19, pbreq.CmdReqRefreshArtifactSummoningId, roleId, p)
			Push(pbreq.GrpId19, pbreq.CmdReqBuyArtifactId, roleId, p)
		}
	})
}

func TestReqUpArtifact(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUpArtifact{
			ArtifactId: 3,
		}
		Push(pbreq.GrpId19, pbreq.CmdReqUpArtifactId, roleId, p)
	})
}
