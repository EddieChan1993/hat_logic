package mcst

/**
奖励
*/

type GoodTyp = int32                      //物品类型
type GoodIdNumMap = map[DefaultInt]BagNum //物品id和数量映射

const (
	GoodsDefaultNums    = 100000000
	RandRewardsForLimit = 2000 //随机奖品最大限制循环次数
)

//==================== 奖品所有发放类型 ====================//

const (
	GoodTypItem  GoodTyp = 1 //道具
	GoodTypEquip GoodTyp = 2 //装备
)
