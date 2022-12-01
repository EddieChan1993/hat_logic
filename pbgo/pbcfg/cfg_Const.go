package config

type ConstCfg struct {
	Id        int32
	Attrib    string
	Value     int32
	ValueList []int32
}

type ConstTable struct {
	data map[int32]*ConstCfg
}

var ConstData = &ConstTable{
	data: map[int32]*ConstCfg{},
}

func (table *ConstTable) Get(id int32) *ConstCfg {
	return table.data[id]
}

func (table *ConstTable) GetAll() []int32 {
	return constKeys
}

var constValues = []*ConstCfg{
	{
		Id:     1,
		Attrib: "crittime_base",
		Value:  1500,
	},
	{
		Id:     2,
		Attrib: "hit_base",
		Value:  700,
	},
	{
		Id:     3,
		Attrib: "dodge_base",
		Value:  100,
	},
	{
		Id:     4,
		Attrib: "faction_restraint",
		Value:  250,
	},
	{
		Id:     5,
		Attrib: "critaAll",
		ValueList: []int32{
			0,
			250,
			450,
			700,
			1000,
			1200,
			1500,
		},
	},
	{
		Id:     6,
		Attrib: "ranges",
		ValueList: []int32{
			0,
			70000,
			200000,
			400000,
			800000,
			1500000,
			3000000,
		},
	},
	{
		Id:     7,
		Attrib: "hit_correction",
		Value:  200,
	},
	{
		Id:     8,
		Attrib: "resetConsume",
		ValueList: []int32{
			0,
			100,
			100,
			100,
			200,
			200,
			200,
		},
	},
	{
		Id:     9,
		Attrib: "pickUp",
		ValueList: []int32{
			30,
			30,
		},
	},
	{
		Id:     10,
		Attrib: "worker_num_initial",
		Value:  10,
	},
	{
		Id:     11,
		Attrib: "equipment_sell_control",
		Value:  1,
	},
	{
		Id:     12,
		Attrib: "equipment_sell_price",
		Value:  1,
	},
	{
		Id:     13,
		Attrib: "artifact_unlock",
		Value:  5010,
	},
	{
		Id:     14,
		Attrib: "artifact_refresh",
		ValueList: []int32{
			1,
			1,
			1,
			2,
			2,
			2,
			3,
			3,
			3,
			5,
		},
	},
	{
		Id:     15,
		Attrib: "mercenary_unlock",
		Value:  7010,
	},
	{
		Id:     16,
		Attrib: "mercenary_num_initial",
		Value:  2,
	},
	{
		Id:     17,
		Attrib: "mercenary_price",
		ValueList: []int32{
			1000,
			2000,
			5000,
		},
	},
	{
		Id:     18,
		Attrib: "mercenary_missions",
		Value:  3,
	},
	{
		Id:     19,
		Attrib: "mercenary_refresh1",
		Value:  1,
	},
	{
		Id:     20,
		Attrib: "mercenary_name",
		Value:  71,
	},
	{
		Id:     21,
		Attrib: "attack_interval",
		Value:  100,
	},
	{
		Id:     22,
		Attrib: "slot_position",
		ValueList: []int32{
			1,
			5,
			8,
			10,
		},
	},
	{
		Id:     23,
		Attrib: "hook_time",
		Value:  43200,
	},
	{
		Id:     24,
		Attrib: "settlement_cycle",
		Value:  10,
	},
	{
		Id:     25,
		Attrib: "mercenary_refresh2",
		Value:  20,
	},
	{
		Id:     26,
		Attrib: "instance_unlock",
		Value:  1002,
	},
	{
		Id:     27,
		Attrib: "instanceBoss_dailytime",
		Value:  2,
	},
	{
		Id:     28,
		Attrib: "instanceBoss_rank",
		Value:  100,
	},
	{
		Id:     29,
		Attrib: "instanceTowerType",
		ValueList: []int32{
			1,
			2,
			3,
		},
	},
	{
		Id:     30,
		Attrib: "instance_boss_unlock",
		Value:  1002,
	},
	{
		Id:     31,
		Attrib: "instance_coin_unlock",
		Value:  1002,
	},
	{
		Id:     32,
		Attrib: "instance_resource_unlock",
		Value:  1002,
	},
	{
		Id:     33,
		Attrib: "instance_enhance_unlock",
		Value:  1002,
	},
	{
		Id:     34,
		Attrib: "shop_random",
		Value:  6,
	},
	{
		Id:     35,
		Attrib: "energy_treasure_chest",
		Value:  500,
	},
	{
		Id:     36,
		Attrib: "territory_unlock",
		Value:  2010,
	},
	{
		Id:     37,
		Attrib: "legend_freetime",
		Value:  14400,
	},
	{
		Id:     38,
		Attrib: "restore_goldcc",
		Value:  0,
	},
	{
		Id:     39,
		Attrib: "resources_recovery",
		Value:  0,
	},
	{
		Id:     40,
		Attrib: "strengthen_recovery",
		Value:  0,
	},
	{
		Id:     41,
		Attrib: "equip_attribute_entries",
		ValueList: []int32{
			1,
			1,
			2,
			3,
			4,
		},
	},
	{
		Id:     42,
		Attrib: "change_name_cost",
		Value:  100,
	},
}

var constKeys = []int32{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42,
}

func init() {
	ConstData.data = make(map[int32]*ConstCfg)
	for i := 0; i < len(constKeys); i++ {
		ConstData.data[constKeys[i]] = constValues[i]
	}
}
