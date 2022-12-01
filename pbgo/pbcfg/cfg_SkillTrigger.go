package config

type SkillTriggerCfg struct {
	Id          int32
	CondType    int32
	CondParam   []int32
	OccurProb   int32
	OccurCD     int32
	Obj         int32
	Effect      []int32
	EffectParam [][]int32
}

type SkillTriggerTable struct {
	data map[int32]*SkillTriggerCfg
}

var SkillTriggerData = &SkillTriggerTable{
	data: map[int32]*SkillTriggerCfg{},
}

func (table *SkillTriggerTable) Get(id int32) *SkillTriggerCfg {
	return table.data[id]
}

func (table *SkillTriggerTable) GetAll() []int32 {
	return skillTriggerKeys
}

var skillTriggerValues = []*SkillTriggerCfg{
	{
		Id:       1,
		CondType: 4,
		Obj:      1,
		Effect: []int32{
			5,
		},
		EffectParam: [][]int32{
			{
				4,
			},
		},
	},
	{
		Id:        2,
		CondType:  8,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				100,
			},
		},
	},
	{
		Id:        3,
		CondType:  8,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				140,
			},
		},
	},
	{
		Id:        4,
		CondType:  8,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				180,
			},
		},
	},
	{
		Id:        5,
		CondType:  8,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				220,
			},
		},
	},
	{
		Id:        6,
		CondType:  8,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				260,
			},
		},
	},
	{
		Id:        7,
		CondType:  8,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				300,
			},
		},
	},
	{
		Id:        8,
		CondType:  8,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				340,
			},
		},
	},
	{
		Id:        9,
		CondType:  8,
		OccurProb: 150,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				380,
			},
		},
	},
	{
		Id:        10,
		CondType:  8,
		OccurProb: 150,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				420,
			},
		},
	},
	{
		Id:        11,
		CondType:  8,
		OccurProb: 200,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				460,
			},
		},
	},
	{
		Id:        12,
		CondType:  9,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				100,
			},
		},
	},
	{
		Id:        13,
		CondType:  9,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				140,
			},
		},
	},
	{
		Id:        14,
		CondType:  9,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				180,
			},
		},
	},
	{
		Id:        15,
		CondType:  9,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				220,
			},
		},
	},
	{
		Id:        16,
		CondType:  9,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				260,
			},
		},
	},
	{
		Id:        17,
		CondType:  9,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				300,
			},
		},
	},
	{
		Id:        18,
		CondType:  9,
		OccurProb: 100,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				340,
			},
		},
	},
	{
		Id:        19,
		CondType:  9,
		OccurProb: 150,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				380,
			},
		},
	},
	{
		Id:        20,
		CondType:  9,
		OccurProb: 150,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				420,
			},
		},
	},
	{
		Id:        21,
		CondType:  9,
		OccurProb: 200,
		Obj:       1,
		Effect: []int32{
			7,
		},
		EffectParam: [][]int32{
			{
				460,
			},
		},
	},
	{
		Id:        22,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				10,
			},
		},
	},
	{
		Id:        23,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				15,
			},
		},
	},
	{
		Id:        24,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				20,
			},
		},
	},
	{
		Id:        25,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				25,
			},
		},
	},
	{
		Id:        26,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				30,
			},
		},
	},
	{
		Id:        27,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				35,
			},
		},
	},
	{
		Id:        28,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				40,
			},
		},
	},
	{
		Id:        29,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				45,
			},
		},
	},
	{
		Id:        30,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				50,
			},
		},
	},
	{
		Id:        31,
		CondType:  10,
		OccurProb: 50,
		Obj:       1,
		Effect: []int32{
			9,
		},
		EffectParam: [][]int32{
			{
				3,
				55,
			},
		},
	},
	{
		Id:        32,
		CondType:  10,
		OccurProb: 50,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				70,
			},
		},
	},
	{
		Id:        33,
		CondType:  10,
		OccurProb: 100,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				90,
			},
		},
	},
	{
		Id:        34,
		CondType:  10,
		OccurProb: 150,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				110,
			},
		},
	},
	{
		Id:        35,
		CondType:  10,
		OccurProb: 200,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				130,
			},
		},
	},
	{
		Id:        36,
		CondType:  10,
		OccurProb: 250,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				150,
			},
		},
	},
	{
		Id:        37,
		CondType:  10,
		OccurProb: 300,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				170,
			},
		},
	},
	{
		Id:        38,
		CondType:  10,
		OccurProb: 350,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				190,
			},
		},
	},
	{
		Id:        39,
		CondType:  10,
		OccurProb: 400,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				210,
			},
		},
	},
	{
		Id:        40,
		CondType:  10,
		OccurProb: 450,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				230,
			},
		},
	},
	{
		Id:        41,
		CondType:  10,
		OccurProb: 500,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				2,
				250,
			},
		},
	},
	{
		Id:        42,
		CondType:  11,
		OccurProb: 50,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				70,
			},
		},
	},
	{
		Id:        43,
		CondType:  11,
		OccurProb: 100,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				90,
			},
		},
	},
	{
		Id:        44,
		CondType:  11,
		OccurProb: 150,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				110,
			},
		},
	},
	{
		Id:        45,
		CondType:  11,
		OccurProb: 200,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				130,
			},
		},
	},
	{
		Id:        46,
		CondType:  11,
		OccurProb: 250,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				150,
			},
		},
	},
	{
		Id:        47,
		CondType:  11,
		OccurProb: 300,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				170,
			},
		},
	},
	{
		Id:        48,
		CondType:  11,
		OccurProb: 350,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				190,
			},
		},
	},
	{
		Id:        49,
		CondType:  11,
		OccurProb: 400,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				210,
			},
		},
	},
	{
		Id:        50,
		CondType:  11,
		OccurProb: 450,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				230,
			},
		},
	},
	{
		Id:        51,
		CondType:  11,
		OccurProb: 500,
		Obj:       3,
		Effect: []int32{
			3,
		},
		EffectParam: [][]int32{
			{
				1,
				250,
			},
		},
	},
}

var skillTriggerKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51,
}

func init() {
	SkillTriggerData.data = make(map[int32]*SkillTriggerCfg)
	for i := 0; i < len(skillTriggerKeys); i++ {
		SkillTriggerData.data[skillTriggerKeys[i]] = skillTriggerValues[i]
	}
}
