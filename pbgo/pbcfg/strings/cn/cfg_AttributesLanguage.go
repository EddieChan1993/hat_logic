package config

type AttributesLanguageCfg struct {
	Id             int32
	AttribName     string
	AttribDescribe string
}

type AttributesLanguageTable struct {
	data map[int32]*AttributesLanguageCfg
}

var AttributesLanguageData = &AttributesLanguageTable{
	data: map[int32]*AttributesLanguageCfg{},
}

func (table *AttributesLanguageTable) Get(id int32) *AttributesLanguageCfg {
	return table.data[id]
}

func (table *AttributesLanguageTable) GetAll() []int32 {
	return attributesLanguageKeys
}

var attributesLanguageValues = []*AttributesLanguageCfg{
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

var attributesLanguageKeys = []int32{
	10, 11, 12, 13, 14, 20, 21, 22, 23, 24,
	30, 31, 32, 33, 34, 40, 41, 42, 43, 44,
	50, 51, 52, 53, 54, 60, 61, 62, 63, 64,
	70, 71, 72, 73, 74, 80, 81, 82, 83, 84,
	90, 91, 92, 93, 94, 100, 101, 103, 110, 111,
	112, 113, 114, 120, 121, 122, 123, 124, 130, 133,
	140, 143, 150, 151, 152, 153, 154, 160, 161, 162,
	163, 164, 173, 183, 193, 203, 213, 223, 233, 243,
	253, 263, 273, 283, 293, 303,
}

func init() {
	AttributesLanguageData.data = make(map[int32]*AttributesLanguageCfg)
	for i := 0; i < len(attributesLanguageKeys); i++ {
		AttributesLanguageData.data[attributesLanguageKeys[i]] = attributesLanguageValues[i]
	}
}
