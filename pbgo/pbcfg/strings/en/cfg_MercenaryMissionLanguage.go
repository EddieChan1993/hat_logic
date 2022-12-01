package config

type MercenaryMissionLanguageCfg struct {
	Id       int32
	Describe string
}

type MercenaryMissionLanguageTable struct {
	data map[int32]*MercenaryMissionLanguageCfg
}

var MercenaryMissionLanguageData = &MercenaryMissionLanguageTable{
	data: map[int32]*MercenaryMissionLanguageCfg{},
}

func (table *MercenaryMissionLanguageTable) Get(id int32) *MercenaryMissionLanguageCfg {
	return table.data[id]
}

func (table *MercenaryMissionLanguageTable) GetAll() []int32 {
	return mercenaryMissionLanguageKeys
}

var mercenaryMissionLanguageValues = []*MercenaryMissionLanguageCfg{
	{
		Id:       1001,
		Describe: "A 100% chance to obtain Diamond * 50.",
	},
	{
		Id:       1002,
		Describe: "A 50% chance to obtain Diamond * 50.",
	},
	{
		Id:       1003,
		Describe: "A 30% chance to obtain Diamond * 50.",
	},
	{
		Id:       1004,
		Describe: "A 15% chance to obtain Diamond * 50.",
	},
	{
		Id:       1005,
		Describe: "A 10% chance to obtain Diamond * 50.",
	},
	{
		Id:       1006,
		Describe: "A 5% chance to obtain Diamond * 50.",
	},
	{
		Id:       1101,
		Describe: "A 100% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       1102,
		Describe: "A 50% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       1103,
		Describe: "A 30% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       1104,
		Describe: "A 15% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       1105,
		Describe: "A 10% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       1106,
		Describe: "A 5% chance to obtain Coin Bag * 2.",
	},
	{
		Id:       1201,
		Describe: "A 100% chance to obtain Advanced Summon Scroll * 1.",
	},
	{
		Id:       1202,
		Describe: "A 50% chance to obtain Advanced Summon Scroll * 1.",
	},
	{
		Id:       1203,
		Describe: "A 30% chance to obtain Advanced Summon Scroll * 1.",
	},
	{
		Id:       1204,
		Describe: "A 15% chance to obtain Advanced Summon Scroll * 1.",
	},
	{
		Id:       1205,
		Describe: "A 10% chance to obtain Advanced Summon Scroll * 1.",
	},
	{
		Id:       1206,
		Describe: "A 5% chance to obtain Advanced Summon Scroll * 2.",
	},
	{
		Id:       1301,
		Describe: "A 100% chance to obtain Mercenary Commission * 1.",
	},
	{
		Id:       1302,
		Describe: "A 50% chance to obtain Mercenary Commission * 1.",
	},
	{
		Id:       1303,
		Describe: "A 30% chance to obtain Mercenary Commission * 1.",
	},
	{
		Id:       1304,
		Describe: "A 15% chance to obtain Mercenary Commission * 1.",
	},
	{
		Id:       1305,
		Describe: "A 10% chance to obtain Mercenary Commission * 1.",
	},
	{
		Id:       1306,
		Describe: "A 5% chance to obtain Mercenary Commission * 2.",
	},
	{
		Id:       2001,
		Describe: "A 100% chance to obtain Diamond * 100.",
	},
	{
		Id:       2002,
		Describe: "A 50% chance to obtain Diamond * 100.",
	},
	{
		Id:       2003,
		Describe: "A 30% chance to obtain Diamond * 100.",
	},
	{
		Id:       2004,
		Describe: "A 15% chance to obtain Diamond * 100.",
	},
	{
		Id:       2005,
		Describe: "A 10% chance to obtain Diamond * 100.",
	},
	{
		Id:       2006,
		Describe: "A 5% chance to obtain Diamond * 100.",
	},
	{
		Id:       2101,
		Describe: "A 100% chance to obtain Coin Bag * 2.",
	},
	{
		Id:       2102,
		Describe: "A 50% chance to obtain Coin Bag * 2.",
	},
	{
		Id:       2103,
		Describe: "A 30% chance to obtain Coin Bag * 2.",
	},
	{
		Id:       2104,
		Describe: "A 15% chance to obtain Coin Bag * 2.",
	},
	{
		Id:       2105,
		Describe: "A 10% chance to obtain Coin Bag * 2.",
	},
	{
		Id:       2106,
		Describe: "A 5% chance to obtain Coin Bag * 3.",
	},
	{
		Id:       2201,
		Describe: "A 100% chance to obtain Advanced Summon Scroll * 2.",
	},
	{
		Id:       2202,
		Describe: "A 50% chance to obtain Advanced Summon Scroll * 2.",
	},
	{
		Id:       2203,
		Describe: "A 30% chance to obtain Advanced Summon Scroll * 2.",
	},
	{
		Id:       2204,
		Describe: "A 15% chance to obtain Advanced Summon Scroll * 2.",
	},
	{
		Id:       2205,
		Describe: "A 10% chance to obtain Advanced Summon Scroll * 2.",
	},
	{
		Id:       2206,
		Describe: "A 5% chance to obtain Advanced Summon Scroll * 3.",
	},
	{
		Id:       2301,
		Describe: "A 100% chance to obtain Mercenary Commission * 2.",
	},
	{
		Id:       2302,
		Describe: "A 50% chance to obtain Mercenary Commission * 2.",
	},
	{
		Id:       2303,
		Describe: "A 30% chance to obtain Mercenary Commission * 2.",
	},
	{
		Id:       2304,
		Describe: "A 15% chance to obtain Mercenary Commission * 2.",
	},
	{
		Id:       2305,
		Describe: "A 10% chance to obtain Mercenary Commission * 2.",
	},
	{
		Id:       2306,
		Describe: "A 5% chance to obtain Mercenary Commission * 3.",
	},
	{
		Id:       3001,
		Describe: "A 100% chance to obtain Diamond * 150.",
	},
	{
		Id:       3002,
		Describe: "A 50% chance to obtain Diamond * 150.",
	},
	{
		Id:       3003,
		Describe: "A 30% chance to obtain Diamond * 150.",
	},
	{
		Id:       3004,
		Describe: "A 15% chance to obtain Diamond * 150.",
	},
	{
		Id:       3005,
		Describe: "A 10% chance to obtain Diamond * 150.",
	},
	{
		Id:       3006,
		Describe: "A 5% chance to obtain Diamond * 150.",
	},
	{
		Id:       3101,
		Describe: "A 100% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       3102,
		Describe: "A 50% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       3103,
		Describe: "A 30% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       3104,
		Describe: "A 15% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       3105,
		Describe: "A 10% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       3106,
		Describe: "A 5% chance to obtain Coin Bag * 4.",
	},
	{
		Id:       3201,
		Describe: "A 100% chance to obtain Advanced Summon Scroll * 3.",
	},
	{
		Id:       3202,
		Describe: "A 50% chance to obtain Advanced Summon Scroll * 3.",
	},
	{
		Id:       3203,
		Describe: "A 30% chance to obtain Advanced Summon Scroll * 3.",
	},
	{
		Id:       3204,
		Describe: "A 15% chance to obtain Advanced Summon Scroll * 3.",
	},
	{
		Id:       3205,
		Describe: "A 10% chance to obtain Advanced Summon Scroll * 3.",
	},
	{
		Id:       3206,
		Describe: "A 5% chance to obtain Advanced Summon Scroll * 4.",
	},
	{
		Id:       3301,
		Describe: "A 100% chance to obtain Mercenary Commission * 3.",
	},
	{
		Id:       3302,
		Describe: "A 50% chance to obtain Mercenary Commission * 3.",
	},
	{
		Id:       3303,
		Describe: "A 30% chance to obtain Mercenary Commission * 3.",
	},
	{
		Id:       3304,
		Describe: "A 15% chance to obtain Mercenary Commission * 3.",
	},
	{
		Id:       3305,
		Describe: "A 10% chance to obtain Mercenary Commission * 3.",
	},
	{
		Id:       3306,
		Describe: "A 5% chance to obtain Mercenary Commission * 4.",
	},
	{
		Id:       3402,
		Describe: "A 50% chance to obtain Mystic Summon Scroll * 1.",
	},
	{
		Id:       3403,
		Describe: "A 30% chance to obtain Mystic Summon Scroll * 1.",
	},
	{
		Id:       3404,
		Describe: "A 15% chance to obtain Mystic Summon Scroll * 1.",
	},
	{
		Id:       3405,
		Describe: "A 10% chance to obtain Mystic Summon Scroll * 1.",
	},
	{
		Id:       3406,
		Describe: "A 5% chance to obtain Mystic Summon Scroll * 2.",
	},
	{
		Id:       4001,
		Describe: "A 100% chance to obtain Diamond * 300.",
	},
	{
		Id:       4002,
		Describe: "A 50% chance to obtain Diamond * 301.",
	},
	{
		Id:       4003,
		Describe: "A 30% chance to obtain Diamond * 302.",
	},
	{
		Id:       4004,
		Describe: "A 15% chance to obtain Diamond * 303.",
	},
	{
		Id:       4005,
		Describe: "A 10% chance to obtain Diamond * 304.",
	},
	{
		Id:       4006,
		Describe: "A 5% chance to obtain Diamond * 305.",
	},
	{
		Id:       4101,
		Describe: "A 100% chance to obtain Coin Bag * 5.",
	},
	{
		Id:       4102,
		Describe: "A 50% chance to obtain Coin Bag * 5.",
	},
	{
		Id:       4103,
		Describe: "A 30% chance to obtain Coin Bag * 5.",
	},
	{
		Id:       4104,
		Describe: "A 15% chance to obtain Coin Bag * 5.",
	},
	{
		Id:       4105,
		Describe: "A 10% chance to obtain Coin Bag * 5.",
	},
	{
		Id:       4106,
		Describe: "A 5% chance to obtain Coin Bag * 1.",
	},
	{
		Id:       4201,
		Describe: "A 100% chance to obtain Advanced Summon Scroll * 5.",
	},
	{
		Id:       4202,
		Describe: "A 50% chance to obtain Advanced Summon Scroll * 5.",
	},
	{
		Id:       4203,
		Describe: "A 30% chance to obtain Advanced Summon Scroll * 5.",
	},
	{
		Id:       4204,
		Describe: "A 15% chance to obtain Advanced Summon Scroll * 5.",
	},
	{
		Id:       4205,
		Describe: "A 10% chance to obtain Advanced Summon Scroll * 5.",
	},
	{
		Id:       4206,
		Describe: "A 5% chance to obtain Advanced Summon Scroll * 6.",
	},
	{
		Id:       4301,
		Describe: "A 100% chance to obtain Mercenary Commission * 5.",
	},
	{
		Id:       4302,
		Describe: "A 50% chance to obtain Mercenary Commission * 5.",
	},
	{
		Id:       4303,
		Describe: "A 30% chance to obtain Mercenary Commission * 5.",
	},
	{
		Id:       4304,
		Describe: "A 15% chance to obtain Mercenary Commission * 5.",
	},
	{
		Id:       4305,
		Describe: "A 10% chance to obtain Mercenary Commission * 5.",
	},
	{
		Id:       4306,
		Describe: "A 5% chance to obtain Mercenary Commission * 6.",
	},
	{
		Id:       4401,
		Describe: "A 100% chance to obtain Mystic Summon Scroll * 1.",
	},
	{
		Id:       4402,
		Describe: "A 50% chance to obtain Mystic Summon Scroll * 2.",
	},
	{
		Id:       4403,
		Describe: "A 30% chance to obtain Mystic Summon Scroll * 2.",
	},
	{
		Id:       4404,
		Describe: "A 15% chance to obtain Mystic Summon Scroll * 2.",
	},
	{
		Id:       4405,
		Describe: "A 10% chance to obtain Mystic Summon Scroll * 2.",
	},
	{
		Id:       4406,
		Describe: "A 5% chance to obtain Mystic Summon Scroll * 3.",
	},
	{
		Id:       4501,
		Describe: "A 50% chance to obtain Star-Piercing Bow * 1.",
	},
	{
		Id:       4502,
		Describe: "A 50% chance to obtain Golden Spear * 1.",
	},
	{
		Id:       4503,
		Describe: "A 50% chance to obtain Trident of Raging Wave * 1.",
	},
	{
		Id:       4504,
		Describe: "A 50% chance to obtain Shield of Victory * 1.",
	},
	{
		Id:       4505,
		Describe: "A 50% chance to obtain Double-Snake Staff * 1.",
	},
	{
		Id:       4506,
		Describe: "A 50% chance to obtain Cup of Divine Feast * 1.",
	},
	{
		Id:       4507,
		Describe: "A 50% chance to obtain Scepter of Zeus * 1.",
	},
	{
		Id:       4508,
		Describe: "A 50% chance to obtain Eye of Hades * 1.",
	},
	{
		Id:       4509,
		Describe: "A 50% chance to obtain Golden Armor * 1.",
	},
	{
		Id:       4510,
		Describe: "A 50% chance to obtain Crown of the Sun * 1.",
	},
	{
		Id:       4511,
		Describe: "A 30% chance to obtain Star-Piercing Bow * 1.",
	},
	{
		Id:       4512,
		Describe: "A 30% chance to obtain Golden Spear * 1.",
	},
	{
		Id:       4513,
		Describe: "A 30% chance to obtain Trident of Raging Wave * 1.",
	},
	{
		Id:       4514,
		Describe: "A 30% chance to obtain Shield of Victory * 1.",
	},
	{
		Id:       4515,
		Describe: "A 30% chance to obtain Double-Snake Staff * 1.",
	},
	{
		Id:       4516,
		Describe: "A 30% chance to obtain Cup of Divine Feast * 1.",
	},
	{
		Id:       4517,
		Describe: "A 30% chance to obtain Scepter of Zeus * 1.",
	},
	{
		Id:       4518,
		Describe: "A 30% chance to obtain Eye of Hades * 1.",
	},
	{
		Id:       4519,
		Describe: "A 30% chance to obtain Golden Armor * 1.",
	},
	{
		Id:       4520,
		Describe: "A 30% chance to obtain Crown of the Sun * 1.",
	},
	{
		Id:       4521,
		Describe: "A 15% chance to obtain Star-Piercing Bow * 1.",
	},
	{
		Id:       4522,
		Describe: "A 15% chance to obtain Golden Spear * 1.",
	},
	{
		Id:       4523,
		Describe: "A 15% chance to obtain Trident of Raging Wave * 1.",
	},
	{
		Id:       4524,
		Describe: "A 15% chance to obtain Shield of Victory * 1.",
	},
	{
		Id:       4525,
		Describe: "A 15% chance to obtain Double-Snake Staff * 1.",
	},
	{
		Id:       4526,
		Describe: "A 15% chance to obtain Cup of Divine Feast * 1.",
	},
	{
		Id:       4527,
		Describe: "A 15% chance to obtain Scepter of Zeus * 1.",
	},
	{
		Id:       4528,
		Describe: "A 15% chance to obtain Eye of Hades * 1.",
	},
	{
		Id:       4529,
		Describe: "A 15% chance to obtain Golden Armor * 1.",
	},
	{
		Id:       4530,
		Describe: "A 15% chance to obtain Crown of the Sun * 1.",
	},
	{
		Id:       4531,
		Describe: "A 10% chance to obtain Star-Piercing Bow * 1.",
	},
	{
		Id:       4532,
		Describe: "A 10% chance to obtain Golden Spear * 1.",
	},
	{
		Id:       4533,
		Describe: "A 10% chance to obtain Trident of Raging Wave * 1.",
	},
	{
		Id:       4534,
		Describe: "A 10% chance to obtain Shield of Victory * 1.",
	},
	{
		Id:       4535,
		Describe: "A 10% chance to obtain Double-Snake Staff * 1.",
	},
	{
		Id:       4536,
		Describe: "A 10% chance to obtain Cup of Divine Feast * 1.",
	},
	{
		Id:       4537,
		Describe: "A 10% chance to obtain Scepter of Zeus * 1.",
	},
	{
		Id:       4538,
		Describe: "A 10% chance to obtain Eye of Hades * 1.",
	},
	{
		Id:       4539,
		Describe: "A 10% chance to obtain Golden Armor * 1.",
	},
	{
		Id:       4540,
		Describe: "A 10% chance to obtain Crown of the Sun * 1.",
	},
	{
		Id:       4541,
		Describe: "A 5% chance to obtain Star-Piercing Bow * 2.",
	},
	{
		Id:       4542,
		Describe: "A 5% chance to obtain Golden Spear * 2.",
	},
	{
		Id:       4543,
		Describe: "A 5% chance to obtain Trident of Raging Wave * 2.",
	},
	{
		Id:       4544,
		Describe: "A 5% chance to obtain Shield of Victory * 2.",
	},
	{
		Id:       4545,
		Describe: "A 5% chance to obtain Double-Snake Staff * 2.",
	},
	{
		Id:       4546,
		Describe: "A 5% chance to obtain Cup of Divine Feast * 2.",
	},
	{
		Id:       4547,
		Describe: "A 5% chance to obtain Scepter of Zeus * 2.",
	},
	{
		Id:       4548,
		Describe: "A 5% chance to obtain Eye of Hades * 2.",
	},
	{
		Id:       4549,
		Describe: "A 5% chance to obtain Golden Armor * 2.",
	},
	{
		Id:       4550,
		Describe: "A 5% chance to obtain Crown of the Sun * 2.",
	},
}

var mercenaryMissionLanguageKeys = []int32{
	1001, 1002, 1003, 1004, 1005, 1006, 1101, 1102, 1103, 1104,
	1105, 1106, 1201, 1202, 1203, 1204, 1205, 1206, 1301, 1302,
	1303, 1304, 1305, 1306, 2001, 2002, 2003, 2004, 2005, 2006,
	2101, 2102, 2103, 2104, 2105, 2106, 2201, 2202, 2203, 2204,
	2205, 2206, 2301, 2302, 2303, 2304, 2305, 2306, 3001, 3002,
	3003, 3004, 3005, 3006, 3101, 3102, 3103, 3104, 3105, 3106,
	3201, 3202, 3203, 3204, 3205, 3206, 3301, 3302, 3303, 3304,
	3305, 3306, 3402, 3403, 3404, 3405, 3406, 4001, 4002, 4003,
	4004, 4005, 4006, 4101, 4102, 4103, 4104, 4105, 4106, 4201,
	4202, 4203, 4204, 4205, 4206, 4301, 4302, 4303, 4304, 4305,
	4306, 4401, 4402, 4403, 4404, 4405, 4406, 4501, 4502, 4503,
	4504, 4505, 4506, 4507, 4508, 4509, 4510, 4511, 4512, 4513,
	4514, 4515, 4516, 4517, 4518, 4519, 4520, 4521, 4522, 4523,
	4524, 4525, 4526, 4527, 4528, 4529, 4530, 4531, 4532, 4533,
	4534, 4535, 4536, 4537, 4538, 4539, 4540, 4541, 4542, 4543,
	4544, 4545, 4546, 4547, 4548, 4549, 4550,
}

func init() {
	MercenaryMissionLanguageData.data = make(map[int32]*MercenaryMissionLanguageCfg)
	for i := 0; i < len(mercenaryMissionLanguageKeys); i++ {
		MercenaryMissionLanguageData.data[mercenaryMissionLanguageKeys[i]] = mercenaryMissionLanguageValues[i]
	}
}
