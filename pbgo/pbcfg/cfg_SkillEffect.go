package config

type SkillEffectCfg struct {
	Id                 int32
	SkillObj           []int32
	SkillObjFaction    []int32
	SkillObjOccupation []int32
	Effect             []int32
	EffectParams       [][]int32
	EffectDelay        []int32
	FxTarget           []*Fx
}

type SkillEffectTable struct {
	data map[int32]*SkillEffectCfg
}

var SkillEffectData = &SkillEffectTable{
	data: map[int32]*SkillEffectCfg{},
}

func (table *SkillEffectTable) Get(id int32) *SkillEffectCfg {
	return table.data[id]
}

func (table *SkillEffectTable) GetAll() []int32 {
	return skillEffectKeys
}

var skillEffectValues = []*SkillEffectCfg{
	{
		Id: 1100011,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
	},
	{
		Id: 1100012,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
	},
	{
		Id: 1110011,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
	},
	{
		Id: 1110012,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
	},
	{
		Id: 10010011,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			1,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 10000101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 2,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				1,
				10000,
			},
		},
	},
	{
		Id: 1,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				3,
				10000,
			},
		},
	},
	{
		Id: 10000201,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 1100101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				35000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100102,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				40000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100103,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				45000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100104,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				50000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100105,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				55000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100106,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				60000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100107,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				65000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100108,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				70000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100109,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				75000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100110,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				80000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1100201,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				5,
				0,
			},
		},
	},
	{
		Id: 1100202,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				6,
				0,
			},
		},
	},
	{
		Id: 1100203,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				7,
				0,
			},
		},
	},
	{
		Id: 1100204,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				8,
				0,
			},
		},
	},
	{
		Id: 1100205,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				9,
				0,
			},
		},
	},
	{
		Id: 1100301,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				65,
				0,
			},
		},
	},
	{
		Id: 1100302,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				66,
				0,
			},
		},
	},
	{
		Id: 1100303,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				67,
				0,
			},
		},
	},
	{
		Id: 1100304,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				68,
				0,
			},
		},
	},
	{
		Id: 1100305,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				69,
				0,
			},
		},
	},
	{
		Id: 1100306,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				70,
				0,
			},
		},
	},
	{
		Id: 1100307,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				71,
				0,
			},
		},
	},
	{
		Id: 1100308,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				72,
				0,
			},
		},
	},
	{
		Id: 1100309,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				73,
				0,
			},
		},
	},
	{
		Id: 1100310,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				74,
				0,
			},
		},
	},
	{
		Id: 1100401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				10,
				0,
			},
		},
	},
	{
		Id: 1100402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				11,
				0,
			},
		},
	},
	{
		Id: 1100403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				12,
				0,
			},
		},
	},
	{
		Id: 1100404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				13,
				0,
			},
		},
	},
	{
		Id: 1100405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				14,
				0,
			},
		},
	},
	{
		Id: 1100501,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
			{
				4500,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			4500,
		},
	},
	{
		Id: 1100502,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
			{
				5000,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100503,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
			{
				5500,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100504,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
			{
				6000,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100505,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
			{
				6500,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100506,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
			{
				7000,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100507,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
			{
				7500,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100508,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
			{
				8000,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100509,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
			{
				8500,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100510,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
			{
				9000,
			},
		},
		EffectDelay: []int32{
			333,
			467,
			600,
			733,
			867,
			1000,
			1333,
			1467,
			1600,
			1733,
			1867,
			2000,
		},
	},
	{
		Id: 1100601,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				15,
				0,
			},
		},
	},
	{
		Id: 1100602,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				16,
				0,
			},
		},
	},
	{
		Id: 1100603,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				17,
				0,
			},
		},
	},
	{
		Id: 1100604,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				18,
				0,
			},
		},
	},
	{
		Id: 1100605,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				19,
				0,
			},
		},
	},
	{
		Id: 1100606,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				20,
				0,
			},
		},
	},
	{
		Id: 1100607,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				21,
				0,
			},
		},
	},
	{
		Id: 1100608,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				22,
				0,
			},
		},
	},
	{
		Id: 1100609,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				23,
				0,
			},
		},
	},
	{
		Id: 1100610,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				24,
				0,
			},
		},
	},
	{
		Id: 1100701,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				25,
				10000,
			},
		},
	},
	{
		Id: 1100702,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				26,
				10000,
			},
		},
	},
	{
		Id: 1100703,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				27,
				10000,
			},
		},
	},
	{
		Id: 1100704,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				28,
				10000,
			},
		},
	},
	{
		Id: 1100705,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				29,
				10000,
			},
		},
	},
	{
		Id: 1100706,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				30,
				10000,
			},
		},
	},
	{
		Id: 1100707,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				31,
				10000,
			},
		},
	},
	{
		Id: 1100708,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				32,
				10000,
			},
		},
	},
	{
		Id: 1100709,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				33,
				10000,
			},
		},
	},
	{
		Id: 1100710,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				34,
				10000,
			},
		},
	},
	{
		Id: 1100801,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-50,
				0,
			},
			{
				476,
				0,
			},
		},
	},
	{
		Id: 1100802,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-100,
				0,
			},
			{
				477,
				0,
			},
		},
	},
	{
		Id: 1100803,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-150,
				0,
			},
			{
				478,
				0,
			},
		},
	},
	{
		Id: 1100804,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-200,
				0,
			},
			{
				479,
				0,
			},
		},
	},
	{
		Id: 1100805,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-250,
				0,
			},
			{
				480,
				0,
			},
		},
	},
	{
		Id: 1100806,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-300,
				0,
			},
			{
				481,
				0,
			},
		},
	},
	{
		Id: 1100807,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-350,
				0,
			},
			{
				482,
				0,
			},
		},
	},
	{
		Id: 1100808,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-400,
				0,
			},
			{
				483,
				0,
			},
		},
	},
	{
		Id: 1100809,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-450,
				0,
			},
			{
				484,
				0,
			},
		},
	},
	{
		Id: 1100810,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11005,
				2,
				0,
				-500,
				0,
			},
			{
				485,
				0,
			},
		},
	},
	{
		Id: 1100901,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				35,
				10000,
			},
		},
	},
	{
		Id: 1100902,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				36,
				10000,
			},
		},
	},
	{
		Id: 1100903,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				37,
				10000,
			},
		},
	},
	{
		Id: 1100904,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				38,
				10000,
			},
		},
	},
	{
		Id: 1100905,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				39,
				10000,
			},
		},
	},
	{
		Id: 1100906,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				40,
				10000,
			},
		},
	},
	{
		Id: 1100907,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				41,
				10000,
			},
		},
	},
	{
		Id: 1100908,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				42,
				10000,
			},
		},
	},
	{
		Id: 1100909,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				43,
				10000,
			},
		},
	},
	{
		Id: 1100910,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				44,
				10000,
			},
		},
	},
	{
		Id: 1100911,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			6,
			6,
		},
		EffectParams: [][]int32{
			{
				1,
			},
			{
				3,
			},
		},
	},
	{
		Id: 1101001,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				45,
				0,
			},
		},
	},
	{
		Id: 1101002,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				46,
				0,
			},
		},
	},
	{
		Id: 1101003,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				47,
				0,
			},
		},
	},
	{
		Id: 1101004,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				48,
				0,
			},
		},
	},
	{
		Id: 1101005,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				49,
				0,
			},
		},
	},
	{
		Id: 1101101,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				50,
				0,
			},
		},
	},
	{
		Id: 1101102,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				51,
				0,
			},
		},
	},
	{
		Id: 1101103,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				52,
				0,
			},
		},
	},
	{
		Id: 1101104,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				53,
				0,
			},
		},
	},
	{
		Id: 1101105,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				54,
				0,
			},
		},
	},
	{
		Id: 1101201,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				75,
				0,
			},
		},
	},
	{
		Id: 1101202,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				76,
				0,
			},
		},
	},
	{
		Id: 1101203,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				77,
				0,
			},
		},
	},
	{
		Id: 1101204,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				78,
				0,
			},
		},
	},
	{
		Id: 1101205,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				79,
				0,
			},
		},
	},
	{
		Id: 1101206,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				80,
				0,
			},
		},
	},
	{
		Id: 1101207,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				81,
				0,
			},
		},
	},
	{
		Id: 1101208,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				82,
				0,
			},
		},
	},
	{
		Id: 1101209,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				83,
				0,
			},
		},
	},
	{
		Id: 1101210,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				84,
				0,
			},
		},
	},
	{
		Id: 1101301,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				50000,
			},
			{
				55,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101302,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				55000,
			},
			{
				56,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101303,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				60000,
			},
			{
				57,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101304,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				65000,
			},
			{
				58,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101305,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				70000,
			},
			{
				59,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101306,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				75000,
			},
			{
				60,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101307,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				80000,
			},
			{
				61,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101308,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				85000,
			},
			{
				62,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101309,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				90000,
			},
			{
				63,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101310,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			4,
		},
		EffectParams: [][]int32{
			{
				95000,
			},
			{
				64,
				5000,
			},
		},
		EffectDelay: []int32{
			667,
			0,
		},
	},
	{
		Id: 1101401,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				12000,
			},
			{
				12000,
			},
			{
				12000,
			},
			{
				12000,
			},
			{
				12000,
			},
			{
				12000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101402,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				13000,
			},
			{
				13000,
			},
			{
				13000,
			},
			{
				13000,
			},
			{
				13000,
			},
			{
				13000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101403,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				14000,
			},
			{
				14000,
			},
			{
				14000,
			},
			{
				14000,
			},
			{
				14000,
			},
			{
				14000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101404,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				15000,
			},
			{
				15000,
			},
			{
				15000,
			},
			{
				15000,
			},
			{
				15000,
			},
			{
				15000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101405,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				16000,
			},
			{
				16000,
			},
			{
				16000,
			},
			{
				16000,
			},
			{
				16000,
			},
			{
				16000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101406,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				17000,
			},
			{
				17000,
			},
			{
				17000,
			},
			{
				17000,
			},
			{
				17000,
			},
			{
				17000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101407,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				18000,
			},
			{
				18000,
			},
			{
				18000,
			},
			{
				18000,
			},
			{
				18000,
			},
			{
				18000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101408,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				19000,
			},
			{
				19000,
			},
			{
				19000,
			},
			{
				19000,
			},
			{
				19000,
			},
			{
				19000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101409,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				20000,
			},
			{
				20000,
			},
			{
				20000,
			},
			{
				20000,
			},
			{
				20000,
			},
			{
				20000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101410,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
			1,
			1,
			1,
			1,
			1,
		},
		EffectParams: [][]int32{
			{
				21000,
			},
			{
				21000,
			},
			{
				21000,
			},
			{
				21000,
			},
			{
				21000,
			},
			{
				21000,
			},
		},
		EffectDelay: []int32{
			233,
			400,
			733,
			867,
			1067,
			1167,
		},
	},
	{
		Id: 1101411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101416,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101417,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101418,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101419,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1101420,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				85,
				1900,
			},
		},
	},
	{
		Id: 1102101,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				86,
				0,
			},
		},
	},
	{
		Id: 1102102,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				87,
				0,
			},
		},
	},
	{
		Id: 1102103,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				88,
				0,
			},
		},
	},
	{
		Id: 1102104,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				89,
				0,
			},
		},
	},
	{
		Id: 1102105,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				90,
				0,
			},
		},
	},
	{
		Id: 1102201,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				35000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102202,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				40000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102203,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				45000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102204,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				50000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102205,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				55000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102206,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				60000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102207,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				65000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102208,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				70000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102209,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				75000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102210,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				80000,
			},
		},
		EffectDelay: []int32{
			867,
		},
	},
	{
		Id: 1102301,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				91,
				0,
			},
		},
	},
	{
		Id: 1102302,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				92,
				0,
			},
		},
	},
	{
		Id: 1102303,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				93,
				0,
			},
		},
	},
	{
		Id: 1102304,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				94,
				0,
			},
		},
	},
	{
		Id: 1102305,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				95,
				0,
			},
		},
	},
	{
		Id: 1102306,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				96,
				0,
			},
		},
	},
	{
		Id: 1102307,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				97,
				0,
			},
		},
	},
	{
		Id: 1102308,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				98,
				0,
			},
		},
	},
	{
		Id: 1102309,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				99,
				0,
			},
		},
	},
	{
		Id: 1102310,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				100,
				0,
			},
		},
	},
	{
		Id: 1102401,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				27000,
			},
			{
				27000,
			},
			{
				27000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102402,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				30000,
			},
			{
				30000,
			},
			{
				30000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102403,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				33000,
			},
			{
				33000,
			},
			{
				33000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102404,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				36000,
			},
			{
				36000,
			},
			{
				36000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102405,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				39000,
			},
			{
				39000,
			},
			{
				39000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102406,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				42000,
			},
			{
				42000,
			},
			{
				42000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102407,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				45000,
			},
			{
				45000,
			},
			{
				45000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102408,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				48000,
			},
			{
				48000,
			},
			{
				48000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102409,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				51000,
			},
			{
				51000,
			},
			{
				51000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102410,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				54000,
			},
			{
				54000,
			},
			{
				54000,
			},
		},
		EffectDelay: []int32{
			933,
			1167,
			1867,
		},
	},
	{
		Id: 1102501,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				101,
				0,
			},
		},
	},
	{
		Id: 1102502,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				102,
				0,
			},
		},
	},
	{
		Id: 1102503,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				103,
				0,
			},
		},
	},
	{
		Id: 1102504,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				104,
				0,
			},
		},
	},
	{
		Id: 1102505,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				105,
				0,
			},
		},
	},
	{
		Id: 1102601,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				106,
				0,
			},
		},
	},
	{
		Id: 1102602,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				107,
				0,
			},
		},
	},
	{
		Id: 1102603,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				108,
				0,
			},
		},
	},
	{
		Id: 1102604,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				109,
				0,
			},
		},
	},
	{
		Id: 1102605,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				110,
				0,
			},
		},
	},
	{
		Id: 1102606,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				111,
				0,
			},
		},
	},
	{
		Id: 1102607,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				112,
				0,
			},
		},
	},
	{
		Id: 1102608,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				113,
				0,
			},
		},
	},
	{
		Id: 1102609,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				114,
				0,
			},
		},
	},
	{
		Id: 1102610,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				115,
				0,
			},
		},
	},
	{
		Id: 1102701,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				116,
				5000,
			},
		},
	},
	{
		Id: 1102702,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				117,
				5000,
			},
		},
	},
	{
		Id: 1102703,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				118,
				5000,
			},
		},
	},
	{
		Id: 1102704,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				119,
				5000,
			},
		},
	},
	{
		Id: 1102705,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				120,
				5000,
			},
		},
	},
	{
		Id: 1102706,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				121,
				5000,
			},
		},
	},
	{
		Id: 1102707,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				122,
				5000,
			},
		},
	},
	{
		Id: 1102708,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				123,
				5000,
			},
		},
	},
	{
		Id: 1102709,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				124,
				5000,
			},
		},
	},
	{
		Id: 1102710,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				125,
				5000,
			},
		},
	},
	{
		Id: 1102801,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				126,
				0,
			},
		},
	},
	{
		Id: 1102802,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				127,
				0,
			},
		},
	},
	{
		Id: 1102803,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				128,
				0,
			},
		},
	},
	{
		Id: 1102804,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				129,
				0,
			},
		},
	},
	{
		Id: 1102805,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				130,
				0,
			},
		},
	},
	{
		Id: 1102901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-40,
				0,
			},
			{
				486,
				0,
			},
		},
	},
	{
		Id: 1102902,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-60,
				0,
			},
			{
				487,
				0,
			},
		},
	},
	{
		Id: 1102903,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-80,
				0,
			},
			{
				488,
				0,
			},
		},
	},
	{
		Id: 1102904,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-100,
				0,
			},
			{
				489,
				0,
			},
		},
	},
	{
		Id: 1102905,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-120,
				0,
			},
			{
				490,
				0,
			},
		},
	},
	{
		Id: 1102906,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-140,
				0,
			},
			{
				491,
				0,
			},
		},
	},
	{
		Id: 1102907,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-160,
				0,
			},
			{
				492,
				0,
			},
		},
	},
	{
		Id: 1102908,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-180,
				0,
			},
			{
				493,
				0,
			},
		},
	},
	{
		Id: 1102909,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-200,
				0,
			},
			{
				494,
				0,
			},
		},
	},
	{
		Id: 1102910,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
			4,
		},
		EffectParams: [][]int32{
			{
				11027,
				1,
				0,
				-250,
				0,
			},
			{
				495,
				0,
			},
		},
	},
	{
		Id: 1103001,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				131,
				0,
			},
		},
	},
	{
		Id: 1103002,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				132,
				0,
			},
		},
	},
	{
		Id: 1103003,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				133,
				0,
			},
		},
	},
	{
		Id: 1103004,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				134,
				0,
			},
		},
	},
	{
		Id: 1103005,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				135,
				0,
			},
		},
	},
	{
		Id: 1103006,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				136,
				0,
			},
		},
	},
	{
		Id: 1103007,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				137,
				0,
			},
		},
	},
	{
		Id: 1103008,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				138,
				0,
			},
		},
	},
	{
		Id: 1103009,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				139,
				0,
			},
		},
	},
	{
		Id: 1103010,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				140,
				0,
			},
		},
	},
	{
		Id: 1103101,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				141,
				10000,
			},
		},
	},
	{
		Id: 1103102,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				142,
				10000,
			},
		},
	},
	{
		Id: 1103103,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				143,
				10000,
			},
		},
	},
	{
		Id: 1103104,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				144,
				10000,
			},
		},
	},
	{
		Id: 1103105,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				145,
				10000,
			},
		},
	},
	{
		Id: 1103106,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				146,
				10000,
			},
		},
	},
	{
		Id: 1103107,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				147,
				10000,
			},
		},
	},
	{
		Id: 1103108,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				148,
				10000,
			},
		},
	},
	{
		Id: 1103109,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				149,
				10000,
			},
		},
	},
	{
		Id: 1103110,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				150,
				10000,
			},
		},
	},
	{
		Id: 1103201,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				25000,
			},
			{
				25000,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103202,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				27500,
			},
			{
				27500,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103203,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				30000,
			},
			{
				30000,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103204,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				32500,
			},
			{
				32500,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103205,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				35000,
			},
			{
				35000,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103206,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				37500,
			},
			{
				37500,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103207,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				40000,
			},
			{
				40000,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103208,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				42500,
			},
			{
				42500,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103209,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				45000,
			},
			{
				45000,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103210,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
			2,
		},
		EffectParams: [][]int32{
			{
				47500,
			},
			{
				47500,
			},
		},
		EffectDelay: []int32{
			400,
			633,
		},
	},
	{
		Id: 1103211,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				151,
				10000,
			},
		},
	},
	{
		Id: 1103212,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				152,
				10000,
			},
		},
	},
	{
		Id: 1103213,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				153,
				10000,
			},
		},
	},
	{
		Id: 1103214,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				154,
				10000,
			},
		},
	},
	{
		Id: 1103215,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				155,
				10000,
			},
		},
	},
	{
		Id: 1103216,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				156,
				10000,
			},
		},
	},
	{
		Id: 1103217,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				157,
				10000,
			},
		},
	},
	{
		Id: 1103218,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				158,
				10000,
			},
		},
	},
	{
		Id: 1103219,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				159,
				10000,
			},
		},
	},
	{
		Id: 1103220,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				160,
				10000,
			},
		},
	},
	{
		Id: 1103301,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				496,
				0,
			},
			{
				506,
				0,
			},
			{
				516,
				0,
			},
		},
	},
	{
		Id: 1103302,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				497,
				0,
			},
			{
				507,
				0,
			},
			{
				517,
				0,
			},
		},
	},
	{
		Id: 1103303,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				498,
				0,
			},
			{
				508,
				0,
			},
			{
				518,
				0,
			},
		},
	},
	{
		Id: 1103304,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				499,
				0,
			},
			{
				509,
				0,
			},
			{
				519,
				0,
			},
		},
	},
	{
		Id: 1103305,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				500,
				0,
			},
			{
				510,
				0,
			},
			{
				520,
				0,
			},
		},
	},
	{
		Id: 1103306,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				501,
				0,
			},
			{
				511,
				0,
			},
			{
				521,
				0,
			},
		},
	},
	{
		Id: 1103307,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				502,
				0,
			},
			{
				512,
				0,
			},
			{
				522,
				0,
			},
		},
	},
	{
		Id: 1103308,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				503,
				0,
			},
			{
				513,
				0,
			},
			{
				523,
				0,
			},
		},
	},
	{
		Id: 1103309,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				504,
				0,
			},
			{
				514,
				0,
			},
			{
				524,
				0,
			},
		},
	},
	{
		Id: 1103310,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
			4,
			4,
		},
		EffectParams: [][]int32{
			{
				505,
				0,
			},
			{
				515,
				0,
			},
			{
				525,
				0,
			},
		},
	},
	{
		Id: 1103401,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				62000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103402,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				68000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103403,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				74000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103404,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				80000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103405,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				86000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103406,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				92000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103407,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				98000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103408,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				104000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103409,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				110000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103410,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				116000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 1103411,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				161,
				10000,
			},
		},
	},
	{
		Id: 1103412,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				162,
				10000,
			},
		},
	},
	{
		Id: 1103413,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				163,
				10000,
			},
		},
	},
	{
		Id: 1103414,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				164,
				10000,
			},
		},
	},
	{
		Id: 1103415,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				165,
				10000,
			},
		},
	},
	{
		Id: 1103416,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				166,
				10000,
			},
		},
	},
	{
		Id: 1103417,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				167,
				10000,
			},
		},
	},
	{
		Id: 1103418,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				168,
				10000,
			},
		},
	},
	{
		Id: 1103419,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				169,
				10000,
			},
		},
	},
	{
		Id: 1103420,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				170,
				10000,
			},
		},
	},
	{
		Id: 1104101,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				173,
				0,
			},
		},
	},
	{
		Id: 1104102,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				174,
				0,
			},
		},
	},
	{
		Id: 1104103,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				175,
				0,
			},
		},
	},
	{
		Id: 1104104,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				176,
				0,
			},
		},
	},
	{
		Id: 1104105,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				177,
				0,
			},
		},
	},
	{
		Id: 1104201,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				178,
				0,
			},
		},
	},
	{
		Id: 1104202,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				179,
				0,
			},
		},
	},
	{
		Id: 1104203,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				180,
				0,
			},
		},
	},
	{
		Id: 1104204,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				181,
				0,
			},
		},
	},
	{
		Id: 1104205,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				182,
				0,
			},
		},
	},
	{
		Id: 1104301,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				203,
				0,
			},
		},
	},
	{
		Id: 1104302,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				204,
				0,
			},
		},
	},
	{
		Id: 1104303,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				205,
				0,
			},
		},
	},
	{
		Id: 1104304,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				206,
				0,
			},
		},
	},
	{
		Id: 1104305,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				207,
				0,
			},
		},
	},
	{
		Id: 1104306,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				208,
				0,
			},
		},
	},
	{
		Id: 1104307,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				209,
				0,
			},
		},
	},
	{
		Id: 1104308,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				210,
				0,
			},
		},
	},
	{
		Id: 1104309,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				211,
				0,
			},
		},
	},
	{
		Id: 1104310,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				212,
				0,
			},
		},
	},
	{
		Id: 1104401,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				183,
				0,
			},
		},
	},
	{
		Id: 1104402,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				184,
				0,
			},
		},
	},
	{
		Id: 1104403,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				185,
				0,
			},
		},
	},
	{
		Id: 1104404,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				186,
				0,
			},
		},
	},
	{
		Id: 1104405,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				187,
				0,
			},
		},
	},
	{
		Id: 1104501,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				188,
				0,
			},
		},
	},
	{
		Id: 1104502,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				189,
				0,
			},
		},
	},
	{
		Id: 1104503,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				190,
				0,
			},
		},
	},
	{
		Id: 1104504,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				191,
				0,
			},
		},
	},
	{
		Id: 1104505,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				192,
				0,
			},
		},
	},
	{
		Id: 1104601,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				193,
				0,
			},
		},
	},
	{
		Id: 1104602,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				194,
				0,
			},
		},
	},
	{
		Id: 1104603,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				195,
				0,
			},
		},
	},
	{
		Id: 1104604,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				196,
				0,
			},
		},
	},
	{
		Id: 1104605,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				197,
				0,
			},
		},
	},
	{
		Id: 1104606,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				198,
				0,
			},
		},
	},
	{
		Id: 1104607,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				199,
				0,
			},
		},
	},
	{
		Id: 1104608,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				200,
				0,
			},
		},
	},
	{
		Id: 1104609,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				201,
				0,
			},
		},
	},
	{
		Id: 1104610,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				202,
				0,
			},
		},
	},
	{
		Id: 1104701,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				213,
				0,
			},
		},
	},
	{
		Id: 1104702,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				214,
				0,
			},
		},
	},
	{
		Id: 1104703,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				215,
				0,
			},
		},
	},
	{
		Id: 1104704,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				216,
				0,
			},
		},
	},
	{
		Id: 1104705,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				217,
				0,
			},
		},
	},
	{
		Id: 1104801,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				218,
				0,
			},
		},
	},
	{
		Id: 1104802,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				219,
				0,
			},
		},
	},
	{
		Id: 1104803,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				220,
				0,
			},
		},
	},
	{
		Id: 1104804,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				221,
				0,
			},
		},
	},
	{
		Id: 1104805,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				222,
				0,
			},
		},
	},
	{
		Id: 1104901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				263,
				0,
			},
			{
				273,
				0,
			},
		},
	},
	{
		Id: 1104902,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				264,
				0,
			},
			{
				274,
				0,
			},
		},
	},
	{
		Id: 1104903,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				265,
				0,
			},
			{
				275,
				0,
			},
		},
	},
	{
		Id: 1104904,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				266,
				0,
			},
			{
				276,
				0,
			},
		},
	},
	{
		Id: 1104905,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				267,
				0,
			},
			{
				277,
				0,
			},
		},
	},
	{
		Id: 1104906,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				268,
				0,
			},
			{
				278,
				0,
			},
		},
	},
	{
		Id: 1104907,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				269,
				0,
			},
			{
				279,
				0,
			},
		},
	},
	{
		Id: 1104908,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				270,
				0,
			},
			{
				280,
				0,
			},
		},
	},
	{
		Id: 1104909,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				271,
				0,
			},
			{
				281,
				0,
			},
		},
	},
	{
		Id: 1104910,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				272,
				0,
			},
			{
				282,
				0,
			},
		},
	},
	{
		Id: 1105001,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				223,
				0,
			},
		},
	},
	{
		Id: 1105002,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				224,
				0,
			},
		},
	},
	{
		Id: 1105003,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				225,
				0,
			},
		},
	},
	{
		Id: 1105004,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				226,
				0,
			},
		},
	},
	{
		Id: 1105005,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				227,
				0,
			},
		},
	},
	{
		Id: 1105006,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				228,
				0,
			},
		},
	},
	{
		Id: 1105007,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				229,
				0,
			},
		},
	},
	{
		Id: 1105008,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				230,
				0,
			},
		},
	},
	{
		Id: 1105009,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				231,
				0,
			},
		},
	},
	{
		Id: 1105010,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				232,
				0,
			},
		},
	},
	{
		Id: 1105101,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				248,
				0,
			},
		},
	},
	{
		Id: 1105102,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				249,
				0,
			},
		},
	},
	{
		Id: 1105103,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				250,
				0,
			},
		},
	},
	{
		Id: 1105104,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				251,
				0,
			},
		},
	},
	{
		Id: 1105105,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				252,
				0,
			},
		},
	},
	{
		Id: 1105201,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				253,
				0,
			},
		},
	},
	{
		Id: 1105202,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				254,
				0,
			},
		},
	},
	{
		Id: 1105203,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				255,
				0,
			},
		},
	},
	{
		Id: 1105204,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				256,
				0,
			},
		},
	},
	{
		Id: 1105205,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				257,
				0,
			},
		},
	},
	{
		Id: 1105206,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				258,
				0,
			},
		},
	},
	{
		Id: 1105207,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				259,
				0,
			},
		},
	},
	{
		Id: 1105208,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				260,
				0,
			},
		},
	},
	{
		Id: 1105209,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				261,
				0,
			},
		},
	},
	{
		Id: 1105210,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				262,
				0,
			},
		},
	},
	{
		Id: 1105301,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				233,
				0,
			},
		},
	},
	{
		Id: 1105302,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				234,
				0,
			},
		},
	},
	{
		Id: 1105303,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				235,
				0,
			},
		},
	},
	{
		Id: 1105304,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				236,
				0,
			},
		},
	},
	{
		Id: 1105305,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				237,
				0,
			},
		},
	},
	{
		Id: 1105306,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				238,
				0,
			},
		},
	},
	{
		Id: 1105307,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				239,
				0,
			},
		},
	},
	{
		Id: 1105308,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				240,
				0,
			},
		},
	},
	{
		Id: 1105309,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				241,
				0,
			},
		},
	},
	{
		Id: 1105310,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				242,
				0,
			},
		},
	},
	{
		Id: 1105401,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				243,
				0,
			},
		},
	},
	{
		Id: 1105402,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				244,
				0,
			},
		},
	},
	{
		Id: 1105403,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				245,
				0,
			},
		},
	},
	{
		Id: 1105404,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				246,
				0,
			},
		},
	},
	{
		Id: 1105405,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				247,
				0,
			},
		},
	},
	{
		Id: 10021101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10121101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10222101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10321101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 10421101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10421401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				416,
				0,
			},
		},
	},
	{
		Id: 10421402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				417,
				0,
			},
		},
	},
	{
		Id: 10421403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				418,
				0,
			},
		},
	},
	{
		Id: 10421404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				419,
				0,
			},
		},
	},
	{
		Id: 10421405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				420,
				0,
			},
		},
	},
	{
		Id: 10421406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				421,
				0,
			},
		},
	},
	{
		Id: 10421407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				422,
				0,
			},
		},
	},
	{
		Id: 10421408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				423,
				0,
			},
		},
	},
	{
		Id: 10421409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				424,
				0,
			},
		},
	},
	{
		Id: 10421410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				425,
				0,
			},
		},
	},
	{
		Id: 10421411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				426,
				0,
			},
		},
	},
	{
		Id: 10421412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				427,
				0,
			},
		},
	},
	{
		Id: 10421413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				428,
				0,
			},
		},
	},
	{
		Id: 10421414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				429,
				0,
			},
		},
	},
	{
		Id: 10421415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				430,
				0,
			},
		},
	},
	{
		Id: 10521101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 10622101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 10622401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				431,
				0,
			},
		},
	},
	{
		Id: 10622402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				432,
				0,
			},
		},
	},
	{
		Id: 10622403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				433,
				0,
			},
		},
	},
	{
		Id: 10622404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				434,
				0,
			},
		},
	},
	{
		Id: 10622405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				435,
				0,
			},
		},
	},
	{
		Id: 10622406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				436,
				0,
			},
		},
	},
	{
		Id: 10622407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				437,
				0,
			},
		},
	},
	{
		Id: 10622408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				438,
				0,
			},
		},
	},
	{
		Id: 10622409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				439,
				0,
			},
		},
	},
	{
		Id: 10622410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				440,
				0,
			},
		},
	},
	{
		Id: 10622411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				441,
				0,
			},
		},
	},
	{
		Id: 10622412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				442,
				0,
			},
		},
	},
	{
		Id: 10622413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				443,
				0,
			},
		},
	},
	{
		Id: 10622414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				444,
				0,
			},
		},
	},
	{
		Id: 10622415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				445,
				0,
			},
		},
	},
	{
		Id: 10721101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			400,
		},
	},
	{
		Id: 10822101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			567,
		},
	},
	{
		Id: 10921101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 11021101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 11021401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				446,
				0,
			},
		},
	},
	{
		Id: 11021402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				447,
				0,
			},
		},
	},
	{
		Id: 11021403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				448,
				0,
			},
		},
	},
	{
		Id: 11021404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				449,
				0,
			},
		},
	},
	{
		Id: 11021405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				450,
				0,
			},
		},
	},
	{
		Id: 11021406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				451,
				0,
			},
		},
	},
	{
		Id: 11021407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				452,
				0,
			},
		},
	},
	{
		Id: 11021408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				453,
				0,
			},
		},
	},
	{
		Id: 11021409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				454,
				0,
			},
		},
	},
	{
		Id: 11021410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				455,
				0,
			},
		},
	},
	{
		Id: 11021411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				456,
				0,
			},
		},
	},
	{
		Id: 11021412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				457,
				0,
			},
		},
	},
	{
		Id: 11021413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				458,
				0,
			},
		},
	},
	{
		Id: 11021414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				459,
				0,
			},
		},
	},
	{
		Id: 11021415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				460,
				0,
			},
		},
	},
	{
		Id: 11122101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 11222101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 11222401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				461,
				0,
			},
		},
	},
	{
		Id: 11222402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				462,
				0,
			},
		},
	},
	{
		Id: 11222403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				463,
				0,
			},
		},
	},
	{
		Id: 11222404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				464,
				0,
			},
		},
	},
	{
		Id: 11222405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				465,
				0,
			},
		},
	},
	{
		Id: 11222406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				466,
				0,
			},
		},
	},
	{
		Id: 11222407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				467,
				0,
			},
		},
	},
	{
		Id: 11222408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				468,
				0,
			},
		},
	},
	{
		Id: 11222409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				469,
				0,
			},
		},
	},
	{
		Id: 11222410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				470,
				0,
			},
		},
	},
	{
		Id: 11222411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				471,
				0,
			},
		},
	},
	{
		Id: 11222412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				472,
				0,
			},
		},
	},
	{
		Id: 11222413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				473,
				0,
			},
		},
	},
	{
		Id: 11222414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				474,
				0,
			},
		},
	},
	{
		Id: 11222415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				475,
				0,
			},
		},
	},
	{
		Id: 11322101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			333,
		},
	},
	{
		Id: 11421101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 11511101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 11612101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 11712101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 11812101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 11911101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			400,
		},
	},
	{
		Id: 12011101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 12112101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 12212101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			767,
		},
	},
	{
		Id: 12312101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 12412101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 12511101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 12611101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			667,
		},
	},
	{
		Id: 12711101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 12811101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			733,
		},
	},
	{
		Id: 12911101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			567,
		},
	},
	{
		Id: 10000301,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				171,
				0,
			},
		},
	},
	{
		Id: 10000302,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				172,
				0,
			},
		},
	},
	{
		Id: 10000303,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				287,
				0,
			},
		},
	},
	{
		Id: 10000304,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				288,
				0,
			},
		},
	},
	{
		Id: 10000305,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				289,
				0,
			},
		},
	},
	{
		Id: 10000306,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				290,
				0,
			},
		},
	},
	{
		Id: 10000307,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				291,
				0,
			},
		},
	},
	{
		Id: 10000308,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				292,
				0,
			},
		},
	},
	{
		Id: 10000309,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				293,
				0,
			},
		},
	},
	{
		Id: 10000310,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				294,
				0,
			},
		},
	},
	{
		Id: 10000311,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				295,
				0,
			},
		},
	},
	{
		Id: 10000312,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				296,
				0,
			},
		},
	},
	{
		Id: 10000313,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				297,
				0,
			},
		},
	},
	{
		Id: 10000314,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				298,
				0,
			},
		},
	},
	{
		Id: 10000315,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				299,
				0,
			},
		},
	},
	{
		Id: 10000401,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				356,
				0,
			},
		},
	},
	{
		Id: 10000402,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				357,
				0,
			},
		},
	},
	{
		Id: 10000403,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				358,
				0,
			},
		},
	},
	{
		Id: 10000404,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				359,
				0,
			},
		},
	},
	{
		Id: 10000405,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				360,
				0,
			},
		},
	},
	{
		Id: 10000406,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				361,
				0,
			},
		},
	},
	{
		Id: 10000407,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				362,
				0,
			},
		},
	},
	{
		Id: 10000408,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				363,
				0,
			},
		},
	},
	{
		Id: 10000409,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				364,
				0,
			},
		},
	},
	{
		Id: 10000410,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				365,
				0,
			},
		},
	},
	{
		Id: 10000411,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				366,
				0,
			},
		},
	},
	{
		Id: 10000412,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				367,
				0,
			},
		},
	},
	{
		Id: 10000413,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				368,
				0,
			},
		},
	},
	{
		Id: 10000414,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				369,
				0,
			},
		},
	},
	{
		Id: 10000415,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				370,
				0,
			},
		},
	},
	{
		Id: 10000501,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				371,
				0,
			},
		},
	},
	{
		Id: 10000502,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				372,
				0,
			},
		},
	},
	{
		Id: 10000503,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				373,
				0,
			},
		},
	},
	{
		Id: 10000504,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				374,
				0,
			},
		},
	},
	{
		Id: 10000505,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				375,
				0,
			},
		},
	},
	{
		Id: 10000506,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				376,
				0,
			},
		},
	},
	{
		Id: 10000507,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				377,
				0,
			},
		},
	},
	{
		Id: 10000508,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				378,
				0,
			},
		},
	},
	{
		Id: 10000509,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				379,
				0,
			},
		},
	},
	{
		Id: 10000510,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				380,
				0,
			},
		},
	},
	{
		Id: 10000511,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				381,
				0,
			},
		},
	},
	{
		Id: 10000512,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				382,
				0,
			},
		},
	},
	{
		Id: 10000513,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				383,
				0,
			},
		},
	},
	{
		Id: 10000514,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				384,
				0,
			},
		},
	},
	{
		Id: 10000515,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				385,
				0,
			},
		},
	},
	{
		Id: 10000601,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				386,
				0,
			},
		},
	},
	{
		Id: 10000602,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				387,
				0,
			},
		},
	},
	{
		Id: 10000603,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				388,
				0,
			},
		},
	},
	{
		Id: 10000604,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				389,
				0,
			},
		},
	},
	{
		Id: 10000605,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				390,
				0,
			},
		},
	},
	{
		Id: 10000606,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				391,
				0,
			},
		},
	},
	{
		Id: 10000607,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				392,
				0,
			},
		},
	},
	{
		Id: 10000608,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				393,
				0,
			},
		},
	},
	{
		Id: 10000609,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				394,
				0,
			},
		},
	},
	{
		Id: 10000610,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				395,
				0,
			},
		},
	},
	{
		Id: 10000611,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				396,
				0,
			},
		},
	},
	{
		Id: 10000612,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				397,
				0,
			},
		},
	},
	{
		Id: 10000613,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				398,
				0,
			},
		},
	},
	{
		Id: 10000614,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				399,
				0,
			},
		},
	},
	{
		Id: 10000615,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				400,
				0,
			},
		},
	},
	{
		Id: 10000701,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				401,
				0,
			},
		},
	},
	{
		Id: 10000702,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				402,
				0,
			},
		},
	},
	{
		Id: 10000703,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				403,
				0,
			},
		},
	},
	{
		Id: 10000704,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				404,
				0,
			},
		},
	},
	{
		Id: 10000705,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				405,
				0,
			},
		},
	},
	{
		Id: 10000706,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				406,
				0,
			},
		},
	},
	{
		Id: 10000707,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				407,
				0,
			},
		},
	},
	{
		Id: 10000708,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				408,
				0,
			},
		},
	},
	{
		Id: 10000709,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				409,
				0,
			},
		},
	},
	{
		Id: 10000710,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				410,
				0,
			},
		},
	},
	{
		Id: 10000711,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				411,
				0,
			},
		},
	},
	{
		Id: 10000712,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				412,
				0,
			},
		},
	},
	{
		Id: 10000713,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				413,
				0,
			},
		},
	},
	{
		Id: 10000714,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				414,
				0,
			},
		},
	},
	{
		Id: 10000715,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				415,
				0,
			},
		},
	},
	{
		Id: 10021401,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				283,
				30000,
			},
		},
	},
	{
		Id: 10021402,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				300,
				30000,
			},
		},
	},
	{
		Id: 10021403,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				301,
				30000,
			},
		},
	},
	{
		Id: 10021404,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				302,
				30000,
			},
		},
	},
	{
		Id: 10021405,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				303,
				30000,
			},
		},
	},
	{
		Id: 10021406,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				304,
				30000,
			},
		},
	},
	{
		Id: 10021407,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				305,
				30000,
			},
		},
	},
	{
		Id: 10021408,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				306,
				30000,
			},
		},
	},
	{
		Id: 10021409,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				307,
				30000,
			},
		},
	},
	{
		Id: 10021410,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				308,
				30000,
			},
		},
	},
	{
		Id: 10021411,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				309,
				30000,
			},
		},
	},
	{
		Id: 10021412,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				310,
				30000,
			},
		},
	},
	{
		Id: 10021413,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				311,
				30000,
			},
		},
	},
	{
		Id: 10021414,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				312,
				30000,
			},
		},
	},
	{
		Id: 10021415,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				313,
				30000,
			},
		},
	},
	{
		Id: 10121401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				500,
				30000,
			},
		},
	},
	{
		Id: 10121402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				1000,
				30000,
			},
		},
	},
	{
		Id: 10121403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				1500,
				30000,
			},
		},
	},
	{
		Id: 10121404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				2000,
				30000,
			},
		},
	},
	{
		Id: 10121405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				2500,
				30000,
			},
		},
	},
	{
		Id: 10121406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				3000,
				30000,
			},
		},
	},
	{
		Id: 10121407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				3500,
				30000,
			},
		},
	},
	{
		Id: 10121408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				4000,
				30000,
			},
		},
	},
	{
		Id: 10121409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				4500,
				30000,
			},
		},
	},
	{
		Id: 10121410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				5000,
				30000,
			},
		},
	},
	{
		Id: 10121411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				5500,
				30000,
			},
		},
	},
	{
		Id: 10121412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				6000,
				30000,
			},
		},
	},
	{
		Id: 10121413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				6500,
				30000,
			},
		},
	},
	{
		Id: 10121414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				7000,
				30000,
			},
		},
	},
	{
		Id: 10121415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				1,
				8000,
				30000,
			},
		},
	},
	{
		Id: 10222401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-500,
			},
		},
	},
	{
		Id: 10222402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-570,
			},
		},
	},
	{
		Id: 10222403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-600,
			},
		},
	},
	{
		Id: 10222404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-620,
			},
		},
	},
	{
		Id: 10222405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-650,
			},
		},
	},
	{
		Id: 10222406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-680,
			},
		},
	},
	{
		Id: 10222407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-700,
			},
		},
	},
	{
		Id: 10222408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-730,
			},
		},
	},
	{
		Id: 10222409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-750,
			},
		},
	},
	{
		Id: 10222410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-780,
			},
		},
	},
	{
		Id: 10222411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-800,
			},
		},
	},
	{
		Id: 10222412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-830,
			},
		},
	},
	{
		Id: 10222413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-850,
			},
		},
	},
	{
		Id: 10222414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-880,
			},
		},
	},
	{
		Id: 10222415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			8,
		},
		EffectParams: [][]int32{
			{
				-900,
			},
		},
	},
	{
		Id: 10321401,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				284,
				30000,
			},
		},
	},
	{
		Id: 10321402,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				314,
				30000,
			},
		},
	},
	{
		Id: 10321403,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				315,
				30000,
			},
		},
	},
	{
		Id: 10321404,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				316,
				30000,
			},
		},
	},
	{
		Id: 10321405,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				317,
				30000,
			},
		},
	},
	{
		Id: 10321406,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				318,
				30000,
			},
		},
	},
	{
		Id: 10321407,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				319,
				30000,
			},
		},
	},
	{
		Id: 10321408,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				320,
				30000,
			},
		},
	},
	{
		Id: 10321409,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				321,
				30000,
			},
		},
	},
	{
		Id: 10321410,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				322,
				30000,
			},
		},
	},
	{
		Id: 10321411,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				323,
				30000,
			},
		},
	},
	{
		Id: 10321412,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				324,
				30000,
			},
		},
	},
	{
		Id: 10321413,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				325,
				30000,
			},
		},
	},
	{
		Id: 10321414,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				326,
				30000,
			},
		},
	},
	{
		Id: 10321415,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				327,
				30000,
			},
		},
	},
	{
		Id: 10521401,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				285,
				30000,
			},
		},
	},
	{
		Id: 10521402,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				328,
				30000,
			},
		},
	},
	{
		Id: 10521403,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				329,
				30000,
			},
		},
	},
	{
		Id: 10521404,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				330,
				30000,
			},
		},
	},
	{
		Id: 10521405,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				331,
				30000,
			},
		},
	},
	{
		Id: 10521406,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				332,
				30000,
			},
		},
	},
	{
		Id: 10521407,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				333,
				30000,
			},
		},
	},
	{
		Id: 10521408,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				334,
				30000,
			},
		},
	},
	{
		Id: 10521409,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				335,
				30000,
			},
		},
	},
	{
		Id: 10521410,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				336,
				30000,
			},
		},
	},
	{
		Id: 10521411,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				337,
				30000,
			},
		},
	},
	{
		Id: 10521412,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				338,
				30000,
			},
		},
	},
	{
		Id: 10521413,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				339,
				30000,
			},
		},
	},
	{
		Id: 10521414,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				340,
				30000,
			},
		},
	},
	{
		Id: 10521415,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				341,
				30000,
			},
		},
	},
	{
		Id: 10822401,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				286,
				30000,
			},
		},
	},
	{
		Id: 10822402,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				342,
				30000,
			},
		},
	},
	{
		Id: 10822403,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				343,
				30000,
			},
		},
	},
	{
		Id: 10822404,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				344,
				30000,
			},
		},
	},
	{
		Id: 10822405,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				345,
				30000,
			},
		},
	},
	{
		Id: 10822406,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				346,
				30000,
			},
		},
	},
	{
		Id: 10822407,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				347,
				30000,
			},
		},
	},
	{
		Id: 10822408,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				348,
				30000,
			},
		},
	},
	{
		Id: 10822409,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				349,
				30000,
			},
		},
	},
	{
		Id: 10822410,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				350,
				30000,
			},
		},
	},
	{
		Id: 10822411,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				351,
				30000,
			},
		},
	},
	{
		Id: 10822412,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				352,
				30000,
			},
		},
	},
	{
		Id: 10822413,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				353,
				30000,
			},
		},
	},
	{
		Id: 10822414,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				354,
				30000,
			},
		},
	},
	{
		Id: 10822415,
		SkillObj: []int32{
			2,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				355,
				30000,
			},
		},
	},
	{
		Id: 10181901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10013901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			767,
		},
	},
	{
		Id: 10161901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			667,
		},
	},
	{
		Id: 10023901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			667,
		},
	},
	{
		Id: 10043901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 10151901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			400,
		},
	},
	{
		Id: 10191901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10093901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			400,
		},
	},
	{
		Id: 10172901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			567,
		},
	},
	{
		Id: 10211901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10251901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			700,
		},
	},
	{
		Id: 10033901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			633,
		},
	},
	{
		Id: 10053901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10063901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10073901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			800,
		},
	},
	{
		Id: 10083901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			567,
		},
	},
	{
		Id: 10103901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 10112901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 10131901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 10142901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
	},
	{
		Id: 10202901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			567,
		},
	},
	{
		Id: 10221901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			433,
		},
	},
	{
		Id: 10232901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			633,
		},
	},
	{
		Id: 10241901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 10262901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			667,
		},
	},
	{
		Id: 10271901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			667,
		},
	},
	{
		Id: 10281901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10292901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10301901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			600,
		},
	},
	{
		Id: 10311901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			733,
		},
	},
	{
		Id: 10322901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 10331901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			400,
		},
	},
	{
		Id: 10341901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 10352901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10361901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			433,
		},
	},
	{
		Id: 10371901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			433,
		},
	},
	{
		Id: 10382901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			467,
		},
	},
	{
		Id: 10391901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			533,
		},
	},
	{
		Id: 10401901,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			2,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			500,
		},
	},
	{
		Id: 11001001,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			1,
		},
		EffectParams: [][]int32{
			{
				1000,
			},
		},
		EffectDelay: []int32{
			767,
		},
	},
	{
		Id: 90100101,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				100,
				0,
			},
		},
	},
	{
		Id: 90100102,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				200,
				0,
			},
		},
	},
	{
		Id: 90100103,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				300,
				0,
			},
		},
	},
	{
		Id: 90100104,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				400,
				0,
			},
		},
	},
	{
		Id: 90100105,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				500,
				0,
			},
		},
	},
	{
		Id: 90100106,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				600,
				0,
			},
		},
	},
	{
		Id: 90100107,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				700,
				0,
			},
		},
	},
	{
		Id: 90100108,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				800,
				0,
			},
		},
	},
	{
		Id: 90100109,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				900,
				0,
			},
		},
	},
	{
		Id: 90100110,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				3,
				1000,
				0,
			},
		},
	},
	{
		Id: 90100201,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				100,
				0,
			},
		},
	},
	{
		Id: 90100202,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				200,
				0,
			},
		},
	},
	{
		Id: 90100203,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				300,
				0,
			},
		},
	},
	{
		Id: 90100204,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				400,
				0,
			},
		},
	},
	{
		Id: 90100205,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				500,
				0,
			},
		},
	},
	{
		Id: 90100206,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				600,
				0,
			},
		},
	},
	{
		Id: 90100207,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				700,
				0,
			},
		},
	},
	{
		Id: 90100208,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				800,
				0,
			},
		},
	},
	{
		Id: 90100209,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				900,
				0,
			},
		},
	},
	{
		Id: 90100210,
		SkillObj: []int32{
			8,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			11,
		},
		EffectParams: [][]int32{
			{
				2,
				1000,
				0,
			},
		},
	},
	{
		Id: 90100301,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				526,
				0,
			},
		},
	},
	{
		Id: 90100302,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				527,
				0,
			},
		},
	},
	{
		Id: 90100303,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				528,
				0,
			},
		},
	},
	{
		Id: 90100304,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				529,
				0,
			},
		},
	},
	{
		Id: 90100305,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				530,
				0,
			},
		},
	},
	{
		Id: 90100306,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				531,
				0,
			},
		},
	},
	{
		Id: 90100307,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				532,
				0,
			},
		},
	},
	{
		Id: 90100308,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				533,
				0,
			},
		},
	},
	{
		Id: 90100309,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				534,
				0,
			},
		},
	},
	{
		Id: 90100310,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				535,
				0,
			},
		},
	},
	{
		Id: 90100401,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				576,
				0,
			},
		},
	},
	{
		Id: 90100402,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				577,
				0,
			},
		},
	},
	{
		Id: 90100403,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				578,
				0,
			},
		},
	},
	{
		Id: 90100404,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				579,
				0,
			},
		},
	},
	{
		Id: 90100405,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				580,
				0,
			},
		},
	},
	{
		Id: 90100406,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				581,
				0,
			},
		},
	},
	{
		Id: 90100407,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				582,
				0,
			},
		},
	},
	{
		Id: 90100408,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				583,
				0,
			},
		},
	},
	{
		Id: 90100409,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				584,
				0,
			},
		},
	},
	{
		Id: 90100410,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				585,
				0,
			},
		},
	},
	{
		Id: 90100501,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				536,
				0,
			},
		},
	},
	{
		Id: 90100502,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				537,
				0,
			},
		},
	},
	{
		Id: 90100503,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				538,
				0,
			},
		},
	},
	{
		Id: 90100504,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				539,
				0,
			},
		},
	},
	{
		Id: 90100505,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				540,
				0,
			},
		},
	},
	{
		Id: 90100506,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				541,
				0,
			},
		},
	},
	{
		Id: 90100507,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				542,
				0,
			},
		},
	},
	{
		Id: 90100508,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				543,
				0,
			},
		},
	},
	{
		Id: 90100509,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				544,
				0,
			},
		},
	},
	{
		Id: 90100510,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				545,
				0,
			},
		},
	},
	{
		Id: 90100601,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				546,
				0,
			},
		},
	},
	{
		Id: 90100602,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				547,
				0,
			},
		},
	},
	{
		Id: 90100603,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				548,
				0,
			},
		},
	},
	{
		Id: 90100604,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				549,
				0,
			},
		},
	},
	{
		Id: 90100605,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				550,
				0,
			},
		},
	},
	{
		Id: 90100606,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				551,
				0,
			},
		},
	},
	{
		Id: 90100607,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				552,
				0,
			},
		},
	},
	{
		Id: 90100608,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				553,
				0,
			},
		},
	},
	{
		Id: 90100609,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				554,
				0,
			},
		},
	},
	{
		Id: 90100610,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				555,
				0,
			},
		},
	},
	{
		Id: 90100701,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				556,
				0,
			},
		},
	},
	{
		Id: 90100702,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				557,
				0,
			},
		},
	},
	{
		Id: 90100703,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				558,
				0,
			},
		},
	},
	{
		Id: 90100704,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				559,
				0,
			},
		},
	},
	{
		Id: 90100705,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				560,
				0,
			},
		},
	},
	{
		Id: 90100706,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				561,
				0,
			},
		},
	},
	{
		Id: 90100707,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				562,
				0,
			},
		},
	},
	{
		Id: 90100708,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				563,
				0,
			},
		},
	},
	{
		Id: 90100709,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				564,
				0,
			},
		},
	},
	{
		Id: 90100710,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				565,
				0,
			},
		},
	},
	{
		Id: 90100801,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				566,
				0,
			},
		},
	},
	{
		Id: 90100802,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				567,
				0,
			},
		},
	},
	{
		Id: 90100803,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				568,
				0,
			},
		},
	},
	{
		Id: 90100804,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				569,
				0,
			},
		},
	},
	{
		Id: 90100805,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				570,
				0,
			},
		},
	},
	{
		Id: 90100806,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				571,
				0,
			},
		},
	},
	{
		Id: 90100807,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				572,
				0,
			},
		},
	},
	{
		Id: 90100808,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				573,
				0,
			},
		},
	},
	{
		Id: 90100809,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				574,
				0,
			},
		},
	},
	{
		Id: 90100810,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				575,
				0,
			},
		},
	},
	{
		Id: 90100901,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				10,
				0,
			},
			{
				2,
				0,
				10,
				0,
			},
		},
	},
	{
		Id: 90100902,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				20,
				0,
			},
			{
				2,
				0,
				20,
				0,
			},
		},
	},
	{
		Id: 90100903,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				30,
				0,
			},
			{
				2,
				0,
				30,
				0,
			},
		},
	},
	{
		Id: 90100904,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				40,
				0,
			},
			{
				2,
				0,
				40,
				0,
			},
		},
	},
	{
		Id: 90100905,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				50,
				0,
			},
			{
				2,
				0,
				50,
				0,
			},
		},
	},
	{
		Id: 90100906,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				60,
				0,
			},
			{
				2,
				0,
				60,
				0,
			},
		},
	},
	{
		Id: 90100907,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				70,
				0,
			},
			{
				2,
				0,
				70,
				0,
			},
		},
	},
	{
		Id: 90100908,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				80,
				0,
			},
			{
				2,
				0,
				80,
				0,
			},
		},
	},
	{
		Id: 90100909,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				90,
				0,
			},
			{
				2,
				0,
				90,
				0,
			},
		},
	},
	{
		Id: 90100910,
		SkillObj: []int32{
			3,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			12,
			12,
		},
		EffectParams: [][]int32{
			{
				1,
				0,
				100,
				0,
			},
			{
				2,
				0,
				100,
				0,
			},
		},
	},
	{
		Id: 90101001,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				20,
				0,
			},
		},
	},
	{
		Id: 90101002,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				40,
				0,
			},
		},
	},
	{
		Id: 90101003,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				60,
				0,
			},
		},
	},
	{
		Id: 90101004,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				80,
				0,
			},
		},
	},
	{
		Id: 90101005,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				100,
				0,
			},
		},
	},
	{
		Id: 90101006,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				120,
				0,
			},
		},
	},
	{
		Id: 90101007,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				140,
				0,
			},
		},
	},
	{
		Id: 90101008,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				160,
				0,
			},
		},
	},
	{
		Id: 90101009,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				180,
				0,
			},
		},
	},
	{
		Id: 90101010,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			13,
		},
		EffectParams: [][]int32{
			{
				0,
				200,
				0,
			},
		},
	},
	{
		Id: 9000101,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				586,
				0,
			},
		},
	},
	{
		Id: 9000201,
		SkillObj: []int32{
			20,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				587,
				0,
			},
		},
	},
	{
		Id: 9000301,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				588,
				0,
			},
		},
	},
	{
		Id: 9000401,
		SkillObj: []int32{
			4,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			4,
		},
		EffectParams: [][]int32{
			{
				589,
				0,
			},
		},
	},
	{
		Id: 9000501,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
		},
		EffectParams: [][]int32{
			{
				0,
				1,
				0,
				-900,
				0,
			},
		},
	},
	{
		Id: 9000601,
		SkillObj: []int32{
			1,
		},
		SkillObjFaction: []int32{
			15,
		},
		SkillObjOccupation: []int32{
			15,
		},
		Effect: []int32{
			10,
		},
		EffectParams: [][]int32{
			{
				0,
				2,
				0,
				-900,
				0,
			},
		},
	},
}

var skillEffectKeys = []int32{
	1100011, 1100012, 1110011, 1110012, 10010011, 10000101, 2, 1, 10000201, 1100101,
	1100102, 1100103, 1100104, 1100105, 1100106, 1100107, 1100108, 1100109, 1100110, 1100201,
	1100202, 1100203, 1100204, 1100205, 1100301, 1100302, 1100303, 1100304, 1100305, 1100306,
	1100307, 1100308, 1100309, 1100310, 1100401, 1100402, 1100403, 1100404, 1100405, 1100501,
	1100502, 1100503, 1100504, 1100505, 1100506, 1100507, 1100508, 1100509, 1100510, 1100601,
	1100602, 1100603, 1100604, 1100605, 1100606, 1100607, 1100608, 1100609, 1100610, 1100701,
	1100702, 1100703, 1100704, 1100705, 1100706, 1100707, 1100708, 1100709, 1100710, 1100801,
	1100802, 1100803, 1100804, 1100805, 1100806, 1100807, 1100808, 1100809, 1100810, 1100901,
	1100902, 1100903, 1100904, 1100905, 1100906, 1100907, 1100908, 1100909, 1100910, 1100911,
	1101001, 1101002, 1101003, 1101004, 1101005, 1101101, 1101102, 1101103, 1101104, 1101105,
	1101201, 1101202, 1101203, 1101204, 1101205, 1101206, 1101207, 1101208, 1101209, 1101210,
	1101301, 1101302, 1101303, 1101304, 1101305, 1101306, 1101307, 1101308, 1101309, 1101310,
	1101401, 1101402, 1101403, 1101404, 1101405, 1101406, 1101407, 1101408, 1101409, 1101410,
	1101411, 1101412, 1101413, 1101414, 1101415, 1101416, 1101417, 1101418, 1101419, 1101420,
	1102101, 1102102, 1102103, 1102104, 1102105, 1102201, 1102202, 1102203, 1102204, 1102205,
	1102206, 1102207, 1102208, 1102209, 1102210, 1102301, 1102302, 1102303, 1102304, 1102305,
	1102306, 1102307, 1102308, 1102309, 1102310, 1102401, 1102402, 1102403, 1102404, 1102405,
	1102406, 1102407, 1102408, 1102409, 1102410, 1102501, 1102502, 1102503, 1102504, 1102505,
	1102601, 1102602, 1102603, 1102604, 1102605, 1102606, 1102607, 1102608, 1102609, 1102610,
	1102701, 1102702, 1102703, 1102704, 1102705, 1102706, 1102707, 1102708, 1102709, 1102710,
	1102801, 1102802, 1102803, 1102804, 1102805, 1102901, 1102902, 1102903, 1102904, 1102905,
	1102906, 1102907, 1102908, 1102909, 1102910, 1103001, 1103002, 1103003, 1103004, 1103005,
	1103006, 1103007, 1103008, 1103009, 1103010, 1103101, 1103102, 1103103, 1103104, 1103105,
	1103106, 1103107, 1103108, 1103109, 1103110, 1103201, 1103202, 1103203, 1103204, 1103205,
	1103206, 1103207, 1103208, 1103209, 1103210, 1103211, 1103212, 1103213, 1103214, 1103215,
	1103216, 1103217, 1103218, 1103219, 1103220, 1103301, 1103302, 1103303, 1103304, 1103305,
	1103306, 1103307, 1103308, 1103309, 1103310, 1103401, 1103402, 1103403, 1103404, 1103405,
	1103406, 1103407, 1103408, 1103409, 1103410, 1103411, 1103412, 1103413, 1103414, 1103415,
	1103416, 1103417, 1103418, 1103419, 1103420, 1104101, 1104102, 1104103, 1104104, 1104105,
	1104201, 1104202, 1104203, 1104204, 1104205, 1104301, 1104302, 1104303, 1104304, 1104305,
	1104306, 1104307, 1104308, 1104309, 1104310, 1104401, 1104402, 1104403, 1104404, 1104405,
	1104501, 1104502, 1104503, 1104504, 1104505, 1104601, 1104602, 1104603, 1104604, 1104605,
	1104606, 1104607, 1104608, 1104609, 1104610, 1104701, 1104702, 1104703, 1104704, 1104705,
	1104801, 1104802, 1104803, 1104804, 1104805, 1104901, 1104902, 1104903, 1104904, 1104905,
	1104906, 1104907, 1104908, 1104909, 1104910, 1105001, 1105002, 1105003, 1105004, 1105005,
	1105006, 1105007, 1105008, 1105009, 1105010, 1105101, 1105102, 1105103, 1105104, 1105105,
	1105201, 1105202, 1105203, 1105204, 1105205, 1105206, 1105207, 1105208, 1105209, 1105210,
	1105301, 1105302, 1105303, 1105304, 1105305, 1105306, 1105307, 1105308, 1105309, 1105310,
	1105401, 1105402, 1105403, 1105404, 1105405, 10021101, 10121101, 10222101, 10321101, 10421101,
	10421401, 10421402, 10421403, 10421404, 10421405, 10421406, 10421407, 10421408, 10421409, 10421410,
	10421411, 10421412, 10421413, 10421414, 10421415, 10521101, 10622101, 10622401, 10622402, 10622403,
	10622404, 10622405, 10622406, 10622407, 10622408, 10622409, 10622410, 10622411, 10622412, 10622413,
	10622414, 10622415, 10721101, 10822101, 10921101, 11021101, 11021401, 11021402, 11021403, 11021404,
	11021405, 11021406, 11021407, 11021408, 11021409, 11021410, 11021411, 11021412, 11021413, 11021414,
	11021415, 11122101, 11222101, 11222401, 11222402, 11222403, 11222404, 11222405, 11222406, 11222407,
	11222408, 11222409, 11222410, 11222411, 11222412, 11222413, 11222414, 11222415, 11322101, 11421101,
	11511101, 11612101, 11712101, 11812101, 11911101, 12011101, 12112101, 12212101, 12312101, 12412101,
	12511101, 12611101, 12711101, 12811101, 12911101, 10000301, 10000302, 10000303, 10000304, 10000305,
	10000306, 10000307, 10000308, 10000309, 10000310, 10000311, 10000312, 10000313, 10000314, 10000315,
	10000401, 10000402, 10000403, 10000404, 10000405, 10000406, 10000407, 10000408, 10000409, 10000410,
	10000411, 10000412, 10000413, 10000414, 10000415, 10000501, 10000502, 10000503, 10000504, 10000505,
	10000506, 10000507, 10000508, 10000509, 10000510, 10000511, 10000512, 10000513, 10000514, 10000515,
	10000601, 10000602, 10000603, 10000604, 10000605, 10000606, 10000607, 10000608, 10000609, 10000610,
	10000611, 10000612, 10000613, 10000614, 10000615, 10000701, 10000702, 10000703, 10000704, 10000705,
	10000706, 10000707, 10000708, 10000709, 10000710, 10000711, 10000712, 10000713, 10000714, 10000715,
	10021401, 10021402, 10021403, 10021404, 10021405, 10021406, 10021407, 10021408, 10021409, 10021410,
	10021411, 10021412, 10021413, 10021414, 10021415, 10121401, 10121402, 10121403, 10121404, 10121405,
	10121406, 10121407, 10121408, 10121409, 10121410, 10121411, 10121412, 10121413, 10121414, 10121415,
	10222401, 10222402, 10222403, 10222404, 10222405, 10222406, 10222407, 10222408, 10222409, 10222410,
	10222411, 10222412, 10222413, 10222414, 10222415, 10321401, 10321402, 10321403, 10321404, 10321405,
	10321406, 10321407, 10321408, 10321409, 10321410, 10321411, 10321412, 10321413, 10321414, 10321415,
	10521401, 10521402, 10521403, 10521404, 10521405, 10521406, 10521407, 10521408, 10521409, 10521410,
	10521411, 10521412, 10521413, 10521414, 10521415, 10822401, 10822402, 10822403, 10822404, 10822405,
	10822406, 10822407, 10822408, 10822409, 10822410, 10822411, 10822412, 10822413, 10822414, 10822415,
	10181901, 10013901, 10161901, 10023901, 10043901, 10151901, 10191901, 10093901, 10172901, 10211901,
	10251901, 10033901, 10053901, 10063901, 10073901, 10083901, 10103901, 10112901, 10131901, 10142901,
	10202901, 10221901, 10232901, 10241901, 10262901, 10271901, 10281901, 10292901, 10301901, 10311901,
	10322901, 10331901, 10341901, 10352901, 10361901, 10371901, 10382901, 10391901, 10401901, 11001001,
	90100101, 90100102, 90100103, 90100104, 90100105, 90100106, 90100107, 90100108, 90100109, 90100110,
	90100201, 90100202, 90100203, 90100204, 90100205, 90100206, 90100207, 90100208, 90100209, 90100210,
	90100301, 90100302, 90100303, 90100304, 90100305, 90100306, 90100307, 90100308, 90100309, 90100310,
	90100401, 90100402, 90100403, 90100404, 90100405, 90100406, 90100407, 90100408, 90100409, 90100410,
	90100501, 90100502, 90100503, 90100504, 90100505, 90100506, 90100507, 90100508, 90100509, 90100510,
	90100601, 90100602, 90100603, 90100604, 90100605, 90100606, 90100607, 90100608, 90100609, 90100610,
	90100701, 90100702, 90100703, 90100704, 90100705, 90100706, 90100707, 90100708, 90100709, 90100710,
	90100801, 90100802, 90100803, 90100804, 90100805, 90100806, 90100807, 90100808, 90100809, 90100810,
	90100901, 90100902, 90100903, 90100904, 90100905, 90100906, 90100907, 90100908, 90100909, 90100910,
	90101001, 90101002, 90101003, 90101004, 90101005, 90101006, 90101007, 90101008, 90101009, 90101010,
	9000101, 9000201, 9000301, 9000401, 9000501, 9000601,
}

func init() {
	SkillEffectData.data = make(map[int32]*SkillEffectCfg)
	for i := 0; i < len(skillEffectKeys); i++ {
		SkillEffectData.data[skillEffectKeys[i]] = skillEffectValues[i]
	}
}
