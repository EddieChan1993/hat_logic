package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestDailyDoTask(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqLookAdv{}
		Push(pbreq.GrpId33, pbreq.CmdReqLookAdvId, roleId, p)

		//p1 := &pbreq.ReqDrawDailyTask{TaskId: 0}
		//Push(pbreq.GrpId33, pbreq.CmdReqDrawDailyTaskId, roleId, p1)
	})
}
