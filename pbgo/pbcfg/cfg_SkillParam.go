package config

type SkillParamCfg struct {
	Id int32
}

type SkillParamTable struct {
	data map[int32]*SkillParamCfg
}

var SkillParamData = &SkillParamTable{
	data: map[int32]*SkillParamCfg{},
}

func (table *SkillParamTable) Get(id int32) *SkillParamCfg {
	return table.data[id]
}

func (table *SkillParamTable) GetAll() []int32 {
	return skillParamKeys
}

var skillParamValues = []*SkillParamCfg{
	{
		Id: 1,
	},
	{
		Id: 2,
	},
	{
		Id: 3,
	},
	{
		Id: 4,
	},
	{
		Id: 5,
	},
	{
		Id: 6,
	},
	{
		Id: 7,
	},
	{
		Id: 8,
	},
	{
		Id: 9,
	},
	{
		Id: 10,
	},
	{
		Id: 11,
	},
	{
		Id: 12,
	},
	{
		Id: 13,
	},
}

var skillParamKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13,
}

func init() {
	SkillParamData.data = make(map[int32]*SkillParamCfg)
	for i := 0; i < len(skillParamKeys); i++ {
		SkillParamData.data[skillParamKeys[i]] = skillParamValues[i]
	}
}
