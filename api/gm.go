package api

import (
	"hat_logic/core"
	"hat_logic/pbgo/pbapi"
)

type Gm struct {
}

func (g *Gm) GetRoleData(req *pbapi.ReqGetRoleData) (*pbapi.RspGetRoleData, error) {
	res := &pbapi.RspGetRoleData{}
	btData, err := core.GetRoleData(req.RoleId)
	if err != nil {
		res.Status = 1
		res.ErrMsg = err.Error()
		return res, err
	}
	res.Data = string(btData)
	return res, nil
}

func (g *Gm) SetRoleData(req *pbapi.ReqSetRoleData) (*pbapi.RspSetRoleData, error) {
	res := &pbapi.RspSetRoleData{}
	err := core.KickByRoleId(req.RoleId)
	if err != nil {
		res.Status = 1
		res.ErrMsg = err.Error()
		return res, err
	}
	err = core.SetRoleData(req.RoleId, req.Data)
	if err != nil {
		res.Status = 2
		res.ErrMsg = err.Error()
		return res, err
	}
	return res, nil
}

func (g *Gm) KickBySid(req *pbapi.ReqKickBySid) (*pbapi.RspKickBySid, error) {
	res := &pbapi.RspKickBySid{}
	core.KickBySid(req.Sid)
	return res, nil
}

func (g *Gm) KickByRoleId(req *pbapi.ReqKickByRoleId) (*pbapi.RspKickByRoleId, error) {
	res := &pbapi.RspKickByRoleId{}
	err := core.KickByRoleId(req.RoleId)
	if err != nil {
		res.Status = 1
		res.ErrMsg = err.Error()
		return res, err
	}
	return res, nil
}
