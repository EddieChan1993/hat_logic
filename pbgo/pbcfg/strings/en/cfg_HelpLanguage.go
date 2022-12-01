package config

type HelpLanguageCfg struct {
	Id   int32
	Text string
}

type HelpLanguageTable struct {
	data map[int32]*HelpLanguageCfg
}

var HelpLanguageData = &HelpLanguageTable{
	data: map[int32]*HelpLanguageCfg{},
}

func (table *HelpLanguageTable) Get(id int32) *HelpLanguageCfg {
	return table.data[id]
}

func (table *HelpLanguageTable) GetAll() []int32 {
	return helpLanguageKeys
}

var helpLanguageValues = []*HelpLanguageCfg{
	{
		Id:   1,
		Text: "Drop Rate:\\n\\nAdvanced Summon:\\n\\nAdvancement Stone *30: 25%\\nAdvancement Stone *50: 15%\\nAdvancement Stone *70: 10%\\nEpic Hero *1: 30%\\nLegendary Hero *1: 1%\\nLegendary Equipment *1: 2%\\nEpic Accessory *1: 2%\\nMythical Equipment *1: 0.1%\\nArtifact *1: 14.9%\\n\\nMystic Summon:\\n\\nAdvancement Stone *90: 4%\\nAdvancement Stone *150: 4%\\nAdvancement Stone *210: 5%\\nEpic Hero *1: 45%\\nLegendary Hero *1: 10%\\nMythical Equipment or Accessory *1: 2%\\nArtifact *1: 30%",
	},
}

var helpLanguageKeys = []int32{
	1,
}

func init() {
	HelpLanguageData.data = make(map[int32]*HelpLanguageCfg)
	for i := 0; i < len(helpLanguageKeys); i++ {
		HelpLanguageData.data[helpLanguageKeys[i]] = helpLanguageValues[i]
	}
}
