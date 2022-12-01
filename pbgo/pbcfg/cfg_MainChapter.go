package config

type MainChapterCfg struct {
	Id       int32
	LevelNum int32
	Scene    string
}

type MainChapterTable struct {
	data map[int32]*MainChapterCfg
}

var MainChapterData = &MainChapterTable{
	data: map[int32]*MainChapterCfg{},
}

func (table *MainChapterTable) Get(id int32) *MainChapterCfg {
	return table.data[id]
}

func (table *MainChapterTable) GetAll() []int32 {
	return mainChapterKeys
}

var mainChapterValues = []*MainChapterCfg{
	{
		Id:       1,
		LevelNum: 10,
	},
	{
		Id:       2,
		LevelNum: 10,
	},
	{
		Id:       3,
		LevelNum: 10,
	},
	{
		Id:       4,
		LevelNum: 10,
	},
	{
		Id:       5,
		LevelNum: 10,
	},
	{
		Id:       6,
		LevelNum: 10,
	},
	{
		Id:       7,
		LevelNum: 10,
	},
	{
		Id:       8,
		LevelNum: 10,
	},
	{
		Id:       9,
		LevelNum: 10,
	},
	{
		Id:       10,
		LevelNum: 10,
	},
	{
		Id:       11,
		LevelNum: 10,
	},
	{
		Id:       12,
		LevelNum: 10,
	},
	{
		Id:       13,
		LevelNum: 10,
	},
	{
		Id:       14,
		LevelNum: 10,
	},
	{
		Id:       15,
		LevelNum: 10,
	},
	{
		Id:       16,
		LevelNum: 10,
	},
	{
		Id:       17,
		LevelNum: 10,
	},
	{
		Id:       18,
		LevelNum: 10,
	},
	{
		Id:       19,
		LevelNum: 10,
	},
	{
		Id:       20,
		LevelNum: 10,
	},
	{
		Id:       21,
		LevelNum: 10,
	},
	{
		Id:       22,
		LevelNum: 10,
	},
	{
		Id:       23,
		LevelNum: 10,
	},
	{
		Id:       24,
		LevelNum: 10,
	},
	{
		Id:       25,
		LevelNum: 10,
	},
	{
		Id:       26,
		LevelNum: 10,
	},
	{
		Id:       27,
		LevelNum: 10,
	},
	{
		Id:       28,
		LevelNum: 10,
	},
	{
		Id:       29,
		LevelNum: 10,
	},
	{
		Id:       30,
		LevelNum: 10,
	},
	{
		Id:       31,
		LevelNum: 10,
	},
	{
		Id:       32,
		LevelNum: 10,
	},
	{
		Id:       33,
		LevelNum: 10,
	},
	{
		Id:       34,
		LevelNum: 10,
	},
	{
		Id:       35,
		LevelNum: 10,
	},
	{
		Id:       36,
		LevelNum: 10,
	},
	{
		Id:       37,
		LevelNum: 10,
	},
	{
		Id:       38,
		LevelNum: 10,
	},
	{
		Id:       39,
		LevelNum: 10,
	},
	{
		Id:       40,
		LevelNum: 10,
	},
	{
		Id:       41,
		LevelNum: 10,
	},
	{
		Id:       42,
		LevelNum: 10,
	},
	{
		Id:       43,
		LevelNum: 10,
	},
	{
		Id:       44,
		LevelNum: 10,
	},
	{
		Id:       45,
		LevelNum: 10,
	},
	{
		Id:       46,
		LevelNum: 10,
	},
	{
		Id:       47,
		LevelNum: 10,
	},
	{
		Id:       48,
		LevelNum: 10,
	},
	{
		Id:       49,
		LevelNum: 10,
	},
	{
		Id:       50,
		LevelNum: 10,
	},
	{
		Id:       51,
		LevelNum: 10,
	},
	{
		Id:       52,
		LevelNum: 10,
	},
	{
		Id:       53,
		LevelNum: 10,
	},
	{
		Id:       54,
		LevelNum: 10,
	},
	{
		Id:       55,
		LevelNum: 10,
	},
	{
		Id:       56,
		LevelNum: 10,
	},
	{
		Id:       57,
		LevelNum: 10,
	},
	{
		Id:       58,
		LevelNum: 10,
	},
	{
		Id:       59,
		LevelNum: 10,
	},
	{
		Id:       60,
		LevelNum: 10,
	},
	{
		Id:       61,
		LevelNum: 10,
	},
	{
		Id:       62,
		LevelNum: 10,
	},
	{
		Id:       63,
		LevelNum: 10,
	},
	{
		Id:       64,
		LevelNum: 10,
	},
	{
		Id:       65,
		LevelNum: 10,
	},
	{
		Id:       66,
		LevelNum: 10,
	},
	{
		Id:       67,
		LevelNum: 10,
	},
	{
		Id:       68,
		LevelNum: 10,
	},
	{
		Id:       69,
		LevelNum: 10,
	},
	{
		Id:       70,
		LevelNum: 10,
	},
	{
		Id:       71,
		LevelNum: 10,
	},
	{
		Id:       72,
		LevelNum: 10,
	},
	{
		Id:       73,
		LevelNum: 10,
	},
	{
		Id:       74,
		LevelNum: 10,
	},
	{
		Id:       75,
		LevelNum: 10,
	},
	{
		Id:       76,
		LevelNum: 10,
	},
	{
		Id:       77,
		LevelNum: 10,
	},
	{
		Id:       78,
		LevelNum: 10,
	},
	{
		Id:       79,
		LevelNum: 10,
	},
	{
		Id:       80,
		LevelNum: 10,
	},
	{
		Id:       81,
		LevelNum: 10,
	},
	{
		Id:       82,
		LevelNum: 10,
	},
	{
		Id:       83,
		LevelNum: 10,
	},
	{
		Id:       84,
		LevelNum: 10,
	},
	{
		Id:       85,
		LevelNum: 10,
	},
	{
		Id:       86,
		LevelNum: 10,
	},
	{
		Id:       87,
		LevelNum: 10,
	},
	{
		Id:       88,
		LevelNum: 10,
	},
	{
		Id:       89,
		LevelNum: 10,
	},
	{
		Id:       90,
		LevelNum: 10,
	},
	{
		Id:       91,
		LevelNum: 10,
	},
	{
		Id:       92,
		LevelNum: 10,
	},
	{
		Id:       93,
		LevelNum: 10,
	},
	{
		Id:       94,
		LevelNum: 10,
	},
	{
		Id:       95,
		LevelNum: 10,
	},
	{
		Id:       96,
		LevelNum: 10,
	},
	{
		Id:       97,
		LevelNum: 10,
	},
	{
		Id:       98,
		LevelNum: 10,
	},
	{
		Id:       99,
		LevelNum: 10,
	},
	{
		Id:       100,
		LevelNum: 10,
	},
	{
		Id:       101,
		LevelNum: 10,
	},
	{
		Id:       102,
		LevelNum: 10,
	},
	{
		Id:       103,
		LevelNum: 10,
	},
	{
		Id:       104,
		LevelNum: 10,
	},
	{
		Id:       105,
		LevelNum: 10,
	},
	{
		Id:       106,
		LevelNum: 10,
	},
	{
		Id:       107,
		LevelNum: 10,
	},
	{
		Id:       108,
		LevelNum: 10,
	},
	{
		Id:       109,
		LevelNum: 10,
	},
	{
		Id:       110,
		LevelNum: 10,
	},
	{
		Id:       111,
		LevelNum: 10,
	},
	{
		Id:       112,
		LevelNum: 10,
	},
	{
		Id:       113,
		LevelNum: 10,
	},
	{
		Id:       114,
		LevelNum: 10,
	},
	{
		Id:       115,
		LevelNum: 10,
	},
	{
		Id:       116,
		LevelNum: 10,
	},
	{
		Id:       117,
		LevelNum: 10,
	},
	{
		Id:       118,
		LevelNum: 10,
	},
	{
		Id:       119,
		LevelNum: 10,
	},
	{
		Id:       120,
		LevelNum: 10,
	},
	{
		Id:       121,
		LevelNum: 10,
	},
	{
		Id:       122,
		LevelNum: 10,
	},
	{
		Id:       123,
		LevelNum: 10,
	},
	{
		Id:       124,
		LevelNum: 10,
	},
	{
		Id:       125,
		LevelNum: 10,
	},
	{
		Id:       126,
		LevelNum: 10,
	},
	{
		Id:       127,
		LevelNum: 10,
	},
	{
		Id:       128,
		LevelNum: 10,
	},
	{
		Id:       129,
		LevelNum: 10,
	},
	{
		Id:       130,
		LevelNum: 10,
	},
	{
		Id:       131,
		LevelNum: 10,
	},
	{
		Id:       132,
		LevelNum: 10,
	},
	{
		Id:       133,
		LevelNum: 10,
	},
	{
		Id:       134,
		LevelNum: 10,
	},
	{
		Id:       135,
		LevelNum: 10,
	},
	{
		Id:       136,
		LevelNum: 10,
	},
	{
		Id:       137,
		LevelNum: 10,
	},
	{
		Id:       138,
		LevelNum: 10,
	},
	{
		Id:       139,
		LevelNum: 10,
	},
	{
		Id:       140,
		LevelNum: 10,
	},
	{
		Id:       141,
		LevelNum: 10,
	},
	{
		Id:       142,
		LevelNum: 10,
	},
	{
		Id:       143,
		LevelNum: 10,
	},
	{
		Id:       144,
		LevelNum: 10,
	},
	{
		Id:       145,
		LevelNum: 10,
	},
	{
		Id:       146,
		LevelNum: 10,
	},
	{
		Id:       147,
		LevelNum: 10,
	},
	{
		Id:       148,
		LevelNum: 10,
	},
	{
		Id:       149,
		LevelNum: 10,
	},
	{
		Id:       150,
		LevelNum: 10,
	},
	{
		Id:       151,
		LevelNum: 10,
	},
	{
		Id:       152,
		LevelNum: 10,
	},
	{
		Id:       153,
		LevelNum: 10,
	},
	{
		Id:       154,
		LevelNum: 10,
	},
	{
		Id:       155,
		LevelNum: 10,
	},
	{
		Id:       156,
		LevelNum: 10,
	},
	{
		Id:       157,
		LevelNum: 10,
	},
	{
		Id:       158,
		LevelNum: 10,
	},
	{
		Id:       159,
		LevelNum: 10,
	},
	{
		Id:       160,
		LevelNum: 10,
	},
	{
		Id:       161,
		LevelNum: 10,
	},
	{
		Id:       162,
		LevelNum: 10,
	},
	{
		Id:       163,
		LevelNum: 10,
	},
	{
		Id:       164,
		LevelNum: 10,
	},
	{
		Id:       165,
		LevelNum: 10,
	},
	{
		Id:       166,
		LevelNum: 10,
	},
	{
		Id:       167,
		LevelNum: 10,
	},
	{
		Id:       168,
		LevelNum: 10,
	},
	{
		Id:       169,
		LevelNum: 10,
	},
	{
		Id:       170,
		LevelNum: 10,
	},
	{
		Id:       171,
		LevelNum: 10,
	},
	{
		Id:       172,
		LevelNum: 10,
	},
	{
		Id:       173,
		LevelNum: 10,
	},
	{
		Id:       174,
		LevelNum: 10,
	},
	{
		Id:       175,
		LevelNum: 10,
	},
	{
		Id:       176,
		LevelNum: 10,
	},
	{
		Id:       177,
		LevelNum: 10,
	},
	{
		Id:       178,
		LevelNum: 10,
	},
	{
		Id:       179,
		LevelNum: 10,
	},
	{
		Id:       180,
		LevelNum: 10,
	},
	{
		Id:       181,
		LevelNum: 10,
	},
	{
		Id:       182,
		LevelNum: 10,
	},
	{
		Id:       183,
		LevelNum: 10,
	},
	{
		Id:       184,
		LevelNum: 10,
	},
	{
		Id:       185,
		LevelNum: 10,
	},
	{
		Id:       186,
		LevelNum: 10,
	},
	{
		Id:       187,
		LevelNum: 10,
	},
	{
		Id:       188,
		LevelNum: 10,
	},
	{
		Id:       189,
		LevelNum: 10,
	},
	{
		Id:       190,
		LevelNum: 10,
	},
	{
		Id:       191,
		LevelNum: 10,
	},
	{
		Id:       192,
		LevelNum: 10,
	},
	{
		Id:       193,
		LevelNum: 10,
	},
	{
		Id:       194,
		LevelNum: 10,
	},
	{
		Id:       195,
		LevelNum: 10,
	},
	{
		Id:       196,
		LevelNum: 10,
	},
	{
		Id:       197,
		LevelNum: 10,
	},
	{
		Id:       198,
		LevelNum: 10,
	},
	{
		Id:       199,
		LevelNum: 10,
	},
	{
		Id:       200,
		LevelNum: 10,
	},
	{
		Id:       201,
		LevelNum: 10,
	},
	{
		Id:       202,
		LevelNum: 10,
	},
	{
		Id:       203,
		LevelNum: 10,
	},
	{
		Id:       204,
		LevelNum: 10,
	},
	{
		Id:       205,
		LevelNum: 10,
	},
	{
		Id:       206,
		LevelNum: 10,
	},
	{
		Id:       207,
		LevelNum: 10,
	},
	{
		Id:       208,
		LevelNum: 10,
	},
	{
		Id:       209,
		LevelNum: 10,
	},
	{
		Id:       210,
		LevelNum: 10,
	},
	{
		Id:       211,
		LevelNum: 10,
	},
}

var mainChapterKeys = []int32{
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
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 175, 176, 177, 178, 179, 180,
	181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
	191, 192, 193, 194, 195, 196, 197, 198, 199, 200,
	201, 202, 203, 204, 205, 206, 207, 208, 209, 210,
	211,
}

func init() {
	MainChapterData.data = make(map[int32]*MainChapterCfg)
	for i := 0; i < len(mainChapterKeys); i++ {
		MainChapterData.data[mainChapterKeys[i]] = mainChapterValues[i]
	}
}
