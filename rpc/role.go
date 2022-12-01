package rpc

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/gclient/gdb"
	gdbpb "git.dhgames.cn/svr_comm/gclient/gdb/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"hat_logic/rpc/rcst"
	"time"
)

func CreateRole(req *pbgo.CreateRoleReq) (*gdbpb.CreateRoleRsp, error) {
	reqDb := &gdbpb.CreateRoleReq{
		Digest: &gdbpb.RoleDigest{
			Account:         int32(req.Account),
			Name:            req.Name,
			Logo:            req.Logo,
			Sid:             req.Sid,
			LastOfflineTime: 0,
			LastOnlineTime:  0,
			Ctime:           time.Now().Unix(),
			Exp:             0,
			Vexp:            0,
			Lv:              0,
			Language:        "",
			Status:          0,
			GroupName:       "",
			Power:           0,
			HeadFrame:       0,
			HeadTitle:       0,
			UnregisterTime:  0,
		},
	}
	resp, err := gdb.CreateRole(reqDb)
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf(resp.ErrMsg)
	}
	return resp, nil
}
func RoleOnline(roleId int64) (*gdbpb.RoleOnlineRsp, error) {
	return gdb.RoleOnline(&gdbpb.RoleOnlineReq{RoleId: TransRoleId(roleId)})
}

func RoleOffline(roleId int64) error {
	resp, err := gdb.RoleOffline(&gdbpb.RoleOfflineReq{RoleId: TransRoleId(roleId)})
	if err != nil {
		return err
	}
	if resp.ErrNo != 0 {
		return fmt.Errorf(resp.ErrMsg)
	}
	return nil
}

//LoadLogic 加载玩家数据
func LoadLogic(roleId int64) (response *gdbpb.LoadLogicRsp, err error) {
	return gdbpb.Logic.LoadLogic(*rcst.RpcWhoAmIdGDB(), &gdbpb.LoadLogicReq{RoleId: TransRoleId(roleId)})
}

//SaveLogic 存储玩家数据
func SaveLogic(req *gdbpb.SaveLogicReq, opts ...kite.Option) (response *gdbpb.CommonRsp, err error) {
	return gdbpb.Logic.SaveLogic(*rcst.RpcWhoAmIdGDB(), req, opts...)
}

func TransRoleId(roleId int64) int32 {
	return int32(roleId)
}

// RoleDigest 获取角色信息
func RoleDigest(roleId int64, opts ...kite.Option) (*gdbpb.RoleDigestRsp, error) {
	req := &gdbpb.RoleDigestReq{RoleId: int32(roleId)}
	return gdbpb.Digest.RoleDigest(*rcst.RpcWhoAmIdGDB(), req, opts...)
}

//RoleDigestMultiMap 批量获取角色信息
func RoleDigestMultiMap(roleIds []int32, opts ...kite.Option) (response map[int32]*gdbpb.RoleDigest, err error) {
	req := &gdbpb.RolesDigestReq{
		RoleIds: roleIds,
	}
	res, err := gdbpb.Digest.RolesDigest(*rcst.RpcWhoAmIdGDB(), req, opts...)
	if err != nil {
		return nil, err
	}
	response = make(map[int32]*gdbpb.RoleDigest)
	for i, digest := range res.Digest {
		response[digest.RoleId] = res.Digest[i]
	}
	return response, nil
}

//RoleDigestSave 存储角色信息
func RoleDigestSave(req *gdbpb.SaveDigestReq) (err error) {
	resp, err := gdbpb.Digest.SaveDigest(*rcst.RpcWhoAmIdGDB(), req)
	if err != nil {
		return err
	}
	if resp.ErrNo != 0 {
		return fmt.Errorf(resp.ErrMsg)
	}
	return nil
}
