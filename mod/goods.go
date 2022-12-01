package mod

/**
各类物品消耗
*/
import (
	"fmt"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/pbgo/pbreq"
	"hat_logic/util"
	rand2 "math/rand"
)

type Goods struct {
	*pbreq.GameCtx
}

//==================== api ====================//

type IModGoodsFun interface {
	Rewards(res []*config.Reward, multi ...mcst.DefaultInt) (*pbreq.Rewards, error)                       //发送物品
	Consume(res []*config.Consume, multi ...mcst.DefaultInt) error                                        //消耗物品
	TransRandReward(randPool []*config.RandomReward, randTimes mcst.DefaultInt) ([]*config.Reward, error) //randPool在一个奖池随，并不发奖
	TransRandRewardPer(randPool []*config.RandomRewardPer) ([]*config.Reward, error)                      //randPool每个元素单独一个奖池，并不发奖
	TransMultiReward(res []*config.Reward, multiNum mcst.DefaultInt) []*config.Reward                     //处理为N倍奖励结构，并不发奖
	CopyRewards(rewards []*config.Reward) []*config.Reward
}

func GoodsMod(gameCtx *pbreq.GameCtx) IModGoodsFun {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&Goods{}))
	if had {
		return ins.(IModGoodsFun)
	}
	return nil
}

//==================== tool ====================//

//Consume 消耗物品
func (g *Goods) Consume(res []*config.Consume, multi ...mcst.DefaultInt) error {
	multiNum := mcst.DefaultInt(1)
	if len(multi) > 1 {
		return fmt.Errorf("mult 只支持一个元素")
	} else if len(multi) != 0 {
		multiNum = multi[0]
	}
	goodsTypMap := g.transConsumeTyp(res, multiNum)
	for typ, numMap := range goodsTypMap {
		switch typ {
		case mcst.GoodTypItem:
			err := BagMod(g.GameCtx).SubItem(numMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//Rewards 添加物品
func (g *Goods) Rewards(res []*config.Reward, multi ...mcst.DefaultInt) (*pbreq.Rewards, error) {
	if len(res) == 0 {
		return nil, nil
	}
	rewards := &pbreq.Rewards{
		Items:  nil,
		Equips: nil,
	}
	multiNum := mcst.DefaultInt(1)
	if len(multi) > 1 {
		return nil, fmt.Errorf("mult 只支持一个元素")
	} else if len(multi) != 0 {
		multiNum = multi[0]
	}
	goodsTypMap := g.transRewardsTyp(res, multiNum)
	for typ, numMap := range goodsTypMap {
		switch typ {
		case mcst.GoodTypItem:
			//添加道具
			itemRewards, err := BagMod(g.GameCtx).AddItem(numMap)
			if err != nil {
				return nil, err
			}
			//比如英雄转为英雄碎片，rewardsItem已经存在数据，需要和本次道具合并
			rewards.Items = util.MapMergeSum(rewards.Items, itemRewards)
		case mcst.GoodTypEquip:
			//添加装备
			equipRewards, err := BagMod(g.GameCtx).AddEquip(numMap)
			if err != nil {
				return nil, err
			}
			rewards.Equips = equipRewards
		}
	}
	return rewards, nil
}

//TransRandReward randPool在一个奖池随，并不发奖
//randPool 随机池
//randTimes 随机次数
func (g *Goods) TransRandReward(randPool []*config.RandomReward, randTimes mcst.DefaultInt) ([]*config.Reward, error) {
	res := make([]*config.Reward, 0, randTimes)
	maxRand := mcst.DefaultInt(0)
	pool := make([]mcst.DefaultInt, len(randPool))
	for i, p := range randPool {
		maxRand += p.Ratio
		pool[i] = maxRand
	}
	if randTimes <= mcst.RandRewardsForLimit {
		//在循环次数范围内
		res = g.randRewardsPer(randTimes, maxRand, pool, randPool)
	} else {
		return nil, fmt.Errorf("随机次数太大，不方便计算")
	}
	return res, nil
}

//TransRandRewardPer randPool每个元素单独一个奖池，并不发奖
func (g *Goods) TransRandRewardPer(randPool []*config.RandomRewardPer) ([]*config.Reward, error) {
	res := make([]*config.Reward, 0)
	for _, pReward := range randPool {
		rand := util.RandInt32(util.RandTotal)
		if rand > pReward.Ratio {
			continue
		}
		tmp := &config.Reward{
			Type:  pReward.Type,
			Id:    pReward.Id,
			Count: mcst.Int64Typ(pReward.Num),
		}
		res = append(res, tmp)
	}
	return res, nil
}

//randRewardsPer 每次循环都随机一次
//randTimes 随机次数
//maxRand 最大随机范围
//pool奖池
//randPool 随机配置奖池
func (g *Goods) randRewardsPer(randTimes, maxRand mcst.DefaultInt, pool []mcst.DefaultInt, randPool []*config.RandomReward) (res []*config.Reward) {
	for i := mcst.DefaultInt(0); i < randTimes; i++ {
		rand := rand2.Int31n(maxRand)
		addReward := &config.Reward{}
		for index, p := range pool {
			addReward = &config.Reward{
				Type:  randPool[index].Type,
				Id:    randPool[index].Id,
				Count: mcst.BagNum(randPool[index].Num),
			}
			if p >= rand {
				break
			}
		}
		res = append(res, addReward)
	}
	return res
}

func (g *Goods) transConsumeTyp(res []*config.Consume, multiNum mcst.DefaultInt) map[mcst.GoodTyp]mcst.GoodIdNumMap {
	data := make(map[mcst.GoodTyp]mcst.GoodIdNumMap)
	for _, re := range res {
		if old, had := data[re.Type]; had {
			old[re.Id] = old[re.Id] + mcst.BagNum(re.Count*multiNum)
			data[re.Type] = old
		} else {
			data[re.Type] = mcst.GoodIdNumMap{
				re.Id: mcst.BagNum(re.Count * multiNum),
			}
		}
	}
	return data
}

func (g *Goods) transRewardsTyp(res []*config.Reward, multiNum mcst.DefaultInt) map[mcst.GoodTyp]mcst.GoodIdNumMap {
	data := make(map[mcst.GoodTyp]mcst.GoodIdNumMap)
	for _, re := range res {
		if re == nil {
			continue
		}
		if old, had := data[re.Type]; had {
			old[re.Id] = old[re.Id] + re.Count*mcst.BagNum(multiNum)
			data[re.Type] = old
		} else {
			data[re.Type] = mcst.GoodIdNumMap{
				re.Id: re.Count * mcst.BagNum(multiNum),
			}
		}
	}
	return data
}

// TransMultiReward 奖励配置倍数转换
func (g *Goods) TransMultiReward(res []*config.Reward, multiNum mcst.DefaultInt) []*config.Reward {
	data := make([]*config.Reward, 0)
	for _, re := range res {
		if re == nil {
			continue
		}
		tmp := &config.Reward{
			Type:  re.Type,
			Id:    re.Id,
			Count: re.Count * mcst.BagNum(multiNum),
		}
		data = append(data, tmp)
	}
	return data
}

//CopyRewards 复制rewards，防止改到配置本身
func (g *Goods) CopyRewards(rewards []*config.Reward) []*config.Reward {
	rewardsRes := make([]*config.Reward, len(rewards))
	for i, reward := range rewards {
		rewardsRes[i] = &config.Reward{
			Type:  reward.Type,
			Id:    reward.Id,
			Count: reward.Count,
		}
	}
	return rewardsRes
}

//==================== base ====================//

func (g *Goods) InitMod(gameCtx *pbreq.GameCtx) {
	g.GameCtx = gameCtx
}

func (g *Goods) RspSync(rsp *pbreq.RspSync) {
	return
}

func NewGoods() *Goods {
	return &Goods{}
}
