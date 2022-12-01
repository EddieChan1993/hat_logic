package config

type MercenaryMissionCfg struct {
	Id          int32
	SuccessProb int32
	Reward      *Reward
	MissionTime int32
}

type MercenaryMissionTable struct {
	data map[int32]*MercenaryMissionCfg
}

var MercenaryMissionData = &MercenaryMissionTable{
	data: map[int32]*MercenaryMissionCfg{},
}

func (table *MercenaryMissionTable) Get(id int32) *MercenaryMissionCfg {
	return table.data[id]
}

func (table *MercenaryMissionTable) GetAll() []int32 {
	return mercenaryMissionKeys
}

var mercenaryMissionValues = []*MercenaryMissionCfg{
	{
		Id:          1001,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 50,
		},
		MissionTime: 14400000,
	},
	{
		Id:          1002,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 50,
		},
		MissionTime: 7200000,
	},
	{
		Id:          1003,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 50,
		},
		MissionTime: 5400000,
	},
	{
		Id:          1004,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 50,
		},
		MissionTime: 3600000,
	},
	{
		Id:          1005,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 50,
		},
		MissionTime: 1800000,
	},
	{
		Id:          1006,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 50,
		},
		MissionTime: 600000,
	},
	{
		Id:          1101,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 1,
		},
		MissionTime: 14400000,
	},
	{
		Id:          1102,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          1103,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          1104,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          1105,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          1106,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          1201,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 1,
		},
		MissionTime: 14400000,
	},
	{
		Id:          1202,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          1203,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          1204,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          1205,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          1206,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          1301,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 1,
		},
		MissionTime: 14400000,
	},
	{
		Id:          1302,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          1303,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          1304,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          1305,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          1306,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          2001,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 100,
		},
		MissionTime: 14400000,
	},
	{
		Id:          2002,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 100,
		},
		MissionTime: 7200000,
	},
	{
		Id:          2003,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 100,
		},
		MissionTime: 5400000,
	},
	{
		Id:          2004,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 100,
		},
		MissionTime: 3600000,
	},
	{
		Id:          2005,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 100,
		},
		MissionTime: 1800000,
	},
	{
		Id:          2006,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 100,
		},
		MissionTime: 600000,
	},
	{
		Id:          2101,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 2,
		},
		MissionTime: 14400000,
	},
	{
		Id:          2102,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 2,
		},
		MissionTime: 7200000,
	},
	{
		Id:          2103,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 2,
		},
		MissionTime: 5400000,
	},
	{
		Id:          2104,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 2,
		},
		MissionTime: 3600000,
	},
	{
		Id:          2105,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 2,
		},
		MissionTime: 1800000,
	},
	{
		Id:          2106,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 3,
		},
		MissionTime: 600000,
	},
	{
		Id:          2201,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 2,
		},
		MissionTime: 14400000,
	},
	{
		Id:          2202,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 2,
		},
		MissionTime: 7200000,
	},
	{
		Id:          2203,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 2,
		},
		MissionTime: 5400000,
	},
	{
		Id:          2204,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 2,
		},
		MissionTime: 3600000,
	},
	{
		Id:          2205,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 2,
		},
		MissionTime: 1800000,
	},
	{
		Id:          2206,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 3,
		},
		MissionTime: 600000,
	},
	{
		Id:          2301,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 2,
		},
		MissionTime: 14400000,
	},
	{
		Id:          2302,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 2,
		},
		MissionTime: 7200000,
	},
	{
		Id:          2303,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 2,
		},
		MissionTime: 5400000,
	},
	{
		Id:          2304,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 2,
		},
		MissionTime: 3600000,
	},
	{
		Id:          2305,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 2,
		},
		MissionTime: 1800000,
	},
	{
		Id:          2306,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 3,
		},
		MissionTime: 600000,
	},
	{
		Id:          3001,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 150,
		},
		MissionTime: 14400000,
	},
	{
		Id:          3002,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 150,
		},
		MissionTime: 7200000,
	},
	{
		Id:          3003,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 150,
		},
		MissionTime: 5400000,
	},
	{
		Id:          3004,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 150,
		},
		MissionTime: 3600000,
	},
	{
		Id:          3005,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 150,
		},
		MissionTime: 1800000,
	},
	{
		Id:          3006,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 150,
		},
		MissionTime: 600000,
	},
	{
		Id:          3101,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    1002,
			Count: 1,
		},
		MissionTime: 14400000,
	},
	{
		Id:          3102,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    1002,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          3103,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    1002,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          3104,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    1002,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          3105,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    1002,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          3106,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 4,
		},
		MissionTime: 600000,
	},
	{
		Id:          3201,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 3,
		},
		MissionTime: 14400000,
	},
	{
		Id:          3202,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 3,
		},
		MissionTime: 7200000,
	},
	{
		Id:          3203,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 3,
		},
		MissionTime: 5400000,
	},
	{
		Id:          3204,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 3,
		},
		MissionTime: 3600000,
	},
	{
		Id:          3205,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 3,
		},
		MissionTime: 1800000,
	},
	{
		Id:          3206,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 4,
		},
		MissionTime: 600000,
	},
	{
		Id:          3301,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 3,
		},
		MissionTime: 14400000,
	},
	{
		Id:          3302,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 3,
		},
		MissionTime: 7200000,
	},
	{
		Id:          3303,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 3,
		},
		MissionTime: 5400000,
	},
	{
		Id:          3304,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 3,
		},
		MissionTime: 3600000,
	},
	{
		Id:          3305,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 3,
		},
		MissionTime: 1800000,
	},
	{
		Id:          3306,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 4,
		},
		MissionTime: 600000,
	},
	{
		Id:          3402,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          3403,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          3404,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          3405,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          3406,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4001,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 300,
		},
		MissionTime: 14400000,
	},
	{
		Id:          4002,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 300,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4003,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 300,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4004,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 300,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4005,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 300,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4006,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    3,
			Count: 300,
		},
		MissionTime: 600000,
	},
	{
		Id:          4101,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 5,
		},
		MissionTime: 14400000,
	},
	{
		Id:          4102,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 5,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4103,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 5,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4104,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 5,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4105,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    1001,
			Count: 5,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4106,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    1003,
			Count: 1,
		},
		MissionTime: 600000,
	},
	{
		Id:          4201,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 5,
		},
		MissionTime: 14400000,
	},
	{
		Id:          4202,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 5,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4203,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 5,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4204,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 5,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4205,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 5,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4206,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    50,
			Count: 6,
		},
		MissionTime: 600000,
	},
	{
		Id:          4301,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 5,
		},
		MissionTime: 14400000,
	},
	{
		Id:          4302,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 5,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4303,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 5,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4304,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 5,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4305,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 5,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4306,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    41,
			Count: 6,
		},
		MissionTime: 600000,
	},
	{
		Id:          4401,
		SuccessProb: 1000,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 1,
		},
		MissionTime: 14400000,
	},
	{
		Id:          4402,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 2,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4403,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 2,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4404,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 2,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4405,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 2,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4406,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  1,
			Id:    51,
			Count: 3,
		},
		MissionTime: 600000,
	},
	{
		Id:          4501,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    1,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4502,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    2,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4503,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    3,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4504,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    4,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4505,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    5,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4506,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    6,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4507,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    7,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4508,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    8,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4509,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    9,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4510,
		SuccessProb: 500,
		Reward: &Reward{
			Type:  4,
			Id:    10,
			Count: 1,
		},
		MissionTime: 7200000,
	},
	{
		Id:          4511,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    1,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4512,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    2,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4513,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    3,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4514,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    4,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4515,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    5,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4516,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    6,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4517,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    7,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4518,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    8,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4519,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    9,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4520,
		SuccessProb: 300,
		Reward: &Reward{
			Type:  4,
			Id:    10,
			Count: 1,
		},
		MissionTime: 5400000,
	},
	{
		Id:          4521,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    1,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4522,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    2,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4523,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    3,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4524,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    4,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4525,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    5,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4526,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    6,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4527,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    7,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4528,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    8,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4529,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    9,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4530,
		SuccessProb: 150,
		Reward: &Reward{
			Type:  4,
			Id:    10,
			Count: 1,
		},
		MissionTime: 3600000,
	},
	{
		Id:          4531,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    1,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4532,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    2,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4533,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    3,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4534,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    4,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4535,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    5,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4536,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    6,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4537,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    7,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4538,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    8,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4539,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    9,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4540,
		SuccessProb: 100,
		Reward: &Reward{
			Type:  4,
			Id:    10,
			Count: 1,
		},
		MissionTime: 1800000,
	},
	{
		Id:          4541,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    1,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4542,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    2,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4543,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    3,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4544,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    4,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4545,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    5,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4546,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    6,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4547,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    7,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4548,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    8,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4549,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    9,
			Count: 2,
		},
		MissionTime: 600000,
	},
	{
		Id:          4550,
		SuccessProb: 50,
		Reward: &Reward{
			Type:  4,
			Id:    10,
			Count: 2,
		},
		MissionTime: 600000,
	},
}

var mercenaryMissionKeys = []int32{
	1001, 1002, 1003, 1004, 1005, 1006, 1101, 1102, 1103, 1104,
	1105, 1106, 1201, 1202, 1203, 1204, 1205, 1206, 1301, 1302,
	1303, 1304, 1305, 1306, 2001, 2002, 2003, 2004, 2005, 2006,
	2101, 2102, 2103, 2104, 2105, 2106, 2201, 2202, 2203, 2204,
	2205, 2206, 2301, 2302, 2303, 2304, 2305, 2306, 3001, 3002,
	3003, 3004, 3005, 3006, 3101, 3102, 3103, 3104, 3105, 3106,
	3201, 3202, 3203, 3204, 3205, 3206, 3301, 3302, 3303, 3304,
	3305, 3306, 3402, 3403, 3404, 3405, 3406, 4001, 4002, 4003,
	4004, 4005, 4006, 4101, 4102, 4103, 4104, 4105, 4106, 4201,
	4202, 4203, 4204, 4205, 4206, 4301, 4302, 4303, 4304, 4305,
	4306, 4401, 4402, 4403, 4404, 4405, 4406, 4501, 4502, 4503,
	4504, 4505, 4506, 4507, 4508, 4509, 4510, 4511, 4512, 4513,
	4514, 4515, 4516, 4517, 4518, 4519, 4520, 4521, 4522, 4523,
	4524, 4525, 4526, 4527, 4528, 4529, 4530, 4531, 4532, 4533,
	4534, 4535, 4536, 4537, 4538, 4539, 4540, 4541, 4542, 4543,
	4544, 4545, 4546, 4547, 4548, 4549, 4550,
}

func init() {
	MercenaryMissionData.data = make(map[int32]*MercenaryMissionCfg)
	for i := 0; i < len(mercenaryMissionKeys); i++ {
		MercenaryMissionData.data[mercenaryMissionKeys[i]] = mercenaryMissionValues[i]
	}
}
