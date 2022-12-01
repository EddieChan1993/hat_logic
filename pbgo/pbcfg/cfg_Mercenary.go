package config

type MercenaryCfg struct {
	Id      int32
	Mission []int32
	Random  int32
}

type MercenaryTable struct {
	data map[int32]*MercenaryCfg
}

var MercenaryData = &MercenaryTable{
	data: map[int32]*MercenaryCfg{},
}

func (table *MercenaryTable) Get(id int32) *MercenaryCfg {
	return table.data[id]
}

func (table *MercenaryTable) GetAll() []int32 {
	return mercenaryKeys
}

var mercenaryValues = []*MercenaryCfg{
	{
		Id: 1,
		Mission: []int32{
			4001,
			4002,
			4003,
			4004,
			4005,
			4006,
			4101,
			4102,
			4103,
			4104,
			4105,
			4106,
			4201,
			4202,
			4203,
			4204,
			4205,
			4206,
			4301,
			4302,
			4303,
			4304,
			4305,
			4306,
			4401,
			4402,
			4403,
			4404,
			4405,
			4406,
			4501,
			4502,
			4503,
			4504,
			4505,
			4506,
			4507,
			4508,
			4509,
			4510,
			4511,
			4512,
			4513,
			4514,
			4515,
			4516,
			4517,
			4518,
			4519,
			4520,
			4521,
			4522,
			4523,
			4524,
			4525,
			4526,
			4527,
			4528,
			4529,
			4530,
			4531,
			4532,
			4533,
			4534,
			4535,
			4536,
			4537,
			4538,
			4539,
			4540,
			4541,
			4542,
			4543,
			4544,
			4545,
			4546,
			4547,
			4548,
			4549,
			4550,
		},
		Random: 200,
	},
	{
		Id: 2,
		Mission: []int32{
			3001,
			3002,
			3003,
			3004,
			3005,
			3006,
			3101,
			3102,
			3103,
			3104,
			3105,
			3106,
			3201,
			3202,
			3203,
			3204,
			3205,
			3206,
			3301,
			3302,
			3303,
			3304,
			3305,
			3306,
			3402,
			3403,
			3404,
			3405,
			3406,
		},
		Random: 1000,
	},
	{
		Id: 3,
		Mission: []int32{
			2001,
			2002,
			2003,
			2004,
			2005,
			2006,
			2101,
			2102,
			2103,
			2104,
			2105,
			2106,
			2201,
			2202,
			2203,
			2204,
			2205,
			2206,
			2301,
			2302,
			2303,
			2304,
			2305,
			2306,
		},
		Random: 3800,
	},
	{
		Id: 4,
		Mission: []int32{
			1001,
			1002,
			1003,
			1004,
			1005,
			1006,
			1101,
			1102,
			1103,
			1104,
			1105,
			1106,
			1201,
			1202,
			1203,
			1204,
			1205,
			1206,
			1301,
			1302,
			1303,
			1304,
			1305,
			1306,
		},
		Random: 5000,
	},
}

var mercenaryKeys = []int32{
	1, 2, 3, 4,
}

func init() {
	MercenaryData.data = make(map[int32]*MercenaryCfg)
	for i := 0; i < len(mercenaryKeys); i++ {
		MercenaryData.data[mercenaryKeys[i]] = mercenaryValues[i]
	}
}
