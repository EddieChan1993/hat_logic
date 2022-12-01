package config

type RechargeCfg struct {
	Id             int32
	PayId          string
	Name           string
	ActType        int32
	Price          float64
	PriceStr       string
	VouchersDollar int32
	PriceCn        string
	PriceCnStr     string
	Diamonds       int32
	Icon           string
	DailyGems      []*Reward
	Days           int32
	EventName      string
	AppsFlyerId    int32
}

type RechargeTable struct {
	data                    map[int32]*RechargeCfg
	rechargePayIdIndexMap   map[string][]int32
	rechargeActTypeIndexMap map[int32][]int32
}

var RechargeData = &RechargeTable{
	data: map[int32]*RechargeCfg{},
	rechargePayIdIndexMap: map[string][]int32{
		"com.droidhang.re.diamond1": {1},
		"com.droidhang.re.diamond2": {2},
		"com.droidhang.re.diamond3": {3},
		"com.droidhang.re.diamond4": {4},
		"com.droidhang.re.diamond5": {5},
		"com.droidhang.re.diamond6": {6},
		"com.droidhang.re.diamond7": {7},
		"com.droidhang.re.diamond8": {8},
		"com.droidhang.re.diamond9": {9},
	},
	rechargeActTypeIndexMap: map[int32][]int32{
		1: {4, 5, 6, 7, 8, 9},
		2: {1, 2},
		3: {3},
	},
}

func (table *RechargeTable) Get(id int32) *RechargeCfg {
	return table.data[id]
}

func (table *RechargeTable) GetAll() []int32 {
	return rechargeKeys
}

var rechargeValues = []*RechargeCfg{
	{
		Id:             1,
		PayId:          "com.droidhang.re.diamond1",
		Name:           "高级月卡",
		ActType:        2,
		Price:          4.99,
		PriceStr:       "$4.99",
		VouchersDollar: 5,
		PriceCn:        "30",
		PriceCnStr:     "¥30",
		Diamonds:       300,
		DailyGems: []*Reward{
			{
				Type:  1,
				Id:    3,
				Count: 150,
			},
		},
		Days:        30,
		EventName:   "normal_monthly_card",
		AppsFlyerId: 2,
	},
	{
		Id:             2,
		PayId:          "com.droidhang.re.diamond2",
		Name:           "终身月卡",
		ActType:        2,
		Price:          14.99,
		PriceStr:       "$14.99",
		VouchersDollar: 15,
		PriceCn:        "98",
		PriceCnStr:     "¥98",
		Diamonds:       900,
		DailyGems: []*Reward{
			{
				Type:  1,
				Id:    3,
				Count: 300,
			},
		},
		Days:        9999,
		EventName:   "super_monthly_card",
		AppsFlyerId: 8,
	},
	{
		Id:             3,
		PayId:          "com.droidhang.re.diamond3",
		Name:           "勇士令",
		ActType:        3,
		Price:          19.99,
		PriceStr:       "$19.99",
		VouchersDollar: 20,
		PriceCn:        "128",
		PriceCnStr:     "128",
		EventName:      "19d99_purchase",
		AppsFlyerId:    9,
	},
	{
		Id:             4,
		PayId:          "com.droidhang.re.diamond4",
		Name:           "80钻石",
		ActType:        1,
		Price:          0.99,
		PriceStr:       "$0.99",
		VouchersDollar: 1,
		PriceCn:        "6",
		PriceCnStr:     "¥6",
		Diamonds:       80,
		Icon:           "shop_diamond1",
		EventName:      "0d99_purchase",
		AppsFlyerId:    60,
	},
	{
		Id:             5,
		PayId:          "com.droidhang.re.diamond5",
		Name:           "500钻石",
		ActType:        1,
		Price:          4.99,
		PriceStr:       "$4.99",
		VouchersDollar: 5,
		PriceCn:        "30",
		PriceCnStr:     "¥30",
		Diamonds:       500,
		Icon:           "shop_diamond2",
		EventName:      "4d99_purchase",
		AppsFlyerId:    2,
	},
	{
		Id:             6,
		PayId:          "com.droidhang.re.diamond6",
		Name:           "1200钻石",
		ActType:        1,
		Price:          9.99,
		PriceStr:       "$9.99",
		VouchersDollar: 10,
		PriceCn:        "68",
		PriceCnStr:     "¥68",
		Diamonds:       1200,
		Icon:           "shop_diamond3",
		EventName:      "9d99_purchase",
		AppsFlyerId:    3,
	},
	{
		Id:             7,
		PayId:          "com.droidhang.re.diamond7",
		Name:           "2500钻石",
		ActType:        1,
		Price:          19.99,
		PriceStr:       "$19.99",
		VouchersDollar: 20,
		PriceCn:        "128",
		PriceCnStr:     "¥128",
		Diamonds:       2500,
		Icon:           "shop_diamond4",
		EventName:      "19d99_purchase",
		AppsFlyerId:    4,
	},
	{
		Id:             8,
		PayId:          "com.droidhang.re.diamond8",
		Name:           "6500钻石",
		ActType:        1,
		Price:          49.99,
		PriceStr:       "$49.99",
		VouchersDollar: 50,
		PriceCn:        "328",
		PriceCnStr:     "¥328",
		Diamonds:       6500,
		Icon:           "shop_diamond5",
		EventName:      "49d99_purchase",
		AppsFlyerId:    6,
	},
	{
		Id:             9,
		PayId:          "com.droidhang.re.diamond9",
		Name:           "14000钻石",
		ActType:        1,
		Price:          99.99,
		PriceStr:       "$99.99",
		VouchersDollar: 100,
		PriceCn:        "648",
		PriceCnStr:     "¥648",
		Diamonds:       14000,
		Icon:           "shop_diamond6",
		EventName:      "99d99_purchase",
		AppsFlyerId:    7,
	},
}

var rechargeKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
}

func init() {
	RechargeData.data = make(map[int32]*RechargeCfg)
	for i := 0; i < len(rechargeKeys); i++ {
		RechargeData.data[rechargeKeys[i]] = rechargeValues[i]
	}
}

func (table *RechargeTable) GetByPayId(PayId string) (res []*RechargeCfg) {
	for _, i := range table.rechargePayIdIndexMap[PayId] {
		res = append(res, table.Get(i))
	}
	return
}

func (table *RechargeTable) GetByActType(ActType int32) (res []*RechargeCfg) {
	for _, i := range table.rechargeActTypeIndexMap[ActType] {
		res = append(res, table.Get(i))
	}
	return
}
