package config

type InstanceBossCfg struct {
	Id               int32
	MonsterId        int32
	SpecialCondition []*SpecialCondition
	ConditionObj     int32
	Scene            string
}

type InstanceBossTable struct {
	data map[int32]*InstanceBossCfg
}

var InstanceBossData = &InstanceBossTable{
	data: map[int32]*InstanceBossCfg{},
}

func (table *InstanceBossTable) Get(id int32) *InstanceBossCfg {
	return table.data[id]
}

func (table *InstanceBossTable) GetAll() []int32 {
	return instanceBossKeys
}

var instanceBossValues = []*InstanceBossCfg{
	{
		Id:        1,
		MonsterId: 1,
		SpecialCondition: []*SpecialCondition{
			{
				Skill:      90004,
				SkillLevel: 1,
			},
		},
		ConditionObj: 2,
	},
}

var instanceBossKeys = []int32{
	1,
}

func init() {
	InstanceBossData.data = make(map[int32]*InstanceBossCfg)
	for i := 0; i < len(instanceBossKeys); i++ {
		InstanceBossData.data[instanceBossKeys[i]] = instanceBossValues[i]
	}
}
