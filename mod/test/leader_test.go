package test

import (
	"fmt"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"hat_logic/util"
	"testing"
	"time"
)

func TestWearLeader(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqWearEquip{
			EquipIds: []int32{1, 2},
		}
		Push(pbreq.GrpId13, pbreq.CmdReqWearEquipId, roleId, p)
	})
}

func TestUpLeader(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqUpLvLeader{
			Num: 32,
		}
		Push(pbreq.GrpId13, pbreq.CmdReqUpLvLeaderId, roleId, p)
	})
}

func TestSetSkill(t *testing.T) {
	now := util.TimeNowUnixMilli()
	time.Sleep(10 * time.Second)
	last := util.TimeNowUnixMilli()
	fmt.Println(last - now)

	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqSetSkill{
			SkillId: 11001,
		}
		Push(pbreq.GrpId13, pbreq.CmdReqSetSkillId, roleId, p)
	})
}
