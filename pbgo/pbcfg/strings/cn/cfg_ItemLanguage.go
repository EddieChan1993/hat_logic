package config

type ItemLanguageCfg struct {
	Id      int32
	Name    string
	Explain string
}

type ItemLanguageTable struct {
	data map[int32]*ItemLanguageCfg
}

var ItemLanguageData = &ItemLanguageTable{
	data: map[int32]*ItemLanguageCfg{},
}

func (table *ItemLanguageTable) Get(id int32) *ItemLanguageCfg {
	return table.data[id]
}

func (table *ItemLanguageTable) GetAll() []int32 {
	return itemLanguageKeys
}

var itemLanguageValues = []*ItemLanguageCfg{
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

var itemLanguageKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 50, 51, 101,
	1001, 1002, 1003, 1004, 1005, 2001, 2002, 2003, 2004, 2005,
	10001, 10002, 10003,
}

func init() {
	ItemLanguageData.data = make(map[int32]*ItemLanguageCfg)
	for i := 0; i < len(itemLanguageKeys); i++ {
		ItemLanguageData.data[itemLanguageKeys[i]] = itemLanguageValues[i]
	}
}
