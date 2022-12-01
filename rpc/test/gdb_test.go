package test

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"math/rand"
	"hat_logic/rpc"
	"strconv"
	"testing"
)

func init() {
	initServer()
}
func TestCreateRole(t *testing.T) {
	for i := 0; i < 1; i++ {
		gdbRsp, err := rpc.CreateRole(&pbgo.CreateRoleReq{
			Name:    "aaa2" + strconv.Itoa(i),
			Logo:    0,
			Sid:     2,
			Account: rand.Int63n(1000000000),
			Client:  nil,
		})
		fmt.Println(gdbRsp, err)
	}

}
