package config

type ArtifactLanguageCfg struct {
	Id   int32
	Name string
}

type ArtifactLanguageTable struct {
	data map[int32]*ArtifactLanguageCfg
}

var ArtifactLanguageData = &ArtifactLanguageTable{
	data: map[int32]*ArtifactLanguageCfg{},
}

func (table *ArtifactLanguageTable) Get(id int32) *ArtifactLanguageCfg {
	return table.data[id]
}

func (table *ArtifactLanguageTable) GetAll() []int32 {
	return artifactLanguageKeys
}

var artifactLanguageValues = []*ArtifactLanguageCfg{
	{
		Id:   1,
		Name: "贯星长弓",
	},
	{
		Id:   2,
		Name: "黄金长矛",
	},
	{
		Id:   3,
		Name: "怒涛三叉戟",
	},
	{
		Id:   4,
		Name: "胜利之盾",
	},
	{
		Id:   5,
		Name: "双蛇杖",
	},
	{
		Id:   6,
		Name: "飨宴角杯",
	},
	{
		Id:   7,
		Name: "宙斯的权杖",
	},
	{
		Id:   8,
		Name: "哈迪斯之眼",
	},
	{
		Id:   9,
		Name: "黄金盔甲",
	},
	{
		Id:   10,
		Name: "太阳王冠",
	},
}

var artifactLanguageKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func init() {
	ArtifactLanguageData.data = make(map[int32]*ArtifactLanguageCfg)
	for i := 0; i < len(artifactLanguageKeys); i++ {
		ArtifactLanguageData.data[artifactLanguageKeys[i]] = artifactLanguageValues[i]
	}
}
