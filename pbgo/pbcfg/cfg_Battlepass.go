package config

type BattlepassCfg struct {
	Id           int32
	Type         int32
	Level        int32
	NormalReward []*Reward
	BraveReward  []*Reward
	None         string
	RewardLv     int32
}

type BattlepassTable struct {
	data map[int32]*BattlepassCfg
}

var BattlepassData = &BattlepassTable{
	data: map[int32]*BattlepassCfg{},
}

func (table *BattlepassTable) Get(id int32) *BattlepassCfg {
	return table.data[id]
}

func (table *BattlepassTable) GetAll() []int32 {
	return battlepassKeys
}

var battlepassValues = []*BattlepassCfg{
	{
		Id:    1,
		Type:  1,
		Level: 1010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 1,
	},
	{
		Id:    2,
		Type:  1,
		Level: 2010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    2005,
				Count: 1,
			},
		},
		RewardLv: 2,
	},
	{
		Id:    3,
		Type:  1,
		Level: 3010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 3,
	},
	{
		Id:    4,
		Type:  1,
		Level: 4010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    2005,
				Count: 1,
			},
		},
		RewardLv: 4,
	},
	{
		Id:    5,
		Type:  1,
		Level: 5010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 5,
	},
	{
		Id:    6,
		Type:  1,
		Level: 6010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		RewardLv: 6,
	},
	{
		Id:    7,
		Type:  1,
		Level: 7010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 7,
	},
	{
		Id:    8,
		Type:  1,
		Level: 8010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 8,
	},
	{
		Id:    9,
		Type:  1,
		Level: 9010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 9,
	},
	{
		Id:    10,
		Type:  1,
		Level: 10010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 10,
	},
	{
		Id:    11,
		Type:  1,
		Level: 11010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 11,
	},
	{
		Id:    12,
		Type:  1,
		Level: 12010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 12,
	},
	{
		Id:    13,
		Type:  1,
		Level: 13010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 13,
	},
	{
		Id:    14,
		Type:  1,
		Level: 14010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 14,
	},
	{
		Id:    15,
		Type:  1,
		Level: 15010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 15,
	},
	{
		Id:    16,
		Type:  1,
		Level: 16010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 16,
	},
	{
		Id:    17,
		Type:  1,
		Level: 17010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 17,
	},
	{
		Id:    18,
		Type:  1,
		Level: 18010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 18,
	},
	{
		Id:    19,
		Type:  1,
		Level: 19010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 19,
	},
	{
		Id:    20,
		Type:  1,
		Level: 20010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 20,
	},
	{
		Id:    21,
		Type:  1,
		Level: 21010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 21,
	},
	{
		Id:    22,
		Type:  1,
		Level: 22010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 22,
	},
	{
		Id:    23,
		Type:  1,
		Level: 23010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 23,
	},
	{
		Id:    24,
		Type:  1,
		Level: 24010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 24,
	},
	{
		Id:    25,
		Type:  1,
		Level: 25010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 25,
	},
	{
		Id:    26,
		Type:  1,
		Level: 26010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 26,
	},
	{
		Id:    27,
		Type:  1,
		Level: 27010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 27,
	},
	{
		Id:    28,
		Type:  1,
		Level: 28010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 28,
	},
	{
		Id:    29,
		Type:  1,
		Level: 29010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 29,
	},
	{
		Id:    30,
		Type:  1,
		Level: 30010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 30,
	},
	{
		Id:    31,
		Type:  1,
		Level: 31010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 31,
	},
	{
		Id:    32,
		Type:  1,
		Level: 32010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 32,
	},
	{
		Id:    33,
		Type:  1,
		Level: 33010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 33,
	},
	{
		Id:    34,
		Type:  1,
		Level: 34010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 34,
	},
	{
		Id:    35,
		Type:  1,
		Level: 35010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 35,
	},
	{
		Id:    36,
		Type:  1,
		Level: 36010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 36,
	},
	{
		Id:    37,
		Type:  1,
		Level: 37010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 37,
	},
	{
		Id:    38,
		Type:  1,
		Level: 38010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 38,
	},
	{
		Id:    39,
		Type:  1,
		Level: 39010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 39,
	},
	{
		Id:    40,
		Type:  1,
		Level: 40010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 40,
	},
	{
		Id:    41,
		Type:  1,
		Level: 41010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 41,
	},
	{
		Id:    42,
		Type:  1,
		Level: 42010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 42,
	},
	{
		Id:    43,
		Type:  1,
		Level: 43010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 43,
	},
	{
		Id:    44,
		Type:  1,
		Level: 44010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 44,
	},
	{
		Id:    45,
		Type:  1,
		Level: 45010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 45,
	},
	{
		Id:    46,
		Type:  1,
		Level: 46010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 46,
	},
	{
		Id:    47,
		Type:  1,
		Level: 47010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 47,
	},
	{
		Id:    48,
		Type:  1,
		Level: 48010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 48,
	},
	{
		Id:    49,
		Type:  1,
		Level: 49010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 10,
			},
		},
		RewardLv: 49,
	},
	{
		Id:    50,
		Type:  1,
		Level: 50010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 50,
	},
	{
		Id:    51,
		Type:  1,
		Level: 51010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 51,
	},
	{
		Id:    52,
		Type:  1,
		Level: 52010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 52,
	},
	{
		Id:    53,
		Type:  1,
		Level: 53010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 53,
	},
	{
		Id:    54,
		Type:  1,
		Level: 54010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 54,
	},
	{
		Id:    55,
		Type:  1,
		Level: 55010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 55,
	},
	{
		Id:    56,
		Type:  1,
		Level: 56010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 56,
	},
	{
		Id:    57,
		Type:  1,
		Level: 57010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 57,
	},
	{
		Id:    58,
		Type:  1,
		Level: 58010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 58,
	},
	{
		Id:    59,
		Type:  1,
		Level: 59010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 59,
	},
	{
		Id:    60,
		Type:  1,
		Level: 60010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 60,
	},
	{
		Id:    61,
		Type:  1,
		Level: 61010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 61,
	},
	{
		Id:    62,
		Type:  1,
		Level: 62010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 62,
	},
	{
		Id:    63,
		Type:  1,
		Level: 63010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 63,
	},
	{
		Id:    64,
		Type:  1,
		Level: 64010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 64,
	},
	{
		Id:    65,
		Type:  1,
		Level: 65010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 65,
	},
	{
		Id:    66,
		Type:  1,
		Level: 66010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 66,
	},
	{
		Id:    67,
		Type:  1,
		Level: 67010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 67,
	},
	{
		Id:    68,
		Type:  1,
		Level: 68010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 68,
	},
	{
		Id:    69,
		Type:  1,
		Level: 69010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 69,
	},
	{
		Id:    70,
		Type:  1,
		Level: 70010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 70,
	},
	{
		Id:    71,
		Type:  1,
		Level: 71010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 71,
	},
	{
		Id:    72,
		Type:  1,
		Level: 72010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 72,
	},
	{
		Id:    73,
		Type:  1,
		Level: 73010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 73,
	},
	{
		Id:    74,
		Type:  1,
		Level: 74010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 74,
	},
	{
		Id:    75,
		Type:  1,
		Level: 75010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 75,
	},
	{
		Id:    76,
		Type:  1,
		Level: 76010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 76,
	},
	{
		Id:    77,
		Type:  1,
		Level: 77010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 77,
	},
	{
		Id:    78,
		Type:  1,
		Level: 78010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 78,
	},
	{
		Id:    79,
		Type:  1,
		Level: 79010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 79,
	},
	{
		Id:    80,
		Type:  1,
		Level: 80010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 80,
	},
	{
		Id:    81,
		Type:  1,
		Level: 81010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 81,
	},
	{
		Id:    82,
		Type:  1,
		Level: 82010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 82,
	},
	{
		Id:    83,
		Type:  1,
		Level: 83010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 83,
	},
	{
		Id:    84,
		Type:  1,
		Level: 84010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 84,
	},
	{
		Id:    85,
		Type:  1,
		Level: 85010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 85,
	},
	{
		Id:    86,
		Type:  1,
		Level: 86010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 86,
	},
	{
		Id:    87,
		Type:  1,
		Level: 87010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 87,
	},
	{
		Id:    88,
		Type:  1,
		Level: 88010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 88,
	},
	{
		Id:    89,
		Type:  1,
		Level: 89010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 89,
	},
	{
		Id:    90,
		Type:  1,
		Level: 90010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 90,
	},
	{
		Id:    91,
		Type:  1,
		Level: 91010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
		RewardLv: 91,
	},
	{
		Id:    92,
		Type:  1,
		Level: 92010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 92,
	},
	{
		Id:    93,
		Type:  1,
		Level: 93010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 93,
	},
	{
		Id:    94,
		Type:  1,
		Level: 94010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 94,
	},
	{
		Id:    95,
		Type:  1,
		Level: 95010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 5,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		RewardLv: 95,
	},
	{
		Id:    96,
		Type:  1,
		Level: 96010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 96,
	},
	{
		Id:    97,
		Type:  1,
		Level: 97010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10001,
				Count: 1,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		RewardLv: 97,
	},
	{
		Id:    98,
		Type:  1,
		Level: 98010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
		RewardLv: 98,
	},
	{
		Id:    99,
		Type:  1,
		Level: 99010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 2,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 2,
			},
		},
		RewardLv: 99,
	},
	{
		Id:    100,
		Type:  1,
		Level: 100010,
		NormalReward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		BraveReward: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 10,
			},
		},
		RewardLv: 100,
	},
}

var battlepassKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
}

func init() {
	BattlepassData.data = make(map[int32]*BattlepassCfg)
	for i := 0; i < len(battlepassKeys); i++ {
		BattlepassData.data[battlepassKeys[i]] = battlepassValues[i]
	}
}
