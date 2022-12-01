package config

type BuildingResourceLvCfg struct {
	Id             int32
	BuildingId     int32
	Lv             int32
	LevelUpConsume []*Consume
	MaxSave        int32
	MaxWorker      int32
}

type BuildingResourceLvTable struct {
	data                                      map[int32]*BuildingResourceLvCfg
	buildingResourceLvBuildingIdIndexMap      map[int32][]int32
	buildingResourceLvBuildingIdAndLvIndexMap map[buildingResourceLvBuildingIdAndLvIndex][]int32
}

var BuildingResourceLvData = &BuildingResourceLvTable{
	data: map[int32]*BuildingResourceLvCfg{},
	buildingResourceLvBuildingIdIndexMap: map[int32][]int32{
		1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		2: {101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150},
		3: {201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250},
		4: {301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350},
	},
	buildingResourceLvBuildingIdAndLvIndexMap: map[buildingResourceLvBuildingIdAndLvIndex][]int32{
		{1, 10}: {10},
		{1, 11}: {11},
		{1, 12}: {12},
		{1, 13}: {13},
		{1, 14}: {14},
		{1, 15}: {15},
		{1, 16}: {16},
		{1, 17}: {17},
		{1, 18}: {18},
		{1, 19}: {19},
		{1, 1}:  {1},
		{1, 20}: {20},
		{1, 21}: {21},
		{1, 22}: {22},
		{1, 23}: {23},
		{1, 24}: {24},
		{1, 25}: {25},
		{1, 26}: {26},
		{1, 27}: {27},
		{1, 28}: {28},
		{1, 29}: {29},
		{1, 2}:  {2},
		{1, 30}: {30},
		{1, 31}: {31},
		{1, 32}: {32},
		{1, 33}: {33},
		{1, 34}: {34},
		{1, 35}: {35},
		{1, 36}: {36},
		{1, 37}: {37},
		{1, 38}: {38},
		{1, 39}: {39},
		{1, 3}:  {3},
		{1, 40}: {40},
		{1, 41}: {41},
		{1, 42}: {42},
		{1, 43}: {43},
		{1, 44}: {44},
		{1, 45}: {45},
		{1, 46}: {46},
		{1, 47}: {47},
		{1, 48}: {48},
		{1, 49}: {49},
		{1, 4}:  {4},
		{1, 50}: {50},
		{1, 5}:  {5},
		{1, 6}:  {6},
		{1, 7}:  {7},
		{1, 8}:  {8},
		{1, 9}:  {9},
		{2, 10}: {110},
		{2, 11}: {111},
		{2, 12}: {112},
		{2, 13}: {113},
		{2, 14}: {114},
		{2, 15}: {115},
		{2, 16}: {116},
		{2, 17}: {117},
		{2, 18}: {118},
		{2, 19}: {119},
		{2, 1}:  {101},
		{2, 20}: {120},
		{2, 21}: {121},
		{2, 22}: {122},
		{2, 23}: {123},
		{2, 24}: {124},
		{2, 25}: {125},
		{2, 26}: {126},
		{2, 27}: {127},
		{2, 28}: {128},
		{2, 29}: {129},
		{2, 2}:  {102},
		{2, 30}: {130},
		{2, 31}: {131},
		{2, 32}: {132},
		{2, 33}: {133},
		{2, 34}: {134},
		{2, 35}: {135},
		{2, 36}: {136},
		{2, 37}: {137},
		{2, 38}: {138},
		{2, 39}: {139},
		{2, 3}:  {103},
		{2, 40}: {140},
		{2, 41}: {141},
		{2, 42}: {142},
		{2, 43}: {143},
		{2, 44}: {144},
		{2, 45}: {145},
		{2, 46}: {146},
		{2, 47}: {147},
		{2, 48}: {148},
		{2, 49}: {149},
		{2, 4}:  {104},
		{2, 50}: {150},
		{2, 5}:  {105},
		{2, 6}:  {106},
		{2, 7}:  {107},
		{2, 8}:  {108},
		{2, 9}:  {109},
		{3, 10}: {210},
		{3, 11}: {211},
		{3, 12}: {212},
		{3, 13}: {213},
		{3, 14}: {214},
		{3, 15}: {215},
		{3, 16}: {216},
		{3, 17}: {217},
		{3, 18}: {218},
		{3, 19}: {219},
		{3, 1}:  {201},
		{3, 20}: {220},
		{3, 21}: {221},
		{3, 22}: {222},
		{3, 23}: {223},
		{3, 24}: {224},
		{3, 25}: {225},
		{3, 26}: {226},
		{3, 27}: {227},
		{3, 28}: {228},
		{3, 29}: {229},
		{3, 2}:  {202},
		{3, 30}: {230},
		{3, 31}: {231},
		{3, 32}: {232},
		{3, 33}: {233},
		{3, 34}: {234},
		{3, 35}: {235},
		{3, 36}: {236},
		{3, 37}: {237},
		{3, 38}: {238},
		{3, 39}: {239},
		{3, 3}:  {203},
		{3, 40}: {240},
		{3, 41}: {241},
		{3, 42}: {242},
		{3, 43}: {243},
		{3, 44}: {244},
		{3, 45}: {245},
		{3, 46}: {246},
		{3, 47}: {247},
		{3, 48}: {248},
		{3, 49}: {249},
		{3, 4}:  {204},
		{3, 50}: {250},
		{3, 5}:  {205},
		{3, 6}:  {206},
		{3, 7}:  {207},
		{3, 8}:  {208},
		{3, 9}:  {209},
		{4, 10}: {310},
		{4, 11}: {311},
		{4, 12}: {312},
		{4, 13}: {313},
		{4, 14}: {314},
		{4, 15}: {315},
		{4, 16}: {316},
		{4, 17}: {317},
		{4, 18}: {318},
		{4, 19}: {319},
		{4, 1}:  {301},
		{4, 20}: {320},
		{4, 21}: {321},
		{4, 22}: {322},
		{4, 23}: {323},
		{4, 24}: {324},
		{4, 25}: {325},
		{4, 26}: {326},
		{4, 27}: {327},
		{4, 28}: {328},
		{4, 29}: {329},
		{4, 2}:  {302},
		{4, 30}: {330},
		{4, 31}: {331},
		{4, 32}: {332},
		{4, 33}: {333},
		{4, 34}: {334},
		{4, 35}: {335},
		{4, 36}: {336},
		{4, 37}: {337},
		{4, 38}: {338},
		{4, 39}: {339},
		{4, 3}:  {303},
		{4, 40}: {340},
		{4, 41}: {341},
		{4, 42}: {342},
		{4, 43}: {343},
		{4, 44}: {344},
		{4, 45}: {345},
		{4, 46}: {346},
		{4, 47}: {347},
		{4, 48}: {348},
		{4, 49}: {349},
		{4, 4}:  {304},
		{4, 50}: {350},
		{4, 5}:  {305},
		{4, 6}:  {306},
		{4, 7}:  {307},
		{4, 8}:  {308},
		{4, 9}:  {309},
	},
}

func (table *BuildingResourceLvTable) Get(id int32) *BuildingResourceLvCfg {
	return table.data[id]
}

func (table *BuildingResourceLvTable) GetAll() []int32 {
	return buildingResourceLvKeys
}

var buildingResourceLvValues = []*BuildingResourceLvCfg{
	{
		Id:         1,
		BuildingId: 1,
		Lv:         1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3600,
			},
		},
		MaxSave:   20000,
		MaxWorker: 4,
	},
	{
		Id:         2,
		BuildingId: 1,
		Lv:         2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 6480,
			},
		},
		MaxSave:   25000,
		MaxWorker: 5,
	},
	{
		Id:         3,
		BuildingId: 1,
		Lv:         3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 10000,
			},
		},
		MaxSave:   32400,
		MaxWorker: 6,
	},
	{
		Id:         4,
		BuildingId: 1,
		Lv:         4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 14000,
			},
		},
		MaxSave:   50400,
		MaxWorker: 7,
	},
	{
		Id:         5,
		BuildingId: 1,
		Lv:         5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 19000,
			},
		},
		MaxSave:   86400,
		MaxWorker: 8,
	},
	{
		Id:         6,
		BuildingId: 1,
		Lv:         6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 25000,
			},
		},
		MaxSave:   129600,
		MaxWorker: 9,
	},
	{
		Id:         7,
		BuildingId: 1,
		Lv:         7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 31000,
			},
		},
		MaxSave:   180000,
		MaxWorker: 10,
	},
	{
		Id:         8,
		BuildingId: 1,
		Lv:         8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 38000,
			},
		},
		MaxSave:   237600,
		MaxWorker: 11,
	},
	{
		Id:         9,
		BuildingId: 1,
		Lv:         9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 46000,
			},
		},
		MaxSave:   302400,
		MaxWorker: 12,
	},
	{
		Id:         10,
		BuildingId: 1,
		Lv:         10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 55000,
			},
		},
		MaxSave:   421200,
		MaxWorker: 13,
	},
	{
		Id:         11,
		BuildingId: 1,
		Lv:         11,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 69000,
			},
		},
		MaxSave:   554400,
		MaxWorker: 14,
	},
	{
		Id:         12,
		BuildingId: 1,
		Lv:         12,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 84000,
			},
		},
		MaxSave:   702000,
		MaxWorker: 15,
	},
	{
		Id:         13,
		BuildingId: 1,
		Lv:         13,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 100000,
			},
		},
		MaxSave:   864000,
		MaxWorker: 16,
	},
	{
		Id:         14,
		BuildingId: 1,
		Lv:         14,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 110000,
			},
		},
		MaxSave:   1040400,
		MaxWorker: 17,
	},
	{
		Id:         15,
		BuildingId: 1,
		Lv:         15,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 138000,
			},
		},
		MaxSave:   1296000,
		MaxWorker: 18,
	},
	{
		Id:         16,
		BuildingId: 1,
		Lv:         16,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 159000,
			},
		},
		MaxSave:   1573200,
		MaxWorker: 19,
	},
	{
		Id:         17,
		BuildingId: 1,
		Lv:         17,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 181000,
			},
		},
		MaxSave:   1872000,
		MaxWorker: 20,
	},
	{
		Id:         18,
		BuildingId: 1,
		Lv:         18,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 205000,
			},
		},
		MaxSave:   2192400,
		MaxWorker: 21,
	},
	{
		Id:         19,
		BuildingId: 1,
		Lv:         19,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 230000,
			},
		},
		MaxSave:   2534400,
		MaxWorker: 22,
	},
	{
		Id:         20,
		BuildingId: 1,
		Lv:         20,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 257000,
			},
		},
		MaxSave:   2898000,
		MaxWorker: 23,
	},
	{
		Id:         21,
		BuildingId: 1,
		Lv:         21,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 293000,
			},
		},
		MaxSave:   3283200,
		MaxWorker: 24,
	},
	{
		Id:         22,
		BuildingId: 1,
		Lv:         22,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 331000,
			},
		},
		MaxSave:   3690000,
		MaxWorker: 25,
	},
	{
		Id:         23,
		BuildingId: 1,
		Lv:         23,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 371000,
			},
		},
		MaxSave:   4118400,
		MaxWorker: 26,
	},
	{
		Id:         24,
		BuildingId: 1,
		Lv:         24,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 414000,
			},
		},
		MaxSave:   4568400,
		MaxWorker: 27,
	},
	{
		Id:         25,
		BuildingId: 1,
		Lv:         25,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 458000,
			},
		},
		MaxSave:   5040000,
		MaxWorker: 28,
	},
	{
		Id:         26,
		BuildingId: 1,
		Lv:         26,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 505000,
			},
		},
		MaxSave:   5533200,
		MaxWorker: 29,
	},
	{
		Id:         27,
		BuildingId: 1,
		Lv:         27,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 554000,
			},
		},
		MaxSave:   6048000,
		MaxWorker: 30,
	},
	{
		Id:         28,
		BuildingId: 1,
		Lv:         28,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 605000,
			},
		},
		MaxSave:   6584400,
		MaxWorker: 31,
	},
	{
		Id:         29,
		BuildingId: 1,
		Lv:         29,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 658000,
			},
		},
		MaxSave:   7142400,
		MaxWorker: 32,
	},
	{
		Id:         30,
		BuildingId: 1,
		Lv:         30,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 714000,
			},
		},
		MaxSave:   7722000,
		MaxWorker: 33,
	},
	{
		Id:         31,
		BuildingId: 1,
		Lv:         31,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 794000,
			},
		},
		MaxSave:   8323200,
		MaxWorker: 34,
	},
	{
		Id:         32,
		BuildingId: 1,
		Lv:         32,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 879000,
			},
		},
		MaxSave:   8946000,
		MaxWorker: 35,
	},
	{
		Id:         33,
		BuildingId: 1,
		Lv:         33,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 966000,
			},
		},
		MaxSave:   9590400,
		MaxWorker: 36,
	},
	{
		Id:         34,
		BuildingId: 1,
		Lv:         34,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1058000,
			},
		},
		MaxSave:   10256400,
		MaxWorker: 37,
	},
	{
		Id:         35,
		BuildingId: 1,
		Lv:         35,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1153000,
			},
		},
		MaxSave:   10944000,
		MaxWorker: 38,
	},
	{
		Id:         36,
		BuildingId: 1,
		Lv:         36,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1252000,
			},
		},
		MaxSave:   11653200,
		MaxWorker: 39,
	},
	{
		Id:         37,
		BuildingId: 1,
		Lv:         37,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1354000,
			},
		},
		MaxSave:   12384000,
		MaxWorker: 40,
	},
	{
		Id:         38,
		BuildingId: 1,
		Lv:         38,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1460000,
			},
		},
		MaxSave:   13136400,
		MaxWorker: 41,
	},
	{
		Id:         39,
		BuildingId: 1,
		Lv:         39,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1569000,
			},
		},
		MaxSave:   13910400,
		MaxWorker: 42,
	},
	{
		Id:         40,
		BuildingId: 1,
		Lv:         40,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1682000,
			},
		},
		MaxSave:   14706000,
		MaxWorker: 43,
	},
	{
		Id:         41,
		BuildingId: 1,
		Lv:         41,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1874000,
			},
		},
		MaxSave:   15523200,
		MaxWorker: 44,
	},
	{
		Id:         42,
		BuildingId: 1,
		Lv:         42,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2074000,
			},
		},
		MaxSave:   16362000,
		MaxWorker: 45,
	},
	{
		Id:         43,
		BuildingId: 1,
		Lv:         43,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2280000,
			},
		},
		MaxSave:   17222400,
		MaxWorker: 46,
	},
	{
		Id:         44,
		BuildingId: 1,
		Lv:         44,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2494000,
			},
		},
		MaxSave:   18104400,
		MaxWorker: 47,
	},
	{
		Id:         45,
		BuildingId: 1,
		Lv:         45,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2715000,
			},
		},
		MaxSave:   19008000,
		MaxWorker: 48,
	},
	{
		Id:         46,
		BuildingId: 1,
		Lv:         46,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2944000,
			},
		},
		MaxSave:   19933200,
		MaxWorker: 49,
	},
	{
		Id:         47,
		BuildingId: 1,
		Lv:         47,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3179000,
			},
		},
		MaxSave:   20880000,
		MaxWorker: 50,
	},
	{
		Id:         48,
		BuildingId: 1,
		Lv:         48,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3422000,
			},
		},
		MaxSave:   21848400,
		MaxWorker: 51,
	},
	{
		Id:         49,
		BuildingId: 1,
		Lv:         49,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3672000,
			},
		},
		MaxSave:   22838400,
		MaxWorker: 52,
	},
	{
		Id:         50,
		BuildingId: 1,
		Lv:         50,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3929000,
			},
		},
		MaxSave:   23850000,
		MaxWorker: 53,
	},
	{
		Id:         101,
		BuildingId: 2,
		Lv:         1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3600,
			},
		},
		MaxSave:   20000,
		MaxWorker: 2,
	},
	{
		Id:         102,
		BuildingId: 2,
		Lv:         2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 6480,
			},
		},
		MaxSave:   25000,
		MaxWorker: 3,
	},
	{
		Id:         103,
		BuildingId: 2,
		Lv:         3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 10000,
			},
		},
		MaxSave:   32400,
		MaxWorker: 4,
	},
	{
		Id:         104,
		BuildingId: 2,
		Lv:         4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 14000,
			},
		},
		MaxSave:   50400,
		MaxWorker: 5,
	},
	{
		Id:         105,
		BuildingId: 2,
		Lv:         5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 19000,
			},
		},
		MaxSave:   86400,
		MaxWorker: 6,
	},
	{
		Id:         106,
		BuildingId: 2,
		Lv:         6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 25000,
			},
		},
		MaxSave:   129600,
		MaxWorker: 7,
	},
	{
		Id:         107,
		BuildingId: 2,
		Lv:         7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 31000,
			},
		},
		MaxSave:   180000,
		MaxWorker: 8,
	},
	{
		Id:         108,
		BuildingId: 2,
		Lv:         8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 38000,
			},
		},
		MaxSave:   237600,
		MaxWorker: 9,
	},
	{
		Id:         109,
		BuildingId: 2,
		Lv:         9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 46000,
			},
		},
		MaxSave:   302400,
		MaxWorker: 10,
	},
	{
		Id:         110,
		BuildingId: 2,
		Lv:         10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 55000,
			},
		},
		MaxSave:   421200,
		MaxWorker: 11,
	},
	{
		Id:         111,
		BuildingId: 2,
		Lv:         11,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 69000,
			},
		},
		MaxSave:   554400,
		MaxWorker: 12,
	},
	{
		Id:         112,
		BuildingId: 2,
		Lv:         12,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 84000,
			},
		},
		MaxSave:   702000,
		MaxWorker: 13,
	},
	{
		Id:         113,
		BuildingId: 2,
		Lv:         13,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 100000,
			},
		},
		MaxSave:   864000,
		MaxWorker: 14,
	},
	{
		Id:         114,
		BuildingId: 2,
		Lv:         14,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 110000,
			},
		},
		MaxSave:   1040400,
		MaxWorker: 15,
	},
	{
		Id:         115,
		BuildingId: 2,
		Lv:         15,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 138000,
			},
		},
		MaxSave:   1296000,
		MaxWorker: 16,
	},
	{
		Id:         116,
		BuildingId: 2,
		Lv:         16,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 159000,
			},
		},
		MaxSave:   1573200,
		MaxWorker: 17,
	},
	{
		Id:         117,
		BuildingId: 2,
		Lv:         17,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 181000,
			},
		},
		MaxSave:   1872000,
		MaxWorker: 18,
	},
	{
		Id:         118,
		BuildingId: 2,
		Lv:         18,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 205000,
			},
		},
		MaxSave:   2192400,
		MaxWorker: 19,
	},
	{
		Id:         119,
		BuildingId: 2,
		Lv:         19,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 230000,
			},
		},
		MaxSave:   2534400,
		MaxWorker: 20,
	},
	{
		Id:         120,
		BuildingId: 2,
		Lv:         20,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 257000,
			},
		},
		MaxSave:   2898000,
		MaxWorker: 21,
	},
	{
		Id:         121,
		BuildingId: 2,
		Lv:         21,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 293000,
			},
		},
		MaxSave:   3283200,
		MaxWorker: 22,
	},
	{
		Id:         122,
		BuildingId: 2,
		Lv:         22,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 331000,
			},
		},
		MaxSave:   3690000,
		MaxWorker: 23,
	},
	{
		Id:         123,
		BuildingId: 2,
		Lv:         23,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 371000,
			},
		},
		MaxSave:   4118400,
		MaxWorker: 24,
	},
	{
		Id:         124,
		BuildingId: 2,
		Lv:         24,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 414000,
			},
		},
		MaxSave:   4568400,
		MaxWorker: 25,
	},
	{
		Id:         125,
		BuildingId: 2,
		Lv:         25,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 458000,
			},
		},
		MaxSave:   5040000,
		MaxWorker: 26,
	},
	{
		Id:         126,
		BuildingId: 2,
		Lv:         26,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 505000,
			},
		},
		MaxSave:   5533200,
		MaxWorker: 27,
	},
	{
		Id:         127,
		BuildingId: 2,
		Lv:         27,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 554000,
			},
		},
		MaxSave:   6048000,
		MaxWorker: 28,
	},
	{
		Id:         128,
		BuildingId: 2,
		Lv:         28,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 605000,
			},
		},
		MaxSave:   6584400,
		MaxWorker: 29,
	},
	{
		Id:         129,
		BuildingId: 2,
		Lv:         29,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 658000,
			},
		},
		MaxSave:   7142400,
		MaxWorker: 30,
	},
	{
		Id:         130,
		BuildingId: 2,
		Lv:         30,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 714000,
			},
		},
		MaxSave:   7722000,
		MaxWorker: 31,
	},
	{
		Id:         131,
		BuildingId: 2,
		Lv:         31,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 794000,
			},
		},
		MaxSave:   8323200,
		MaxWorker: 32,
	},
	{
		Id:         132,
		BuildingId: 2,
		Lv:         32,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 879000,
			},
		},
		MaxSave:   8946000,
		MaxWorker: 33,
	},
	{
		Id:         133,
		BuildingId: 2,
		Lv:         33,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 966000,
			},
		},
		MaxSave:   9590400,
		MaxWorker: 34,
	},
	{
		Id:         134,
		BuildingId: 2,
		Lv:         34,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1058000,
			},
		},
		MaxSave:   10256400,
		MaxWorker: 35,
	},
	{
		Id:         135,
		BuildingId: 2,
		Lv:         35,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1153000,
			},
		},
		MaxSave:   10944000,
		MaxWorker: 36,
	},
	{
		Id:         136,
		BuildingId: 2,
		Lv:         36,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1252000,
			},
		},
		MaxSave:   11653200,
		MaxWorker: 37,
	},
	{
		Id:         137,
		BuildingId: 2,
		Lv:         37,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1354000,
			},
		},
		MaxSave:   12384000,
		MaxWorker: 38,
	},
	{
		Id:         138,
		BuildingId: 2,
		Lv:         38,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1460000,
			},
		},
		MaxSave:   13136400,
		MaxWorker: 39,
	},
	{
		Id:         139,
		BuildingId: 2,
		Lv:         39,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1569000,
			},
		},
		MaxSave:   13910400,
		MaxWorker: 40,
	},
	{
		Id:         140,
		BuildingId: 2,
		Lv:         40,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1682000,
			},
		},
		MaxSave:   14706000,
		MaxWorker: 41,
	},
	{
		Id:         141,
		BuildingId: 2,
		Lv:         41,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1874000,
			},
		},
		MaxSave:   15523200,
		MaxWorker: 42,
	},
	{
		Id:         142,
		BuildingId: 2,
		Lv:         42,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2074000,
			},
		},
		MaxSave:   16362000,
		MaxWorker: 43,
	},
	{
		Id:         143,
		BuildingId: 2,
		Lv:         43,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2280000,
			},
		},
		MaxSave:   17222400,
		MaxWorker: 44,
	},
	{
		Id:         144,
		BuildingId: 2,
		Lv:         44,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2494000,
			},
		},
		MaxSave:   18104400,
		MaxWorker: 45,
	},
	{
		Id:         145,
		BuildingId: 2,
		Lv:         45,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2715000,
			},
		},
		MaxSave:   19008000,
		MaxWorker: 46,
	},
	{
		Id:         146,
		BuildingId: 2,
		Lv:         46,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2944000,
			},
		},
		MaxSave:   19933200,
		MaxWorker: 47,
	},
	{
		Id:         147,
		BuildingId: 2,
		Lv:         47,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3179000,
			},
		},
		MaxSave:   20880000,
		MaxWorker: 48,
	},
	{
		Id:         148,
		BuildingId: 2,
		Lv:         48,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3422000,
			},
		},
		MaxSave:   21848400,
		MaxWorker: 49,
	},
	{
		Id:         149,
		BuildingId: 2,
		Lv:         49,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3672000,
			},
		},
		MaxSave:   22838400,
		MaxWorker: 50,
	},
	{
		Id:         150,
		BuildingId: 2,
		Lv:         50,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3929000,
			},
		},
		MaxSave:   23850000,
		MaxWorker: 51,
	},
	{
		Id:         201,
		BuildingId: 3,
		Lv:         1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3600,
			},
		},
		MaxSave:   20000,
		MaxWorker: 2,
	},
	{
		Id:         202,
		BuildingId: 3,
		Lv:         2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 6480,
			},
		},
		MaxSave:   25000,
		MaxWorker: 3,
	},
	{
		Id:         203,
		BuildingId: 3,
		Lv:         3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 10000,
			},
		},
		MaxSave:   32400,
		MaxWorker: 4,
	},
	{
		Id:         204,
		BuildingId: 3,
		Lv:         4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 14000,
			},
		},
		MaxSave:   50400,
		MaxWorker: 5,
	},
	{
		Id:         205,
		BuildingId: 3,
		Lv:         5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 19000,
			},
		},
		MaxSave:   86400,
		MaxWorker: 6,
	},
	{
		Id:         206,
		BuildingId: 3,
		Lv:         6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 25000,
			},
		},
		MaxSave:   129600,
		MaxWorker: 7,
	},
	{
		Id:         207,
		BuildingId: 3,
		Lv:         7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 31000,
			},
		},
		MaxSave:   180000,
		MaxWorker: 8,
	},
	{
		Id:         208,
		BuildingId: 3,
		Lv:         8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 38000,
			},
		},
		MaxSave:   237600,
		MaxWorker: 9,
	},
	{
		Id:         209,
		BuildingId: 3,
		Lv:         9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 46000,
			},
		},
		MaxSave:   302400,
		MaxWorker: 10,
	},
	{
		Id:         210,
		BuildingId: 3,
		Lv:         10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 55000,
			},
		},
		MaxSave:   421200,
		MaxWorker: 11,
	},
	{
		Id:         211,
		BuildingId: 3,
		Lv:         11,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 69000,
			},
		},
		MaxSave:   554400,
		MaxWorker: 12,
	},
	{
		Id:         212,
		BuildingId: 3,
		Lv:         12,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 84000,
			},
		},
		MaxSave:   702000,
		MaxWorker: 13,
	},
	{
		Id:         213,
		BuildingId: 3,
		Lv:         13,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 100000,
			},
		},
		MaxSave:   864000,
		MaxWorker: 14,
	},
	{
		Id:         214,
		BuildingId: 3,
		Lv:         14,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 110000,
			},
		},
		MaxSave:   1040400,
		MaxWorker: 15,
	},
	{
		Id:         215,
		BuildingId: 3,
		Lv:         15,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 138000,
			},
		},
		MaxSave:   1296000,
		MaxWorker: 16,
	},
	{
		Id:         216,
		BuildingId: 3,
		Lv:         16,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 159000,
			},
		},
		MaxSave:   1573200,
		MaxWorker: 17,
	},
	{
		Id:         217,
		BuildingId: 3,
		Lv:         17,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 181000,
			},
		},
		MaxSave:   1872000,
		MaxWorker: 18,
	},
	{
		Id:         218,
		BuildingId: 3,
		Lv:         18,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 205000,
			},
		},
		MaxSave:   2192400,
		MaxWorker: 19,
	},
	{
		Id:         219,
		BuildingId: 3,
		Lv:         19,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 230000,
			},
		},
		MaxSave:   2534400,
		MaxWorker: 20,
	},
	{
		Id:         220,
		BuildingId: 3,
		Lv:         20,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 257000,
			},
		},
		MaxSave:   2898000,
		MaxWorker: 21,
	},
	{
		Id:         221,
		BuildingId: 3,
		Lv:         21,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 293000,
			},
		},
		MaxSave:   3283200,
		MaxWorker: 22,
	},
	{
		Id:         222,
		BuildingId: 3,
		Lv:         22,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 331000,
			},
		},
		MaxSave:   3690000,
		MaxWorker: 23,
	},
	{
		Id:         223,
		BuildingId: 3,
		Lv:         23,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 371000,
			},
		},
		MaxSave:   4118400,
		MaxWorker: 24,
	},
	{
		Id:         224,
		BuildingId: 3,
		Lv:         24,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 414000,
			},
		},
		MaxSave:   4568400,
		MaxWorker: 25,
	},
	{
		Id:         225,
		BuildingId: 3,
		Lv:         25,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 458000,
			},
		},
		MaxSave:   5040000,
		MaxWorker: 26,
	},
	{
		Id:         226,
		BuildingId: 3,
		Lv:         26,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 505000,
			},
		},
		MaxSave:   5533200,
		MaxWorker: 27,
	},
	{
		Id:         227,
		BuildingId: 3,
		Lv:         27,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 554000,
			},
		},
		MaxSave:   6048000,
		MaxWorker: 28,
	},
	{
		Id:         228,
		BuildingId: 3,
		Lv:         28,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 605000,
			},
		},
		MaxSave:   6584400,
		MaxWorker: 29,
	},
	{
		Id:         229,
		BuildingId: 3,
		Lv:         29,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 658000,
			},
		},
		MaxSave:   7142400,
		MaxWorker: 30,
	},
	{
		Id:         230,
		BuildingId: 3,
		Lv:         30,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 714000,
			},
		},
		MaxSave:   7722000,
		MaxWorker: 31,
	},
	{
		Id:         231,
		BuildingId: 3,
		Lv:         31,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 794000,
			},
		},
		MaxSave:   8323200,
		MaxWorker: 32,
	},
	{
		Id:         232,
		BuildingId: 3,
		Lv:         32,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 879000,
			},
		},
		MaxSave:   8946000,
		MaxWorker: 33,
	},
	{
		Id:         233,
		BuildingId: 3,
		Lv:         33,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 966000,
			},
		},
		MaxSave:   9590400,
		MaxWorker: 34,
	},
	{
		Id:         234,
		BuildingId: 3,
		Lv:         34,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1058000,
			},
		},
		MaxSave:   10256400,
		MaxWorker: 35,
	},
	{
		Id:         235,
		BuildingId: 3,
		Lv:         35,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1153000,
			},
		},
		MaxSave:   10944000,
		MaxWorker: 36,
	},
	{
		Id:         236,
		BuildingId: 3,
		Lv:         36,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1252000,
			},
		},
		MaxSave:   11653200,
		MaxWorker: 37,
	},
	{
		Id:         237,
		BuildingId: 3,
		Lv:         37,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1354000,
			},
		},
		MaxSave:   12384000,
		MaxWorker: 38,
	},
	{
		Id:         238,
		BuildingId: 3,
		Lv:         38,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1460000,
			},
		},
		MaxSave:   13136400,
		MaxWorker: 39,
	},
	{
		Id:         239,
		BuildingId: 3,
		Lv:         39,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1569000,
			},
		},
		MaxSave:   13910400,
		MaxWorker: 40,
	},
	{
		Id:         240,
		BuildingId: 3,
		Lv:         40,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1682000,
			},
		},
		MaxSave:   14706000,
		MaxWorker: 41,
	},
	{
		Id:         241,
		BuildingId: 3,
		Lv:         41,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1874000,
			},
		},
		MaxSave:   15523200,
		MaxWorker: 42,
	},
	{
		Id:         242,
		BuildingId: 3,
		Lv:         42,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2074000,
			},
		},
		MaxSave:   16362000,
		MaxWorker: 43,
	},
	{
		Id:         243,
		BuildingId: 3,
		Lv:         43,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2280000,
			},
		},
		MaxSave:   17222400,
		MaxWorker: 44,
	},
	{
		Id:         244,
		BuildingId: 3,
		Lv:         44,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2494000,
			},
		},
		MaxSave:   18104400,
		MaxWorker: 45,
	},
	{
		Id:         245,
		BuildingId: 3,
		Lv:         45,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2715000,
			},
		},
		MaxSave:   19008000,
		MaxWorker: 46,
	},
	{
		Id:         246,
		BuildingId: 3,
		Lv:         46,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2944000,
			},
		},
		MaxSave:   19933200,
		MaxWorker: 47,
	},
	{
		Id:         247,
		BuildingId: 3,
		Lv:         47,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3179000,
			},
		},
		MaxSave:   20880000,
		MaxWorker: 48,
	},
	{
		Id:         248,
		BuildingId: 3,
		Lv:         48,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3422000,
			},
		},
		MaxSave:   21848400,
		MaxWorker: 49,
	},
	{
		Id:         249,
		BuildingId: 3,
		Lv:         49,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3672000,
			},
		},
		MaxSave:   22838400,
		MaxWorker: 50,
	},
	{
		Id:         250,
		BuildingId: 3,
		Lv:         50,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3929000,
			},
		},
		MaxSave:   23850000,
		MaxWorker: 51,
	},
	{
		Id:         301,
		BuildingId: 4,
		Lv:         1,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3600,
			},
		},
		MaxSave:   20000,
		MaxWorker: 2,
	},
	{
		Id:         302,
		BuildingId: 4,
		Lv:         2,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 6480,
			},
		},
		MaxSave:   25000,
		MaxWorker: 3,
	},
	{
		Id:         303,
		BuildingId: 4,
		Lv:         3,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 10000,
			},
		},
		MaxSave:   32400,
		MaxWorker: 4,
	},
	{
		Id:         304,
		BuildingId: 4,
		Lv:         4,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 14000,
			},
		},
		MaxSave:   50400,
		MaxWorker: 5,
	},
	{
		Id:         305,
		BuildingId: 4,
		Lv:         5,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 19000,
			},
		},
		MaxSave:   86400,
		MaxWorker: 6,
	},
	{
		Id:         306,
		BuildingId: 4,
		Lv:         6,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 25000,
			},
		},
		MaxSave:   129600,
		MaxWorker: 7,
	},
	{
		Id:         307,
		BuildingId: 4,
		Lv:         7,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 31000,
			},
		},
		MaxSave:   180000,
		MaxWorker: 8,
	},
	{
		Id:         308,
		BuildingId: 4,
		Lv:         8,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 38000,
			},
		},
		MaxSave:   237600,
		MaxWorker: 9,
	},
	{
		Id:         309,
		BuildingId: 4,
		Lv:         9,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 46000,
			},
		},
		MaxSave:   302400,
		MaxWorker: 10,
	},
	{
		Id:         310,
		BuildingId: 4,
		Lv:         10,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 55000,
			},
		},
		MaxSave:   421200,
		MaxWorker: 11,
	},
	{
		Id:         311,
		BuildingId: 4,
		Lv:         11,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 69000,
			},
		},
		MaxSave:   554400,
		MaxWorker: 12,
	},
	{
		Id:         312,
		BuildingId: 4,
		Lv:         12,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 84000,
			},
		},
		MaxSave:   702000,
		MaxWorker: 13,
	},
	{
		Id:         313,
		BuildingId: 4,
		Lv:         13,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 100000,
			},
		},
		MaxSave:   864000,
		MaxWorker: 14,
	},
	{
		Id:         314,
		BuildingId: 4,
		Lv:         14,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 110000,
			},
		},
		MaxSave:   1040400,
		MaxWorker: 15,
	},
	{
		Id:         315,
		BuildingId: 4,
		Lv:         15,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 138000,
			},
		},
		MaxSave:   1296000,
		MaxWorker: 16,
	},
	{
		Id:         316,
		BuildingId: 4,
		Lv:         16,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 159000,
			},
		},
		MaxSave:   1573200,
		MaxWorker: 17,
	},
	{
		Id:         317,
		BuildingId: 4,
		Lv:         17,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 181000,
			},
		},
		MaxSave:   1872000,
		MaxWorker: 18,
	},
	{
		Id:         318,
		BuildingId: 4,
		Lv:         18,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 205000,
			},
		},
		MaxSave:   2192400,
		MaxWorker: 19,
	},
	{
		Id:         319,
		BuildingId: 4,
		Lv:         19,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 230000,
			},
		},
		MaxSave:   2534400,
		MaxWorker: 20,
	},
	{
		Id:         320,
		BuildingId: 4,
		Lv:         20,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 257000,
			},
		},
		MaxSave:   2898000,
		MaxWorker: 21,
	},
	{
		Id:         321,
		BuildingId: 4,
		Lv:         21,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 293000,
			},
		},
		MaxSave:   3283200,
		MaxWorker: 22,
	},
	{
		Id:         322,
		BuildingId: 4,
		Lv:         22,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 331000,
			},
		},
		MaxSave:   3690000,
		MaxWorker: 23,
	},
	{
		Id:         323,
		BuildingId: 4,
		Lv:         23,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 371000,
			},
		},
		MaxSave:   4118400,
		MaxWorker: 24,
	},
	{
		Id:         324,
		BuildingId: 4,
		Lv:         24,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 414000,
			},
		},
		MaxSave:   4568400,
		MaxWorker: 25,
	},
	{
		Id:         325,
		BuildingId: 4,
		Lv:         25,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 458000,
			},
		},
		MaxSave:   5040000,
		MaxWorker: 26,
	},
	{
		Id:         326,
		BuildingId: 4,
		Lv:         26,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 505000,
			},
		},
		MaxSave:   5533200,
		MaxWorker: 27,
	},
	{
		Id:         327,
		BuildingId: 4,
		Lv:         27,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 554000,
			},
		},
		MaxSave:   6048000,
		MaxWorker: 28,
	},
	{
		Id:         328,
		BuildingId: 4,
		Lv:         28,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 605000,
			},
		},
		MaxSave:   6584400,
		MaxWorker: 29,
	},
	{
		Id:         329,
		BuildingId: 4,
		Lv:         29,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 658000,
			},
		},
		MaxSave:   7142400,
		MaxWorker: 30,
	},
	{
		Id:         330,
		BuildingId: 4,
		Lv:         30,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 714000,
			},
		},
		MaxSave:   7722000,
		MaxWorker: 31,
	},
	{
		Id:         331,
		BuildingId: 4,
		Lv:         31,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 794000,
			},
		},
		MaxSave:   8323200,
		MaxWorker: 32,
	},
	{
		Id:         332,
		BuildingId: 4,
		Lv:         32,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 879000,
			},
		},
		MaxSave:   8946000,
		MaxWorker: 33,
	},
	{
		Id:         333,
		BuildingId: 4,
		Lv:         33,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 966000,
			},
		},
		MaxSave:   9590400,
		MaxWorker: 34,
	},
	{
		Id:         334,
		BuildingId: 4,
		Lv:         34,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1058000,
			},
		},
		MaxSave:   10256400,
		MaxWorker: 35,
	},
	{
		Id:         335,
		BuildingId: 4,
		Lv:         35,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1153000,
			},
		},
		MaxSave:   10944000,
		MaxWorker: 36,
	},
	{
		Id:         336,
		BuildingId: 4,
		Lv:         36,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1252000,
			},
		},
		MaxSave:   11653200,
		MaxWorker: 37,
	},
	{
		Id:         337,
		BuildingId: 4,
		Lv:         37,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1354000,
			},
		},
		MaxSave:   12384000,
		MaxWorker: 38,
	},
	{
		Id:         338,
		BuildingId: 4,
		Lv:         38,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1460000,
			},
		},
		MaxSave:   13136400,
		MaxWorker: 39,
	},
	{
		Id:         339,
		BuildingId: 4,
		Lv:         39,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1569000,
			},
		},
		MaxSave:   13910400,
		MaxWorker: 40,
	},
	{
		Id:         340,
		BuildingId: 4,
		Lv:         40,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1682000,
			},
		},
		MaxSave:   14706000,
		MaxWorker: 41,
	},
	{
		Id:         341,
		BuildingId: 4,
		Lv:         41,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 1874000,
			},
		},
		MaxSave:   15523200,
		MaxWorker: 42,
	},
	{
		Id:         342,
		BuildingId: 4,
		Lv:         42,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2074000,
			},
		},
		MaxSave:   16362000,
		MaxWorker: 43,
	},
	{
		Id:         343,
		BuildingId: 4,
		Lv:         43,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2280000,
			},
		},
		MaxSave:   17222400,
		MaxWorker: 44,
	},
	{
		Id:         344,
		BuildingId: 4,
		Lv:         44,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2494000,
			},
		},
		MaxSave:   18104400,
		MaxWorker: 45,
	},
	{
		Id:         345,
		BuildingId: 4,
		Lv:         45,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2715000,
			},
		},
		MaxSave:   19008000,
		MaxWorker: 46,
	},
	{
		Id:         346,
		BuildingId: 4,
		Lv:         46,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 2944000,
			},
		},
		MaxSave:   19933200,
		MaxWorker: 47,
	},
	{
		Id:         347,
		BuildingId: 4,
		Lv:         47,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3179000,
			},
		},
		MaxSave:   20880000,
		MaxWorker: 48,
	},
	{
		Id:         348,
		BuildingId: 4,
		Lv:         48,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3422000,
			},
		},
		MaxSave:   21848400,
		MaxWorker: 49,
	},
	{
		Id:         349,
		BuildingId: 4,
		Lv:         49,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3672000,
			},
		},
		MaxSave:   22838400,
		MaxWorker: 50,
	},
	{
		Id:         350,
		BuildingId: 4,
		Lv:         50,
		LevelUpConsume: []*Consume{
			{
				Type:  1,
				Id:    40,
				Count: 3929000,
			},
		},
		MaxSave:   23850000,
		MaxWorker: 51,
	},
}

var buildingResourceLvKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	201, 202, 203, 204, 205, 206, 207, 208, 209, 210,
	211, 212, 213, 214, 215, 216, 217, 218, 219, 220,
	221, 222, 223, 224, 225, 226, 227, 228, 229, 230,
	231, 232, 233, 234, 235, 236, 237, 238, 239, 240,
	241, 242, 243, 244, 245, 246, 247, 248, 249, 250,
	301, 302, 303, 304, 305, 306, 307, 308, 309, 310,
	311, 312, 313, 314, 315, 316, 317, 318, 319, 320,
	321, 322, 323, 324, 325, 326, 327, 328, 329, 330,
	331, 332, 333, 334, 335, 336, 337, 338, 339, 340,
	341, 342, 343, 344, 345, 346, 347, 348, 349, 350,
}

func init() {
	BuildingResourceLvData.data = make(map[int32]*BuildingResourceLvCfg)
	for i := 0; i < len(buildingResourceLvKeys); i++ {
		BuildingResourceLvData.data[buildingResourceLvKeys[i]] = buildingResourceLvValues[i]
	}
}

func (table *BuildingResourceLvTable) GetByBuildingId(BuildingId int32) (res []*BuildingResourceLvCfg) {
	for _, i := range table.buildingResourceLvBuildingIdIndexMap[BuildingId] {
		res = append(res, table.Get(i))
	}
	return
}

func (table *BuildingResourceLvTable) GetByBuildingIdAndLv(BuildingId int32, Lv int32) (res []*BuildingResourceLvCfg) {
	for _, i := range table.buildingResourceLvBuildingIdAndLvIndexMap[buildingResourceLvBuildingIdAndLvIndex{BuildingId, Lv}] {
		res = append(res, table.Get(i))
	}
	return
}

type buildingResourceLvBuildingIdAndLvIndex struct {
	buildingId int32
	lv         int32
}
