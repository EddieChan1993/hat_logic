package config

type RandomChestCfg struct {
	Id           int32
	Type         int32
	Reward       []*Reward
	RandomReward []*RandomReward
}

type RandomChestTable struct {
	data map[int32]*RandomChestCfg
}

var RandomChestData = &RandomChestTable{
	data: map[int32]*RandomChestCfg{},
}

func (table *RandomChestTable) Get(id int32) *RandomChestCfg {
	return table.data[id]
}

func (table *RandomChestTable) GetAll() []int32 {
	return randomChestKeys
}

var randomChestValues = []*RandomChestCfg{
	{
		Id:   1,
		Type: 2,
		RandomReward: []*RandomReward{
			{
				Type:  2,
				Id:    11001,
				Num:   1,
				Ratio: 600,
			},
			{
				Type:  2,
				Id:    13001,
				Num:   1,
				Ratio: 600,
			},
			{
				Type:  2,
				Id:    44001,
				Num:   1,
				Ratio: 500,
			},
			{
				Type:  2,
				Id:    45001,
				Num:   1,
				Ratio: 400,
			},
		},
	},
	{
		Id:   2,
		Type: 2,
		RandomReward: []*RandomReward{
			{
				Type:  4,
				Id:    1,
				Num:   1,
				Ratio: 600,
			},
			{
				Type:  4,
				Id:    2,
				Num:   1,
				Ratio: 600,
			},
			{
				Type:  4,
				Id:    5,
				Num:   1,
				Ratio: 500,
			},
			{
				Type:  4,
				Id:    6,
				Num:   1,
				Ratio: 400,
			},
		},
	},
	{
		Id:   10001,
		Type: 2,
		RandomReward: []*RandomReward{
			{
				Type:  3,
				Id:    10721,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    10921,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11122,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11322,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11421,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11511,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11612,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11712,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11812,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    11911,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12011,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12112,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12212,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12312,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12412,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12511,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12611,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12711,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12811,
				Num:   1,
				Ratio: 100,
			},
			{
				Type:  3,
				Id:    12911,
				Num:   1,
				Ratio: 100,
			},
		},
	},
	{
		Id:   10002,
		Type: 2,
		RandomReward: []*RandomReward{
			{
				Type:  3,
				Id:    10021,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10121,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10222,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10321,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10421,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10521,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10622,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10721,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    10822,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10921,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11021,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    11122,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11222,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    11322,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11421,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11511,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11612,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11712,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11812,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    11911,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12011,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12112,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12212,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12312,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12412,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12511,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12611,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12711,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12811,
				Num:   1,
				Ratio: 350,
			},
			{
				Type:  3,
				Id:    12911,
				Num:   1,
				Ratio: 350,
			},
		},
	},
	{
		Id:   10003,
		Type: 2,
		RandomReward: []*RandomReward{
			{
				Type:  3,
				Id:    10021,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10121,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10222,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10321,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10421,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10521,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10622,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    10822,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    11021,
				Num:   1,
				Ratio: 300,
			},
			{
				Type:  3,
				Id:    11222,
				Num:   1,
				Ratio: 300,
			},
		},
	},
}

var randomChestKeys = []int32{
	1, 2, 10001, 10002, 10003,
}

func init() {
	RandomChestData.data = make(map[int32]*RandomChestCfg)
	for i := 0; i < len(randomChestKeys); i++ {
		RandomChestData.data[randomChestKeys[i]] = randomChestValues[i]
	}
}
