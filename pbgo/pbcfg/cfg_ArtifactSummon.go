package config

type ArtifactSummonCfg struct {
	Id          int32
	Price1      int32
	Price2      int32
	Probability [][]int32
}

type ArtifactSummonTable struct {
	data map[int32]*ArtifactSummonCfg
}

var ArtifactSummonData = &ArtifactSummonTable{
	data: map[int32]*ArtifactSummonCfg{},
}

func (table *ArtifactSummonTable) Get(id int32) *ArtifactSummonCfg {
	return table.data[id]
}

func (table *ArtifactSummonTable) GetAll() []int32 {
	return artifactSummonKeys
}

var artifactSummonValues = []*ArtifactSummonCfg{
	{
		Id:     1,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     2,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     3,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     4,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     5,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     6,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     7,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     8,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     9,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
	{
		Id:     10,
		Price2: 1,
		Probability: [][]int32{
			{
				0,
				200,
			},
			{
				0,
				200,
			},
			{
				0,
				200,
			},
		},
	},
}

var artifactSummonKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func init() {
	ArtifactSummonData.data = make(map[int32]*ArtifactSummonCfg)
	for i := 0; i < len(artifactSummonKeys); i++ {
		ArtifactSummonData.data[artifactSummonKeys[i]] = artifactSummonValues[i]
	}
}
