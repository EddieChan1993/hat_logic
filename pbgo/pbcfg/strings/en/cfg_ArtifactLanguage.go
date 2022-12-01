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
		Name: "Star-Piercing Bow",
	},
	{
		Id:   2,
		Name: "Golden Spear",
	},
	{
		Id:   3,
		Name: "Trident of Raging Wave",
	},
	{
		Id:   4,
		Name: "Shield of Victory",
	},
	{
		Id:   5,
		Name: "Double-Snake Staff",
	},
	{
		Id:   6,
		Name: "Cup of Divine Feast",
	},
	{
		Id:   7,
		Name: "Scepter of Zeus",
	},
	{
		Id:   8,
		Name: "Eye of Hades",
	},
	{
		Id:   9,
		Name: "Golden Armor",
	},
	{
		Id:   10,
		Name: "Crown of the Sun",
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
