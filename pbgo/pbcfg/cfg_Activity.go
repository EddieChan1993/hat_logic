package config

type ActivityCfg struct {
	Id          int32
	RewardsDay1 []*Reward
	RewardsDay2 []*Reward
	RewardsDay3 []*Reward
	LastDays    int32
}

type ActivityTable struct {
	data map[int32]*ActivityCfg
}

var ActivityData = &ActivityTable{
	data: map[int32]*ActivityCfg{},
}

func (table *ActivityTable) Get(id int32) *ActivityCfg {
	return table.data[id]
}

func (table *ActivityTable) GetAll() []int32 {
	return activityKeys
}

var activityValues = []*ActivityCfg{
	{
		Id: 1,
		RewardsDay1: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
			{
				Type:  1,
				Id:    3,
				Count: 666,
			},
		},
		RewardsDay2: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 10,
			},
			{
				Type:  1,
				Id:    3,
				Count: 666,
			},
		},
		RewardsDay3: []*Reward{
			{
				Type:  1,
				Id:    10003,
				Count: 1,
			},
			{
				Type:  1,
				Id:    3,
				Count: 666,
			},
		},
		LastDays: 9999,
	},
}

var activityKeys = []int32{
	1,
}

func init() {
	ActivityData.data = make(map[int32]*ActivityCfg)
	for i := 0; i < len(activityKeys); i++ {
		ActivityData.data[activityKeys[i]] = activityValues[i]
	}
}
