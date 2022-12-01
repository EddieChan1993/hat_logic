package test

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"hat_logic/core"
	"hat_logic/rpc"
	"hat_logic/tool"
	"hat_logic/util"
	"testing"
	"time"
)

func TestCreateServer(t *testing.T) {
	initServer()
	res, err := rpc.UpdateServer(&pbgo.UpdateServerReq{
		Info: &pbgo.ServerInfo{
			Id:        1,
			Recommend: 0,
			Name:      "1",
			Ctime:     time.Now().Unix(),
			Logicid:   1,
			Bfid:      1,
			Status:    1,
			Interval:  false,
		},
	})
	fmt.Println(res, err)
	res2, err := rpc.GetServerList(&pbgo.GetServerListReq{})
	fmt.Println(res2, err)
}

func TestGetServer(t *testing.T) {
	initServer()
	res, err := rpc.GetNodesServer()
	fmt.Println(res, err)
}

func initServer() {
	util.InitDebugForTest()
	tool.InitTool()
	core.InitCore()
	kite.WhoAmI(util.GetLocalClusterArgs(), util.GetLocalServerArgs(), util.GetLocalArgsIndex())
}
