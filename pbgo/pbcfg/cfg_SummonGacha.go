package config

type SummonGachaCfg struct {
	Id             int32
	Type           int32
	ItemCost       int32
	ItemCost1      int32
	BuildingLordLv int32
	AddPower       int32
	MmuGuarantee   int32
	RewardPreview  []*Reward
}

type SummonGachaTable struct {
	data                    map[int32]*SummonGachaCfg
	summonGachaTypeIndexMap map[int32][]int32
}

var SummonGachaData = &SummonGachaTable{
	data: map[int32]*SummonGachaCfg{},
	summonGachaTypeIndexMap: map[int32][]int32{
		1: {101, 102, 103, 104, 105, 106, 107, 108, 109, 110},
		2: {201, 202, 203, 204, 205, 206, 207, 208, 209, 210},
		3: {301},
	},
}

func (table *SummonGachaTable) Get(id int32) *SummonGachaCfg {
	return table.data[id]
}

func (table *SummonGachaTable) GetAll() []int32 {
	return summonGachaKeys
}

var summonGachaValues = []*SummonGachaCfg{
	{
		Id:             101,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 1,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    53001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    63001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    73001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    83001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    93001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    103001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    15001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    25001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    35001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    45001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    14001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    24001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    34001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    44001,
				Count: 0,
			},
		},
	},
	{
		Id:             102,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 10,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    153001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    163001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    173001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    183001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    193001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    203001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    115001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    125001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    135001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    145001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    114001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    124001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    134001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    144001,
				Count: 0,
			},
		},
	},
	{
		Id:             103,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 20,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    253001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    263001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    273001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    283001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    293001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    303001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    215001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    225001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    235001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    245001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    214001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    224001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    234001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    244001,
				Count: 0,
			},
		},
	},
	{
		Id:             104,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 30,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    353001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    363001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    373001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    383001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    393001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    403001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    315001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    325001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    335001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    345001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    314001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    324001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    334001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    344001,
				Count: 0,
			},
		},
	},
	{
		Id:             105,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 40,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    453001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    463001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    473001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    483001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    493001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    503001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    415001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    425001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    435001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    445001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    414001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    424001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    434001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    444001,
				Count: 0,
			},
		},
	},
	{
		Id:             106,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 50,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    553001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    563001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    573001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    583001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    593001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    603001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    515001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    525001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    535001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    545001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    514001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    524001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    534001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    544001,
				Count: 0,
			},
		},
	},
	{
		Id:             107,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 60,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    653001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    663001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    673001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    683001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    693001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    703001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    615001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    625001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    635001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    645001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    614001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    624001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    634001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    644001,
				Count: 0,
			},
		},
	},
	{
		Id:             108,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 75,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    753001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    763001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    773001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    783001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    793001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    803001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    715001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    725001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    735001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    745001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    714001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    724001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    734001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    744001,
				Count: 0,
			},
		},
	},
	{
		Id:             109,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 90,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    853001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    863001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    873001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    883001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    893001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    903001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    815001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    825001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    835001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    845001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    814001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    824001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    834001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    844001,
				Count: 0,
			},
		},
	},
	{
		Id:             110,
		Type:           1,
		ItemCost:       1,
		ItemCost1:      150,
		BuildingLordLv: 100,
		AddPower:       5,
		MmuGuarantee:   20,
		RewardPreview: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    953001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    963001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    973001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    983001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    993001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    1003001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    915001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    925001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    935001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    945001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    914001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    924001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    934001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    944001,
				Count: 0,
			},
		},
	},
	{
		Id:             201,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 1,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    15001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    25001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    35001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    45001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    55001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    65001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    75001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    85001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    95001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    105001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    54001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    64001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    74001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    84001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    94001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    104001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             202,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 10,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    115001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    125001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    135001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    145001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    155001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    165001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    175001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    185001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    195001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    205001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    154001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    164001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    174001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    184001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    194001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    204001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             203,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 20,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    215001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    225001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    235001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    245001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    255001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    265001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    275001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    285001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    295001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    305001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    254001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    264001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    274001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    284001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    294001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    304001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             204,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 30,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    315001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    325001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    335001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    345001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    355001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    365001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    375001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    385001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    395001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    405001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    354001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    364001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    374001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    384001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    394001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    404001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             205,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 40,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    415001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    425001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    435001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    445001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    455001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    465001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    475001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    485001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    495001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    505001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    454001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    464001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    474001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    484001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    494001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    504001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             206,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 50,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    515001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    525001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    535001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    545001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    555001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    565001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    575001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    585001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    595001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    605001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    554001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    564001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    574001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    584001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    594001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    604001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             207,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 60,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    615001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    625001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    635001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    645001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    655001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    665001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    675001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    685001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    695001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    705001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    654001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    664001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    674001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    684001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    694001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    704001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             208,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 75,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    715001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    725001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    735001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    745001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    755001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    765001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    775001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    785001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    795001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    805001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    754001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    764001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    774001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    784001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    794001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    804001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             209,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 90,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    815001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    825001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    835001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    845001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    855001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    865001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    875001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    885001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    895001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    905001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    854001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    864001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    874001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    884001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    894001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    904001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             210,
		Type:           2,
		ItemCost:       1,
		BuildingLordLv: 100,
		AddPower:       10,
		MmuGuarantee:   30,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
			{
				Type:  2,
				Id:    915001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    925001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    935001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    945001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    955001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    965001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    975001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    985001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    995001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    1005001,
				Count: 0,
			},
			{
				Type:  4,
				Id:    1,
				Count: 0,
			},
			{
				Type:  4,
				Id:    2,
				Count: 0,
			},
			{
				Type:  4,
				Id:    3,
				Count: 0,
			},
			{
				Type:  4,
				Id:    4,
				Count: 0,
			},
			{
				Type:  4,
				Id:    5,
				Count: 0,
			},
			{
				Type:  4,
				Id:    6,
				Count: 0,
			},
			{
				Type:  4,
				Id:    7,
				Count: 0,
			},
			{
				Type:  4,
				Id:    8,
				Count: 0,
			},
			{
				Type:  4,
				Id:    9,
				Count: 0,
			},
			{
				Type:  4,
				Id:    10,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10721,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10921,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11122,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11322,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11612,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11712,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11812,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11911,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12011,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12112,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12212,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12312,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12412,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12511,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12611,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12711,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12811,
				Count: 0,
			},
			{
				Type:  3,
				Id:    12911,
				Count: 0,
			},
			{
				Type:  2,
				Id:    954001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    964001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    974001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    984001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    994001,
				Count: 0,
			},
			{
				Type:  2,
				Id:    1004001,
				Count: 0,
			},
			{
				Type:  1,
				Id:    2,
				Count: 0,
			},
		},
	},
	{
		Id:             301,
		Type:           3,
		BuildingLordLv: 1,
		RewardPreview: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10121,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10222,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10321,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10421,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10521,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10622,
				Count: 0,
			},
			{
				Type:  3,
				Id:    10822,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11021,
				Count: 0,
			},
			{
				Type:  3,
				Id:    11222,
				Count: 0,
			},
		},
	},
}

var summonGachaKeys = []int32{
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	201, 202, 203, 204, 205, 206, 207, 208, 209, 210,
	301,
}

func init() {
	SummonGachaData.data = make(map[int32]*SummonGachaCfg)
	for i := 0; i < len(summonGachaKeys); i++ {
		SummonGachaData.data[summonGachaKeys[i]] = summonGachaValues[i]
	}
}

func (table *SummonGachaTable) GetByType(Type int32) (res []*SummonGachaCfg) {
	for _, i := range table.summonGachaTypeIndexMap[Type] {
		res = append(res, table.Get(i))
	}
	return
}
