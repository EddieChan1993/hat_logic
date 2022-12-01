package config

type FunctionOpenLanguageCfg struct {
	Id   int32
	Name string
}

type FunctionOpenLanguageTable struct {
	data map[int32]*FunctionOpenLanguageCfg
}

var FunctionOpenLanguageData = &FunctionOpenLanguageTable{
	data: map[int32]*FunctionOpenLanguageCfg{},
}

func (table *FunctionOpenLanguageTable) Get(id int32) *FunctionOpenLanguageCfg {
	return table.data[id]
}

func (table *FunctionOpenLanguageTable) GetAll() []int32 {
	return functionOpenLanguageKeys
}

var functionOpenLanguageValues = []*FunctionOpenLanguageCfg{
	{
		Id:   1,
		Name: "Market",
	},
	{
		Id:   2,
		Name: "Artifact",
	},
	{
		Id:   3,
		Name: "Manor",
	},
	{
		Id:   4,
		Name: "Mercenary",
	},
	{
		Id:   5,
		Name: "Dungeon",
	},
	{
		Id:   6,
		Name: "Coin Dungeon",
	},
	{
		Id:   7,
		Name: "Resource Dungeon",
	},
	{
		Id:   8,
		Name: "Enhancement Stone Dungeon",
	},
	{
		Id:   9,
		Name: "World Boss",
	},
}

var functionOpenLanguageKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
}

func init() {
	FunctionOpenLanguageData.data = make(map[int32]*FunctionOpenLanguageCfg)
	for i := 0; i < len(functionOpenLanguageKeys); i++ {
		FunctionOpenLanguageData.data[functionOpenLanguageKeys[i]] = functionOpenLanguageValues[i]
	}
}
