package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

//同步请求
func TestGm(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqGm{
			Type: pbreq.GMType_TestGachaShop,
			Data: "101:200000",
		}
		Push(pbreq.GrpId12, pbreq.CmdReqGmId, roleId, p)
	})
}
