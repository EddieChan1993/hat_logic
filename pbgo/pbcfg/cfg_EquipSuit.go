package config

type EquipSuitCfg struct {
	Id       int32
	SuitTwo  int32
	AttType1 []int32
	SuitFour int32
	AttType2 []int32
	SuitSix  int32
	AttType3 []int32
}

type EquipSuitTable struct {
	data map[int32]*EquipSuitCfg
}

var EquipSuitData = &EquipSuitTable{
	data: map[int32]*EquipSuitCfg{},
}

func (table *EquipSuitTable) Get(id int32) *EquipSuitCfg {
	return table.data[id]
}

func (table *EquipSuitTable) GetAll() []int32 {
	return equipSuitKeys
}

var equipSuitValues = []*EquipSuitCfg{
	{
		Id:      11101,
		SuitTwo: 111,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 112,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      11102,
		SuitTwo: 121,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 122,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 123,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      11103,
		SuitTwo: 131,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 132,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 133,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      12101,
		SuitTwo: 211,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 212,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      12102,
		SuitTwo: 221,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 222,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 223,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      12103,
		SuitTwo: 231,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 232,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 233,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      13101,
		SuitTwo: 311,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 312,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      13102,
		SuitTwo: 321,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 322,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 323,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      13103,
		SuitTwo: 331,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 332,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 333,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      14101,
		SuitTwo: 411,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 412,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      14102,
		SuitTwo: 421,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 422,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 423,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      14103,
		SuitTwo: 431,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 432,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 433,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      15101,
		SuitTwo: 511,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 512,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      15102,
		SuitTwo: 521,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 522,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 523,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      15103,
		SuitTwo: 531,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 532,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 533,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      16101,
		SuitTwo: 611,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 612,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      16102,
		SuitTwo: 621,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 622,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 623,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      16103,
		SuitTwo: 631,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 632,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 633,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      17101,
		SuitTwo: 711,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 712,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      17102,
		SuitTwo: 721,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 722,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 723,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      17103,
		SuitTwo: 731,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 732,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 733,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      18101,
		SuitTwo: 811,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 812,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      18102,
		SuitTwo: 821,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 822,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 823,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      18103,
		SuitTwo: 831,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 832,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 833,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      19101,
		SuitTwo: 911,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 912,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      19102,
		SuitTwo: 921,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 922,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 923,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      19103,
		SuitTwo: 931,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 932,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 933,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      20101,
		SuitTwo: 1011,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 1012,
		AttType2: []int32{
			22,
			32,
			63,
		},
	},
	{
		Id:      20102,
		SuitTwo: 1021,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 1022,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 1023,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
	{
		Id:      20103,
		SuitTwo: 1031,
		AttType1: []int32{
			22,
			32,
			103,
		},
		SuitFour: 1032,
		AttType2: []int32{
			22,
			32,
			63,
		},
		SuitSix: 1033,
		AttType3: []int32{
			22,
			32,
			73,
		},
	},
}

var equipSuitKeys = []int32{
	11101, 11102, 11103, 12101, 12102, 12103, 13101, 13102, 13103, 14101,
	14102, 14103, 15101, 15102, 15103, 16101, 16102, 16103, 17101, 17102,
	17103, 18101, 18102, 18103, 19101, 19102, 19103, 20101, 20102, 20103,
}

func init() {
	EquipSuitData.data = make(map[int32]*EquipSuitCfg)
	for i := 0; i < len(equipSuitKeys); i++ {
		EquipSuitData.data[equipSuitKeys[i]] = equipSuitValues[i]
	}
}
