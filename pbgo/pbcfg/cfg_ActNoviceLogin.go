package config

type ActNoviceLoginCfg struct {
	Id     int32
	Reward []*Reward
}

type ActNoviceLoginTable struct {
	data map[int32]*ActNoviceLoginCfg
}

var ActNoviceLoginData = &ActNoviceLoginTable{
	data: map[int32]*ActNoviceLoginCfg{},
}

func (table *ActNoviceLoginTable) Get(id int32) *ActNoviceLoginCfg {
	return table.data[id]
}

func (table *ActNoviceLoginTable) GetAll() []int32 {
	return actNoviceLoginKeys
}

var actNoviceLoginValues = []*ActNoviceLoginCfg{
	{
		Id: 1,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    3,
				Count: 1000,
			},
		},
	},
	{
		Id: 2,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
	},
	{
		Id: 3,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    124001,
				Count: 1,
			},
		},
	},
	{
		Id: 4,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 10,
			},
		},
	},
	{
		Id: 5,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 10,
			},
		},
	},
	{
		Id: 6,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 10,
			},
		},
	},
	{
		Id: 7,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    225001,
				Count: 1,
			},
		},
	},
}

var actNoviceLoginKeys = []int32{
	1, 2, 3, 4, 5, 6, 7,
}

func init() {
	ActNoviceLoginData.data = make(map[int32]*ActNoviceLoginCfg)
	for i := 0; i < len(actNoviceLoginKeys); i++ {
		ActNoviceLoginData.data[actNoviceLoginKeys[i]] = actNoviceLoginValues[i]
	}
}
