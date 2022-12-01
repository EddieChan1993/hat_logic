package mod

/**
背包道具
*/

import (
	"fmt"
	"git.dhgames.cn/svr_comm/kite/utils/cast"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/pbgo/pbdb"
	"hat_logic/pbgo/pbreq"
	"hat_logic/tool/bi"
	"hat_logic/util"
	"math"
)

type Bag struct {
	bagMgr *pbdb.BagMgr
	pbreq.TModApi12
}

//==================== var ====================//

const defaultEquipLv = 0 //默认装备等级

//======================其他模块可调用的接口==============//

type IModBagFun interface {
	EquipInfo(id mcst.EquipId) (*pbdb.Equip, bool)                             //是否有该装备
	SubItem(items mcst.ItemIdMapNum) error                                     //消耗道具
	GetItem(id mcst.ItemId) mcst.BagNum                                        //道具数量
	AddItem(items mcst.ItemIdMapNum) (mcst.ItemIdMapNum, error)                //添加道具
	AddEquip(equips mcst.EquipIdMapNum) (map[mcst.EquipId]*pbreq.Equip, error) //添加装备
	DelEquip(eids []mcst.DefaultInt) (delEquips []*pbdb.Equip)                 //删除装备
	BiSnapshot()
}

func BagMod(gameCtx *pbreq.GameCtx) IModBagFun {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&Bag{}))
	if had {
		return ins.(IModBagFun)
	}
	return nil
}

//========================req api========================//

//ReqUseItem 道具使用
func (b *Bag) ReqUseItem(req *pbreq.ReqUseItem) *pbreq.RspUseItem {
	res := &pbreq.RspUseItem{
		ItemId: req.ItemId,
		Num:    req.Num,
	}
	itemCfg := config.ItemData.Get(req.ItemId)
	if itemCfg == nil {
		res.Status = 1 //道具配置不存在
		return res
	}
	var err error
	var rewardsCfg []*config.Reward
	switch itemCfg.ShowType {
	case mcst.ItemTypCoin:
		rewardsCfg = GoodsMod(b.GameCtx).TransMultiReward(itemCfg.Use, req.Num)
	}
	if itemCfg.RandomChest != 0 {
		//随机宝箱表
		rewardsRandCfg, err := b.useRandItem(itemCfg, req.Num)
		if err != nil {
			klog.Error(b.Log(), err)
			res.Status = 3
			return res
		}
		rewardsCfg = append(rewardsCfg, rewardsRandCfg...)
	}
	err = BagMod(b.GameCtx).SubItem(mcst.ItemIdMapNum{itemCfg.Id: mcst.BagNum(req.Num)})
	if err != nil {
		klog.Error(b.Log(), err)
		res.Status = 2
		return res
	}
	rewards, err := GoodsMod(b.GameCtx).Rewards(rewardsCfg)
	if err != nil {
		klog.Warn(b.Log(), err)
		res.Status = 4
		return res
	}
	res.Rewards = rewards
	klog.Infof("%s ReqUseItem ItemId %d Num %d", b.Log(), req.ItemId, req.Num)
	return res
}

//========================tool========================//

//装备最大等级
func (b *Bag) equipMaxLevel(eid mcst.DefaultInt) mcst.DefaultInt {
	equipCfg := config.EquipData.Get(eid)
	return mcst.DefaultInt(len(equipCfg.EquipAttrbTemp) - 1)
}

func (b *Bag) GetItem(id mcst.ItemId) mcst.BagNum {
	return b.bagMgr.Items[id]
}

//useRandItem 使用随机道具
func (b *Bag) useRandItem(itemCfg *config.ItemCfg, num mcst.DefaultInt) ([]*config.Reward, error) {
	if num > mcst.RandRewardsForLimit {
		return nil, fmt.Errorf("随机次数太大，不方便计算")
	}
	randomChestCfg := config.RandomChestData.Get(itemCfg.RandomChest)
	if randomChestCfg == nil {
		return nil, fmt.Errorf("randomChestCfg 配置不存在 id %d", itemCfg.RandomChest)
	}
	baseRewards := GoodsMod(b.GameCtx).TransMultiReward(randomChestCfg.Reward, num)
	randRewards, err := GoodsMod(b.GameCtx).TransRandReward(randomChestCfg.RandomReward, num)
	if err != nil {
		return nil, err
	}
	rewardsAll := append(baseRewards, randRewards...)
	return rewardsAll, nil
}

//EquipInfo 装备信息
func (b *Bag) EquipInfo(id mcst.EquipId) (*pbdb.Equip, bool) {
	equip, had := b.bagMgr.Equips[id]
	return equip, had
}

//AddEquip 添加装备
func (b *Bag) AddEquip(equips mcst.EquipIdMapNum) (map[mcst.EquipId]*pbreq.Equip, error) {
	equipRewards := make(map[mcst.EquipId]*pbreq.Equip)
	oldMaxEqId := b.bagMgr.MaxEqId
	tmpEquip := make(map[mcst.EquipId]*pbdb.Equip)
	logs := make([]*bi.ItemLog, 0, len(equips))
	for Id, Nums := range equips {
		for i := mcst.BagNum(0); i < Nums; i++ {
			cfgEquip := config.EquipData.Get(Id)
			if cfgEquip == nil {
				return nil, fmt.Errorf("装备不存在 id %d", Id)
			}
			oldMaxEqId++
			randExtAttrs := b.randExtAttrEquip(cfgEquip)
			tmpEquip[oldMaxEqId] = &pbdb.Equip{
				Id:         Id,
				Lv:         defaultEquipLv,
				ExtAttrIds: randExtAttrs,
			}
			equipRewards[oldMaxEqId] = &pbreq.Equip{
				Id:         Id,
				Lv:         defaultEquipLv,
				ExtAttrIds: randExtAttrs,
			}
		}
		log := &bi.ItemLog{
			ProductNum: Nums,
			ProductId:  mcst.BiPreEquip + cast.ToString(Id),
			LogType:    mcst.BiAdd,
			CreateTs:   bi.CreateTs(),
			Balance:    Nums,
		}
		logs = append(logs, log)
	}
	BiMod(b.GameCtx).AddProductLog(logs)
	b.bagMgr.MaxEqId = oldMaxEqId
	for eid := range tmpEquip {
		b.bagMgr.Equips[eid] = tmpEquip[eid]
	}
	return equipRewards, nil
}

//AddItem 添加道具
func (b *Bag) AddItem(items mcst.ItemIdMapNum) (mcst.ItemIdMapNum, error) {
	itemRewards := make(mcst.ItemIdMapNum)
	tmpItem := make(mcst.ItemIdMapNum)
	logs := make([]*bi.ItemLog, 0, len(items))
	for id, num := range items {
		ItemInfo := config.ItemData.Get(id)
		if ItemInfo == nil {
			return nil, fmt.Errorf("所加道具不存在 itemId %d", id)
		}
		if num < 0 {
			return nil, fmt.Errorf("数量不能为负")
		}
		if num == 0 {
			continue
		}
		oldNum := b.bagMgr.Items[id]
		if numRes, isOver := util.Cal64Safe(oldNum, num); isOver {
			err := fmt.Errorf("数据溢出 max %d left %d right %d", math.MaxInt64, b.bagMgr.Items[id], num)
			return nil, err
		} else {
			tmpItem[id] = numRes
			itemRewards[id] = itemRewards[id] + num
			log := &bi.ItemLog{
				ProductNum: num,
				ProductId:  mcst.BiPreItem + cast.ToString(id),
				LogType:    mcst.BiAdd,
				CreateTs:   bi.CreateTs(),
				Balance:    numRes,
			}
			logs = append(logs, log)
		}
	}
	BiMod(b.GameCtx).AddProductLog(logs)
	for id, num := range tmpItem {
		b.bagMgr.Items[id] = num
	}
	return itemRewards, nil
}

//SubItem 消耗道具
func (b *Bag) SubItem(items mcst.ItemIdMapNum) error {
	tmpItem := make(mcst.ItemIdMapNum)
	logs := make([]*bi.ItemLog, 0, len(items))
	for id, num := range items {
		oldNum := b.bagMgr.Items[id]
		if num < 0 {
			return fmt.Errorf("数量不能为负数")
		}
		if num == 0 {
			continue
		}
		if oldNum < num {
			return fmt.Errorf("item no enoght itemId %d need %d had %d", id, num, b.bagMgr.Items[id])
		}
		if numRes, isOver := util.Cal64Safe(oldNum, -num); isOver {
			err := fmt.Errorf("date overflow min %d left %d right %d", math.MinInt64, b.bagMgr.Items[id], num)
			return err
		} else {
			tmpItem[id] = numRes
			log := &bi.ItemLog{
				ProductNum: num,
				ProductId:  mcst.BiPreItem + cast.ToString(id),
				LogType:    mcst.BiSub,
				CreateTs:   bi.CreateTs(),
				Balance:    numRes,
			}
			logs = append(logs, log)
		}
	}
	BiMod(b.GameCtx).AddProductLog(logs)
	for id, num := range tmpItem {
		b.bagMgr.Items[id] = num
	}
	return nil
}

//beginnerRewards 仅用于前期测试，背包拉满
func (b *Bag) beginnerRewards() {
	items := make([]*config.Reward, 0)
	if len(b.bagMgr.Items) == 0 {
		for _, id := range config.ItemData.GetAll() {
			tmp := &config.Reward{
				Type:  mcst.GoodTypItem,
				Id:    id,
				Count: mcst.GoodsDefaultNums,
			}
			items = append(items, tmp)
		}
	}
	if len(b.bagMgr.Equips) == 0 {
		for _, cfgId := range config.EquipData.GetAll() {
			tmpCfg := config.EquipData.Get(cfgId)
			tmp := &config.Reward{
				Type:  mcst.GoodTypEquip,
				Id:    tmpCfg.Id,
				Count: 1,
			}
			items = append(items, tmp)
		}
	}
	GoodsMod(b.GameCtx).Rewards(items)
}

//DelEquip 删除装备
func (b *Bag) DelEquip(eids []mcst.DefaultInt) (delEquips []*pbdb.Equip) {
	if len(eids) == 0 {
		return delEquips
	}
	logs := make([]*bi.ItemLog, 0, len(eids))
	for _, eid := range eids {
		equip, had := b.bagMgr.Equips[eid]
		if !had {
			continue
		}
		delete(b.bagMgr.Equips, eid)
		delEquips = append(delEquips, equip)
		logs = append(logs, &bi.ItemLog{
			ProductNum: 1,
			ProductId:  mcst.BiPreEquip + cast.ToString(equip.Id),
			LogType:    mcst.BiSub,
			CreateTs:   bi.CreateTs(),
			Balance:    0,
		})
	}
	BiMod(b.GameCtx).AddProductLog(logs)
	return delEquips
}

//randExtAttrEquip 获取随机装备属性
func (b *Bag) randExtAttrEquip(cfgEquip *config.EquipCfg) []mcst.DefaultInt {
	CstEquipAttributeEntries := config.ConstData.Get(mcst.CstEquipAttributeEntries)
	randNums := CstEquipAttributeEntries.ValueList[cfgEquip.Quality-1]
	randMap := make(util.RandPoolTyp)
	for _, ran := range cfgEquip.AddProp {
		randMap[ran.Id] = ran.Ratio
	}
	return util.RandMultiNoRepeat(randMap, randNums)
}

func (b *Bag) BiSnapshot() {
	bagMap := make(bi.EventValJson)
	itemSize := 1
	for id, num := range b.bagMgr.Items {
		bagMap[cast.ToString(itemSize)] = bi.EventValJson{
			"product_id":  mcst.BiPreItem + cast.ToString(id),
			"product_num": num,
		}
		itemSize++
	}
	equipMap := make(map[int32]int64)
	for _, equip := range b.bagMgr.Equips {
		equipMap[equip.Id]++
	}
	for id, num := range equipMap {
		bagMap[cast.ToString(itemSize)] = bi.EventValJson{
			"product_id":  mcst.BiPreEquip + cast.ToString(id),
			"product_num": num,
		}
		itemSize++
	}
	eventInfo := bi.EventValJson{
		"items": bagMap,
	}
	BiMod(b.GameCtx).Upload(bi.Bag, eventInfo)
}

//========================mod base========================//

func (b *Bag) InitDbRoleData(db *pbdb.RoleData) {
	db.BagMgr = &pbdb.BagMgr{
		Items:  make(map[mcst.ItemId]mcst.BagNum),
		Equips: make(map[mcst.EquipId]*pbdb.Equip),
	}
}

func (b *Bag) InitLoadDb(db *pbdb.RoleData) {
	b.bagMgr = db.BagMgr
}

func (b *Bag) RspSync(rsp *pbreq.RspSync) {
	equips := make(map[mcst.EquipId]*pbreq.Equip)
	for i, e := range b.bagMgr.Equips {
		equips[i] = &pbreq.Equip{
			Id:         e.Id,
			Lv:         e.Lv,
			ExtAttrIds: e.ExtAttrIds,
		}
	}
	rsp.Bag = &pbreq.Bag{
		Items:  b.bagMgr.Items,
		Equips: equips,
	}
	return
}

func (b *Bag) Save(any *pbdb.RoleData) {
	any.BagMgr = b.bagMgr
}

func NewBag() *Bag {
	ins := &Bag{}
	ins.InitModApi(ins)
	return ins
}
