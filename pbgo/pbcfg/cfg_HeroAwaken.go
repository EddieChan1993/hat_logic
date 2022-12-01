package config

type HeroAwakenCfg struct {
	Id        int32
	MaxLv     int32
	AtkHabit1 []*AddAttribute
	AtkHabit2 []*AddAttribute
	AtkHabit3 []*AddAttribute
}

type HeroAwakenTable struct {
	data map[int32]*HeroAwakenCfg
}

var HeroAwakenData = &HeroAwakenTable{
	data: map[int32]*HeroAwakenCfg{},
}

func (table *HeroAwakenTable) Get(id int32) *HeroAwakenCfg {
	return table.data[id]
}

func (table *HeroAwakenTable) GetAll() []int32 {
	return heroAwakenKeys
}

var heroAwakenValues = []*HeroAwakenCfg{
	{
		Id:    1,
		MaxLv: 100,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 0,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 0,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 0,
			},
			{
				Type:  2,
				Id:    32,
				Count: 0,
			},
		},
	},
	{
		Id:    2,
		MaxLv: 300,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 3500,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 3500,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 1750,
			},
			{
				Type:  2,
				Id:    32,
				Count: 1750,
			},
		},
	},
	{
		Id:    3,
		MaxLv: 600,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 7500,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 7500,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 3750,
			},
			{
				Type:  2,
				Id:    32,
				Count: 3750,
			},
		},
	},
	{
		Id:    4,
		MaxLv: 900,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 12000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 12000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 6000,
			},
			{
				Type:  2,
				Id:    32,
				Count: 6000,
			},
		},
	},
	{
		Id:    5,
		MaxLv: 1200,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 17000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 17000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 8500,
			},
			{
				Type:  2,
				Id:    32,
				Count: 8500,
			},
		},
	},
	{
		Id:    6,
		MaxLv: 1500,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 22500,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 22500,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 11250,
			},
			{
				Type:  2,
				Id:    32,
				Count: 11250,
			},
		},
	},
	{
		Id:    7,
		MaxLv: 1800,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 28500,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 28500,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 14250,
			},
			{
				Type:  2,
				Id:    32,
				Count: 14250,
			},
		},
	},
	{
		Id:    8,
		MaxLv: 2100,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 35000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 35000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 17500,
			},
			{
				Type:  2,
				Id:    32,
				Count: 17500,
			},
		},
	},
	{
		Id:    9,
		MaxLv: 2400,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 42000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 42000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 21000,
			},
			{
				Type:  2,
				Id:    32,
				Count: 21000,
			},
		},
	},
	{
		Id:    10,
		MaxLv: 2700,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 49500,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 49500,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 24750,
			},
			{
				Type:  2,
				Id:    32,
				Count: 24750,
			},
		},
	},
	{
		Id:    11,
		MaxLv: 3000,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 57500,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 57500,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 28750,
			},
			{
				Type:  2,
				Id:    32,
				Count: 28750,
			},
		},
	},
	{
		Id:    12,
		MaxLv: 3300,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 66000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 66000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 33000,
			},
			{
				Type:  2,
				Id:    32,
				Count: 33000,
			},
		},
	},
	{
		Id:    13,
		MaxLv: 3600,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 75000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 75000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 37500,
			},
			{
				Type:  2,
				Id:    32,
				Count: 37500,
			},
		},
	},
	{
		Id:    14,
		MaxLv: 3900,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 85000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 85000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 42500,
			},
			{
				Type:  2,
				Id:    32,
				Count: 42500,
			},
		},
	},
	{
		Id:    15,
		MaxLv: 4200,
		AtkHabit1: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 95000,
			},
		},
		AtkHabit2: []*AddAttribute{
			{
				Type:  2,
				Id:    32,
				Count: 95000,
			},
		},
		AtkHabit3: []*AddAttribute{
			{
				Type:  2,
				Id:    22,
				Count: 47500,
			},
			{
				Type:  2,
				Id:    32,
				Count: 47500,
			},
		},
	},
}

var heroAwakenKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15,
}

func init() {
	HeroAwakenData.data = make(map[int32]*HeroAwakenCfg)
	for i := 0; i < len(heroAwakenKeys); i++ {
		HeroAwakenData.data[heroAwakenKeys[i]] = heroAwakenValues[i]
	}
}
