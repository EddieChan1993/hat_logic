package test

import (
	"hat_logic/core/cst"
	"hat_logic/mod/mcst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqSetBuildWorker(t *testing.T) {
	RunTestKeepAlive(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqSetBuildWorker{
			Nums:       1,
			BuildingId: mcst.BuildIdFarm,
		}
		Push(pbreq.GrpId18, pbreq.CmdReqSetBuildWorkerId, roleId, p)
		p = &pbreq.ReqSetBuildWorker{
			Nums:       1,
			BuildingId: mcst.BuildIdMine,
		}
	})
}
