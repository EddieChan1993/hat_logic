package config

type BuildingLordLvCfg struct {
	Id             int32
	StageLv        int32
	Lv             int32
	LevelUpConsume []*Consume
	LevelUpPercent int32
	MainLevelId    int32
	GetTalent      int32
	AttribAdd      []*AddAttribute
}

type BuildingLordLvTable struct {
	data map[int32]*BuildingLordLvCfg
}

var BuildingLordLvData = &BuildingLordLvTable{
	data: map[int32]*BuildingLordLvCfg{},
}

func (table *BuildingLordLvTable) Get(id int32) *BuildingLordLvCfg {
	return table.data[id]
}

func (table *BuildingLordLvTable) GetAll() []int32 {
	return buildingLordLvKeys
}

var buildingLordLvValues = []*BuildingLordLvCfg{
	{
		Id:      1,
		StageLv: 1,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 14000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 14000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 14000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 14000,
			},
		},
		LevelUpPercent: 1000,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 50,
			},
			{
				Type:  1,
				Id:    42,
				Count: 50,
			},
			{
				Type:  1,
				Id:    52,
				Count: 50,
			},
			{
				Type:  3,
				Id:    22,
				Count: 50,
			},
			{
				Type:  3,
				Id:    32,
				Count: 50,
			},
		},
	},
	{
		Id:      2,
		StageLv: 1,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 15000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 15000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 15000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 15000,
			},
		},
		LevelUpPercent: 1000,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 100,
			},
			{
				Type:  1,
				Id:    42,
				Count: 100,
			},
			{
				Type:  1,
				Id:    52,
				Count: 100,
			},
			{
				Type:  3,
				Id:    22,
				Count: 100,
			},
			{
				Type:  3,
				Id:    32,
				Count: 100,
			},
		},
	},
	{
		Id:      3,
		StageLv: 1,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 17000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 17000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 17000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 17000,
			},
		},
		LevelUpPercent: 1000,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 150,
			},
		},
	},
	{
		Id:      4,
		StageLv: 1,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 18000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 18000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 18000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 18000,
			},
		},
		LevelUpPercent: 1000,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 200,
			},
			{
				Type:  1,
				Id:    42,
				Count: 200,
			},
			{
				Type:  1,
				Id:    52,
				Count: 200,
			},
			{
				Type:  3,
				Id:    22,
				Count: 200,
			},
			{
				Type:  3,
				Id:    32,
				Count: 200,
			},
		},
	},
	{
		Id:      5,
		StageLv: 1,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 32000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 32000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 32000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 32000,
			},
		},
		LevelUpPercent: 1000,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 250,
			},
			{
				Type:  1,
				Id:    42,
				Count: 250,
			},
			{
				Type:  1,
				Id:    52,
				Count: 250,
			},
			{
				Type:  3,
				Id:    22,
				Count: 250,
			},
			{
				Type:  3,
				Id:    32,
				Count: 250,
			},
		},
	},
	{
		Id:      6,
		StageLv: 1,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 34000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 34000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 34000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 34000,
			},
		},
		LevelUpPercent: 1000,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 300,
			},
			{
				Type:  1,
				Id:    42,
				Count: 300,
			},
			{
				Type:  1,
				Id:    52,
				Count: 300,
			},
			{
				Type:  3,
				Id:    22,
				Count: 300,
			},
			{
				Type:  3,
				Id:    32,
				Count: 300,
			},
		},
	},
	{
		Id:      7,
		StageLv: 1,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 36000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 36000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 36000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 36000,
			},
		},
		LevelUpPercent: 950,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 350,
			},
		},
	},
	{
		Id:      8,
		StageLv: 1,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 38000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 38000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 38000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 38000,
			},
		},
		LevelUpPercent: 950,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 400,
			},
			{
				Type:  1,
				Id:    42,
				Count: 400,
			},
			{
				Type:  1,
				Id:    52,
				Count: 400,
			},
			{
				Type:  3,
				Id:    22,
				Count: 400,
			},
			{
				Type:  3,
				Id:    32,
				Count: 400,
			},
		},
	},
	{
		Id:      9,
		StageLv: 1,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 57000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 57000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 57000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 57000,
			},
		},
		LevelUpPercent: 700,
		GetTalent:      5,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 450,
			},
			{
				Type:  1,
				Id:    42,
				Count: 450,
			},
			{
				Type:  1,
				Id:    52,
				Count: 450,
			},
			{
				Type:  3,
				Id:    22,
				Count: 450,
			},
			{
				Type:  3,
				Id:    32,
				Count: 450,
			},
		},
	},
	{
		Id:      10,
		StageLv: 1,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 60000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 60000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 60000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 60000,
			},
		},
		LevelUpPercent: 950,
		MainLevelId:    2010,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 650,
			},
			{
				Type:  1,
				Id:    42,
				Count: 650,
			},
			{
				Type:  1,
				Id:    52,
				Count: 650,
			},
			{
				Type:  3,
				Id:    22,
				Count: 650,
			},
			{
				Type:  3,
				Id:    32,
				Count: 650,
			},
		},
	},
	{
		Id:      11,
		StageLv: 2,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 63000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 63000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 63000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 63000,
			},
		},
		LevelUpPercent: 950,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 750,
			},
		},
	},
	{
		Id:      12,
		StageLv: 2,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 66000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 66000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 66000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 66000,
			},
		},
		LevelUpPercent: 950,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 850,
			},
		},
	},
	{
		Id:      13,
		StageLv: 2,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 86000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 86000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 86000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 86000,
			},
		},
		LevelUpPercent: 900,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 950,
			},
		},
	},
	{
		Id:      14,
		StageLv: 2,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 90000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 90000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 90000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 90000,
			},
		},
		LevelUpPercent: 900,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1050,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1050,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1050,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1050,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1050,
			},
		},
	},
	{
		Id:      15,
		StageLv: 2,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 93000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 93000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 93000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 93000,
			},
		},
		LevelUpPercent: 900,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1150,
			},
		},
	},
	{
		Id:      16,
		StageLv: 2,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 97000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 97000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 97000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 97000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1250,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1250,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1250,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1250,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1250,
			},
		},
	},
	{
		Id:      17,
		StageLv: 2,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 120000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 120000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 120000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 120000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1350,
			},
		},
	},
	{
		Id:      18,
		StageLv: 2,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 125000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 125000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 125000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 125000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1450,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1450,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1450,
			},
		},
	},
	{
		Id:      19,
		StageLv: 2,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 129000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 129000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 129000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 129000,
			},
		},
		LevelUpPercent: 600,
		GetTalent:      7,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1550,
			},
		},
	},
	{
		Id:      20,
		StageLv: 2,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 133000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 133000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 133000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 133000,
			},
		},
		LevelUpPercent: 900,
		MainLevelId:    3010,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 1950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 1950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 1950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 1950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 1950,
			},
		},
	},
	{
		Id:      21,
		StageLv: 3,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 161000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 161000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 161000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 161000,
			},
		},
		LevelUpPercent: 900,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2050,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2050,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2050,
			},
		},
	},
	{
		Id:      22,
		StageLv: 3,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 166000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 166000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 166000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 166000,
			},
		},
		LevelUpPercent: 900,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2150,
			},
		},
	},
	{
		Id:      23,
		StageLv: 3,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 171000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 171000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 171000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 171000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2250,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2250,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2250,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2250,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2250,
			},
		},
	},
	{
		Id:      24,
		StageLv: 3,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 176000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 176000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 176000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 176000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2350,
			},
		},
	},
	{
		Id:      25,
		StageLv: 3,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 207000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 207000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 207000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 207000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2450,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2450,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2450,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2450,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2450,
			},
		},
	},
	{
		Id:      26,
		StageLv: 3,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 213000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 213000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 213000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 213000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2550,
			},
		},
	},
	{
		Id:      27,
		StageLv: 3,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 218000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 218000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 218000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 218000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2650,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2650,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2650,
			},
		},
	},
	{
		Id:      28,
		StageLv: 3,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 224000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 224000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 224000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 224000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2750,
			},
		},
	},
	{
		Id:      29,
		StageLv: 3,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 259000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 259000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 259000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 259000,
			},
		},
		LevelUpPercent: 500,
		GetTalent:      10,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 2850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 2850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 2850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 2850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 2850,
			},
		},
	},
	{
		Id:      30,
		StageLv: 3,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 265000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 265000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 265000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 265000,
			},
		},
		LevelUpPercent: 850,
		MainLevelId:    4010,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 3250,
			},
			{
				Type:  1,
				Id:    42,
				Count: 3250,
			},
			{
				Type:  1,
				Id:    52,
				Count: 3250,
			},
			{
				Type:  3,
				Id:    22,
				Count: 3250,
			},
			{
				Type:  3,
				Id:    32,
				Count: 3250,
			},
		},
	},
	{
		Id:      31,
		StageLv: 4,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 272000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 272000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 272000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 272000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    42,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    52,
				Count: 3400,
			},
			{
				Type:  3,
				Id:    22,
				Count: 3400,
			},
			{
				Type:  3,
				Id:    32,
				Count: 3400,
			},
		},
	},
	{
		Id:      32,
		StageLv: 4,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 278000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 278000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 278000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 278000,
			},
		},
		LevelUpPercent: 850,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 3550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 3550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 3550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 3550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 3550,
			},
		},
	},
	{
		Id:      33,
		StageLv: 4,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 316000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 316000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 316000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 316000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 3700,
			},
			{
				Type:  1,
				Id:    42,
				Count: 3700,
			},
			{
				Type:  1,
				Id:    52,
				Count: 3700,
			},
			{
				Type:  3,
				Id:    22,
				Count: 3700,
			},
			{
				Type:  3,
				Id:    32,
				Count: 3700,
			},
		},
	},
	{
		Id:      34,
		StageLv: 4,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 324000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 324000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 324000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 324000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 3850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 3850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 3850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 3850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 3850,
			},
		},
	},
	{
		Id:      35,
		StageLv: 4,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 331000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 331000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 331000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 331000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    42,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    52,
				Count: 4000,
			},
			{
				Type:  3,
				Id:    22,
				Count: 4000,
			},
			{
				Type:  3,
				Id:    32,
				Count: 4000,
			},
		},
	},
	{
		Id:      36,
		StageLv: 4,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 338000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 338000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 338000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 338000,
			},
		},
		LevelUpPercent: 750,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 4150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 4150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 4150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 4150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 4150,
			},
		},
	},
	{
		Id:      37,
		StageLv: 4,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 380000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 380000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 380000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 380000,
			},
		},
		LevelUpPercent: 750,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 4300,
			},
			{
				Type:  1,
				Id:    42,
				Count: 4300,
			},
			{
				Type:  1,
				Id:    52,
				Count: 4300,
			},
			{
				Type:  3,
				Id:    22,
				Count: 4300,
			},
			{
				Type:  3,
				Id:    32,
				Count: 4300,
			},
		},
	},
	{
		Id:      38,
		StageLv: 4,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 388000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 388000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 388000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 388000,
			},
		},
		LevelUpPercent: 750,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 4450,
			},
			{
				Type:  1,
				Id:    42,
				Count: 4450,
			},
			{
				Type:  1,
				Id:    52,
				Count: 4450,
			},
			{
				Type:  3,
				Id:    22,
				Count: 4450,
			},
			{
				Type:  3,
				Id:    32,
				Count: 4450,
			},
		},
	},
	{
		Id:      39,
		StageLv: 4,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 396000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 396000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 396000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 396000,
			},
		},
		LevelUpPercent: 400,
		GetTalent:      15,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 4600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 4600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 4600,
			},
		},
	},
	{
		Id:      40,
		StageLv: 4,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 403000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 403000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 403000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 403000,
			},
		},
		LevelUpPercent: 800,
		MainLevelId:    5010,
		GetTalent:      50,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    42,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    52,
				Count: 5200,
			},
			{
				Type:  3,
				Id:    22,
				Count: 5200,
			},
			{
				Type:  3,
				Id:    32,
				Count: 5200,
			},
		},
	},
	{
		Id:      41,
		StageLv: 5,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 449000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 449000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 449000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 449000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 5350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 5350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 5350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 5350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 5350,
			},
		},
	},
	{
		Id:      42,
		StageLv: 5,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 457000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 457000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 457000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 457000,
			},
		},
		LevelUpPercent: 800,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 5500,
			},
			{
				Type:  1,
				Id:    42,
				Count: 5500,
			},
			{
				Type:  1,
				Id:    52,
				Count: 5500,
			},
			{
				Type:  3,
				Id:    22,
				Count: 5500,
			},
			{
				Type:  3,
				Id:    32,
				Count: 5500,
			},
		},
	},
	{
		Id:      43,
		StageLv: 5,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 466000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 466000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 466000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 466000,
			},
		},
		LevelUpPercent: 750,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 5650,
			},
			{
				Type:  1,
				Id:    42,
				Count: 5650,
			},
			{
				Type:  1,
				Id:    52,
				Count: 5650,
			},
			{
				Type:  3,
				Id:    22,
				Count: 5650,
			},
			{
				Type:  3,
				Id:    32,
				Count: 5650,
			},
		},
	},
	{
		Id:      44,
		StageLv: 5,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 475000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 475000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 475000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 475000,
			},
		},
		LevelUpPercent: 750,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    42,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    52,
				Count: 5800,
			},
			{
				Type:  3,
				Id:    22,
				Count: 5800,
			},
			{
				Type:  3,
				Id:    32,
				Count: 5800,
			},
		},
	},
	{
		Id:      45,
		StageLv: 5,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 524000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 524000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 524000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 524000,
			},
		},
		LevelUpPercent: 750,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 5950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 5950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 5950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 5950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 5950,
			},
		},
	},
	{
		Id:      46,
		StageLv: 5,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 533000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 533000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 533000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 533000,
			},
		},
		LevelUpPercent: 700,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 6100,
			},
			{
				Type:  1,
				Id:    42,
				Count: 6100,
			},
			{
				Type:  1,
				Id:    52,
				Count: 6100,
			},
			{
				Type:  3,
				Id:    22,
				Count: 6100,
			},
			{
				Type:  3,
				Id:    32,
				Count: 6100,
			},
		},
	},
	{
		Id:      47,
		StageLv: 5,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 542000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 542000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 542000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 542000,
			},
		},
		LevelUpPercent: 700,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 6250,
			},
			{
				Type:  1,
				Id:    42,
				Count: 6250,
			},
			{
				Type:  1,
				Id:    52,
				Count: 6250,
			},
			{
				Type:  3,
				Id:    22,
				Count: 6250,
			},
			{
				Type:  3,
				Id:    32,
				Count: 6250,
			},
		},
	},
	{
		Id:      48,
		StageLv: 5,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 552000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 552000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 552000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 552000,
			},
		},
		LevelUpPercent: 700,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    42,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    52,
				Count: 6400,
			},
			{
				Type:  3,
				Id:    22,
				Count: 6400,
			},
			{
				Type:  3,
				Id:    32,
				Count: 6400,
			},
		},
	},
	{
		Id:      49,
		StageLv: 5,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 604000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 604000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 604000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 604000,
			},
		},
		LevelUpPercent: 300,
		GetTalent:      20,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 6550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 6550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 6550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 6550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 6550,
			},
		},
	},
	{
		Id:      50,
		StageLv: 5,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 614000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 614000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 614000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 614000,
			},
		},
		LevelUpPercent: 700,
		MainLevelId:    6010,
		GetTalent:      60,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 7150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 7150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 7150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 7150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 7150,
			},
		},
	},
	{
		Id:      51,
		StageLv: 6,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 624000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 624000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 624000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 624000,
			},
		},
		LevelUpPercent: 700,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 7350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 7350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 7350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 7350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 7350,
			},
		},
	},
	{
		Id:      52,
		StageLv: 6,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 635000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 635000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 635000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 635000,
			},
		},
		LevelUpPercent: 700,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 7550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 7550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 7550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 7550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 7550,
			},
		},
	},
	{
		Id:      53,
		StageLv: 6,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 691000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 691000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 691000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 691000,
			},
		},
		LevelUpPercent: 650,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 7750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 7750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 7750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 7750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 7750,
			},
		},
	},
	{
		Id:      54,
		StageLv: 6,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 701000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 701000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 701000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 701000,
			},
		},
		LevelUpPercent: 650,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 7950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 7950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 7950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 7950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 7950,
			},
		},
	},
	{
		Id:      55,
		StageLv: 6,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 712000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 712000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 712000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 712000,
			},
		},
		LevelUpPercent: 650,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 8150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 8150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 8150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 8150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 8150,
			},
		},
	},
	{
		Id:      56,
		StageLv: 6,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 723000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 723000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 723000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 723000,
			},
		},
		LevelUpPercent: 600,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 8350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 8350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 8350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 8350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 8350,
			},
		},
	},
	{
		Id:      57,
		StageLv: 6,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 783000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 783000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 783000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 783000,
			},
		},
		LevelUpPercent: 600,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 8550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 8550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 8550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 8550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 8550,
			},
		},
	},
	{
		Id:      58,
		StageLv: 6,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 794000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 794000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 794000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 794000,
			},
		},
		LevelUpPercent: 600,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 8750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 8750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 8750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 8750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 8750,
			},
		},
	},
	{
		Id:      59,
		StageLv: 6,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 806000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 806000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 806000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 806000,
			},
		},
		LevelUpPercent: 200,
		GetTalent:      25,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 8950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 8950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 8950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 8950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 8950,
			},
		},
	},
	{
		Id:      60,
		StageLv: 6,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 817000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 817000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 817000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 817000,
			},
		},
		LevelUpPercent: 600,
		MainLevelId:    7010,
		GetTalent:      80,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 9750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 9750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 9750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 9750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 9750,
			},
		},
	},
	{
		Id:      61,
		StageLv: 7,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 881000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 881000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 881000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 881000,
			},
		},
		LevelUpPercent: 600,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 9950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 9950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 9950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 9950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 9950,
			},
		},
	},
	{
		Id:      62,
		StageLv: 7,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 893000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 893000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 893000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 893000,
			},
		},
		LevelUpPercent: 600,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 10150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 10150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 10150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 10150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 10150,
			},
		},
	},
	{
		Id:      63,
		StageLv: 7,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 905000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 905000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 905000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 905000,
			},
		},
		LevelUpPercent: 550,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 10350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 10350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 10350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 10350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 10350,
			},
		},
	},
	{
		Id:      64,
		StageLv: 7,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 917000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 917000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 917000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 917000,
			},
		},
		LevelUpPercent: 550,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 10550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 10550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 10550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 10550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 10550,
			},
		},
	},
	{
		Id:      65,
		StageLv: 7,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 984000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 984000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 984000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 984000,
			},
		},
		LevelUpPercent: 550,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 10750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 10750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 10750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 10750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 10750,
			},
		},
	},
	{
		Id:      66,
		StageLv: 7,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 997000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 997000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 997000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 997000,
			},
		},
		LevelUpPercent: 500,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 10950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 10950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 10950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 10950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 10950,
			},
		},
	},
	{
		Id:      67,
		StageLv: 7,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1010000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1010000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1010000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1010000,
			},
		},
		LevelUpPercent: 500,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 11150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 11150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 11150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 11150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 11150,
			},
		},
	},
	{
		Id:      68,
		StageLv: 7,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1023000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1023000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1023000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1023000,
			},
		},
		LevelUpPercent: 500,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 11350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 11350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 11350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 11350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 11350,
			},
		},
	},
	{
		Id:      69,
		StageLv: 7,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1094000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1094000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1094000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1094000,
			},
		},
		LevelUpPercent: 100,
		GetTalent:      30,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 11550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 11550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 11550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 11550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 11550,
			},
		},
	},
	{
		Id:      70,
		StageLv: 7,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1108000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1108000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1108000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1108000,
			},
		},
		LevelUpPercent: 500,
		MainLevelId:    8010,
		GetTalent:      90,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 12350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 12350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 12350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 12350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 12350,
			},
		},
	},
	{
		Id:      71,
		StageLv: 8,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1121000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1121000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1121000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1121000,
			},
		},
		LevelUpPercent: 500,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 12600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 12600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 12600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 12600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 12600,
			},
		},
	},
	{
		Id:      72,
		StageLv: 8,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1135000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1135000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1135000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1135000,
			},
		},
		LevelUpPercent: 500,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 12850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 12850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 12850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 12850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 12850,
			},
		},
	},
	{
		Id:      73,
		StageLv: 8,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1209000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1209000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1209000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1209000,
			},
		},
		LevelUpPercent: 450,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 13100,
			},
			{
				Type:  1,
				Id:    42,
				Count: 13100,
			},
			{
				Type:  1,
				Id:    52,
				Count: 13100,
			},
			{
				Type:  3,
				Id:    22,
				Count: 13100,
			},
			{
				Type:  3,
				Id:    32,
				Count: 13100,
			},
		},
	},
	{
		Id:      74,
		StageLv: 8,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1224000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1224000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1224000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1224000,
			},
		},
		LevelUpPercent: 450,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 13350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 13350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 13350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 13350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 13350,
			},
		},
	},
	{
		Id:      75,
		StageLv: 8,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1238000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1238000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1238000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1238000,
			},
		},
		LevelUpPercent: 450,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 13600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 13600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 13600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 13600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 13600,
			},
		},
	},
	{
		Id:      76,
		StageLv: 8,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1252000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1252000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1252000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1252000,
			},
		},
		LevelUpPercent: 400,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 13850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 13850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 13850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 13850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 13850,
			},
		},
	},
	{
		Id:      77,
		StageLv: 8,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1330000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1330000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1330000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1330000,
			},
		},
		LevelUpPercent: 400,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 14100,
			},
			{
				Type:  1,
				Id:    42,
				Count: 14100,
			},
			{
				Type:  1,
				Id:    52,
				Count: 14100,
			},
			{
				Type:  3,
				Id:    22,
				Count: 14100,
			},
			{
				Type:  3,
				Id:    32,
				Count: 14100,
			},
		},
	},
	{
		Id:      78,
		StageLv: 8,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1345000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1345000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1345000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1345000,
			},
		},
		LevelUpPercent: 400,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 14350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 14350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 14350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 14350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 14350,
			},
		},
	},
	{
		Id:      79,
		StageLv: 8,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1663000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1663000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1663000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1663000,
			},
		},
		LevelUpPercent: 50,
		GetTalent:      35,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 14600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 14600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 14600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 14600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 14600,
			},
		},
	},
	{
		Id:      80,
		StageLv: 8,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1814000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1814000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1814000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1814000,
			},
		},
		LevelUpPercent: 400,
		MainLevelId:    9010,
		GetTalent:      105,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 15600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 15600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 15600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 15600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 15600,
			},
		},
	},
	{
		Id:      81,
		StageLv: 9,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1900000,
			},
		},
		LevelUpPercent: 400,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 15850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 15850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 15850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 15850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 15850,
			},
		},
	},
	{
		Id:      82,
		StageLv: 9,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1900000,
			},
		},
		LevelUpPercent: 400,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 16100,
			},
			{
				Type:  1,
				Id:    42,
				Count: 16100,
			},
			{
				Type:  1,
				Id:    52,
				Count: 16100,
			},
			{
				Type:  3,
				Id:    22,
				Count: 16100,
			},
			{
				Type:  3,
				Id:    32,
				Count: 16100,
			},
		},
	},
	{
		Id:      83,
		StageLv: 9,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1900000,
			},
		},
		LevelUpPercent: 350,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 16350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 16350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 16350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 16350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 16350,
			},
		},
	},
	{
		Id:      84,
		StageLv: 9,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1900000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1900000,
			},
		},
		LevelUpPercent: 350,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 16600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 16600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 16600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 16600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 16600,
			},
		},
	},
	{
		Id:      85,
		StageLv: 9,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1987000,
			},
		},
		LevelUpPercent: 350,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 16850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 16850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 16850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 16850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 16850,
			},
		},
	},
	{
		Id:      86,
		StageLv: 9,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1987000,
			},
		},
		LevelUpPercent: 300,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 17100,
			},
			{
				Type:  1,
				Id:    42,
				Count: 17100,
			},
			{
				Type:  1,
				Id:    52,
				Count: 17100,
			},
			{
				Type:  3,
				Id:    22,
				Count: 17100,
			},
			{
				Type:  3,
				Id:    32,
				Count: 17100,
			},
		},
	},
	{
		Id:      87,
		StageLv: 9,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1987000,
			},
		},
		LevelUpPercent: 300,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 17350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 17350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 17350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 17350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 17350,
			},
		},
	},
	{
		Id:      88,
		StageLv: 9,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 1987000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 1987000,
			},
		},
		LevelUpPercent: 300,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 17600,
			},
			{
				Type:  1,
				Id:    42,
				Count: 17600,
			},
			{
				Type:  1,
				Id:    52,
				Count: 17600,
			},
			{
				Type:  3,
				Id:    22,
				Count: 17600,
			},
			{
				Type:  3,
				Id:    32,
				Count: 17600,
			},
		},
	},
	{
		Id:      89,
		StageLv: 9,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2070000,
			},
		},
		LevelUpPercent: 40,
		GetTalent:      40,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 17850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 17850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 17850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 17850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 17850,
			},
		},
	},
	{
		Id:      90,
		StageLv: 9,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2070000,
			},
		},
		LevelUpPercent: 300,
		MainLevelId:    10010,
		GetTalent:      120,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 18850,
			},
			{
				Type:  1,
				Id:    42,
				Count: 18850,
			},
			{
				Type:  1,
				Id:    52,
				Count: 18850,
			},
			{
				Type:  3,
				Id:    22,
				Count: 18850,
			},
			{
				Type:  3,
				Id:    32,
				Count: 18850,
			},
		},
	},
	{
		Id:      91,
		StageLv: 10,
		Lv:      1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2070000,
			},
		},
		LevelUpPercent: 300,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 19150,
			},
			{
				Type:  1,
				Id:    42,
				Count: 19150,
			},
			{
				Type:  1,
				Id:    52,
				Count: 19150,
			},
			{
				Type:  3,
				Id:    22,
				Count: 19150,
			},
			{
				Type:  3,
				Id:    32,
				Count: 19150,
			},
		},
	},
	{
		Id:      92,
		StageLv: 10,
		Lv:      2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2070000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2070000,
			},
		},
		LevelUpPercent: 300,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 19450,
			},
			{
				Type:  1,
				Id:    42,
				Count: 19450,
			},
			{
				Type:  1,
				Id:    52,
				Count: 19450,
			},
			{
				Type:  3,
				Id:    22,
				Count: 19450,
			},
			{
				Type:  3,
				Id:    32,
				Count: 19450,
			},
		},
	},
	{
		Id:      93,
		StageLv: 10,
		Lv:      3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2160000,
			},
		},
		LevelUpPercent: 250,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 19750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 19750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 19750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 19750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 19750,
			},
		},
	},
	{
		Id:      94,
		StageLv: 10,
		Lv:      4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2160000,
			},
		},
		LevelUpPercent: 250,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 20050,
			},
			{
				Type:  1,
				Id:    42,
				Count: 20050,
			},
			{
				Type:  1,
				Id:    52,
				Count: 20050,
			},
			{
				Type:  3,
				Id:    22,
				Count: 20050,
			},
			{
				Type:  3,
				Id:    32,
				Count: 20050,
			},
		},
	},
	{
		Id:      95,
		StageLv: 10,
		Lv:      5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2160000,
			},
		},
		LevelUpPercent: 250,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 20350,
			},
			{
				Type:  1,
				Id:    42,
				Count: 20350,
			},
			{
				Type:  1,
				Id:    52,
				Count: 20350,
			},
			{
				Type:  3,
				Id:    22,
				Count: 20350,
			},
			{
				Type:  3,
				Id:    32,
				Count: 20350,
			},
		},
	},
	{
		Id:      96,
		StageLv: 10,
		Lv:      6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2160000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2160000,
			},
		},
		LevelUpPercent: 200,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 20650,
			},
			{
				Type:  1,
				Id:    42,
				Count: 20650,
			},
			{
				Type:  1,
				Id:    52,
				Count: 20650,
			},
			{
				Type:  3,
				Id:    22,
				Count: 20650,
			},
			{
				Type:  3,
				Id:    32,
				Count: 20650,
			},
		},
	},
	{
		Id:      97,
		StageLv: 10,
		Lv:      7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2240000,
			},
		},
		LevelUpPercent: 200,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 20950,
			},
			{
				Type:  1,
				Id:    42,
				Count: 20950,
			},
			{
				Type:  1,
				Id:    52,
				Count: 20950,
			},
			{
				Type:  3,
				Id:    22,
				Count: 20950,
			},
			{
				Type:  3,
				Id:    32,
				Count: 20950,
			},
		},
	},
	{
		Id:      98,
		StageLv: 10,
		Lv:      8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2240000,
			},
		},
		LevelUpPercent: 200,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 21250,
			},
			{
				Type:  1,
				Id:    42,
				Count: 21250,
			},
			{
				Type:  1,
				Id:    52,
				Count: 21250,
			},
			{
				Type:  3,
				Id:    22,
				Count: 21250,
			},
			{
				Type:  3,
				Id:    32,
				Count: 21250,
			},
		},
	},
	{
		Id:      99,
		StageLv: 10,
		Lv:      9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2240000,
			},
		},
		LevelUpPercent: 20,
		GetTalent:      45,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 21550,
			},
			{
				Type:  1,
				Id:    42,
				Count: 21550,
			},
			{
				Type:  1,
				Id:    52,
				Count: 21550,
			},
			{
				Type:  3,
				Id:    22,
				Count: 21550,
			},
			{
				Type:  3,
				Id:    32,
				Count: 21550,
			},
		},
	},
	{
		Id:      100,
		StageLv: 10,
		Lv:      10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    37,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    38,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    39,
				Count: 2240000,
			},
			{
				Type:  1,
				Id:    40,
				Count: 2240000,
			},
		},
		LevelUpPercent: 20,
		GetTalent:      140,
		AttribAdd: []*AddAttribute{
			{
				Type:  1,
				Id:    12,
				Count: 22750,
			},
			{
				Type:  1,
				Id:    42,
				Count: 22750,
			},
			{
				Type:  1,
				Id:    52,
				Count: 22750,
			},
			{
				Type:  3,
				Id:    22,
				Count: 22750,
			},
			{
				Type:  3,
				Id:    32,
				Count: 22750,
			},
		},
	},
}

var buildingLordLvKeys = []int32{
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
	BuildingLordLvData.data = make(map[int32]*BuildingLordLvCfg)
	for i := 0; i < len(buildingLordLvKeys); i++ {
		BuildingLordLvData.data[buildingLordLvKeys[i]] = buildingLordLvValues[i]
	}
}
