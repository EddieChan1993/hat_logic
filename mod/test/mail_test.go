package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestReqMailReward(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqMailReward{
			MailIds: []int32{229},
		}
		Push(pbreq.GrpId23, pbreq.CmdReqMailRewardId, roleId, p)
	})
}
