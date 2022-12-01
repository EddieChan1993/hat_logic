package config

type MercenaryLanguageCfg struct {
	Id            int32
	MercenaryName string
}

type MercenaryLanguageTable struct {
	data map[int32]*MercenaryLanguageCfg
}

var MercenaryLanguageData = &MercenaryLanguageTable{
	data: map[int32]*MercenaryLanguageCfg{},
}

func (table *MercenaryLanguageTable) Get(id int32) *MercenaryLanguageCfg {
	return table.data[id]
}

func (table *MercenaryLanguageTable) GetAll() []int32 {
	return mercenaryLanguageKeys
}

var mercenaryLanguageValues = []*MercenaryLanguageCfg{
	{
		Id: 1,
	},
	{
		Id: 2,
	},
	{
		Id: 3,
	},
	{
		Id: 4,
	},
	{
		Id: 5,
	},
	{
		Id: 6,
	},
	{
		Id: 7,
	},
	{
		Id: 8,
	},
	{
		Id: 9,
	},
	{
		Id: 10,
	},
	{
		Id: 11,
	},
	{
		Id: 12,
	},
	{
		Id: 13,
	},
	{
		Id: 14,
	},
	{
		Id: 15,
	},
	{
		Id: 16,
	},
	{
		Id: 17,
	},
	{
		Id: 18,
	},
	{
		Id: 19,
	},
	{
		Id: 20,
	},
	{
		Id: 21,
	},
	{
		Id: 22,
	},
	{
		Id: 23,
	},
	{
		Id: 24,
	},
	{
		Id: 25,
	},
	{
		Id: 26,
	},
	{
		Id: 27,
	},
	{
		Id: 28,
	},
	{
		Id: 29,
	},
	{
		Id: 30,
	},
	{
		Id: 31,
	},
	{
		Id: 32,
	},
	{
		Id: 33,
	},
	{
		Id: 34,
	},
	{
		Id: 35,
	},
	{
		Id: 36,
	},
	{
		Id: 37,
	},
	{
		Id: 38,
	},
	{
		Id: 39,
	},
	{
		Id: 40,
	},
	{
		Id: 41,
	},
	{
		Id: 42,
	},
	{
		Id: 43,
	},
	{
		Id: 44,
	},
	{
		Id: 45,
	},
	{
		Id: 46,
	},
	{
		Id: 47,
	},
	{
		Id: 48,
	},
	{
		Id: 49,
	},
	{
		Id: 50,
	},
	{
		Id: 51,
	},
	{
		Id: 52,
	},
	{
		Id: 53,
	},
	{
		Id: 54,
	},
	{
		Id: 55,
	},
	{
		Id: 56,
	},
	{
		Id: 57,
	},
	{
		Id: 58,
	},
	{
		Id: 59,
	},
	{
		Id: 60,
	},
	{
		Id: 61,
	},
	{
		Id: 62,
	},
	{
		Id: 63,
	},
	{
		Id: 64,
	},
	{
		Id: 65,
	},
	{
		Id: 66,
	},
	{
		Id: 67,
	},
	{
		Id: 68,
	},
	{
		Id: 69,
	},
	{
		Id: 70,
	},
	{
		Id: 71,
	},
}

var mercenaryLanguageKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71,
}

func init() {
	MercenaryLanguageData.data = make(map[int32]*MercenaryLanguageCfg)
	for i := 0; i < len(mercenaryLanguageKeys); i++ {
		MercenaryLanguageData.data[mercenaryLanguageKeys[i]] = mercenaryLanguageValues[i]
	}
}
