package config

type BuildingCfg struct {
	Id           int32
	BaseOutput   int32
	BaseReward   *Reward
	OutputTime   int32
	FoodCost     int32
	FoodCostTime int32
	Spine        string
}

type BuildingTable struct {
	data map[int32]*BuildingCfg
}

var BuildingData = &BuildingTable{
	data: map[int32]*BuildingCfg{},
}

func (table *BuildingTable) Get(id int32) *BuildingCfg {
	return table.data[id]
}

func (table *BuildingTable) GetAll() []int32 {
	return buildingKeys
}

var buildingValues = []*BuildingCfg{
	{
		Id:         1,
		BaseOutput: 10,
		BaseReward: &Reward{
			Type:  1,
			Id:    37,
			Count: 10,
		},
		OutputTime:   10000,
		FoodCost:     0,
		FoodCostTime: 0,
	},
	{
		Id:         2,
		BaseOutput: 10,
		BaseReward: &Reward{
			Type:  1,
			Id:    38,
			Count: 10,
		},
		OutputTime:   10000,
		FoodCost:     10,
		FoodCostTime: 10000,
	},
	{
		Id:         3,
		BaseOutput: 10,
		BaseReward: &Reward{
			Type:  1,
			Id:    39,
			Count: 10,
		},
		OutputTime:   10000,
		FoodCost:     10,
		FoodCostTime: 10000,
	},
	{
		Id:         4,
		BaseOutput: 10,
		BaseReward: &Reward{
			Type:  1,
			Id:    40,
			Count: 10,
		},
		OutputTime:   10000,
		FoodCost:     5,
		FoodCostTime: 10000,
	},
}

var buildingKeys = []int32{
	1, 2, 3, 4,
}

func init() {
	BuildingData.data = make(map[int32]*BuildingCfg)
	for i := 0; i < len(buildingKeys); i++ {
		BuildingData.data[buildingKeys[i]] = buildingValues[i]
	}
}
