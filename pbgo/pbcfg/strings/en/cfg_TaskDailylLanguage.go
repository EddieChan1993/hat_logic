package config

type TaskDailylLanguageCfg struct {
	Id    int32
	Title string
}

type TaskDailylLanguageTable struct {
	data map[int32]*TaskDailylLanguageCfg
}

var TaskDailylLanguageData = &TaskDailylLanguageTable{
	data: map[int32]*TaskDailylLanguageCfg{},
}

func (table *TaskDailylLanguageTable) Get(id int32) *TaskDailylLanguageCfg {
	return table.data[id]
}

func (table *TaskDailylLanguageTable) GetAll() []int32 {
	return taskDailylLanguageKeys
}

var taskDailylLanguageValues = []*TaskDailylLanguageCfg{
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
}

var taskDailylLanguageKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8,
}

func init() {
	TaskDailylLanguageData.data = make(map[int32]*TaskDailylLanguageCfg)
	for i := 0; i < len(taskDailylLanguageKeys); i++ {
		TaskDailylLanguageData.data[taskDailylLanguageKeys[i]] = taskDailylLanguageValues[i]
	}
}
