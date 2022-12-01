package mcst

/**
背包
*/

type EquipPart = int32 //装备部位
type ItemType = int32  //道具类型
type ItemId = int32    //道具id
type EquipId = int32   //装备id
type BagNum = Int64Typ //需要用大数值，int32无法满足要求背包各类道具数量类型
type ItemIdMapNum = map[DefaultInt]BagNum
type EquipIdMapNum = map[DefaultInt]BagNum

//==================== 道具id ====================//
const (
	ItemIdCOIN        ItemId = 1 //金币
	ItemIdHoly        ItemId = 2 //圣水
	ItemIdDIAMOND     ItemId = 3 //钻石
	ItemIdTalentPoint ItemId = 4 //天赋点

	ItemIdFarm    ItemId = 37 //农场
	ItemIdMine    ItemId = 38 //铁矿场
	ItemIdStone   ItemId = 39 //石头场
	ItemIdSawmill ItemId = 40 //木料场

	ItemUpEquipStone ItemId = 101 //强化石头
)

//==================== 装备类型 ====================//
const (
	EquipPartHead     EquipPart = 1 //头部装备
	EquipPartClothes  EquipPart = 2 //衣服装备
	EquipPartShoes    EquipPart = 3 //鞋子装备
	EquipPartWeapon   EquipPart = 4 //武器装备
	EquipPartRing     EquipPart = 5 //戒子装备
	EquipPartNecklace EquipPart = 6 //项链装备
)

//==================== 道具类型 ====================//
const (
	ItemTypCoin        ItemType = 2 //金币袋
	ItemTypCoinOffline ItemType = 3 //金币袋（离线收益）
	ItemTypFarm        ItemType = 4 //食材资源
	ItemTypSawmill     ItemType = 5 //木材资源
	ItemTypStone       ItemType = 6 //石材资源
	ItemTypMine        ItemType = 7 //铁矿资源
	ItemTypResBuild    ItemType = 8 //资源包
)
