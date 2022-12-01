package config

type MailLanguageCfg struct {
	Id      int32
	Title   string
	Content string
	From    string
}

type MailLanguageTable struct {
	data map[int32]*MailLanguageCfg
}

var MailLanguageData = &MailLanguageTable{
	data: map[int32]*MailLanguageCfg{},
}

func (table *MailLanguageTable) Get(id int32) *MailLanguageCfg {
	return table.data[id]
}

func (table *MailLanguageTable) GetAll() []int32 {
	return mailLanguageKeys
}

var mailLanguageValues = []*MailLanguageCfg{
	{},
	{},
	{},
	{},
	{},
}

var mailLanguageKeys = []int32{
	1, 2, 3, 4, 5,
}

func init() {
	MailLanguageData.data = make(map[int32]*MailLanguageCfg)
	for i := 0; i < len(mailLanguageKeys); i++ {
		MailLanguageData.data[mailLanguageKeys[i]] = mailLanguageValues[i]
	}
}
