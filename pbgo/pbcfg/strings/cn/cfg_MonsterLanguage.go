package config

type MonsterLanguageCfg struct {
	Id          int32
	MonsterName string
}

type MonsterLanguageTable struct {
	data map[int32]*MonsterLanguageCfg
}

var MonsterLanguageData = &MonsterLanguageTable{
	data: map[int32]*MonsterLanguageCfg{},
}

func (table *MonsterLanguageTable) Get(id int32) *MonsterLanguageCfg {
	return table.data[id]
}

func (table *MonsterLanguageTable) GetAll() []int32 {
	return monsterLanguageKeys
}

var monsterLanguageValues = []*MonsterLanguageCfg{
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
}

var monsterLanguageKeys = []int32{
	10013, 10015, 10023, 10025, 10033, 10043, 10045, 10053, 10055, 10063,
	10073, 10075, 10083, 10085, 10093, 10095, 10103, 10105, 10112, 10131,
	10142, 10145, 10151, 10161, 10165, 10172, 10175, 10181, 10185, 10191,
	10195, 10202, 10205, 10211, 10215, 10221, 10225, 10232, 10235, 10241,
	10245, 10251, 10262, 10271, 10281, 10292, 10295, 10301, 10305, 10311,
	10322, 10325, 10331, 10341, 10345, 10352, 10361, 10371, 10375, 10382,
	10385, 10391, 10395, 10401, 1000001,
}

func init() {
	MonsterLanguageData.data = make(map[int32]*MonsterLanguageCfg)
	for i := 0; i < len(monsterLanguageKeys); i++ {
		MonsterLanguageData.data[monsterLanguageKeys[i]] = monsterLanguageValues[i]
	}
}
