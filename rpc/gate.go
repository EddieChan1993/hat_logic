package rpc

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gate"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"google.golang.org/protobuf/proto"
	"hat_logic/core/cst"
	"hat_logic/rpc/rcst"
)

func SendPush(sessionId int64, grp, cmd cst.GrpIdTyp, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		klog.Error(err)
		return
	}
	req := &pbgo.SendBySessionReq{Session: sessionId, Grp: grp, Cmd: cmd, Bin: data}
	rsp, err := SendBySession(req)
	if err != nil {
		klog.Error(err)
		return
	}
	if rsp.ErrNo != 0 {
		klog.Error(rsp.ErrMsg)
		return
	}
}
func SendBySession(req *pbgo.SendBySessionReq) (*pbgo.CommonRsp, error) {
	return gate.SendBySession(req, kite.Cast)
}

// KickBySession 踢玩家下线
func KickBySession(session int64, reason string, args ...kite.Option) (*pbgo.CommonRsp, error) {
	req := &pbgo.KickBySessionReq{Session: session, Reason: reason}
	return gate.KickBySession(req, args...)
}

// GetServerList 获取服务器列表
func GetServerList(req *pbgo.GetServerListReq) (*pbgo.GetServerListRsp, error) {
	return pbgo.Gate.GetServerList(*rcst.RpcWhoAmIdGate(), req)
}

//GetNodesServer 获取当前节点下的服务列表
func GetNodesServer() (servers []int32, err error) {
	_, _, _, logicId := kite.GetWhoAmI()
	req := &pbgo.GetNodeServersReq{
		LogicId: int32(logicId),
	}
	res, err := pbgo.Gate.GetNodeServers(*rcst.RpcWhoAmIdGate(), req)
	if err != nil {
		return nil, err
	}
	if res.ErrNo != 0 {
		return nil, fmt.Errorf(res.ErrMsg)
	}
	return res.Servers, nil
}

// CreateServer 创建服务器
func CreateServer(req *pbgo.RegServerReq) (*pbgo.RegServerRsp, error) {
	return pbgo.Gate.RegServer(*rcst.RpcWhoAmIdGate(), req)
}

// UpdateServer 修改服务器信息
func UpdateServer(req *pbgo.UpdateServerReq) (*pbgo.UpdateServerRsp, error) {
	return pbgo.Gate.UpdateServer(*rcst.RpcWhoAmIdGate(), req)
}

// DelServer 删除服务器
func DelServer(req *pbgo.DeleteServerReq) (*pbgo.DeleteServerRsp, error) {
	return pbgo.Gate.DeleteServer(*rcst.RpcWhoAmIdGate(), req)
}
