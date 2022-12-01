package rpc

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gpay/pbgo"
	jsoniter "github.com/json-iterator/go"
	"hat_logic/rpc/rcst"
	"hat_logic/tool/consul/static"
	"hat_logic/tool/restyHttp"
)

const (
	PrePayPath = "/logic_prepay"
	RewardPath = "/reward"
)

func PrePay(req *pbgo.PrePayReq) (resp *pbgo.PrePayRsp, err error) {
	resp, err = pbgo.Payment.PrePay(rcst.RpcWhoAmIdGPay(), req)
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf("error %d", resp.ErrNo)
	}
	return resp, nil
}

func PrePayHttp(req *pbgo.PrePayReq) (resp *pbgo.PrePayRsp, err error) {
	bodyBy, err := jsoniter.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := restyHttp.Post(static.StaticGPayUrl()+PrePayPath, bodyBy)
	if err != nil {
		return nil, err
	}
	resp = &pbgo.PrePayRsp{}
	err = jsoniter.Unmarshal(res, resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf("error %d reqId %s", resp.ErrNo, resp.ReqId)
	}
	return resp, nil
}

func PayReward(req *pbgo.RewardReq) (resp *pbgo.RewardRsp, err error) {
	resp, err = pbgo.Payment.Reward(rcst.RpcWhoAmIdGPay(), req)
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf("error %d", resp.ErrNo)
	}
	return resp, nil
}

func PayRewardHttp(req *pbgo.RewardReq) (resp *pbgo.RewardRsp, err error) {
	bodyBy, err := jsoniter.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := restyHttp.Post(static.StaticGPayUrl()+RewardPath, bodyBy)
	if err != nil {
		return nil, err
	}
	resp = &pbgo.RewardRsp{}
	err = jsoniter.Unmarshal(res, resp)
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf("error %d reqId %s", resp.ErrNo, resp.ReqId)
	}
	return resp, nil
}

func PayLoginReward(req *pbgo.LoginRewardReq) (resp *pbgo.LoginRewardRsp, err error) {
	resp, err = pbgo.Payment.LoginReward(rcst.RpcWhoAmIdGPay(), req)
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf("error %d", resp.ErrNo)
	}
	return resp, nil
}
