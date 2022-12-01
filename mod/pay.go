package mod

/**
支付
*/
import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gpay/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/cast"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
	"hat_logic/tool/bi"
	"hat_logic/util"
)

const priceRate = 1000000

type Pay struct {
	pbreq.TModApi16
	modPays map[mcst.RechargeActTyp]IModPay
}

//==================== api ====================//

//IModPay 具有支付功能的模块特征
type IModPay interface {
	PreCheck(id mcst.RechargeId) error                                   //支付前检测
	PayRewards(id mcst.RechargeId) (Rewards []*config.Reward, err error) //支付后发奖
}

type IModPayFun interface {
	RegisterPayMod(id mcst.RechargeActTyp, mod IModPay) //注册支付Mod
}

func PayMod(gameCtx *pbreq.GameCtx) IModPayFun {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&Pay{}))
	if had {
		return ins.(IModPayFun)
	}
	return nil
}

//==================== req ====================//

func (p *Pay) ReqPrePay(req *pbreq.ReqPrePay) *pbreq.RspPrePay {
	klog.Infof("%s ReqPrePay start RechargeId %d", p.Log(), req.RechargeId)
	res := &pbreq.RspPrePay{}
	rechargeData := config.RechargeData.Get(req.RechargeId)
	modPay, had := p.modPays[rechargeData.ActType]
	if !had {
		klog.Warnf("%s rechargeId %d actType %d 没有找到注册的modpay", p.Log(), req.RechargeId, rechargeData.ActType)
		res.Status = 1
		return res
	}
	err := modPay.PreCheck(req.RechargeId)
	if err != nil {
		res.Status = 2
		klog.Errorf("%s rechargeId %d %v", p.Log(), req.RechargeId, err)
		return res
	}
	account := RoleMod(p.GameCtx).RoleAccount()
	rpcResp, err := rpc.PrePayHttp(&pbgo.PrePayReq{
		Account:  cast.ToString(account),
		Uid:      cast.ToString(p.GetRoleId()),
		Sid:      p.GetSid(),
		Package:  req.App.Package,
		PayType:  req.App.PayType,
		StoreId:  rechargeData.PayId,
		Desc:     "",
		PayPrice: rechargeData.PriceCn + "00", //分为单位
		DeviceId: req.App.DeviceId,
		HawkeyeInfo: &pbgo.HawkeyeInfo{
			GameCd:   util.GameCd,
			SmId:     "",
			Ip:       "",
			ServerId: cast.ToString(p.GetSid()),
		},
		ProductName: rechargeData.Name,
	})
	if err != nil {
		res.Status = 3
		klog.Errorf("%s rechargeId %d rpc %v", p.Log(), req.RechargeId, err)
		return res
	}
	res.OrderId = rpcResp.OrderId
	klog.Infof("%s ReqPrePay end  RechargeId %d PriceCn %s PayId %s", p.Log(), req.RechargeId, rechargeData.PriceCn, rechargeData.PayId)
	return res
}

func (p *Pay) ReqRewardPay(req *pbreq.ReqRewardPay) *pbreq.RspRewardPay {
	klog.Infof("%s ReqRewardPay start OrderId %s StoreId %s Price %d", p.Log(), req.OrderId, req.StoreId, req.CurrencyInfo.Price/priceRate)
	res := &pbreq.RspRewardPay{}
	rewardsCfg, err := p.payReward(req.StoreId, req.OrderId, req.App.Package, req.CurrencyInfo)
	if err != nil {
		klog.Errorf("%s payReward req.OrderId %s req.StoreId %s, %v", p.Log(), req.OrderId, req.StoreId, err)
		res.Status = 1
		return res
	}
	rewardsRes, err := GoodsMod(p.GameCtx).Rewards(rewardsCfg)
	if err != nil {
		klog.Errorf("%s Rewards req.OrderId %s req.StoreId %s, %v", p.Log(), req.OrderId, req.StoreId, err)
		res.Status = 4
		return res
	}
	res.Rewards = rewardsRes
	klog.Infof("%s ReqRewardPay end OrderId %s StoreId %s Price %d", p.Log(), req.OrderId, req.StoreId, req.CurrencyInfo.Price/priceRate)
	return res
}

func (p *Pay) ReqLoginRewardPay(req *pbreq.ReqLoginRewardPay) *pbreq.RspLoginRewardPay {
	res := &pbreq.RspLoginRewardPay{
		StoreId: req.StoreId,
	}
	rewardsCfg, err := p.payReward(req.StoreId, req.OrderId, req.App.Package, req.CurrencyInfo)
	if err != nil {
		klog.Error(p.Log(), err)
		res.Status = 1
		return res
	}
	err = rpc.SendMail(&rpc.MailContent{
		RoleId:  p.GetRoleId(),
		MailId:  mcst.MailIdSupplementOrder,
		Rewards: rewardsCfg,
		Args:    "",
	})
	if err != nil {
		klog.Errorf("%s req.OrderId %s req.StoreId %s, %v", p.Log(), req.OrderId, req.StoreId, err)
		res.Status = 2
		return res
	}
	return res
}

//==================== tool ====================//

//payReward 支付处理
func (p *Pay) payReward(storeId, orderId, pkg string, currencyInfo *pbreq.CurrencyInfo) ([]*config.Reward, error) {
	account := RoleMod(p.GameCtx).RoleAccount()
	rpcResp, err := rpc.PayRewardHttp(&pbgo.RewardReq{
		Sid:     p.GetSid(),
		StoreId: storeId,
		OrderId: orderId,
		Uid:     cast.ToString(p.GetRoleId()),
		Package: pkg,
		Account: cast.ToString(account),
	})
	if err != nil {
		return nil, fmt.Errorf("rpc.PayReward storeId %s orderId %s err %v", storeId, orderId, err)
	}
	rechargeData := config.RechargeData.GetByPayId(rpcResp.Order.StoreId)
	if len(rechargeData) == 0 {
		return nil, fmt.Errorf("RechargeData.GetByPayId storeId %s orderId %s nil", storeId, orderId)
	}
	rechargeActTyp := rechargeData[0].ActType
	modPay, had := p.modPays[rechargeActTyp]
	if !had {
		return nil, fmt.Errorf("modPays storeId %s orderId %s rechargeActTyp %d none", storeId, orderId, rechargeActTyp)
	}
	rechargeId := rechargeData[0].Id
	rewardsCfg, err := modPay.PayRewards(rechargeId)
	if err != nil {
		return nil, fmt.Errorf("modPay.PayRewards storeId %s orderId %s rechargeId %d err %v", storeId, orderId, rechargeId, err)
	}
	p.uploadBi(rpcResp, currencyInfo)
	return rewardsCfg, nil
}

//uploadBi bi上报
func (p *Pay) uploadBi(rpcResp *pbgo.RewardRsp, currencyInfo *pbreq.CurrencyInfo) {
	biMod := BiMod(p.GameCtx)
	rechargeData := config.RechargeData.GetByPayId(rpcResp.Order.StoreId)
	biMod.Upload(bi.Pay, bi.EventValJson{"order_no": rpcResp.Order.ThirdOrder,
		"pay_state":      0,
		"pay_type":       rpcResp.Order.SandBox,
		"pay_price":      rechargeData[0].Price,
		"pay_price_rmb":  rechargeData[0].PriceCn,
		"pay_channel":    999,
		"store_id":       rechargeData[0].Id,
		"currency_pay":   currencyInfo.Type,
		"currency_price": currencyInfo.Price / priceRate,
	})
	afRevenue := cast.ToString(rechargeData[0].Price)
	afCurrency := "USD"
	biMod.AppsflyerUpload(bi.AfPurchase, bi.EventValJson{
		"af_revenue":  afRevenue,
		"af_currency": afCurrency,
	})
}

//RegisterPayMod 注册pay mod
func (p *Pay) RegisterPayMod(reActTyp mcst.RechargeActTyp, mod IModPay) {
	_, had := p.modPays[reActTyp]
	if had {
		klog.Panic(fmt.Sprintf("%s 重复注册支付mod rechargeActTyp=%d", p.Log(), reActTyp))
		return
	}
	p.modPays[reActTyp] = mod
}

//========================mod base========================//

func (p *Pay) InitMod(gameCtx *pbreq.GameCtx) {
	p.GameCtx = gameCtx
	p.modPays = make(map[mcst.DefaultInt]IModPay)
}

func (p *Pay) RspSync(rsp *pbreq.RspSync) {
}

func NewPay() *Pay {
	res := &Pay{}
	res.InitModApi(res)
	return res
}
