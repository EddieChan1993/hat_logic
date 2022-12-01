package config

type FunctionOpenCfg struct {
	Id     int32
	Level  int32
	Stage  int32
	Page   int32
	Icon   string
	Iftips int32
}

type FunctionOpenTable struct {
	data map[int32]*FunctionOpenCfg
}

var FunctionOpenData = &FunctionOpenTable{
	data: map[int32]*FunctionOpenCfg{},
}

func (table *FunctionOpenTable) Get(id int32) *FunctionOpenCfg {
	return table.data[id]
}

func (table *FunctionOpenTable) GetAll() []int32 {
	return functionOpenKeys
}

var functionOpenValues = []*FunctionOpenCfg{
	{
		Id:     1,
		Level:  1,
		Stage:  1002,
		Iftips: 1,
	},
	{
		Id:     2,
		Level:  1,
		Stage:  5005,
		Page:   1,
		Icon:   "common_icon_artifact",
		Iftips: 1,
	},
	{
		Id:     3,
		Level:  1,
		Stage:  2005,
		Page:   1,
		Icon:   "common_icon_territory",
		Iftips: 1,
	},
	{
		Id:     4,
		Level:  1,
		Stage:  3005,
		Page:   1,
		Icon:   "common_icon_missions",
		Iftips: 1,
	},
	{
		Id:     5,
		Level:  1,
		Stage:  4005,
		Page:   1,
		Icon:   "common_icon_instance",
		Iftips: 1,
	},
	{
		Id:     6,
		Level:  1,
		Stage:  4005,
		Icon:   "common_icon_instance",
		Iftips: 1,
	},
	{
		Id:     7,
		Level:  1,
		Stage:  6005,
		Page:   1,
		Icon:   "common_icon_instance",
		Iftips: 1,
	},
	{
		Id:     8,
		Level:  1,
		Stage:  7005,
		Page:   1,
		Icon:   "common_icon_instance",
		Iftips: 1,
	},
	{
		Id:     9,
		Level:  1,
		Stage:  9005,
		Page:   1,
		Icon:   "common_icon_instance",
		Iftips: 1,
	},
	{
		Id:     10,
		Level:  1,
		Stage:  2001,
		Iftips: 1,
	},
}

var functionOpenKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func init() {
	FunctionOpenData.data = make(map[int32]*FunctionOpenCfg)
	for i := 0; i < len(functionOpenKeys); i++ {
		FunctionOpenData.data[functionOpenKeys[i]] = functionOpenValues[i]
	}
}
