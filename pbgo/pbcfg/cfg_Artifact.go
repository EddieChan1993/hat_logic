package config

type ArtifactCfg struct {
	Id      int32
	Attrib  []int32
	Target  int32
	SkillId []int32
	Picture string
	Icon    string
}

type ArtifactTable struct {
	data map[int32]*ArtifactCfg
}

var ArtifactData = &ArtifactTable{
	data: map[int32]*ArtifactCfg{},
}

func (table *ArtifactTable) Get(id int32) *ArtifactCfg {
	return table.data[id]
}

func (table *ArtifactTable) GetAll() []int32 {
	return artifactKeys
}

var artifactValues = []*ArtifactCfg{
	{
		Id: 1,
		Attrib: []int32{
			22,
			83,
		},
		Target: 3,
		SkillId: []int32{
			901001,
		},
	},
	{
		Id: 2,
		Attrib: []int32{
			22,
			153,
		},
		Target: 3,
		SkillId: []int32{
			901002,
		},
	},
	{
		Id: 3,
		Attrib: []int32{
			32,
			93,
		},
		Target: 3,
		SkillId: []int32{
			901003,
		},
	},
	{
		Id: 4,
		Attrib: []int32{
			12,
			42,
		},
		Target: 3,
		SkillId: []int32{
			901004,
		},
	},
	{
		Id: 5,
		Attrib: []int32{
			12,
			52,
		},
		Target: 3,
		SkillId: []int32{
			901005,
		},
	},
	{
		Id: 6,
		Attrib: []int32{
			12,
			163,
		},
		Target: 3,
		SkillId: []int32{
			901006,
		},
	},
	{
		Id: 7,
		Attrib: []int32{
			32,
			63,
		},
		Target: 3,
		SkillId: []int32{
			901007,
		},
	},
	{
		Id: 8,
		Attrib: []int32{
			22,
			103,
		},
		Target: 3,
		SkillId: []int32{
			901008,
		},
	},
	{
		Id: 9,
		Attrib: []int32{
			32,
			133,
		},
		Target: 3,
		SkillId: []int32{
			901009,
		},
	},
	{
		Id: 10,
		Attrib: []int32{
			12,
			143,
		},
		Target: 3,
		SkillId: []int32{
			901010,
		},
	},
}

var artifactKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func init() {
	ArtifactData.data = make(map[int32]*ArtifactCfg)
	for i := 0; i < len(artifactKeys); i++ {
		ArtifactData.data[artifactKeys[i]] = artifactValues[i]
	}
}
