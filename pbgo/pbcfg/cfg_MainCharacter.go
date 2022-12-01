package config

type MainCharacterCfg struct {
	Id               int32
	Group            int32
	Job              int32
	SkillGroup       []int32
	ActiveSkillGroup []int32
	Talent           []int32
	Interval         int32
	BattleModel      string
	Model            string
	Card             string
	HeadIcon         string
}

type MainCharacterTable struct {
	data map[int32]*MainCharacterCfg
}

var MainCharacterData = &MainCharacterTable{
	data: map[int32]*MainCharacterCfg{},
}

func (table *MainCharacterTable) Get(id int32) *MainCharacterCfg {
	return table.data[id]
}

func (table *MainCharacterTable) GetAll() []int32 {
	return mainCharacterKeys
}

var mainCharacterValues = []*MainCharacterCfg{
	{
		Id:    11,
		Group: 1,
		SkillGroup: []int32{
			101,
			111,
		},
		ActiveSkillGroup: []int32{
			11001,
			11005,
			11007,
			11009,
			11013,
			11014,
			11022,
			11024,
			11027,
			11031,
			11032,
			11034,
		},
		Talent: []int32{
			1,
			2,
			3,
		},
		Interval: 2000,
	},
}

var mainCharacterKeys = []int32{
	11,
}

func init() {
	MainCharacterData.data = make(map[int32]*MainCharacterCfg)
	for i := 0; i < len(mainCharacterKeys); i++ {
		MainCharacterData.data[mainCharacterKeys[i]] = mainCharacterValues[i]
	}
}
