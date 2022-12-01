package test

import (
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"github.com/golang/protobuf/proto"
	"log"
	"hat_logic/core/cst"
)

const (
	roleId    = 75889416
	sessionId = 1000
)

const (
	dc      = "dc1"
	cluster = "re_dev"
	server  = "hat_logic"
	index   = 5
)

type defaultFn func(roleId cst.RoleId)

var bin []byte
var err error

//RunTest 登陆，发消息，断开
func RunTest(roleId cst.RoleId, fn defaultFn) {
	klog.ResetToDevelopment()
	launch(roleId)
	fn(roleId)
	waitReconnect()
}

//RunTestKeepAlive 登陆，发消息
func RunTestKeepAlive(roleId cst.RoleId, fn defaultFn) {
	klog.ResetToDevelopment()
	launch(roleId)
	fn(roleId)
}

func launch(roleId cst.RoleId) {
	_, err := pbgo.Logic.Launch(
		ServiceIndex(),
		&pbgo.LaunchReq{
			Session: sessionId,
			RoleId:  roleId,
			Sid:     index,
			Client: &pbgo.LogicClientInfo{
				Adid:        "",
				Idfv:        "",
				Dhid:        "",
				Imei:        "",
				AndroidId:   "",
				AppsflyerId: "",
				DeviceToken: "",
				MacAddress:  "",
				DeviceModel: "",
				DviceName:   "",
				OsVersion:   "",
				Language:    "",
				NetworkType: "",
				Reserved_2:  "",
				BundleId:    "",
				Country:     "",
				AppVersion:  "",
				Platform:    "",
				Channel:     "",
				Att:         "",
				SubPackage:  "",
				Ip:          "",
				SessionId:   "",
				Openid:      "",
				Openkey:     "",
			},
		},
	)
	if err != nil {
		klog.Warn(err)
	}
}

func waitReconnect() {
	pbgo.Logic.WaitReconnect(
		ServiceIndex(),
		&pbgo.WaitReconnectReq{
			Session: sessionId,
			RoleId:  roleId,
		},
	)
}

func Push(grpId, cmdId cst.GrpIdTyp, roleId cst.RoleId, p proto.Message) {
	bin, err = proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	_, err := pbgo.Logic.Push(
		ServiceIndex(),
		&pbgo.PushReq{
			Grp:     grpId,
			Cmd:     cmdId,
			Bin:     bin,
			Session: sessionId,
			RoleId:  roleId,
		},
	)
	if err != nil {
		klog.Warn(err)
	}
}

func ServiceIndex() kite.Destination {
	return kite.Destination{
		DC:      dc,
		Cluster: cluster,
		Service: server,
		Index:   index,
	}
}
