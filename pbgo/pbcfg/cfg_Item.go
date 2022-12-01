package config

type ItemCfg struct {
	Id          int32
	ShowType    int32
	Qlt         int32
	Recovery    int32
	Use         []*Reward
	Icon        string
	ProfitTime  int32
	HeroId      int32
	RandomChest int32
	BapType     int32
}

type ItemTable struct {
	data map[int32]*ItemCfg
}

var ItemData = &ItemTable{
	data: map[int32]*ItemCfg{},
}

func (table *ItemTable) Get(id int32) *ItemCfg {
	return table.data[id]
}

func (table *ItemTable) GetAll() []int32 {
	return itemKeys
}

var itemValues = []*ItemCfg{
	{
		Id:       1,
		ShowType: 1,
	},
	{
		Id:       2,
		ShowType: 1,
	},
	{
		Id:       3,
		ShowType: 1,
	},
	{
		Id:       4,
		ShowType: 1,
	},
	{
		Id:       5,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10021,
	},
	{
		Id:       6,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10121,
	},
	{
		Id:       7,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10222,
	},
	{
		Id:       8,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10321,
	},
	{
		Id:       9,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10421,
	},
	{
		Id:       10,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10521,
	},
	{
		Id:       11,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10622,
	},
	{
		Id:       12,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10721,
	},
	{
		Id:       13,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10822,
	},
	{
		Id:       14,
		ShowType: 2,
		Recovery: 1,
		HeroId:   10921,
	},
	{
		Id:       15,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11021,
	},
	{
		Id:       16,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11122,
	},
	{
		Id:       17,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11222,
	},
	{
		Id:       18,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11322,
	},
	{
		Id:       19,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11421,
	},
	{
		Id:       20,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11511,
	},
	{
		Id:       21,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11612,
	},
	{
		Id:       22,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11712,
	},
	{
		Id:       23,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11812,
	},
	{
		Id:       24,
		ShowType: 2,
		Recovery: 1,
		HeroId:   11911,
	},
	{
		Id:       25,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12011,
	},
	{
		Id:       26,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12112,
	},
	{
		Id:       27,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12212,
	},
	{
		Id:       28,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12312,
	},
	{
		Id:       29,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12412,
	},
	{
		Id:       30,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12511,
	},
	{
		Id:       31,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12611,
	},
	{
		Id:       32,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12711,
	},
	{
		Id:       33,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12811,
	},
	{
		Id:       34,
		ShowType: 2,
		Recovery: 1,
		HeroId:   12911,
	},
	{
		Id:       35,
		ShowType: 1,
		Recovery: 1,
		Use: []*Reward{
			{
				Type:  1,
				Id:    1,
				Count: 10000,
			},
		},
	},
	{
		Id:          36,
		ShowType:    1,
		RandomChest: 1,
	},
	{
		Id:       37,
		ShowType: 1,
	},
	{
		Id:       38,
		ShowType: 1,
	},
	{
		Id:       39,
		ShowType: 1,
	},
	{
		Id:       40,
		ShowType: 1,
	},
	{
		Id:       41,
		ShowType: 1,
	},
	{
		Id:          42,
		ShowType:    1,
		RandomChest: 2,
	},
	{
		Id:       43,
		ShowType: 1,
	},
	{
		Id:       44,
		ShowType: 1,
	},
	{
		Id:       45,
		ShowType: 1,
	},
	{
		Id:       46,
		ShowType: 1,
	},
	{
		Id:       47,
		ShowType: 1,
	},
	{
		Id:       50,
		ShowType: 1,
	},
	{
		Id:       51,
		ShowType: 1,
	},
	{
		Id:       101,
		ShowType: 1,
	},
	{
		Id:         1001,
		ShowType:   3,
		ProfitTime: 3600,
	},
	{
		Id:         1002,
		ShowType:   3,
		ProfitTime: 10800,
	},
	{
		Id:         1003,
		ShowType:   3,
		ProfitTime: 21600,
	},
	{
		Id:         1004,
		ShowType:   3,
		ProfitTime: 43200,
	},
	{
		Id:         1005,
		ShowType:   3,
		ProfitTime: 86400,
	},
	{
		Id:         2001,
		ShowType:   4,
		ProfitTime: 3600,
	},
	{
		Id:         2002,
		ShowType:   5,
		ProfitTime: 3600,
	},
	{
		Id:         2003,
		ShowType:   6,
		ProfitTime: 3600,
	},
	{
		Id:         2004,
		ShowType:   7,
		ProfitTime: 3600,
	},
	{
		Id:         2005,
		ShowType:   8,
		ProfitTime: 3600,
	},
	{
		Id:          10001,
		ShowType:    1,
		RandomChest: 10001,
	},
	{
		Id:          10002,
		ShowType:    1,
		RandomChest: 10002,
	},
	{
		Id:          10003,
		ShowType:    1,
		RandomChest: 10003,
	},
}

var itemKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 50, 51, 101,
	1001, 1002, 1003, 1004, 1005, 2001, 2002, 2003, 2004, 2005,
	10001, 10002, 10003,
}

func init() {
	ItemData.data = make(map[int32]*ItemCfg)
	for i := 0; i < len(itemKeys); i++ {
		ItemData.data[itemKeys[i]] = itemValues[i]
	}
}
