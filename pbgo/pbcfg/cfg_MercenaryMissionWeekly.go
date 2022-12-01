package config

type MercenaryMissionWeeklyCfg struct {
	Id           int32
	Progress     int32
	WeeklyReward []*Reward
}

type MercenaryMissionWeeklyTable struct {
	data map[int32]*MercenaryMissionWeeklyCfg
}

var MercenaryMissionWeeklyData = &MercenaryMissionWeeklyTable{
	data: map[int32]*MercenaryMissionWeeklyCfg{},
}

func (table *MercenaryMissionWeeklyTable) Get(id int32) *MercenaryMissionWeeklyCfg {
	return table.data[id]
}

func (table *MercenaryMissionWeeklyTable) GetAll() []int32 {
	return mercenaryMissionWeeklyKeys
}

var mercenaryMissionWeeklyValues = []*MercenaryMissionWeeklyCfg{
	{
		Id:       1,
		Progress: 10,
		WeeklyReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
	},
	{
		Id:       2,
		Progress: 20,
		WeeklyReward: []*Reward{
			{
				Type:  1,
				Id:    10002,
				Count: 1,
			},
		},
	},
	{
		Id:       3,
		Progress: 30,
		WeeklyReward: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
		},
	},
}

var mercenaryMissionWeeklyKeys = []int32{
	1, 2, 3,
}

func init() {
	MercenaryMissionWeeklyData.data = make(map[int32]*MercenaryMissionWeeklyCfg)
	for i := 0; i < len(mercenaryMissionWeeklyKeys); i++ {
		MercenaryMissionWeeklyData.data[mercenaryMissionWeeklyKeys[i]] = mercenaryMissionWeeklyValues[i]
	}
}
