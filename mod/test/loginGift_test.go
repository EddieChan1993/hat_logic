package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

func TestDrawLoginGift(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqDrawLoginGift{
			GiftId: 3,
		}
		Push(pbreq.GrpId32, pbreq.CmdReqDrawLoginGiftId, roleId, p)
	})
}
