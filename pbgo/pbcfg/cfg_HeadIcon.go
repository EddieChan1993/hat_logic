package config

type HeadIconCfg struct {
	Id       int32
	HeroId   int32
	HeadIcon string
}

type HeadIconTable struct {
	data map[int32]*HeadIconCfg
}

var HeadIconData = &HeadIconTable{
	data: map[int32]*HeadIconCfg{},
}

func (table *HeadIconTable) Get(id int32) *HeadIconCfg {
	return table.data[id]
}

func (table *HeadIconTable) GetAll() []int32 {
	return headIconKeys
}

var headIconValues = []*HeadIconCfg{
	{
		Id:     1,
		HeroId: 0,
	},
	{
		Id:     2,
		HeroId: 10021,
	},
	{
		Id:     3,
		HeroId: 10121,
	},
	{
		Id:     4,
		HeroId: 10222,
	},
	{
		Id:     5,
		HeroId: 10321,
	},
	{
		Id:     6,
		HeroId: 10421,
	},
	{
		Id:     7,
		HeroId: 10521,
	},
	{
		Id:     8,
		HeroId: 10622,
	},
	{
		Id:     9,
		HeroId: 10721,
	},
	{
		Id:     10,
		HeroId: 10822,
	},
	{
		Id:     11,
		HeroId: 10921,
	},
	{
		Id:     12,
		HeroId: 11021,
	},
	{
		Id:     13,
		HeroId: 11122,
	},
	{
		Id:     14,
		HeroId: 11222,
	},
	{
		Id:     15,
		HeroId: 11322,
	},
	{
		Id:     16,
		HeroId: 11421,
	},
	{
		Id:     17,
		HeroId: 11511,
	},
	{
		Id:     18,
		HeroId: 11612,
	},
	{
		Id:     19,
		HeroId: 11712,
	},
	{
		Id:     20,
		HeroId: 11812,
	},
	{
		Id:     21,
		HeroId: 11911,
	},
	{
		Id:     22,
		HeroId: 12011,
	},
	{
		Id:     23,
		HeroId: 12112,
	},
	{
		Id:     24,
		HeroId: 12212,
	},
	{
		Id:     25,
		HeroId: 12312,
	},
	{
		Id:     26,
		HeroId: 12412,
	},
	{
		Id:     27,
		HeroId: 12511,
	},
	{
		Id:     28,
		HeroId: 12611,
	},
	{
		Id:     29,
		HeroId: 12711,
	},
	{
		Id:     30,
		HeroId: 12811,
	},
	{
		Id:     31,
		HeroId: 12911,
	},
}

var headIconKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31,
}

func init() {
	HeadIconData.data = make(map[int32]*HeadIconCfg)
	for i := 0; i < len(headIconKeys); i++ {
		HeadIconData.data[headIconKeys[i]] = headIconValues[i]
	}
}
