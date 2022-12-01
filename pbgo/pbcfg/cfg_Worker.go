package config

type WorkerCfg struct {
	Id       int32
	Consume1 int32
	Consume2 int32
}

type WorkerTable struct {
	data map[int32]*WorkerCfg
}

var WorkerData = &WorkerTable{
	data: map[int32]*WorkerCfg{},
}

func (table *WorkerTable) Get(id int32) *WorkerCfg {
	return table.data[id]
}

func (table *WorkerTable) GetAll() []int32 {
	return workerKeys
}

var workerValues = []*WorkerCfg{
	{
		Id:       1,
		Consume1: 10000,
		Consume2: 0,
	},
	{
		Id:       2,
		Consume1: 100000,
		Consume2: 0,
	},
	{
		Id:       3,
		Consume1: 400000,
		Consume2: 0,
	},
	{
		Id:       4,
		Consume1: 1000000,
		Consume2: 0,
	},
	{
		Id:       5,
		Consume1: 3000000,
		Consume2: 0,
	},
	{
		Id:       6,
		Consume1: 0,
		Consume2: 100,
	},
	{
		Id:       7,
		Consume1: 0,
		Consume2: 100,
	},
	{
		Id:       8,
		Consume1: 0,
		Consume2: 100,
	},
	{
		Id:       9,
		Consume1: 0,
		Consume2: 100,
	},
	{
		Id:       10,
		Consume1: 0,
		Consume2: 100,
	},
	{
		Id:       11,
		Consume1: 0,
		Consume2: 300,
	},
	{
		Id:       12,
		Consume1: 0,
		Consume2: 300,
	},
	{
		Id:       13,
		Consume1: 0,
		Consume2: 300,
	},
	{
		Id:       14,
		Consume1: 0,
		Consume2: 300,
	},
	{
		Id:       15,
		Consume1: 0,
		Consume2: 300,
	},
	{
		Id:       16,
		Consume1: 0,
		Consume2: 500,
	},
	{
		Id:       17,
		Consume1: 0,
		Consume2: 500,
	},
	{
		Id:       18,
		Consume1: 0,
		Consume2: 500,
	},
	{
		Id:       19,
		Consume1: 0,
		Consume2: 500,
	},
	{
		Id:       20,
		Consume1: 0,
		Consume2: 500,
	},
	{
		Id:       21,
		Consume1: 0,
		Consume2: 700,
	},
	{
		Id:       22,
		Consume1: 0,
		Consume2: 700,
	},
	{
		Id:       23,
		Consume1: 0,
		Consume2: 700,
	},
	{
		Id:       24,
		Consume1: 0,
		Consume2: 700,
	},
	{
		Id:       25,
		Consume1: 0,
		Consume2: 700,
	},
	{
		Id:       26,
		Consume1: 0,
		Consume2: 900,
	},
	{
		Id:       27,
		Consume1: 0,
		Consume2: 900,
	},
	{
		Id:       28,
		Consume1: 0,
		Consume2: 900,
	},
	{
		Id:       29,
		Consume1: 0,
		Consume2: 900,
	},
	{
		Id:       30,
		Consume1: 0,
		Consume2: 900,
	},
	{
		Id:       31,
		Consume1: 0,
		Consume2: 1100,
	},
	{
		Id:       32,
		Consume1: 0,
		Consume2: 1100,
	},
	{
		Id:       33,
		Consume1: 0,
		Consume2: 1100,
	},
	{
		Id:       34,
		Consume1: 0,
		Consume2: 1100,
	},
	{
		Id:       35,
		Consume1: 0,
		Consume2: 1100,
	},
	{
		Id:       36,
		Consume1: 0,
		Consume2: 1300,
	},
	{
		Id:       37,
		Consume1: 0,
		Consume2: 1300,
	},
	{
		Id:       38,
		Consume1: 0,
		Consume2: 1300,
	},
	{
		Id:       39,
		Consume1: 0,
		Consume2: 1300,
	},
	{
		Id:       40,
		Consume1: 0,
		Consume2: 1300,
	},
	{
		Id:       41,
		Consume1: 0,
		Consume2: 1500,
	},
	{
		Id:       42,
		Consume1: 0,
		Consume2: 1500,
	},
	{
		Id:       43,
		Consume1: 0,
		Consume2: 1500,
	},
	{
		Id:       44,
		Consume1: 0,
		Consume2: 1500,
	},
	{
		Id:       45,
		Consume1: 0,
		Consume2: 1500,
	},
	{
		Id:       46,
		Consume1: 0,
		Consume2: 1700,
	},
	{
		Id:       47,
		Consume1: 0,
		Consume2: 1700,
	},
	{
		Id:       48,
		Consume1: 0,
		Consume2: 1700,
	},
	{
		Id:       49,
		Consume1: 0,
		Consume2: 1700,
	},
	{
		Id:       50,
		Consume1: 0,
		Consume2: 1700,
	},
	{
		Id:       51,
		Consume1: 0,
		Consume2: 1900,
	},
	{
		Id:       52,
		Consume1: 0,
		Consume2: 1900,
	},
	{
		Id:       53,
		Consume1: 0,
		Consume2: 1900,
	},
	{
		Id:       54,
		Consume1: 0,
		Consume2: 1900,
	},
	{
		Id:       55,
		Consume1: 0,
		Consume2: 1900,
	},
	{
		Id:       56,
		Consume1: 0,
		Consume2: 2100,
	},
	{
		Id:       57,
		Consume1: 0,
		Consume2: 2100,
	},
	{
		Id:       58,
		Consume1: 0,
		Consume2: 2100,
	},
	{
		Id:       59,
		Consume1: 0,
		Consume2: 2100,
	},
	{
		Id:       60,
		Consume1: 0,
		Consume2: 2100,
	},
	{
		Id:       61,
		Consume1: 0,
		Consume2: 2300,
	},
	{
		Id:       62,
		Consume1: 0,
		Consume2: 2300,
	},
	{
		Id:       63,
		Consume1: 0,
		Consume2: 2300,
	},
	{
		Id:       64,
		Consume1: 0,
		Consume2: 2300,
	},
	{
		Id:       65,
		Consume1: 0,
		Consume2: 2300,
	},
	{
		Id:       66,
		Consume1: 0,
		Consume2: 2500,
	},
	{
		Id:       67,
		Consume1: 0,
		Consume2: 2500,
	},
	{
		Id:       68,
		Consume1: 0,
		Consume2: 2500,
	},
	{
		Id:       69,
		Consume1: 0,
		Consume2: 2500,
	},
	{
		Id:       70,
		Consume1: 0,
		Consume2: 2500,
	},
	{
		Id:       71,
		Consume1: 0,
		Consume2: 2700,
	},
	{
		Id:       72,
		Consume1: 0,
		Consume2: 2700,
	},
	{
		Id:       73,
		Consume1: 0,
		Consume2: 2700,
	},
	{
		Id:       74,
		Consume1: 0,
		Consume2: 2700,
	},
	{
		Id:       75,
		Consume1: 0,
		Consume2: 2700,
	},
	{
		Id:       76,
		Consume1: 0,
		Consume2: 2900,
	},
	{
		Id:       77,
		Consume1: 0,
		Consume2: 2900,
	},
	{
		Id:       78,
		Consume1: 0,
		Consume2: 2900,
	},
	{
		Id:       79,
		Consume1: 0,
		Consume2: 2900,
	},
	{
		Id:       80,
		Consume1: 0,
		Consume2: 2900,
	},
	{
		Id:       81,
		Consume1: 0,
		Consume2: 3100,
	},
	{
		Id:       82,
		Consume1: 0,
		Consume2: 3100,
	},
	{
		Id:       83,
		Consume1: 0,
		Consume2: 3100,
	},
	{
		Id:       84,
		Consume1: 0,
		Consume2: 3100,
	},
	{
		Id:       85,
		Consume1: 0,
		Consume2: 3100,
	},
	{
		Id:       86,
		Consume1: 0,
		Consume2: 3300,
	},
	{
		Id:       87,
		Consume1: 0,
		Consume2: 3300,
	},
	{
		Id:       88,
		Consume1: 0,
		Consume2: 3300,
	},
	{
		Id:       89,
		Consume1: 0,
		Consume2: 3300,
	},
	{
		Id:       90,
		Consume1: 0,
		Consume2: 3300,
	},
	{
		Id:       91,
		Consume1: 0,
		Consume2: 3500,
	},
	{
		Id:       92,
		Consume1: 0,
		Consume2: 3500,
	},
	{
		Id:       93,
		Consume1: 0,
		Consume2: 3500,
	},
	{
		Id:       94,
		Consume1: 0,
		Consume2: 3500,
	},
	{
		Id:       95,
		Consume1: 0,
		Consume2: 3500,
	},
	{
		Id:       96,
		Consume1: 0,
		Consume2: 3700,
	},
	{
		Id:       97,
		Consume1: 0,
		Consume2: 3700,
	},
	{
		Id:       98,
		Consume1: 0,
		Consume2: 3700,
	},
	{
		Id:       99,
		Consume1: 0,
		Consume2: 3700,
	},
	{
		Id:       100,
		Consume1: 0,
		Consume2: 3700,
	},
}

var workerKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
}

func init() {
	WorkerData.data = make(map[int32]*WorkerCfg)
	for i := 0; i < len(workerKeys); i++ {
		WorkerData.data[workerKeys[i]] = workerValues[i]
	}
}
