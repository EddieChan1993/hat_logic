package config

type BuildingLanguageCfg struct {
	Id             int32
	BuildingName   string
	BuildingNameLv string
}

type BuildingLanguageTable struct {
	data map[int32]*BuildingLanguageCfg
}

var BuildingLanguageData = &BuildingLanguageTable{
	data: map[int32]*BuildingLanguageCfg{},
}

func (table *BuildingLanguageTable) Get(id int32) *BuildingLanguageCfg {
	return table.data[id]
}

func (table *BuildingLanguageTable) GetAll() []int32 {
	return buildingLanguageKeys
}

var buildingLanguageValues = []*BuildingLanguageCfg{
	{
		Id:             1,
		BuildingName:   "农场",
		BuildingNameLv: "农场 {0}级",
	},
	{
		Id:             2,
		BuildingName:   "采矿场",
		BuildingNameLv: "采矿场 {0}级",
	},
	{
		Id:             3,
		BuildingName:   "采石场",
		BuildingNameLv: "采石场 {0}级",
	},
	{
		Id:             4,
		BuildingName:   "木材场",
		BuildingNameLv: "木材场 {0}级",
	},
}

var buildingLanguageKeys = []int32{
	1, 2, 3, 4,
}

func init() {
	BuildingLanguageData.data = make(map[int32]*BuildingLanguageCfg)
	for i := 0; i < len(buildingLanguageKeys); i++ {
		BuildingLanguageData.data[buildingLanguageKeys[i]] = buildingLanguageValues[i]
	}
}
