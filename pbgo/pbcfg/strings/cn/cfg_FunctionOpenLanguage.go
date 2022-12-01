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
		Name: "活动（商城）",
	},
	{
		Id:   2,
		Name: "神器",
	},
	{
		Id:   3,
		Name: "领地",
	},
	{
		Id:   4,
		Name: "佣兵",
	},
	{
		Id:   5,
		Name: "副本",
	},
	{
		Id:   6,
		Name: "金币副本",
	},
	{
		Id:   7,
		Name: "资源副本",
	},
	{
		Id:   8,
		Name: "装备副本",
	},
	{
		Id:   9,
		Name: "世界BOSS",
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
