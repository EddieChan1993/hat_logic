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
		Text: "概率公示：\\n\\n高级召唤：\\n\\n英雄进阶石*30：25%\\n英雄进阶石*50：15%\\n英雄进阶石*70：10%\\n史诗英雄*1：30%\\n传说英雄*1：1%\\n传说装备*1：2%\\n史诗饰品*1：2%\\n神话装备*1：0.1%\\n神器*1：14.9%\\n\\n神秘召唤：\\n\\n英雄进阶石*90：4%\\n英雄进阶石*150：4%\\n英雄进阶石*210：5%\\n史诗英雄*1：45%\\n传说英雄*1：10%\\n神话装备（包含饰品）*1：2%\\n神器*1：30%",
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
