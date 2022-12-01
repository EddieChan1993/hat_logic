package config

type HeroLanguageCfg struct {
	Id   int32
	Name string
}

type HeroLanguageTable struct {
	data map[int32]*HeroLanguageCfg
}

var HeroLanguageData = &HeroLanguageTable{
	data: map[int32]*HeroLanguageCfg{},
}

func (table *HeroLanguageTable) Get(id int32) *HeroLanguageCfg {
	return table.data[id]
}

func (table *HeroLanguageTable) GetAll() []int32 {
	return heroLanguageKeys
}

var heroLanguageValues = []*HeroLanguageCfg{
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

var heroLanguageKeys = []int32{
	10021, 10121, 10222, 10321, 10421, 10521, 10622, 10721, 10822, 10921,
	11021, 11122, 11222, 11322, 11421, 11511, 11612, 11712, 11812, 11911,
	12011, 12112, 12212, 12312, 12412, 12511, 12611, 12711, 12811, 12911,
}

func init() {
	HeroLanguageData.data = make(map[int32]*HeroLanguageCfg)
	for i := 0; i < len(heroLanguageKeys); i++ {
		HeroLanguageData.data[heroLanguageKeys[i]] = heroLanguageValues[i]
	}
}
