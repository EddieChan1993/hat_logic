package api

import (
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/core"
	"hat_logic/core/cst"
	"hat_logic/rpc"
	"hat_logic/tool/bi"
	"hat_logic/tool/consul/static"
)

type GoGateApi struct{}

const gateLog = "[GATE]"

func (this_ *GoGateApi) Push(req *pbgo.PushReq) (*pbgo.Empty, error) {
	rsp := &pbgo.Empty{}
	core.Push(req)
	return rsp, nil
}

func (this_ *GoGateApi) CreateRole(req *pbgo.CreateRoleReq) (*pbgo.CreateRoleRsp, error) {
	rsp := &pbgo.CreateRoleRsp{}
	gdbRsp, err := rpc.CreateRole(req)
	if err != nil {
		klog.Error(err)
		rsp.ErrNo = 1
		rsp.ErrMsg = err.Error()
		return rsp, err
	}
	bi.PushLogSingle(bi.Register, bi.RegisterInfo(req, rsp))
	rsp.RoleId = gdbRsp.RoleId
	rsp.Name = gdbRsp.Name
	rsp.Ctime = gdbRsp.Ctime
	return rsp, nil
}

func (this_ *GoGateApi) Launch(req *pbgo.LaunchReq) (*pbgo.LaunchRes, error) {
	klog.Infof("%s角色:%d, session:%d 登录", gateLog, req.RoleId, req.Session)
	err := core.Login(req)
	if err != nil {
		klog.Warnf("角色:%d, session:%d 登录角色失败 %v", req.RoleId, req.Session, err)
		return &pbgo.LaunchRes{
			ErrNo:  cst.LogicErr,
			ErrMsg: err.Error(),
		}, nil
	}
	return &pbgo.LaunchRes{}, nil

}

// WaitReconnect 下线，断线重连，返回期待下次重连的时间
func (this_ *GoGateApi) WaitReconnect(req *pbgo.WaitReconnectReq) (*pbgo.WaitReconnectRsp, error) {
	klog.Infof("%s角色:%d, session:%d 网络连接断开", gateLog, req.RoleId, req.Session)
	rsp := &pbgo.WaitReconnectRsp{
		WaitTime: static.StaticWaitReconnectWaitTime(),
	}
	err := core.WaitReconnect(req)
	if err != nil {
		klog.Warnf("角色:%d, session:%d 网络连接断开失败:%v", req.RoleId, req.Session, err)
		return &pbgo.WaitReconnectRsp{
			ErrNo:    cst.LogicErr,
			WaitTime: static.StaticWaitReconnectWaitTime(),
		}, nil
	}
	return rsp, nil
}

func (this_ *GoGateApi) Reconnect(req *pbgo.ReconnectReq) (*pbgo.ReconnectRes, error) {
	klog.Infof("%s角色:%d, session:%d 重连", gateLog, req.RoleId, req.Session)
	err := core.Reconnect(req)
	if err != nil {
		klog.Warnf("角色:%d, session:%d 重连失败", req.RoleId, req.Session, err)
		return &pbgo.ReconnectRes{
			ErrNo:  cst.LogicErr,
			ErrMsg: err.Error(),
		}, nil
	}
	return &pbgo.ReconnectRes{}, nil
}
