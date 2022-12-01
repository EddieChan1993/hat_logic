package test

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"math/rand"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
	"strconv"
	"testing"
	"time"
)

//同步请求
func TestReqSync(t *testing.T) {
	RunTestKeepAlive(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqSync{}
		Push(pbreq.GrpId10, pbreq.CmdReqSyncId, roleId, p)
	})
}

func TestLoginMore(t *testing.T) {
	kite.WhoAmI(cluster, server, index)
	for i := 0; i < 10000; i++ {
		time.Sleep(50 * time.Millisecond)
		gdbRsp, err := rpc.CreateRole(&pbgo.CreateRoleReq{
			Name:    "aaa2" + strconv.Itoa(i),
			Logo:    0,
			Sid:     5,
			Account: rand.Int63(),
			Client:  nil,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}

		RunTest(cst.RoleId(gdbRsp.RoleId), func(roleId cst.RoleId) {
			p := &pbreq.ReqSync{}
			fmt.Println("create ", gdbRsp.RoleId)
			Push(pbreq.GrpId10, pbreq.CmdReqSyncId, roleId, p)
		})
	}
}

func TestReqChangeAvatar(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqChangeAvatar{
			AvatarId: 1,
		}
		Push(pbreq.GrpId30, pbreq.CmdReqChangeAvatarId, roleId, p)
	})
}

func TestReqChangeName(t *testing.T) {
	RunTest(roleId, func(roleId cst.RoleId) {
		p := &pbreq.ReqChangeName{
			Name: "打开快发发发递",
		}
		fmt.Println(len([]rune(p.Name)))
		Push(pbreq.GrpId30, pbreq.CmdReqChangeNameId, roleId, p)
	})
}
