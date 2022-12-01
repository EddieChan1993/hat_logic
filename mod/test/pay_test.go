package test

import (
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"testing"
)

//支付购买
func TestPay(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		//rechargeId := mcst.DefaultInt(5)
		//p := &pbreq.ReqPrePay{
		//	RechargeId: rechargeId,
		//	App: &pbreq.AppInfo{
		//		Plat:     "",
		//		Package:  "",
		//		DeviceId: "",
		//		PayType:  "",
		//	},
		//}
		//Push(pbreq.GrpId24, pbreq.CmdReqPrePayId, roleId, p)
		p2 := &pbreq.ReqRewardPay{
			OrderId: "",
			StoreId: "",
			App: &pbreq.AppInfo{
				Plat:     "",
				Package:  "",
				DeviceId: "",
				PayType:  "",
			},
		}
		Push(pbreq.GrpId24, pbreq.CmdReqRewardPayId, roleId, p2)
	})
}

//月卡领取
func TestReqMonthCardDraw(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqMonthCardDraw{
			CardTyp: 2,
		}
		Push(pbreq.GrpId27, pbreq.CmdReqMonthCardDrawId, roleId, p)
	})
}

//勇士令领取
func TestReqWarriorOrderDraw(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqWarriorOrderDraw{
			GoodId: 2,
		}
		Push(pbreq.GrpId25, pbreq.CmdReqWarriorOrderDrawId, roleId, p)
	})
}

//首充
func TestReqInitialChargeDraw(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqInitialChargeDraw{
			Day: 1,
		}
		Push(pbreq.GrpId26, pbreq.CmdReqInitialChargeDrawId, roleId, p)
	})
}

//商店
func TestReqShopBuy(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqShopBuy{
			GoodId: 20005,
		}
		Push(pbreq.GrpId29, pbreq.CmdReqShopBuyId, roleId, p)
	})
}
