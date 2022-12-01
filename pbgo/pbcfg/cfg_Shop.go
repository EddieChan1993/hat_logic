package config

type ShopCfg struct {
	Id          int32
	Type        int32
	Item        []*Reward
	PriceDimond int32
	PriceGold   int32
	Limit       int32
	Rount       int32
	Lodlevel    int32
}

type ShopTable struct {
	data map[int32]*ShopCfg
}

var ShopData = &ShopTable{
	data: map[int32]*ShopCfg{},
}

func (table *ShopTable) Get(id int32) *ShopCfg {
	return table.data[id]
}

func (table *ShopTable) GetAll() []int32 {
	return shopKeys
}

var shopValues = []*ShopCfg{
	{
		Id:   10001,
		Type: 1,
		Item: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 1,
			},
		},
		PriceDimond: 140,
		Limit:       10,
		Rount:       100,
	},
	{
		Id:   10002,
		Type: 1,
		Item: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 10,
			},
		},
		PriceDimond: 1400,
		Limit:       5,
		Rount:       100,
	},
	{
		Id:   10003,
		Type: 1,
		Item: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 1,
			},
		},
		PriceDimond: 280,
		Limit:       10,
		Rount:       100,
	},
	{
		Id:   10004,
		Type: 1,
		Item: []*Reward{
			{
				Type:  1,
				Id:    51,
				Count: 10,
			},
		},
		PriceDimond: 2800,
		Limit:       5,
		Rount:       100,
	},
	{
		Id:   10005,
		Type: 1,
		Item: []*Reward{
			{
				Type:  1,
				Id:    45,
				Count: 1,
			},
		},
		PriceDimond: 100,
		Limit:       99,
		Rount:       100,
	},
	{
		Id:   10006,
		Type: 1,
		Item: []*Reward{
			{
				Type:  1,
				Id:    2005,
				Count: 1,
			},
		},
		PriceDimond: 200,
		Limit:       10,
		Rount:       100,
	},
	{
		Id:   20001,
		Type: 2,
		Item: []*Reward{
			{
				Type:  1,
				Id:    44,
				Count: 1,
			},
		},
		Limit: 5,
		Rount: 100,
	},
	{
		Id:   20002,
		Type: 2,
		Item: []*Reward{
			{
				Type:  1,
				Id:    46,
				Count: 1,
			},
		},
		Limit: 5,
		Rount: 100,
	},
	{
		Id:   20003,
		Type: 2,
		Item: []*Reward{
			{
				Type:  1,
				Id:    47,
				Count: 1,
			},
		},
		Limit: 5,
		Rount: 100,
	},
	{
		Id:   20004,
		Type: 2,
		Item: []*Reward{
			{
				Type:  1,
				Id:    1001,
				Count: 1,
			},
		},
		Limit: 5,
		Rount: 100,
	},
	{
		Id:   20005,
		Type: 2,
		Item: []*Reward{
			{
				Type:  1,
				Id:    2005,
				Count: 1,
			},
		},
		Limit: 3,
		Rount: 100,
	},
	{
		Id:   20006,
		Type: 2,
		Item: []*Reward{
			{
				Type:  1,
				Id:    50,
				Count: 1,
			},
		},
		Limit: 3,
		Rount: 100,
	},
}

var shopKeys = []int32{
	10001, 10002, 10003, 10004, 10005, 10006, 20001, 20002, 20003, 20004,
	20005, 20006,
}

func init() {
	ShopData.data = make(map[int32]*ShopCfg)
	for i := 0; i < len(shopKeys); i++ {
		ShopData.data[shopKeys[i]] = shopValues[i]
	}
}
