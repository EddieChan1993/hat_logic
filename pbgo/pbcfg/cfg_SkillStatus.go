package config

type SkillStatusCfg struct {
	Id               int32
	Interval         int32
	Effect           []int32
	EffectParam      [][]int32
	Trigger          []int32
	SpecialTag       []int32
	Attrib           []*Attribute
	ChangeParam      []int32
	StatusContinuous int32
	FxSelf           *Fx
}

type SkillStatusTable struct {
	data map[int32]*SkillStatusCfg
}

var SkillStatusData = &SkillStatusTable{
	data: map[int32]*SkillStatusCfg{},
}

func (table *SkillStatusTable) Get(id int32) *SkillStatusCfg {
	return table.data[id]
}

func (table *SkillStatusTable) GetAll() []int32 {
	return skillStatusKeys
}

var skillStatusValues = []*SkillStatusCfg{
	{
		Id: 1,
		Attrib: []*Attribute{
			{
				Id:    23,
				Count: 100,
			},
			{
				Id:    33,
				Count: 100,
			},
		},
	},
	{
		Id: 2,
		Attrib: []*Attribute{
			{
				Id:    32,
				Count: 100,
			},
		},
	},
	{
		Id:       3,
		Interval: 5000,
		Effect: []int32{
			4,
		},
		EffectParam: [][]int32{
			{
				4,
				0,
			},
		},
		Trigger: []int32{
			1,
		},
	},
	{
		Id: 4,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 1000000,
			},
			{
				Id:    92,
				Count: 1000000,
			},
		},
	},
	{
		Id: 5,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 50,
			},
		},
	},
	{
		Id: 6,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 100,
			},
		},
	},
	{
		Id: 7,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 200,
			},
		},
	},
	{
		Id: 8,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 250,
			},
		},
	},
	{
		Id: 9,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 400,
			},
		},
	},
	{
		Id: 10,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: 50,
			},
		},
	},
	{
		Id: 11,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: 100,
			},
		},
	},
	{
		Id: 12,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: 200,
			},
		},
	},
	{
		Id: 13,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: 250,
			},
		},
	},
	{
		Id: 14,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: 400,
			},
		},
	},
	{
		Id: 15,
		Trigger: []int32{
			2,
		},
	},
	{
		Id: 16,
		Trigger: []int32{
			3,
		},
	},
	{
		Id: 17,
		Trigger: []int32{
			4,
		},
	},
	{
		Id: 18,
		Trigger: []int32{
			5,
		},
	},
	{
		Id: 19,
		Trigger: []int32{
			6,
		},
	},
	{
		Id: 20,
		Trigger: []int32{
			7,
		},
	},
	{
		Id: 21,
		Trigger: []int32{
			8,
		},
	},
	{
		Id: 22,
		Trigger: []int32{
			9,
		},
	},
	{
		Id: 23,
		Trigger: []int32{
			10,
		},
	},
	{
		Id: 24,
		Trigger: []int32{
			11,
		},
	},
	{
		Id:       25,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				300,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 30,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       26,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				400,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 50,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       27,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				500,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 70,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       28,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				700,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 100,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       29,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				800,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 120,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       30,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				900,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 140,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       31,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				1000,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 160,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       32,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				1200,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 180,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       33,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				1300,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 200,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id:       34,
		Interval: 500,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				1,
				1500,
			},
		},
		Attrib: []*Attribute{
			{
				Id:    173,
				Count: 220,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 35,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 1500,
			},
			{
				Id:    74,
				Count: 50,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 36,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 3000,
			},
			{
				Id:    74,
				Count: 100,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 37,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 4500,
			},
			{
				Id:    74,
				Count: 150,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 38,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 7000,
			},
			{
				Id:    74,
				Count: 200,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 39,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 8500,
			},
			{
				Id:    74,
				Count: 250,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 40,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 10000,
			},
			{
				Id:    74,
				Count: 300,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 41,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 11500,
			},
			{
				Id:    74,
				Count: 350,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 42,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 14000,
			},
			{
				Id:    74,
				Count: 400,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 43,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 15500,
			},
			{
				Id:    74,
				Count: 450,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 44,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 18000,
			},
			{
				Id:    74,
				Count: 500,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 45,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 50,
			},
		},
	},
	{
		Id: 46,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 100,
			},
		},
	},
	{
		Id: 47,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 200,
			},
		},
	},
	{
		Id: 48,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 250,
			},
		},
	},
	{
		Id: 49,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 400,
			},
		},
	},
	{
		Id: 50,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 6000,
			},
		},
	},
	{
		Id: 51,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 10000,
			},
		},
	},
	{
		Id: 52,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 14000,
			},
		},
	},
	{
		Id: 53,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 20000,
			},
		},
	},
	{
		Id: 54,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 28000,
			},
		},
	},
	{
		Id: 55,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -100,
			},
		},
	},
	{
		Id: 56,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -120,
			},
		},
	},
	{
		Id: 57,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -140,
			},
		},
	},
	{
		Id: 58,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -160,
			},
		},
	},
	{
		Id: 59,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -180,
			},
		},
	},
	{
		Id: 60,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -200,
			},
		},
	},
	{
		Id: 61,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -220,
			},
		},
	},
	{
		Id: 62,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -240,
			},
		},
	},
	{
		Id: 63,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -260,
			},
		},
	},
	{
		Id: 64,
		Attrib: []*Attribute{
			{
				Id:    44,
				Count: -280,
			},
		},
	},
	{
		Id: 65,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 50,
			},
		},
	},
	{
		Id: 66,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 75,
			},
		},
	},
	{
		Id: 67,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 100,
			},
		},
	},
	{
		Id: 68,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 150,
			},
		},
	},
	{
		Id: 69,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 175,
			},
		},
	},
	{
		Id: 70,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 200,
			},
		},
	},
	{
		Id: 71,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 250,
			},
		},
	},
	{
		Id: 72,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 275,
			},
		},
	},
	{
		Id: 73,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 300,
			},
		},
	},
	{
		Id: 74,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 350,
			},
		},
	},
	{
		Id: 75,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 25,
			},
		},
	},
	{
		Id: 76,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 50,
			},
		},
	},
	{
		Id: 77,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 75,
			},
		},
	},
	{
		Id: 78,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 100,
			},
		},
	},
	{
		Id: 79,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 125,
			},
		},
	},
	{
		Id: 80,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 150,
			},
		},
	},
	{
		Id: 81,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 175,
			},
		},
	},
	{
		Id: 82,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 200,
			},
		},
	},
	{
		Id: 83,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 225,
			},
		},
	},
	{
		Id: 84,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 250,
			},
		},
	},
	{
		Id: 85,
		SpecialTag: []int32{
			1,
			2,
		},
	},
	{
		Id: 86,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 50,
			},
		},
	},
	{
		Id: 87,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 100,
			},
		},
	},
	{
		Id: 88,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 200,
			},
		},
	},
	{
		Id: 89,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 250,
			},
		},
	},
	{
		Id: 90,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 400,
			},
		},
	},
	{
		Id: 91,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 50,
			},
		},
	},
	{
		Id: 92,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 75,
			},
		},
	},
	{
		Id: 93,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 100,
			},
		},
	},
	{
		Id: 94,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 150,
			},
		},
	},
	{
		Id: 95,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 175,
			},
		},
	},
	{
		Id: 96,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 200,
			},
		},
	},
	{
		Id: 97,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 250,
			},
		},
	},
	{
		Id: 98,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 275,
			},
		},
	},
	{
		Id: 99,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 300,
			},
		},
	},
	{
		Id: 100,
		Attrib: []*Attribute{
			{
				Id:    223,
				Count: 350,
			},
		},
	},
	{
		Id: 101,
		Attrib: []*Attribute{
			{
				Id:    54,
				Count: 50,
			},
		},
	},
	{
		Id: 102,
		Attrib: []*Attribute{
			{
				Id:    54,
				Count: 100,
			},
		},
	},
	{
		Id: 103,
		Attrib: []*Attribute{
			{
				Id:    54,
				Count: 200,
			},
		},
	},
	{
		Id: 104,
		Attrib: []*Attribute{
			{
				Id:    54,
				Count: 250,
			},
		},
	},
	{
		Id: 105,
		Attrib: []*Attribute{
			{
				Id:    54,
				Count: 400,
			},
		},
	},
	{
		Id: 106,
		Trigger: []int32{
			12,
		},
	},
	{
		Id: 107,
		Trigger: []int32{
			13,
		},
	},
	{
		Id: 108,
		Trigger: []int32{
			14,
		},
	},
	{
		Id: 109,
		Trigger: []int32{
			15,
		},
	},
	{
		Id: 110,
		Trigger: []int32{
			16,
		},
	},
	{
		Id: 111,
		Trigger: []int32{
			17,
		},
	},
	{
		Id: 112,
		Trigger: []int32{
			18,
		},
	},
	{
		Id: 113,
		Trigger: []int32{
			19,
		},
	},
	{
		Id: 114,
		Trigger: []int32{
			20,
		},
	},
	{
		Id: 115,
		Trigger: []int32{
			21,
		},
	},
	{
		Id: 116,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 20,
			},
			{
				Id:    203,
				Count: 20,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 117,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 40,
			},
			{
				Id:    203,
				Count: 40,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 118,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 60,
			},
			{
				Id:    203,
				Count: 60,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 119,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 80,
			},
			{
				Id:    203,
				Count: 80,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 120,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 100,
			},
			{
				Id:    203,
				Count: 100,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 121,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 120,
			},
			{
				Id:    203,
				Count: 120,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 122,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 140,
			},
			{
				Id:    203,
				Count: 140,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 123,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 160,
			},
			{
				Id:    203,
				Count: 160,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 124,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 180,
			},
			{
				Id:    203,
				Count: 180,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 125,
		Attrib: []*Attribute{
			{
				Id:    183,
				Count: 200,
			},
			{
				Id:    203,
				Count: 200,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 126,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 40,
			},
			{
				Id:    93,
				Count: 6000,
			},
		},
	},
	{
		Id: 127,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 80,
			},
			{
				Id:    93,
				Count: 10000,
			},
		},
	},
	{
		Id: 128,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 120,
			},
			{
				Id:    93,
				Count: 14000,
			},
		},
	},
	{
		Id: 129,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 160,
			},
			{
				Id:    93,
				Count: 20000,
			},
		},
	},
	{
		Id: 130,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 200,
			},
			{
				Id:    93,
				Count: 28000,
			},
		},
	},
	{
		Id: 131,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 25,
			},
		},
	},
	{
		Id: 132,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 50,
			},
		},
	},
	{
		Id: 133,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 75,
			},
		},
	},
	{
		Id: 134,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 100,
			},
		},
	},
	{
		Id: 135,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 125,
			},
		},
	},
	{
		Id: 136,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 150,
			},
		},
	},
	{
		Id: 137,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 175,
			},
		},
	},
	{
		Id: 138,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 200,
			},
		},
	},
	{
		Id: 139,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 225,
			},
		},
	},
	{
		Id: 140,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 250,
			},
		},
	},
	{
		Id: 141,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 1500,
			},
			{
				Id:    103,
				Count: 100,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 142,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 3000,
			},
			{
				Id:    103,
				Count: 200,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 143,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 4500,
			},
			{
				Id:    103,
				Count: 300,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 144,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 7000,
			},
			{
				Id:    103,
				Count: 400,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 145,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 8500,
			},
			{
				Id:    103,
				Count: 500,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 146,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 10000,
			},
			{
				Id:    103,
				Count: 600,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 147,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 11500,
			},
			{
				Id:    103,
				Count: 700,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 148,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 14000,
			},
			{
				Id:    103,
				Count: 800,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 149,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 15500,
			},
			{
				Id:    103,
				Count: 900,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 150,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: 18000,
			},
			{
				Id:    103,
				Count: 1000,
			},
		},
		StatusContinuous: 1,
	},
	{
		Id: 151,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -500,
			},
		},
	},
	{
		Id: 152,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -3000,
			},
		},
	},
	{
		Id: 153,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -8000,
			},
		},
	},
	{
		Id: 154,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -15000,
			},
		},
	},
	{
		Id: 155,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -50000,
			},
		},
	},
	{
		Id: 156,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -100000,
			},
		},
	},
	{
		Id: 157,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -200000,
			},
		},
	},
	{
		Id: 158,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -500000,
			},
		},
	},
	{
		Id: 159,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -1200000,
			},
		},
	},
	{
		Id: 160,
		Attrib: []*Attribute{
			{
				Id:    53,
				Count: -2000000,
			},
		},
	},
	{
		Id: 161,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -1000,
			},
			{
				Id:    54,
				Count: -100,
			},
		},
	},
	{
		Id: 162,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -2000,
			},
			{
				Id:    54,
				Count: -120,
			},
		},
	},
	{
		Id: 163,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -3000,
			},
			{
				Id:    54,
				Count: -140,
			},
		},
	},
	{
		Id: 164,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -4000,
			},
			{
				Id:    54,
				Count: -160,
			},
		},
	},
	{
		Id: 165,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -5000,
			},
			{
				Id:    54,
				Count: -180,
			},
		},
	},
	{
		Id: 166,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -6000,
			},
			{
				Id:    54,
				Count: -200,
			},
		},
	},
	{
		Id: 167,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -7000,
			},
			{
				Id:    54,
				Count: -220,
			},
		},
	},
	{
		Id: 168,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -8000,
			},
			{
				Id:    54,
				Count: -240,
			},
		},
	},
	{
		Id: 169,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -9000,
			},
			{
				Id:    54,
				Count: -260,
			},
		},
	},
	{
		Id: 170,
		Attrib: []*Attribute{
			{
				Id:    93,
				Count: -10000,
			},
			{
				Id:    54,
				Count: -280,
			},
		},
	},
	{
		Id: 171,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 10,
			},
			{
				Id:    32,
				Count: 10,
			},
		},
	},
	{
		Id: 172,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 20,
			},
			{
				Id:    32,
				Count: 20,
			},
		},
	},
	{
		Id: 173,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 100,
			},
		},
	},
	{
		Id: 174,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 200,
			},
		},
	},
	{
		Id: 175,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 350,
			},
		},
	},
	{
		Id: 176,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 600,
			},
		},
	},
	{
		Id: 177,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 800,
			},
		},
	},
	{
		Id: 178,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 100,
			},
		},
	},
	{
		Id: 179,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 200,
			},
		},
	},
	{
		Id: 180,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 350,
			},
		},
	},
	{
		Id: 181,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 600,
			},
		},
	},
	{
		Id: 182,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 800,
			},
		},
	},
	{
		Id: 183,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 100,
			},
		},
	},
	{
		Id: 184,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 200,
			},
		},
	},
	{
		Id: 185,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 350,
			},
		},
	},
	{
		Id: 186,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 600,
			},
		},
	},
	{
		Id: 187,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 800,
			},
		},
	},
	{
		Id: 188,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 100,
			},
		},
	},
	{
		Id: 189,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 200,
			},
		},
	},
	{
		Id: 190,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 350,
			},
		},
	},
	{
		Id: 191,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 600,
			},
		},
	},
	{
		Id: 192,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 800,
			},
		},
	},
	{
		Id: 193,
		Trigger: []int32{
			22,
		},
	},
	{
		Id: 194,
		Trigger: []int32{
			23,
		},
	},
	{
		Id: 195,
		Trigger: []int32{
			24,
		},
	},
	{
		Id: 196,
		Trigger: []int32{
			25,
		},
	},
	{
		Id: 197,
		Trigger: []int32{
			26,
		},
	},
	{
		Id: 198,
		Trigger: []int32{
			27,
		},
	},
	{
		Id: 199,
		Trigger: []int32{
			28,
		},
	},
	{
		Id: 200,
		Trigger: []int32{
			29,
		},
	},
	{
		Id: 201,
		Trigger: []int32{
			30,
		},
	},
	{
		Id: 202,
		Trigger: []int32{
			31,
		},
	},
	{
		Id: 203,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 20,
			},
			{
				Id:    223,
				Count: 20,
			},
		},
	},
	{
		Id: 204,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 40,
			},
			{
				Id:    223,
				Count: 40,
			},
		},
	},
	{
		Id: 205,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 60,
			},
			{
				Id:    223,
				Count: 60,
			},
		},
	},
	{
		Id: 206,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 80,
			},
			{
				Id:    223,
				Count: 80,
			},
		},
	},
	{
		Id: 207,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 100,
			},
			{
				Id:    223,
				Count: 100,
			},
		},
	},
	{
		Id: 208,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 120,
			},
			{
				Id:    223,
				Count: 120,
			},
		},
	},
	{
		Id: 209,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 140,
			},
			{
				Id:    223,
				Count: 140,
			},
		},
	},
	{
		Id: 210,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 160,
			},
			{
				Id:    223,
				Count: 160,
			},
		},
	},
	{
		Id: 211,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 180,
			},
			{
				Id:    223,
				Count: 180,
			},
		},
	},
	{
		Id: 212,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 200,
			},
			{
				Id:    223,
				Count: 200,
			},
		},
	},
	{
		Id: 213,
		Attrib: []*Attribute{
			{
				Id:    63,
				Count: 40,
			},
		},
	},
	{
		Id: 214,
		Attrib: []*Attribute{
			{
				Id:    63,
				Count: 80,
			},
		},
	},
	{
		Id: 215,
		Attrib: []*Attribute{
			{
				Id:    63,
				Count: 120,
			},
		},
	},
	{
		Id: 216,
		Attrib: []*Attribute{
			{
				Id:    63,
				Count: 160,
			},
		},
	},
	{
		Id: 217,
		Attrib: []*Attribute{
			{
				Id:    63,
				Count: 200,
			},
		},
	},
	{
		Id: 218,
		Attrib: []*Attribute{
			{
				Id:    153,
				Count: 4000,
			},
		},
	},
	{
		Id: 219,
		Attrib: []*Attribute{
			{
				Id:    153,
				Count: 8000,
			},
		},
	},
	{
		Id: 220,
		Attrib: []*Attribute{
			{
				Id:    153,
				Count: 12000,
			},
		},
	},
	{
		Id: 221,
		Attrib: []*Attribute{
			{
				Id:    153,
				Count: 18000,
			},
		},
	},
	{
		Id: 222,
		Attrib: []*Attribute{
			{
				Id:    153,
				Count: 25000,
			},
		},
	},
	{
		Id: 223,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 25,
			},
		},
	},
	{
		Id: 224,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 50,
			},
		},
	},
	{
		Id: 225,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 75,
			},
		},
	},
	{
		Id: 226,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 100,
			},
		},
	},
	{
		Id: 227,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 125,
			},
		},
	},
	{
		Id: 228,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 150,
			},
		},
	},
	{
		Id: 229,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 175,
			},
		},
	},
	{
		Id: 230,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 200,
			},
		},
	},
	{
		Id: 231,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 225,
			},
		},
	},
	{
		Id: 232,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 250,
			},
		},
	},
	{
		Id: 233,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 25,
			},
		},
	},
	{
		Id: 234,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 50,
			},
		},
	},
	{
		Id: 235,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 75,
			},
		},
	},
	{
		Id: 236,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 100,
			},
		},
	},
	{
		Id: 237,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 125,
			},
		},
	},
	{
		Id: 238,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 150,
			},
		},
	},
	{
		Id: 239,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 175,
			},
		},
	},
	{
		Id: 240,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 200,
			},
		},
	},
	{
		Id: 241,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 225,
			},
		},
	},
	{
		Id: 242,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 250,
			},
		},
	},
	{
		Id: 243,
		Attrib: []*Attribute{
			{
				Id:    74,
				Count: 100,
			},
		},
	},
	{
		Id: 244,
		Attrib: []*Attribute{
			{
				Id:    74,
				Count: 200,
			},
		},
	},
	{
		Id: 245,
		Attrib: []*Attribute{
			{
				Id:    74,
				Count: 300,
			},
		},
	},
	{
		Id: 246,
		Attrib: []*Attribute{
			{
				Id:    74,
				Count: 400,
			},
		},
	},
	{
		Id: 247,
		Attrib: []*Attribute{
			{
				Id:    74,
				Count: 500,
			},
		},
	},
	{
		Id: 248,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 100,
			},
		},
	},
	{
		Id: 249,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 200,
			},
		},
	},
	{
		Id: 250,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 300,
			},
		},
	},
	{
		Id: 251,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 400,
			},
		},
	},
	{
		Id: 252,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 500,
			},
		},
	},
	{
		Id: 253,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 20,
			},
			{
				Id:    263,
				Count: 20,
			},
		},
	},
	{
		Id: 254,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 40,
			},
			{
				Id:    263,
				Count: 40,
			},
		},
	},
	{
		Id: 255,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 60,
			},
			{
				Id:    263,
				Count: 60,
			},
		},
	},
	{
		Id: 256,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 80,
			},
			{
				Id:    263,
				Count: 80,
			},
		},
	},
	{
		Id: 257,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 100,
			},
			{
				Id:    263,
				Count: 100,
			},
		},
	},
	{
		Id: 258,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 120,
			},
			{
				Id:    263,
				Count: 120,
			},
		},
	},
	{
		Id: 259,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 140,
			},
			{
				Id:    263,
				Count: 140,
			},
		},
	},
	{
		Id: 260,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 160,
			},
			{
				Id:    263,
				Count: 160,
			},
		},
	},
	{
		Id: 261,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 180,
			},
			{
				Id:    263,
				Count: 180,
			},
		},
	},
	{
		Id: 262,
		Attrib: []*Attribute{
			{
				Id:    253,
				Count: 200,
			},
			{
				Id:    263,
				Count: 200,
			},
		},
	},
	{
		Id: 263,
		Trigger: []int32{
			32,
		},
	},
	{
		Id: 264,
		Trigger: []int32{
			33,
		},
	},
	{
		Id: 265,
		Trigger: []int32{
			34,
		},
	},
	{
		Id: 266,
		Trigger: []int32{
			35,
		},
	},
	{
		Id: 267,
		Trigger: []int32{
			36,
		},
	},
	{
		Id: 268,
		Trigger: []int32{
			37,
		},
	},
	{
		Id: 269,
		Trigger: []int32{
			38,
		},
	},
	{
		Id: 270,
		Trigger: []int32{
			39,
		},
	},
	{
		Id: 271,
		Trigger: []int32{
			40,
		},
	},
	{
		Id: 272,
		Trigger: []int32{
			41,
		},
	},
	{
		Id: 273,
		Trigger: []int32{
			42,
		},
	},
	{
		Id: 274,
		Trigger: []int32{
			43,
		},
	},
	{
		Id: 275,
		Trigger: []int32{
			44,
		},
	},
	{
		Id: 276,
		Trigger: []int32{
			45,
		},
	},
	{
		Id: 277,
		Trigger: []int32{
			46,
		},
	},
	{
		Id: 278,
		Trigger: []int32{
			47,
		},
	},
	{
		Id: 279,
		Trigger: []int32{
			48,
		},
	},
	{
		Id: 280,
		Trigger: []int32{
			49,
		},
	},
	{
		Id: 281,
		Trigger: []int32{
			50,
		},
	},
	{
		Id: 282,
		Trigger: []int32{
			51,
		},
	},
	{
		Id: 283,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 50,
			},
			{
				Id:    203,
				Count: 50,
			},
		},
	},
	{
		Id: 284,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 200,
			},
			{
				Id:    223,
				Count: 200,
			},
		},
	},
	{
		Id: 285,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 25,
			},
		},
	},
	{
		Id: 286,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 2500,
			},
			{
				Id:    93,
				Count: 2500,
			},
		},
	},
	{
		Id: 287,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 30,
			},
			{
				Id:    32,
				Count: 30,
			},
		},
	},
	{
		Id: 288,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 40,
			},
			{
				Id:    32,
				Count: 40,
			},
		},
	},
	{
		Id: 289,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 50,
			},
			{
				Id:    32,
				Count: 50,
			},
		},
	},
	{
		Id: 290,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 60,
			},
			{
				Id:    32,
				Count: 60,
			},
		},
	},
	{
		Id: 291,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 70,
			},
			{
				Id:    32,
				Count: 70,
			},
		},
	},
	{
		Id: 292,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 80,
			},
			{
				Id:    32,
				Count: 80,
			},
		},
	},
	{
		Id: 293,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 90,
			},
			{
				Id:    32,
				Count: 90,
			},
		},
	},
	{
		Id: 294,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 100,
			},
			{
				Id:    32,
				Count: 100,
			},
		},
	},
	{
		Id: 295,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 110,
			},
			{
				Id:    32,
				Count: 110,
			},
		},
	},
	{
		Id: 296,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 120,
			},
			{
				Id:    32,
				Count: 120,
			},
		},
	},
	{
		Id: 297,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 130,
			},
			{
				Id:    32,
				Count: 130,
			},
		},
	},
	{
		Id: 298,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 140,
			},
			{
				Id:    32,
				Count: 140,
			},
		},
	},
	{
		Id: 299,
		Attrib: []*Attribute{
			{
				Id:    22,
				Count: 150,
			},
			{
				Id:    32,
				Count: 150,
			},
		},
	},
	{
		Id: 300,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 70,
			},
			{
				Id:    203,
				Count: 70,
			},
		},
	},
	{
		Id: 301,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 90,
			},
			{
				Id:    203,
				Count: 90,
			},
		},
	},
	{
		Id: 302,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 110,
			},
			{
				Id:    203,
				Count: 110,
			},
		},
	},
	{
		Id: 303,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 130,
			},
			{
				Id:    203,
				Count: 130,
			},
		},
	},
	{
		Id: 304,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 150,
			},
			{
				Id:    203,
				Count: 150,
			},
		},
	},
	{
		Id: 305,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 170,
			},
			{
				Id:    203,
				Count: 170,
			},
		},
	},
	{
		Id: 306,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 190,
			},
			{
				Id:    203,
				Count: 190,
			},
		},
	},
	{
		Id: 307,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 210,
			},
			{
				Id:    203,
				Count: 210,
			},
		},
	},
	{
		Id: 308,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 230,
			},
			{
				Id:    203,
				Count: 230,
			},
		},
	},
	{
		Id: 309,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 250,
			},
			{
				Id:    203,
				Count: 250,
			},
		},
	},
	{
		Id: 310,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 270,
			},
			{
				Id:    203,
				Count: 270,
			},
		},
	},
	{
		Id: 311,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 290,
			},
			{
				Id:    203,
				Count: 290,
			},
		},
	},
	{
		Id: 312,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 310,
			},
			{
				Id:    203,
				Count: 310,
			},
		},
	},
	{
		Id: 313,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 330,
			},
			{
				Id:    203,
				Count: 330,
			},
		},
	},
	{
		Id: 314,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 400,
			},
			{
				Id:    223,
				Count: 400,
			},
		},
	},
	{
		Id: 315,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 600,
			},
			{
				Id:    223,
				Count: 600,
			},
		},
	},
	{
		Id: 316,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 800,
			},
			{
				Id:    223,
				Count: 800,
			},
		},
	},
	{
		Id: 317,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 1000,
			},
			{
				Id:    223,
				Count: 1000,
			},
		},
	},
	{
		Id: 318,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 1200,
			},
			{
				Id:    223,
				Count: 1200,
			},
		},
	},
	{
		Id: 319,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 1400,
			},
			{
				Id:    223,
				Count: 1400,
			},
		},
	},
	{
		Id: 320,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 1600,
			},
			{
				Id:    223,
				Count: 1600,
			},
		},
	},
	{
		Id: 321,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 1800,
			},
			{
				Id:    223,
				Count: 1800,
			},
		},
	},
	{
		Id: 322,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 2000,
			},
			{
				Id:    223,
				Count: 2000,
			},
		},
	},
	{
		Id: 323,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 2200,
			},
			{
				Id:    223,
				Count: 2200,
			},
		},
	},
	{
		Id: 324,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 2400,
			},
			{
				Id:    223,
				Count: 2400,
			},
		},
	},
	{
		Id: 325,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 2600,
			},
			{
				Id:    223,
				Count: 2600,
			},
		},
	},
	{
		Id: 326,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 2800,
			},
			{
				Id:    223,
				Count: 2800,
			},
		},
	},
	{
		Id: 327,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 3000,
			},
			{
				Id:    223,
				Count: 3000,
			},
		},
	},
	{
		Id: 328,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 50,
			},
		},
	},
	{
		Id: 329,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 75,
			},
		},
	},
	{
		Id: 330,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 100,
			},
		},
	},
	{
		Id: 331,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 125,
			},
		},
	},
	{
		Id: 332,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 150,
			},
		},
	},
	{
		Id: 333,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 175,
			},
		},
	},
	{
		Id: 334,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 200,
			},
		},
	},
	{
		Id: 335,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 225,
			},
		},
	},
	{
		Id: 336,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 250,
			},
		},
	},
	{
		Id: 337,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 275,
			},
		},
	},
	{
		Id: 338,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 300,
			},
		},
	},
	{
		Id: 339,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 325,
			},
		},
	},
	{
		Id: 340,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 350,
			},
		},
	},
	{
		Id: 341,
		Attrib: []*Attribute{
			{
				Id:    103,
				Count: 375,
			},
		},
	},
	{
		Id: 342,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 5000,
			},
			{
				Id:    93,
				Count: 5000,
			},
		},
	},
	{
		Id: 343,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 7500,
			},
			{
				Id:    93,
				Count: 7500,
			},
		},
	},
	{
		Id: 344,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 10000,
			},
			{
				Id:    93,
				Count: 10000,
			},
		},
	},
	{
		Id: 345,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 12500,
			},
			{
				Id:    93,
				Count: 12500,
			},
		},
	},
	{
		Id: 346,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 15000,
			},
			{
				Id:    93,
				Count: 15000,
			},
		},
	},
	{
		Id: 347,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 17500,
			},
			{
				Id:    93,
				Count: 17500,
			},
		},
	},
	{
		Id: 348,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 20000,
			},
			{
				Id:    93,
				Count: 20000,
			},
		},
	},
	{
		Id: 349,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 22500,
			},
			{
				Id:    93,
				Count: 22500,
			},
		},
	},
	{
		Id: 350,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 25000,
			},
			{
				Id:    93,
				Count: 25000,
			},
		},
	},
	{
		Id: 351,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 27500,
			},
			{
				Id:    93,
				Count: 27500,
			},
		},
	},
	{
		Id: 352,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 30000,
			},
			{
				Id:    93,
				Count: 30000,
			},
		},
	},
	{
		Id: 353,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 32500,
			},
			{
				Id:    93,
				Count: 32500,
			},
		},
	},
	{
		Id: 354,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 35000,
			},
			{
				Id:    93,
				Count: 35000,
			},
		},
	},
	{
		Id: 355,
		Attrib: []*Attribute{
			{
				Id:    83,
				Count: 37500,
			},
			{
				Id:    93,
				Count: 37500,
			},
		},
	},
	{
		Id: 356,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 1,
			},
		},
	},
	{
		Id: 357,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 2,
			},
		},
	},
	{
		Id: 358,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 3,
			},
		},
	},
	{
		Id: 359,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 4,
			},
		},
	},
	{
		Id: 360,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 5,
			},
		},
	},
	{
		Id: 361,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 6,
			},
		},
	},
	{
		Id: 362,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 7,
			},
		},
	},
	{
		Id: 363,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 8,
			},
		},
	},
	{
		Id: 364,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 9,
			},
		},
	},
	{
		Id: 365,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 10,
			},
		},
	},
	{
		Id: 366,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 11,
			},
		},
	},
	{
		Id: 367,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 12,
			},
		},
	},
	{
		Id: 368,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 13,
			},
		},
	},
	{
		Id: 369,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 14,
			},
		},
	},
	{
		Id: 370,
		Attrib: []*Attribute{
			{
				Id:    92,
				Count: 15,
			},
		},
	},
	{
		Id: 371,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 1,
			},
		},
	},
	{
		Id: 372,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 2,
			},
		},
	},
	{
		Id: 373,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 3,
			},
		},
	},
	{
		Id: 374,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 4,
			},
		},
	},
	{
		Id: 375,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 5,
			},
		},
	},
	{
		Id: 376,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 6,
			},
		},
	},
	{
		Id: 377,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 7,
			},
		},
	},
	{
		Id: 378,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 8,
			},
		},
	},
	{
		Id: 379,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 9,
			},
		},
	},
	{
		Id: 380,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 10,
			},
		},
	},
	{
		Id: 381,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 11,
			},
		},
	},
	{
		Id: 382,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 12,
			},
		},
	},
	{
		Id: 383,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 13,
			},
		},
	},
	{
		Id: 384,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 14,
			},
		},
	},
	{
		Id: 385,
		Attrib: []*Attribute{
			{
				Id:    82,
				Count: 15,
			},
		},
	},
	{
		Id: 386,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 10,
			},
		},
	},
	{
		Id: 387,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 20,
			},
		},
	},
	{
		Id: 388,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 30,
			},
		},
	},
	{
		Id: 389,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 40,
			},
		},
	},
	{
		Id: 390,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 50,
			},
		},
	},
	{
		Id: 391,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 60,
			},
		},
	},
	{
		Id: 392,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 70,
			},
		},
	},
	{
		Id: 393,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 80,
			},
		},
	},
	{
		Id: 394,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 90,
			},
		},
	},
	{
		Id: 395,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 100,
			},
		},
	},
	{
		Id: 396,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 110,
			},
		},
	},
	{
		Id: 397,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 120,
			},
		},
	},
	{
		Id: 398,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 130,
			},
		},
	},
	{
		Id: 399,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 140,
			},
		},
	},
	{
		Id: 400,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 150,
			},
		},
	},
	{
		Id: 401,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 10,
			},
		},
	},
	{
		Id: 402,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 20,
			},
		},
	},
	{
		Id: 403,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 30,
			},
		},
	},
	{
		Id: 404,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 40,
			},
		},
	},
	{
		Id: 405,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 50,
			},
		},
	},
	{
		Id: 406,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 60,
			},
		},
	},
	{
		Id: 407,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 70,
			},
		},
	},
	{
		Id: 408,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 80,
			},
		},
	},
	{
		Id: 409,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 90,
			},
		},
	},
	{
		Id: 410,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 100,
			},
		},
	},
	{
		Id: 411,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 110,
			},
		},
	},
	{
		Id: 412,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 120,
			},
		},
	},
	{
		Id: 413,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 130,
			},
		},
	},
	{
		Id: 414,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 140,
			},
		},
	},
	{
		Id: 415,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 150,
			},
		},
	},
	{
		Id: 416,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 20,
			},
			{
				Id:    323,
				Count: 20,
			},
		},
	},
	{
		Id: 417,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 40,
			},
			{
				Id:    323,
				Count: 40,
			},
		},
	},
	{
		Id: 418,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 60,
			},
			{
				Id:    323,
				Count: 60,
			},
		},
	},
	{
		Id: 419,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 80,
			},
			{
				Id:    323,
				Count: 80,
			},
		},
	},
	{
		Id: 420,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 100,
			},
			{
				Id:    323,
				Count: 100,
			},
		},
	},
	{
		Id: 421,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 120,
			},
			{
				Id:    323,
				Count: 120,
			},
		},
	},
	{
		Id: 422,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 140,
			},
			{
				Id:    323,
				Count: 140,
			},
		},
	},
	{
		Id: 423,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 160,
			},
			{
				Id:    323,
				Count: 160,
			},
		},
	},
	{
		Id: 424,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 180,
			},
			{
				Id:    323,
				Count: 180,
			},
		},
	},
	{
		Id: 425,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 200,
			},
			{
				Id:    323,
				Count: 200,
			},
		},
	},
	{
		Id: 426,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 220,
			},
			{
				Id:    323,
				Count: 220,
			},
		},
	},
	{
		Id: 427,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 240,
			},
			{
				Id:    323,
				Count: 240,
			},
		},
	},
	{
		Id: 428,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 260,
			},
			{
				Id:    323,
				Count: 260,
			},
		},
	},
	{
		Id: 429,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 280,
			},
			{
				Id:    323,
				Count: 280,
			},
		},
	},
	{
		Id: 430,
		Attrib: []*Attribute{
			{
				Id:    313,
				Count: 300,
			},
			{
				Id:    323,
				Count: 300,
			},
		},
	},
	{
		Id: 431,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 5,
			},
		},
	},
	{
		Id: 432,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 10,
			},
		},
	},
	{
		Id: 433,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 15,
			},
		},
	},
	{
		Id: 434,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 20,
			},
		},
	},
	{
		Id: 435,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 25,
			},
		},
	},
	{
		Id: 436,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 30,
			},
		},
	},
	{
		Id: 437,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 35,
			},
		},
	},
	{
		Id: 438,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 40,
			},
		},
	},
	{
		Id: 439,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 45,
			},
		},
	},
	{
		Id: 440,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 50,
			},
		},
	},
	{
		Id: 441,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 55,
			},
		},
	},
	{
		Id: 442,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 60,
			},
		},
	},
	{
		Id: 443,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 65,
			},
		},
	},
	{
		Id: 444,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 70,
			},
		},
	},
	{
		Id: 445,
		Attrib: []*Attribute{
			{
				Id:    24,
				Count: 80,
			},
		},
	},
	{
		Id: 446,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 5,
			},
		},
	},
	{
		Id: 447,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 10,
			},
		},
	},
	{
		Id: 448,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 15,
			},
		},
	},
	{
		Id: 449,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 20,
			},
		},
	},
	{
		Id: 450,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 25,
			},
		},
	},
	{
		Id: 451,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 30,
			},
		},
	},
	{
		Id: 452,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 35,
			},
		},
	},
	{
		Id: 453,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 40,
			},
		},
	},
	{
		Id: 454,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 45,
			},
		},
	},
	{
		Id: 455,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 50,
			},
		},
	},
	{
		Id: 456,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 55,
			},
		},
	},
	{
		Id: 457,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 60,
			},
		},
	},
	{
		Id: 458,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 65,
			},
		},
	},
	{
		Id: 459,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 70,
			},
		},
	},
	{
		Id: 460,
		Attrib: []*Attribute{
			{
				Id:    34,
				Count: 80,
			},
		},
	},
	{
		Id: 461,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 5,
			},
		},
	},
	{
		Id: 462,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 10,
			},
		},
	},
	{
		Id: 463,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 15,
			},
		},
	},
	{
		Id: 464,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 20,
			},
		},
	},
	{
		Id: 465,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 25,
			},
		},
	},
	{
		Id: 466,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 30,
			},
		},
	},
	{
		Id: 467,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 35,
			},
		},
	},
	{
		Id: 468,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 40,
			},
		},
	},
	{
		Id: 469,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 45,
			},
		},
	},
	{
		Id: 470,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 50,
			},
		},
	},
	{
		Id: 471,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 55,
			},
		},
	},
	{
		Id: 472,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 60,
			},
		},
	},
	{
		Id: 473,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 65,
			},
		},
	},
	{
		Id: 474,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 70,
			},
		},
	},
	{
		Id: 475,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: 80,
			},
		},
	},
	{
		Id: 476,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			50,
		},
	},
	{
		Id: 477,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			70,
		},
	},
	{
		Id: 478,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			90,
		},
	},
	{
		Id: 479,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			120,
		},
	},
	{
		Id: 480,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			140,
		},
	},
	{
		Id: 481,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			160,
		},
	},
	{
		Id: 482,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			180,
		},
	},
	{
		Id: 483,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			200,
		},
	},
	{
		Id: 484,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			220,
		},
	},
	{
		Id: 485,
		ChangeParam: []int32{
			11005,
			1,
			0,
			0,
			0,
			250,
		},
	},
	{
		Id: 486,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			50,
		},
	},
	{
		Id: 487,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			80,
		},
	},
	{
		Id: 488,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			110,
		},
	},
	{
		Id: 489,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			140,
		},
	},
	{
		Id: 490,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			170,
		},
	},
	{
		Id: 491,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			200,
		},
	},
	{
		Id: 492,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			230,
		},
	},
	{
		Id: 493,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			260,
		},
	},
	{
		Id: 494,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			290,
		},
	},
	{
		Id: 495,
		ChangeParam: []int32{
			11027,
			1,
			1,
			2,
			0,
			320,
		},
	},
	{
		Id: 496,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			30,
		},
	},
	{
		Id: 497,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			45,
		},
	},
	{
		Id: 498,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			60,
		},
	},
	{
		Id: 499,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			75,
		},
	},
	{
		Id: 500,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			90,
		},
	},
	{
		Id: 501,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			110,
		},
	},
	{
		Id: 502,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			125,
		},
	},
	{
		Id: 503,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			140,
		},
	},
	{
		Id: 504,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			160,
		},
	},
	{
		Id: 505,
		ChangeParam: []int32{
			11022,
			1,
			1,
			1,
			0,
			180,
		},
	},
	{
		Id: 506,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			30,
		},
	},
	{
		Id: 507,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			45,
		},
	},
	{
		Id: 508,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			60,
		},
	},
	{
		Id: 509,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			75,
		},
	},
	{
		Id: 510,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			90,
		},
	},
	{
		Id: 511,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			110,
		},
	},
	{
		Id: 512,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			125,
		},
	},
	{
		Id: 513,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			140,
		},
	},
	{
		Id: 514,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			160,
		},
	},
	{
		Id: 515,
		ChangeParam: []int32{
			11024,
			1,
			0,
			0,
			0,
			180,
		},
	},
	{
		Id: 516,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			30,
		},
	},
	{
		Id: 517,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			45,
		},
	},
	{
		Id: 518,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			60,
		},
	},
	{
		Id: 519,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			75,
		},
	},
	{
		Id: 520,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			90,
		},
	},
	{
		Id: 521,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			110,
		},
	},
	{
		Id: 522,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			125,
		},
	},
	{
		Id: 523,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			140,
		},
	},
	{
		Id: 524,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			160,
		},
	},
	{
		Id: 525,
		ChangeParam: []int32{
			11032,
			1,
			0,
			0,
			0,
			180,
		},
	},
	{
		Id: 526,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 10,
			},
		},
	},
	{
		Id: 527,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 20,
			},
		},
	},
	{
		Id: 528,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 30,
			},
		},
	},
	{
		Id: 529,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 40,
			},
		},
	},
	{
		Id: 530,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 50,
			},
		},
	},
	{
		Id: 531,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 60,
			},
		},
	},
	{
		Id: 532,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 70,
			},
		},
	},
	{
		Id: 533,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 80,
			},
		},
	},
	{
		Id: 534,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 90,
			},
		},
	},
	{
		Id: 535,
		Attrib: []*Attribute{
			{
				Id:    193,
				Count: 100,
			},
		},
	},
	{
		Id: 536,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 10,
			},
		},
	},
	{
		Id: 537,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 20,
			},
		},
	},
	{
		Id: 538,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 30,
			},
		},
	},
	{
		Id: 539,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 40,
			},
		},
	},
	{
		Id: 540,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 50,
			},
		},
	},
	{
		Id: 541,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 60,
			},
		},
	},
	{
		Id: 542,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 70,
			},
		},
	},
	{
		Id: 543,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 80,
			},
		},
	},
	{
		Id: 544,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 90,
			},
		},
	},
	{
		Id: 545,
		Attrib: []*Attribute{
			{
				Id:    203,
				Count: 100,
			},
		},
	},
	{
		Id: 546,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 10,
			},
			{
				Id:    283,
				Count: 10,
			},
		},
	},
	{
		Id: 547,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 20,
			},
			{
				Id:    283,
				Count: 20,
			},
		},
	},
	{
		Id: 548,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 30,
			},
			{
				Id:    283,
				Count: 30,
			},
		},
	},
	{
		Id: 549,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 40,
			},
			{
				Id:    283,
				Count: 40,
			},
		},
	},
	{
		Id: 550,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 50,
			},
			{
				Id:    283,
				Count: 50,
			},
		},
	},
	{
		Id: 551,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 60,
			},
			{
				Id:    283,
				Count: 60,
			},
		},
	},
	{
		Id: 552,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 70,
			},
			{
				Id:    283,
				Count: 70,
			},
		},
	},
	{
		Id: 553,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 80,
			},
			{
				Id:    283,
				Count: 80,
			},
		},
	},
	{
		Id: 554,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 90,
			},
			{
				Id:    283,
				Count: 90,
			},
		},
	},
	{
		Id: 555,
		Attrib: []*Attribute{
			{
				Id:    273,
				Count: 100,
			},
			{
				Id:    283,
				Count: 100,
			},
		},
	},
	{
		Id: 556,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 10,
			},
			{
				Id:    243,
				Count: 10,
			},
		},
	},
	{
		Id: 557,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 20,
			},
			{
				Id:    243,
				Count: 20,
			},
		},
	},
	{
		Id: 558,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 30,
			},
			{
				Id:    243,
				Count: 30,
			},
		},
	},
	{
		Id: 559,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 40,
			},
			{
				Id:    243,
				Count: 40,
			},
		},
	},
	{
		Id: 560,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 50,
			},
			{
				Id:    243,
				Count: 50,
			},
		},
	},
	{
		Id: 561,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 60,
			},
			{
				Id:    243,
				Count: 60,
			},
		},
	},
	{
		Id: 562,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 70,
			},
			{
				Id:    243,
				Count: 70,
			},
		},
	},
	{
		Id: 563,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 80,
			},
			{
				Id:    243,
				Count: 80,
			},
		},
	},
	{
		Id: 564,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 90,
			},
			{
				Id:    243,
				Count: 90,
			},
		},
	},
	{
		Id: 565,
		Attrib: []*Attribute{
			{
				Id:    233,
				Count: 100,
			},
			{
				Id:    243,
				Count: 100,
			},
		},
	},
	{
		Id: 566,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 50,
			},
			{
				Id:    223,
				Count: 50,
			},
		},
	},
	{
		Id: 567,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 100,
			},
			{
				Id:    223,
				Count: 100,
			},
		},
	},
	{
		Id: 568,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 150,
			},
			{
				Id:    223,
				Count: 150,
			},
		},
	},
	{
		Id: 569,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 200,
			},
			{
				Id:    223,
				Count: 200,
			},
		},
	},
	{
		Id: 570,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 250,
			},
			{
				Id:    223,
				Count: 250,
			},
		},
	},
	{
		Id: 571,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 300,
			},
			{
				Id:    223,
				Count: 300,
			},
		},
	},
	{
		Id: 572,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 350,
			},
			{
				Id:    223,
				Count: 350,
			},
		},
	},
	{
		Id: 573,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 400,
			},
			{
				Id:    223,
				Count: 400,
			},
		},
	},
	{
		Id: 574,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 450,
			},
			{
				Id:    223,
				Count: 450,
			},
		},
	},
	{
		Id: 575,
		Attrib: []*Attribute{
			{
				Id:    213,
				Count: 500,
			},
			{
				Id:    223,
				Count: 500,
			},
		},
	},
	{
		Id: 576,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -10,
			},
		},
	},
	{
		Id: 577,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -20,
			},
		},
	},
	{
		Id: 578,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -30,
			},
		},
	},
	{
		Id: 579,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -40,
			},
		},
	},
	{
		Id: 580,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -50,
			},
		},
	},
	{
		Id: 581,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -60,
			},
		},
	},
	{
		Id: 582,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -70,
			},
		},
	},
	{
		Id: 583,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -80,
			},
		},
	},
	{
		Id: 584,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -90,
			},
		},
	},
	{
		Id: 585,
		Attrib: []*Attribute{
			{
				Id:    14,
				Count: -100,
			},
		},
	},
	{
		Id: 586,
		SpecialTag: []int32{
			3,
		},
	},
	{
		Id: 587,
		SpecialTag: []int32{
			4,
		},
	},
	{
		Id: 588,
		Attrib: []*Attribute{
			{
				Id:    163,
				Count: 1,
			},
		},
	},
	{
		Id: 589,
		SpecialTag: []int32{
			6,
		},
	},
}

var skillStatusKeys = []int32{
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
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 175, 176, 177, 178, 179, 180,
	181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
	191, 192, 193, 194, 195, 196, 197, 198, 199, 200,
	201, 202, 203, 204, 205, 206, 207, 208, 209, 210,
	211, 212, 213, 214, 215, 216, 217, 218, 219, 220,
	221, 222, 223, 224, 225, 226, 227, 228, 229, 230,
	231, 232, 233, 234, 235, 236, 237, 238, 239, 240,
	241, 242, 243, 244, 245, 246, 247, 248, 249, 250,
	251, 252, 253, 254, 255, 256, 257, 258, 259, 260,
	261, 262, 263, 264, 265, 266, 267, 268, 269, 270,
	271, 272, 273, 274, 275, 276, 277, 278, 279, 280,
	281, 282, 283, 284, 285, 286, 287, 288, 289, 290,
	291, 292, 293, 294, 295, 296, 297, 298, 299, 300,
	301, 302, 303, 304, 305, 306, 307, 308, 309, 310,
	311, 312, 313, 314, 315, 316, 317, 318, 319, 320,
	321, 322, 323, 324, 325, 326, 327, 328, 329, 330,
	331, 332, 333, 334, 335, 336, 337, 338, 339, 340,
	341, 342, 343, 344, 345, 346, 347, 348, 349, 350,
	351, 352, 353, 354, 355, 356, 357, 358, 359, 360,
	361, 362, 363, 364, 365, 366, 367, 368, 369, 370,
	371, 372, 373, 374, 375, 376, 377, 378, 379, 380,
	381, 382, 383, 384, 385, 386, 387, 388, 389, 390,
	391, 392, 393, 394, 395, 396, 397, 398, 399, 400,
	401, 402, 403, 404, 405, 406, 407, 408, 409, 410,
	411, 412, 413, 414, 415, 416, 417, 418, 419, 420,
	421, 422, 423, 424, 425, 426, 427, 428, 429, 430,
	431, 432, 433, 434, 435, 436, 437, 438, 439, 440,
	441, 442, 443, 444, 445, 446, 447, 448, 449, 450,
	451, 452, 453, 454, 455, 456, 457, 458, 459, 460,
	461, 462, 463, 464, 465, 466, 467, 468, 469, 470,
	471, 472, 473, 474, 475, 476, 477, 478, 479, 480,
	481, 482, 483, 484, 485, 486, 487, 488, 489, 490,
	491, 492, 493, 494, 495, 496, 497, 498, 499, 500,
	501, 502, 503, 504, 505, 506, 507, 508, 509, 510,
	511, 512, 513, 514, 515, 516, 517, 518, 519, 520,
	521, 522, 523, 524, 525, 526, 527, 528, 529, 530,
	531, 532, 533, 534, 535, 536, 537, 538, 539, 540,
	541, 542, 543, 544, 545, 546, 547, 548, 549, 550,
	551, 552, 553, 554, 555, 556, 557, 558, 559, 560,
	561, 562, 563, 564, 565, 566, 567, 568, 569, 570,
	571, 572, 573, 574, 575, 576, 577, 578, 579, 580,
	581, 582, 583, 584, 585, 586, 587, 588, 589,
}

func init() {
	SkillStatusData.data = make(map[int32]*SkillStatusCfg)
	for i := 0; i < len(skillStatusKeys); i++ {
		SkillStatusData.data[skillStatusKeys[i]] = skillStatusValues[i]
	}
}
