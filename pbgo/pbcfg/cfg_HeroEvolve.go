package config

type HeroEvolveCfg struct {
	Id              int32
	Fragment        []*Consume
	NaterialScience []*Consume
}

type HeroEvolveTable struct {
	data map[int32]*HeroEvolveCfg
}

var HeroEvolveData = &HeroEvolveTable{
	data: map[int32]*HeroEvolveCfg{},
}

func (table *HeroEvolveTable) Get(id int32) *HeroEvolveCfg {
	return table.data[id]
}

func (table *HeroEvolveTable) GetAll() []int32 {
	return heroEvolveKeys
}

var heroEvolveValues = []*HeroEvolveCfg{
	{
		Id: 10021,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    5,
				Count: 50,
			},
			{
				Type:  1,
				Id:    5,
				Count: 50,
			},
			{
				Type:  1,
				Id:    5,
				Count: 50,
			},
			{
				Type:  1,
				Id:    5,
				Count: 100,
			},
			{
				Type:  1,
				Id:    5,
				Count: 100,
			},
			{
				Type:  1,
				Id:    5,
				Count: 150,
			},
			{
				Type:  1,
				Id:    5,
				Count: 150,
			},
			{
				Type:  1,
				Id:    5,
				Count: 200,
			},
			{
				Type:  1,
				Id:    5,
				Count: 200,
			},
			{
				Type:  1,
				Id:    5,
				Count: 200,
			},
			{
				Type:  1,
				Id:    5,
				Count: 300,
			},
			{
				Type:  1,
				Id:    5,
				Count: 300,
			},
			{
				Type:  1,
				Id:    5,
				Count: 300,
			},
			{
				Type:  1,
				Id:    5,
				Count: 300,
			},
			{
				Type:  1,
				Id:    5,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10121,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    6,
				Count: 50,
			},
			{
				Type:  1,
				Id:    6,
				Count: 50,
			},
			{
				Type:  1,
				Id:    6,
				Count: 50,
			},
			{
				Type:  1,
				Id:    6,
				Count: 100,
			},
			{
				Type:  1,
				Id:    6,
				Count: 100,
			},
			{
				Type:  1,
				Id:    6,
				Count: 150,
			},
			{
				Type:  1,
				Id:    6,
				Count: 150,
			},
			{
				Type:  1,
				Id:    6,
				Count: 200,
			},
			{
				Type:  1,
				Id:    6,
				Count: 200,
			},
			{
				Type:  1,
				Id:    6,
				Count: 200,
			},
			{
				Type:  1,
				Id:    6,
				Count: 300,
			},
			{
				Type:  1,
				Id:    6,
				Count: 300,
			},
			{
				Type:  1,
				Id:    6,
				Count: 300,
			},
			{
				Type:  1,
				Id:    6,
				Count: 300,
			},
			{
				Type:  1,
				Id:    6,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10222,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    7,
				Count: 50,
			},
			{
				Type:  1,
				Id:    7,
				Count: 50,
			},
			{
				Type:  1,
				Id:    7,
				Count: 50,
			},
			{
				Type:  1,
				Id:    7,
				Count: 100,
			},
			{
				Type:  1,
				Id:    7,
				Count: 100,
			},
			{
				Type:  1,
				Id:    7,
				Count: 150,
			},
			{
				Type:  1,
				Id:    7,
				Count: 150,
			},
			{
				Type:  1,
				Id:    7,
				Count: 200,
			},
			{
				Type:  1,
				Id:    7,
				Count: 200,
			},
			{
				Type:  1,
				Id:    7,
				Count: 200,
			},
			{
				Type:  1,
				Id:    7,
				Count: 300,
			},
			{
				Type:  1,
				Id:    7,
				Count: 300,
			},
			{
				Type:  1,
				Id:    7,
				Count: 300,
			},
			{
				Type:  1,
				Id:    7,
				Count: 300,
			},
			{
				Type:  1,
				Id:    7,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10321,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    8,
				Count: 50,
			},
			{
				Type:  1,
				Id:    8,
				Count: 50,
			},
			{
				Type:  1,
				Id:    8,
				Count: 50,
			},
			{
				Type:  1,
				Id:    8,
				Count: 100,
			},
			{
				Type:  1,
				Id:    8,
				Count: 100,
			},
			{
				Type:  1,
				Id:    8,
				Count: 150,
			},
			{
				Type:  1,
				Id:    8,
				Count: 150,
			},
			{
				Type:  1,
				Id:    8,
				Count: 200,
			},
			{
				Type:  1,
				Id:    8,
				Count: 200,
			},
			{
				Type:  1,
				Id:    8,
				Count: 200,
			},
			{
				Type:  1,
				Id:    8,
				Count: 300,
			},
			{
				Type:  1,
				Id:    8,
				Count: 300,
			},
			{
				Type:  1,
				Id:    8,
				Count: 300,
			},
			{
				Type:  1,
				Id:    8,
				Count: 300,
			},
			{
				Type:  1,
				Id:    8,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10421,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    9,
				Count: 50,
			},
			{
				Type:  1,
				Id:    9,
				Count: 50,
			},
			{
				Type:  1,
				Id:    9,
				Count: 50,
			},
			{
				Type:  1,
				Id:    9,
				Count: 100,
			},
			{
				Type:  1,
				Id:    9,
				Count: 100,
			},
			{
				Type:  1,
				Id:    9,
				Count: 150,
			},
			{
				Type:  1,
				Id:    9,
				Count: 150,
			},
			{
				Type:  1,
				Id:    9,
				Count: 200,
			},
			{
				Type:  1,
				Id:    9,
				Count: 200,
			},
			{
				Type:  1,
				Id:    9,
				Count: 200,
			},
			{
				Type:  1,
				Id:    9,
				Count: 300,
			},
			{
				Type:  1,
				Id:    9,
				Count: 300,
			},
			{
				Type:  1,
				Id:    9,
				Count: 300,
			},
			{
				Type:  1,
				Id:    9,
				Count: 300,
			},
			{
				Type:  1,
				Id:    9,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10521,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    10,
				Count: 50,
			},
			{
				Type:  1,
				Id:    10,
				Count: 50,
			},
			{
				Type:  1,
				Id:    10,
				Count: 50,
			},
			{
				Type:  1,
				Id:    10,
				Count: 100,
			},
			{
				Type:  1,
				Id:    10,
				Count: 100,
			},
			{
				Type:  1,
				Id:    10,
				Count: 150,
			},
			{
				Type:  1,
				Id:    10,
				Count: 150,
			},
			{
				Type:  1,
				Id:    10,
				Count: 200,
			},
			{
				Type:  1,
				Id:    10,
				Count: 200,
			},
			{
				Type:  1,
				Id:    10,
				Count: 200,
			},
			{
				Type:  1,
				Id:    10,
				Count: 300,
			},
			{
				Type:  1,
				Id:    10,
				Count: 300,
			},
			{
				Type:  1,
				Id:    10,
				Count: 300,
			},
			{
				Type:  1,
				Id:    10,
				Count: 300,
			},
			{
				Type:  1,
				Id:    10,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10622,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    11,
				Count: 50,
			},
			{
				Type:  1,
				Id:    11,
				Count: 50,
			},
			{
				Type:  1,
				Id:    11,
				Count: 50,
			},
			{
				Type:  1,
				Id:    11,
				Count: 100,
			},
			{
				Type:  1,
				Id:    11,
				Count: 100,
			},
			{
				Type:  1,
				Id:    11,
				Count: 150,
			},
			{
				Type:  1,
				Id:    11,
				Count: 150,
			},
			{
				Type:  1,
				Id:    11,
				Count: 200,
			},
			{
				Type:  1,
				Id:    11,
				Count: 200,
			},
			{
				Type:  1,
				Id:    11,
				Count: 200,
			},
			{
				Type:  1,
				Id:    11,
				Count: 300,
			},
			{
				Type:  1,
				Id:    11,
				Count: 300,
			},
			{
				Type:  1,
				Id:    11,
				Count: 300,
			},
			{
				Type:  1,
				Id:    11,
				Count: 300,
			},
			{
				Type:  1,
				Id:    11,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10721,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    12,
				Count: 50,
			},
			{
				Type:  1,
				Id:    12,
				Count: 50,
			},
			{
				Type:  1,
				Id:    12,
				Count: 50,
			},
			{
				Type:  1,
				Id:    12,
				Count: 100,
			},
			{
				Type:  1,
				Id:    12,
				Count: 100,
			},
			{
				Type:  1,
				Id:    12,
				Count: 150,
			},
			{
				Type:  1,
				Id:    12,
				Count: 150,
			},
			{
				Type:  1,
				Id:    12,
				Count: 200,
			},
			{
				Type:  1,
				Id:    12,
				Count: 200,
			},
			{
				Type:  1,
				Id:    12,
				Count: 200,
			},
			{
				Type:  1,
				Id:    12,
				Count: 300,
			},
			{
				Type:  1,
				Id:    12,
				Count: 300,
			},
			{
				Type:  1,
				Id:    12,
				Count: 300,
			},
			{
				Type:  1,
				Id:    12,
				Count: 300,
			},
			{
				Type:  1,
				Id:    12,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10822,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    13,
				Count: 50,
			},
			{
				Type:  1,
				Id:    13,
				Count: 50,
			},
			{
				Type:  1,
				Id:    13,
				Count: 50,
			},
			{
				Type:  1,
				Id:    13,
				Count: 100,
			},
			{
				Type:  1,
				Id:    13,
				Count: 100,
			},
			{
				Type:  1,
				Id:    13,
				Count: 150,
			},
			{
				Type:  1,
				Id:    13,
				Count: 150,
			},
			{
				Type:  1,
				Id:    13,
				Count: 200,
			},
			{
				Type:  1,
				Id:    13,
				Count: 200,
			},
			{
				Type:  1,
				Id:    13,
				Count: 200,
			},
			{
				Type:  1,
				Id:    13,
				Count: 300,
			},
			{
				Type:  1,
				Id:    13,
				Count: 300,
			},
			{
				Type:  1,
				Id:    13,
				Count: 300,
			},
			{
				Type:  1,
				Id:    13,
				Count: 300,
			},
			{
				Type:  1,
				Id:    13,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 10921,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    14,
				Count: 50,
			},
			{
				Type:  1,
				Id:    14,
				Count: 50,
			},
			{
				Type:  1,
				Id:    14,
				Count: 50,
			},
			{
				Type:  1,
				Id:    14,
				Count: 100,
			},
			{
				Type:  1,
				Id:    14,
				Count: 100,
			},
			{
				Type:  1,
				Id:    14,
				Count: 150,
			},
			{
				Type:  1,
				Id:    14,
				Count: 150,
			},
			{
				Type:  1,
				Id:    14,
				Count: 200,
			},
			{
				Type:  1,
				Id:    14,
				Count: 200,
			},
			{
				Type:  1,
				Id:    14,
				Count: 200,
			},
			{
				Type:  1,
				Id:    14,
				Count: 300,
			},
			{
				Type:  1,
				Id:    14,
				Count: 300,
			},
			{
				Type:  1,
				Id:    14,
				Count: 300,
			},
			{
				Type:  1,
				Id:    14,
				Count: 300,
			},
			{
				Type:  1,
				Id:    14,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11021,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    15,
				Count: 50,
			},
			{
				Type:  1,
				Id:    15,
				Count: 50,
			},
			{
				Type:  1,
				Id:    15,
				Count: 50,
			},
			{
				Type:  1,
				Id:    15,
				Count: 100,
			},
			{
				Type:  1,
				Id:    15,
				Count: 100,
			},
			{
				Type:  1,
				Id:    15,
				Count: 150,
			},
			{
				Type:  1,
				Id:    15,
				Count: 150,
			},
			{
				Type:  1,
				Id:    15,
				Count: 200,
			},
			{
				Type:  1,
				Id:    15,
				Count: 200,
			},
			{
				Type:  1,
				Id:    15,
				Count: 200,
			},
			{
				Type:  1,
				Id:    15,
				Count: 300,
			},
			{
				Type:  1,
				Id:    15,
				Count: 300,
			},
			{
				Type:  1,
				Id:    15,
				Count: 300,
			},
			{
				Type:  1,
				Id:    15,
				Count: 300,
			},
			{
				Type:  1,
				Id:    15,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11122,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    16,
				Count: 50,
			},
			{
				Type:  1,
				Id:    16,
				Count: 50,
			},
			{
				Type:  1,
				Id:    16,
				Count: 50,
			},
			{
				Type:  1,
				Id:    16,
				Count: 100,
			},
			{
				Type:  1,
				Id:    16,
				Count: 100,
			},
			{
				Type:  1,
				Id:    16,
				Count: 150,
			},
			{
				Type:  1,
				Id:    16,
				Count: 150,
			},
			{
				Type:  1,
				Id:    16,
				Count: 200,
			},
			{
				Type:  1,
				Id:    16,
				Count: 200,
			},
			{
				Type:  1,
				Id:    16,
				Count: 200,
			},
			{
				Type:  1,
				Id:    16,
				Count: 300,
			},
			{
				Type:  1,
				Id:    16,
				Count: 300,
			},
			{
				Type:  1,
				Id:    16,
				Count: 300,
			},
			{
				Type:  1,
				Id:    16,
				Count: 300,
			},
			{
				Type:  1,
				Id:    16,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11222,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    17,
				Count: 50,
			},
			{
				Type:  1,
				Id:    17,
				Count: 50,
			},
			{
				Type:  1,
				Id:    17,
				Count: 50,
			},
			{
				Type:  1,
				Id:    17,
				Count: 100,
			},
			{
				Type:  1,
				Id:    17,
				Count: 100,
			},
			{
				Type:  1,
				Id:    17,
				Count: 150,
			},
			{
				Type:  1,
				Id:    17,
				Count: 150,
			},
			{
				Type:  1,
				Id:    17,
				Count: 200,
			},
			{
				Type:  1,
				Id:    17,
				Count: 200,
			},
			{
				Type:  1,
				Id:    17,
				Count: 200,
			},
			{
				Type:  1,
				Id:    17,
				Count: 300,
			},
			{
				Type:  1,
				Id:    17,
				Count: 300,
			},
			{
				Type:  1,
				Id:    17,
				Count: 300,
			},
			{
				Type:  1,
				Id:    17,
				Count: 300,
			},
			{
				Type:  1,
				Id:    17,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11322,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    18,
				Count: 50,
			},
			{
				Type:  1,
				Id:    18,
				Count: 50,
			},
			{
				Type:  1,
				Id:    18,
				Count: 50,
			},
			{
				Type:  1,
				Id:    18,
				Count: 100,
			},
			{
				Type:  1,
				Id:    18,
				Count: 100,
			},
			{
				Type:  1,
				Id:    18,
				Count: 150,
			},
			{
				Type:  1,
				Id:    18,
				Count: 150,
			},
			{
				Type:  1,
				Id:    18,
				Count: 200,
			},
			{
				Type:  1,
				Id:    18,
				Count: 200,
			},
			{
				Type:  1,
				Id:    18,
				Count: 200,
			},
			{
				Type:  1,
				Id:    18,
				Count: 300,
			},
			{
				Type:  1,
				Id:    18,
				Count: 300,
			},
			{
				Type:  1,
				Id:    18,
				Count: 300,
			},
			{
				Type:  1,
				Id:    18,
				Count: 300,
			},
			{
				Type:  1,
				Id:    18,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11421,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    19,
				Count: 50,
			},
			{
				Type:  1,
				Id:    19,
				Count: 50,
			},
			{
				Type:  1,
				Id:    19,
				Count: 50,
			},
			{
				Type:  1,
				Id:    19,
				Count: 100,
			},
			{
				Type:  1,
				Id:    19,
				Count: 100,
			},
			{
				Type:  1,
				Id:    19,
				Count: 150,
			},
			{
				Type:  1,
				Id:    19,
				Count: 150,
			},
			{
				Type:  1,
				Id:    19,
				Count: 200,
			},
			{
				Type:  1,
				Id:    19,
				Count: 200,
			},
			{
				Type:  1,
				Id:    19,
				Count: 200,
			},
			{
				Type:  1,
				Id:    19,
				Count: 300,
			},
			{
				Type:  1,
				Id:    19,
				Count: 300,
			},
			{
				Type:  1,
				Id:    19,
				Count: 300,
			},
			{
				Type:  1,
				Id:    19,
				Count: 300,
			},
			{
				Type:  1,
				Id:    19,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11511,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    20,
				Count: 50,
			},
			{
				Type:  1,
				Id:    20,
				Count: 50,
			},
			{
				Type:  1,
				Id:    20,
				Count: 50,
			},
			{
				Type:  1,
				Id:    20,
				Count: 100,
			},
			{
				Type:  1,
				Id:    20,
				Count: 100,
			},
			{
				Type:  1,
				Id:    20,
				Count: 150,
			},
			{
				Type:  1,
				Id:    20,
				Count: 150,
			},
			{
				Type:  1,
				Id:    20,
				Count: 200,
			},
			{
				Type:  1,
				Id:    20,
				Count: 200,
			},
			{
				Type:  1,
				Id:    20,
				Count: 200,
			},
			{
				Type:  1,
				Id:    20,
				Count: 300,
			},
			{
				Type:  1,
				Id:    20,
				Count: 300,
			},
			{
				Type:  1,
				Id:    20,
				Count: 300,
			},
			{
				Type:  1,
				Id:    20,
				Count: 300,
			},
			{
				Type:  1,
				Id:    20,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11612,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    21,
				Count: 50,
			},
			{
				Type:  1,
				Id:    21,
				Count: 50,
			},
			{
				Type:  1,
				Id:    21,
				Count: 50,
			},
			{
				Type:  1,
				Id:    21,
				Count: 100,
			},
			{
				Type:  1,
				Id:    21,
				Count: 100,
			},
			{
				Type:  1,
				Id:    21,
				Count: 150,
			},
			{
				Type:  1,
				Id:    21,
				Count: 150,
			},
			{
				Type:  1,
				Id:    21,
				Count: 200,
			},
			{
				Type:  1,
				Id:    21,
				Count: 200,
			},
			{
				Type:  1,
				Id:    21,
				Count: 200,
			},
			{
				Type:  1,
				Id:    21,
				Count: 300,
			},
			{
				Type:  1,
				Id:    21,
				Count: 300,
			},
			{
				Type:  1,
				Id:    21,
				Count: 300,
			},
			{
				Type:  1,
				Id:    21,
				Count: 300,
			},
			{
				Type:  1,
				Id:    21,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11712,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    22,
				Count: 50,
			},
			{
				Type:  1,
				Id:    22,
				Count: 50,
			},
			{
				Type:  1,
				Id:    22,
				Count: 50,
			},
			{
				Type:  1,
				Id:    22,
				Count: 100,
			},
			{
				Type:  1,
				Id:    22,
				Count: 100,
			},
			{
				Type:  1,
				Id:    22,
				Count: 150,
			},
			{
				Type:  1,
				Id:    22,
				Count: 150,
			},
			{
				Type:  1,
				Id:    22,
				Count: 200,
			},
			{
				Type:  1,
				Id:    22,
				Count: 200,
			},
			{
				Type:  1,
				Id:    22,
				Count: 200,
			},
			{
				Type:  1,
				Id:    22,
				Count: 300,
			},
			{
				Type:  1,
				Id:    22,
				Count: 300,
			},
			{
				Type:  1,
				Id:    22,
				Count: 300,
			},
			{
				Type:  1,
				Id:    22,
				Count: 300,
			},
			{
				Type:  1,
				Id:    22,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11812,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    23,
				Count: 50,
			},
			{
				Type:  1,
				Id:    23,
				Count: 50,
			},
			{
				Type:  1,
				Id:    23,
				Count: 50,
			},
			{
				Type:  1,
				Id:    23,
				Count: 100,
			},
			{
				Type:  1,
				Id:    23,
				Count: 100,
			},
			{
				Type:  1,
				Id:    23,
				Count: 150,
			},
			{
				Type:  1,
				Id:    23,
				Count: 150,
			},
			{
				Type:  1,
				Id:    23,
				Count: 200,
			},
			{
				Type:  1,
				Id:    23,
				Count: 200,
			},
			{
				Type:  1,
				Id:    23,
				Count: 200,
			},
			{
				Type:  1,
				Id:    23,
				Count: 300,
			},
			{
				Type:  1,
				Id:    23,
				Count: 300,
			},
			{
				Type:  1,
				Id:    23,
				Count: 300,
			},
			{
				Type:  1,
				Id:    23,
				Count: 300,
			},
			{
				Type:  1,
				Id:    23,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 11911,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    24,
				Count: 50,
			},
			{
				Type:  1,
				Id:    24,
				Count: 50,
			},
			{
				Type:  1,
				Id:    24,
				Count: 50,
			},
			{
				Type:  1,
				Id:    24,
				Count: 100,
			},
			{
				Type:  1,
				Id:    24,
				Count: 100,
			},
			{
				Type:  1,
				Id:    24,
				Count: 150,
			},
			{
				Type:  1,
				Id:    24,
				Count: 150,
			},
			{
				Type:  1,
				Id:    24,
				Count: 200,
			},
			{
				Type:  1,
				Id:    24,
				Count: 200,
			},
			{
				Type:  1,
				Id:    24,
				Count: 200,
			},
			{
				Type:  1,
				Id:    24,
				Count: 300,
			},
			{
				Type:  1,
				Id:    24,
				Count: 300,
			},
			{
				Type:  1,
				Id:    24,
				Count: 300,
			},
			{
				Type:  1,
				Id:    24,
				Count: 300,
			},
			{
				Type:  1,
				Id:    24,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12011,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    25,
				Count: 50,
			},
			{
				Type:  1,
				Id:    25,
				Count: 50,
			},
			{
				Type:  1,
				Id:    25,
				Count: 50,
			},
			{
				Type:  1,
				Id:    25,
				Count: 100,
			},
			{
				Type:  1,
				Id:    25,
				Count: 100,
			},
			{
				Type:  1,
				Id:    25,
				Count: 150,
			},
			{
				Type:  1,
				Id:    25,
				Count: 150,
			},
			{
				Type:  1,
				Id:    25,
				Count: 200,
			},
			{
				Type:  1,
				Id:    25,
				Count: 200,
			},
			{
				Type:  1,
				Id:    25,
				Count: 200,
			},
			{
				Type:  1,
				Id:    25,
				Count: 300,
			},
			{
				Type:  1,
				Id:    25,
				Count: 300,
			},
			{
				Type:  1,
				Id:    25,
				Count: 300,
			},
			{
				Type:  1,
				Id:    25,
				Count: 300,
			},
			{
				Type:  1,
				Id:    25,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12112,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    26,
				Count: 50,
			},
			{
				Type:  1,
				Id:    26,
				Count: 50,
			},
			{
				Type:  1,
				Id:    26,
				Count: 50,
			},
			{
				Type:  1,
				Id:    26,
				Count: 100,
			},
			{
				Type:  1,
				Id:    26,
				Count: 100,
			},
			{
				Type:  1,
				Id:    26,
				Count: 150,
			},
			{
				Type:  1,
				Id:    26,
				Count: 150,
			},
			{
				Type:  1,
				Id:    26,
				Count: 200,
			},
			{
				Type:  1,
				Id:    26,
				Count: 200,
			},
			{
				Type:  1,
				Id:    26,
				Count: 200,
			},
			{
				Type:  1,
				Id:    26,
				Count: 300,
			},
			{
				Type:  1,
				Id:    26,
				Count: 300,
			},
			{
				Type:  1,
				Id:    26,
				Count: 300,
			},
			{
				Type:  1,
				Id:    26,
				Count: 300,
			},
			{
				Type:  1,
				Id:    26,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12212,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    27,
				Count: 50,
			},
			{
				Type:  1,
				Id:    27,
				Count: 50,
			},
			{
				Type:  1,
				Id:    27,
				Count: 50,
			},
			{
				Type:  1,
				Id:    27,
				Count: 100,
			},
			{
				Type:  1,
				Id:    27,
				Count: 100,
			},
			{
				Type:  1,
				Id:    27,
				Count: 150,
			},
			{
				Type:  1,
				Id:    27,
				Count: 150,
			},
			{
				Type:  1,
				Id:    27,
				Count: 200,
			},
			{
				Type:  1,
				Id:    27,
				Count: 200,
			},
			{
				Type:  1,
				Id:    27,
				Count: 200,
			},
			{
				Type:  1,
				Id:    27,
				Count: 300,
			},
			{
				Type:  1,
				Id:    27,
				Count: 300,
			},
			{
				Type:  1,
				Id:    27,
				Count: 300,
			},
			{
				Type:  1,
				Id:    27,
				Count: 300,
			},
			{
				Type:  1,
				Id:    27,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12312,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    28,
				Count: 50,
			},
			{
				Type:  1,
				Id:    28,
				Count: 50,
			},
			{
				Type:  1,
				Id:    28,
				Count: 50,
			},
			{
				Type:  1,
				Id:    28,
				Count: 100,
			},
			{
				Type:  1,
				Id:    28,
				Count: 100,
			},
			{
				Type:  1,
				Id:    28,
				Count: 150,
			},
			{
				Type:  1,
				Id:    28,
				Count: 150,
			},
			{
				Type:  1,
				Id:    28,
				Count: 200,
			},
			{
				Type:  1,
				Id:    28,
				Count: 200,
			},
			{
				Type:  1,
				Id:    28,
				Count: 200,
			},
			{
				Type:  1,
				Id:    28,
				Count: 300,
			},
			{
				Type:  1,
				Id:    28,
				Count: 300,
			},
			{
				Type:  1,
				Id:    28,
				Count: 300,
			},
			{
				Type:  1,
				Id:    28,
				Count: 300,
			},
			{
				Type:  1,
				Id:    28,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12412,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    29,
				Count: 50,
			},
			{
				Type:  1,
				Id:    29,
				Count: 50,
			},
			{
				Type:  1,
				Id:    29,
				Count: 50,
			},
			{
				Type:  1,
				Id:    29,
				Count: 100,
			},
			{
				Type:  1,
				Id:    29,
				Count: 100,
			},
			{
				Type:  1,
				Id:    29,
				Count: 150,
			},
			{
				Type:  1,
				Id:    29,
				Count: 150,
			},
			{
				Type:  1,
				Id:    29,
				Count: 200,
			},
			{
				Type:  1,
				Id:    29,
				Count: 200,
			},
			{
				Type:  1,
				Id:    29,
				Count: 200,
			},
			{
				Type:  1,
				Id:    29,
				Count: 300,
			},
			{
				Type:  1,
				Id:    29,
				Count: 300,
			},
			{
				Type:  1,
				Id:    29,
				Count: 300,
			},
			{
				Type:  1,
				Id:    29,
				Count: 300,
			},
			{
				Type:  1,
				Id:    29,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12511,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    30,
				Count: 50,
			},
			{
				Type:  1,
				Id:    30,
				Count: 50,
			},
			{
				Type:  1,
				Id:    30,
				Count: 50,
			},
			{
				Type:  1,
				Id:    30,
				Count: 100,
			},
			{
				Type:  1,
				Id:    30,
				Count: 100,
			},
			{
				Type:  1,
				Id:    30,
				Count: 150,
			},
			{
				Type:  1,
				Id:    30,
				Count: 150,
			},
			{
				Type:  1,
				Id:    30,
				Count: 200,
			},
			{
				Type:  1,
				Id:    30,
				Count: 200,
			},
			{
				Type:  1,
				Id:    30,
				Count: 200,
			},
			{
				Type:  1,
				Id:    30,
				Count: 300,
			},
			{
				Type:  1,
				Id:    30,
				Count: 300,
			},
			{
				Type:  1,
				Id:    30,
				Count: 300,
			},
			{
				Type:  1,
				Id:    30,
				Count: 300,
			},
			{
				Type:  1,
				Id:    30,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12611,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    31,
				Count: 50,
			},
			{
				Type:  1,
				Id:    31,
				Count: 50,
			},
			{
				Type:  1,
				Id:    31,
				Count: 50,
			},
			{
				Type:  1,
				Id:    31,
				Count: 100,
			},
			{
				Type:  1,
				Id:    31,
				Count: 100,
			},
			{
				Type:  1,
				Id:    31,
				Count: 150,
			},
			{
				Type:  1,
				Id:    31,
				Count: 150,
			},
			{
				Type:  1,
				Id:    31,
				Count: 200,
			},
			{
				Type:  1,
				Id:    31,
				Count: 200,
			},
			{
				Type:  1,
				Id:    31,
				Count: 200,
			},
			{
				Type:  1,
				Id:    31,
				Count: 300,
			},
			{
				Type:  1,
				Id:    31,
				Count: 300,
			},
			{
				Type:  1,
				Id:    31,
				Count: 300,
			},
			{
				Type:  1,
				Id:    31,
				Count: 300,
			},
			{
				Type:  1,
				Id:    31,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12711,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    32,
				Count: 50,
			},
			{
				Type:  1,
				Id:    32,
				Count: 50,
			},
			{
				Type:  1,
				Id:    32,
				Count: 50,
			},
			{
				Type:  1,
				Id:    32,
				Count: 100,
			},
			{
				Type:  1,
				Id:    32,
				Count: 100,
			},
			{
				Type:  1,
				Id:    32,
				Count: 150,
			},
			{
				Type:  1,
				Id:    32,
				Count: 150,
			},
			{
				Type:  1,
				Id:    32,
				Count: 200,
			},
			{
				Type:  1,
				Id:    32,
				Count: 200,
			},
			{
				Type:  1,
				Id:    32,
				Count: 200,
			},
			{
				Type:  1,
				Id:    32,
				Count: 300,
			},
			{
				Type:  1,
				Id:    32,
				Count: 300,
			},
			{
				Type:  1,
				Id:    32,
				Count: 300,
			},
			{
				Type:  1,
				Id:    32,
				Count: 300,
			},
			{
				Type:  1,
				Id:    32,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12811,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    33,
				Count: 50,
			},
			{
				Type:  1,
				Id:    33,
				Count: 50,
			},
			{
				Type:  1,
				Id:    33,
				Count: 50,
			},
			{
				Type:  1,
				Id:    33,
				Count: 100,
			},
			{
				Type:  1,
				Id:    33,
				Count: 100,
			},
			{
				Type:  1,
				Id:    33,
				Count: 150,
			},
			{
				Type:  1,
				Id:    33,
				Count: 150,
			},
			{
				Type:  1,
				Id:    33,
				Count: 200,
			},
			{
				Type:  1,
				Id:    33,
				Count: 200,
			},
			{
				Type:  1,
				Id:    33,
				Count: 200,
			},
			{
				Type:  1,
				Id:    33,
				Count: 300,
			},
			{
				Type:  1,
				Id:    33,
				Count: 300,
			},
			{
				Type:  1,
				Id:    33,
				Count: 300,
			},
			{
				Type:  1,
				Id:    33,
				Count: 300,
			},
			{
				Type:  1,
				Id:    33,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12911,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    34,
				Count: 50,
			},
			{
				Type:  1,
				Id:    34,
				Count: 50,
			},
			{
				Type:  1,
				Id:    34,
				Count: 50,
			},
			{
				Type:  1,
				Id:    34,
				Count: 100,
			},
			{
				Type:  1,
				Id:    34,
				Count: 100,
			},
			{
				Type:  1,
				Id:    34,
				Count: 150,
			},
			{
				Type:  1,
				Id:    34,
				Count: 150,
			},
			{
				Type:  1,
				Id:    34,
				Count: 200,
			},
			{
				Type:  1,
				Id:    34,
				Count: 200,
			},
			{
				Type:  1,
				Id:    34,
				Count: 200,
			},
			{
				Type:  1,
				Id:    34,
				Count: 300,
			},
			{
				Type:  1,
				Id:    34,
				Count: 300,
			},
			{
				Type:  1,
				Id:    34,
				Count: 300,
			},
			{
				Type:  1,
				Id:    34,
				Count: 300,
			},
			{
				Type:  1,
				Id:    34,
				Count: 300,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12912,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12913,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12914,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12915,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12916,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12917,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12918,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12919,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12920,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12921,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12922,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12923,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12924,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12925,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12926,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12927,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12928,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12929,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12930,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
	{
		Id: 12931,
		Fragment: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
			{
				Type:  1,
				Id:    3,
				Count: 3,
			},
		},
		NaterialScience: []*Consume{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 250,
			},
			{
				Type:  1,
				Id:    2,
				Count: 300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 700,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1100,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1300,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1450,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1750,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 1900,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2050,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2350,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2500,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2650,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 2800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 3800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 4800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 5800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 6800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7000,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7200,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7400,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7600,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
			{
				Type:  1,
				Id:    2,
				Count: 7800,
			},
		},
	},
}

var heroEvolveKeys = []int32{
	10021, 10121, 10222, 10321, 10421, 10521, 10622, 10721, 10822, 10921,
	11021, 11122, 11222, 11322, 11421, 11511, 11612, 11712, 11812, 11911,
	12011, 12112, 12212, 12312, 12412, 12511, 12611, 12711, 12811, 12911,
	12912, 12913, 12914, 12915, 12916, 12917, 12918, 12919, 12920, 12921,
	12922, 12923, 12924, 12925, 12926, 12927, 12928, 12929, 12930, 12931,
}

func init() {
	HeroEvolveData.data = make(map[int32]*HeroEvolveCfg)
	for i := 0; i < len(heroEvolveKeys); i++ {
		HeroEvolveData.data[heroEvolveKeys[i]] = heroEvolveValues[i]
	}
}
