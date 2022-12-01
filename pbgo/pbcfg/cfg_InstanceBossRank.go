package config

type InstanceBossRankCfg struct {
	Id     int32
	Type   int32
	Rank   int32
	Rank2  int32
	Mail   int32
	Reward []*Reward
}

type InstanceBossRankTable struct {
	data map[int32]*InstanceBossRankCfg
}

var InstanceBossRankData = &InstanceBossRankTable{
	data: map[int32]*InstanceBossRankCfg{},
}

func (table *InstanceBossRankTable) Get(id int32) *InstanceBossRankCfg {
	return table.data[id]
}

func (table *InstanceBossRankTable) GetAll() []int32 {
	return instanceBossRankKeys
}

var instanceBossRankValues = []*InstanceBossRankCfg{
	{
		Id:    1,
		Type:  1,
		Rank:  1,
		Rank2: 1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 3,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
		},
	},
	{
		Id:    2,
		Type:  1,
		Rank:  2,
		Rank2: 2,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 4,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 3,
			},
			{
				Type:  1,
				Id:    2,
				Count: 650,
			},
		},
	},
	{
		Id:    3,
		Type:  1,
		Rank:  3,
		Rank2: 3,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 3,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 3,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
		},
	},
	{
		Id:    4,
		Type:  2,
		Rank:  4,
		Rank2: 10,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 3,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 2,
			},
			{
				Type:  1,
				Id:    2,
				Count: 550,
			},
		},
	},
	{
		Id:    5,
		Type:  2,
		Rank:  11,
		Rank2: 20,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 2,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 2,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
		},
	},
	{
		Id:    6,
		Type:  2,
		Rank:  21,
		Rank2: 30,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 2,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
			{
				Type:  1,
				Id:    2,
				Count: 450,
			},
		},
	},
	{
		Id:    7,
		Type:  2,
		Rank:  31,
		Rank2: 50,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 2,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
		},
	},
	{
		Id:    8,
		Type:  2,
		Rank:  51,
		Rank2: 100,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 2,
			},
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
			{
				Type:  1,
				Id:    2,
				Count: 350,
			},
		},
	},
	{
		Id:    9,
		Type:  2,
		Rank:  101,
		Rank2: 300,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 2,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
		},
	},
	{
		Id:    10,
		Type:  2,
		Rank:  301,
		Rank2: -1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 1,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
		},
	},
}

var instanceBossRankKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func init() {
	InstanceBossRankData.data = make(map[int32]*InstanceBossRankCfg)
	for i := 0; i < len(instanceBossRankKeys); i++ {
		InstanceBossRankData.data[instanceBossRankKeys[i]] = instanceBossRankValues[i]
	}
}
