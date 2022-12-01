package config

type MainCharacterLanguageCfg struct {
	Id   int32
	Name string
}

type MainCharacterLanguageTable struct {
	data map[int32]*MainCharacterLanguageCfg
}

var MainCharacterLanguageData = &MainCharacterLanguageTable{
	data: map[int32]*MainCharacterLanguageCfg{},
}

func (table *MainCharacterLanguageTable) Get(id int32) *MainCharacterLanguageCfg {
	return table.data[id]
}

func (table *MainCharacterLanguageTable) GetAll() []int32 {
	return mainCharacterLanguageKeys
}

var mainCharacterLanguageValues = []*MainCharacterLanguageCfg{
	{},
}

var mainCharacterLanguageKeys = []int32{
	11,
}

func init() {
	MainCharacterLanguageData.data = make(map[int32]*MainCharacterLanguageCfg)
	for i := 0; i < len(mainCharacterLanguageKeys); i++ {
		MainCharacterLanguageData.data[mainCharacterLanguageKeys[i]] = mainCharacterLanguageValues[i]
	}
}
