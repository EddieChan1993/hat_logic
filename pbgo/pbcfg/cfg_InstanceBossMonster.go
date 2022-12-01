package config

type InstanceBossMonsterCfg struct {
	Id             int32
	TimeLimit      int32
	MonsterId      int32
	HpB            int64
	AdB            int64
	MpB            int64
	DrB            int64
	MdfB           int64
	HitB           int32
	DodgeB         int32
	AspdB          int32
	AdCritB        int32
	MpCritB        int32
	CritDmgAdd     int32
	AdCritResistB  int32
	MpCritResistB  int32
	LifeStealAdd   int32
	MpLifeStealAdd int32
}

type InstanceBossMonsterTable struct {
	data map[int32]*InstanceBossMonsterCfg
}

var InstanceBossMonsterData = &InstanceBossMonsterTable{
	data: map[int32]*InstanceBossMonsterCfg{},
}

func (table *InstanceBossMonsterTable) Get(id int32) *InstanceBossMonsterCfg {
	return table.data[id]
}

func (table *InstanceBossMonsterTable) GetAll() []int32 {
	return instanceBossMonsterKeys
}

var instanceBossMonsterValues = []*InstanceBossMonsterCfg{
	{
		Id:             1,
		TimeLimit:      30000,
		MonsterId:      1000001,
		HpB:            9000000000000000000,
		AdB:            50,
		MpB:            50,
		DrB:            5,
		MdfB:           5,
		HitB:           100000,
		DodgeB:         0,
		AspdB:          0,
		AdCritB:        0,
		MpCritB:        0,
		CritDmgAdd:     0,
		AdCritResistB:  0,
		MpCritResistB:  0,
		LifeStealAdd:   0,
		MpLifeStealAdd: 0,
	},
}

var instanceBossMonsterKeys = []int32{
	1,
}

func init() {
	InstanceBossMonsterData.data = make(map[int32]*InstanceBossMonsterCfg)
	for i := 0; i < len(instanceBossMonsterKeys); i++ {
		InstanceBossMonsterData.data[instanceBossMonsterKeys[i]] = instanceBossMonsterValues[i]
	}
}
