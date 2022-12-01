package config

type ServerCodeCfg struct {
	Id       int32
	GlobalId string
}

type ServerCodeTable struct {
	data map[int32]*ServerCodeCfg
}

var ServerCodeData = &ServerCodeTable{
	data: map[int32]*ServerCodeCfg{},
}

func (table *ServerCodeTable) Get(id int32) *ServerCodeCfg {
	return table.data[id]
}

func (table *ServerCodeTable) GetAll() []int32 {
	return serverCodeKeys
}

var serverCodeValues = []*ServerCodeCfg{
	{
		Id:       110101,
		GlobalId: "ConfSkillTips6",
	},
	{
		Id:       110102,
		GlobalId: "ConfSkillTips5",
	},
	{
		Id:       110103,
		GlobalId: "ConfCommon6",
	},
	{
		Id:       110104,
		GlobalId: "ConfMagicTips",
	},
	{
		Id:       120101,
		GlobalId: "ConfServerCode2",
	},
	{
		Id:       140201,
		GlobalId: "ConfServerCode3",
	},
	{
		Id:       140202,
		GlobalId: "ConfHeroTips2",
	},
	{
		Id:       140203,
		GlobalId: "ConfCommon4",
	},
	{
		Id:       140301,
		GlobalId: "ConfServerCode3",
	},
	{
		Id:       140302,
		GlobalId: "ConfHeroTips3",
	},
	{
		Id:       140303,
		GlobalId: "ConfCommon4",
	},
	{
		Id:       140101,
		GlobalId: "ConfServerCode3",
	},
	{
		Id:       140102,
		GlobalId: "ConfServerCode4",
	},
	{
		Id:       140103,
		GlobalId: "ConfAscendTips1",
	},
	{
		Id:       140104,
		GlobalId: "ConfServerCode38",
	},
	{
		Id:       130302,
		GlobalId: "ConfServerCode5",
	},
	{
		Id:       130303,
		GlobalId: "ConfServerCode6",
	},
	{
		Id:       130304,
		GlobalId: "ConfServerCode7",
	},
	{
		Id:       130201,
		GlobalId: "ConfServerCode8",
	},
	{
		Id:       130202,
		GlobalId: "ConfAscendTips1",
	},
	{
		Id:       130204,
		GlobalId: "ConfServerCode38",
	},
	{
		Id:       130101,
		GlobalId: "ConfServerCode9",
	},
	{
		Id:       130102,
		GlobalId: "ConfServerCode10",
	},
	{
		Id:       130501,
		GlobalId: "ConfServerCode28",
	},
	{
		Id:       130503,
		GlobalId: "ConfServerCode29",
	},
	{
		Id:       160102,
		GlobalId: "ConfServerCode11",
	},
	{
		Id:       160103,
		GlobalId: "ConfServerCode12",
	},
	{
		Id:       160104,
		GlobalId: "ConfServerCode13",
	},
	{
		Id:       160105,
		GlobalId: "ConfServerCode14",
	},
	{
		Id:       160106,
		GlobalId: "ConfServerCode15",
	},
	{
		Id:       160107,
		GlobalId: "ConfServerCode16",
	},
	{
		Id:       160108,
		GlobalId: "ConfServerCode17",
	},
	{
		Id:       160109,
		GlobalId: "ConfServerCode18",
	},
	{
		Id:       170101,
		GlobalId: "ConfServerCode19",
	},
	{
		Id:       170102,
		GlobalId: "ConfServerCode20",
	},
	{
		Id:       170103,
		GlobalId: "ConfServerCode21",
	},
	{
		Id:       170104,
		GlobalId: "ConfServerCode22",
	},
	{
		Id:       170105,
		GlobalId: "ConfServerCode23",
	},
	{
		Id:       170106,
		GlobalId: "ConfServerCode24",
	},
	{
		Id:       170107,
		GlobalId: "ConfServerCode25",
	},
	{
		Id:       170202,
		GlobalId: "ConfServerCode31",
	},
	{
		Id:       170203,
		GlobalId: "ConfAscendTips1",
	},
	{
		Id:       170304,
		GlobalId: "ConfMagicTips",
	},
	{
		Id:       170305,
		GlobalId: "ConfServerCode55",
	},
	{
		Id:       150103,
		GlobalId: "ConfServerCode27",
	},
	{
		Id:       150104,
		GlobalId: "ConfEquipEnhanceTips",
	},
	{
		Id:       150105,
		GlobalId: "ConfServerCode26",
	},
	{
		Id:       180105,
		GlobalId: "ConfServerCode32",
	},
	{
		Id:       180302,
		GlobalId: "ConfServerCode33",
	},
	{
		Id:       180303,
		GlobalId: "ConfServerCode34",
	},
	{
		Id:       180501,
		GlobalId: "ConfServerCode35",
	},
	{
		Id:       180502,
		GlobalId: "ConfServerCode36",
	},
	{
		Id:       180503,
		GlobalId: "ConfServerCode37",
	},
	{
		Id:       190201,
		GlobalId: "ConfServerCode39",
	},
	{
		Id:       190301,
		GlobalId: "ConfServerCode40",
	},
	{
		Id:       190302,
		GlobalId: "ConfServerCode41",
	},
	{
		Id:       190401,
		GlobalId: "ConfServerCode42",
	},
	{
		Id:       190402,
		GlobalId: "ConfServerCode43",
	},
	{
		Id:       190403,
		GlobalId: "ConfServerCode44",
	},
	{
		Id:       200101,
		GlobalId: "ConfServerCode45",
	},
	{
		Id:       200102,
		GlobalId: "ConfServerCode46",
	},
	{
		Id:       200201,
		GlobalId: "ConfServerCode47",
	},
	{
		Id:       200202,
		GlobalId: "ConfServerCode48",
	},
	{
		Id:       200301,
		GlobalId: "ConfServerCode49",
	},
	{
		Id:       200302,
		GlobalId: "ConfServerCode50",
	},
	{
		Id:       200303,
		GlobalId: "ConfServerCode51",
	},
	{
		Id:       200401,
		GlobalId: "ConfServerCode52",
	},
	{
		Id:       200402,
		GlobalId: "ConfServerCode53",
	},
	{
		Id:       200601,
		GlobalId: "ConfServerCode54",
	},
	{
		Id:       200602,
		GlobalId: "ConfAscendTips1",
	},
	{
		Id:       210103,
		GlobalId: "ConfCommon6",
	},
	{
		Id:       210203,
		GlobalId: "ConfCommon6",
	},
	{
		Id:       220201,
		GlobalId: "ConfRoleSkill3",
	},
	{
		Id:       220202,
		GlobalId: "ConfSkillTips5",
	},
	{
		Id:       220203,
		GlobalId: "ConfSkillTips6",
	},
	{
		Id:       220301,
		GlobalId: "ConfInstanceTips4",
	},
	{
		Id:       220302,
		GlobalId: "ConfInstanceTips2",
	},
	{
		Id:       220303,
		GlobalId: "ConfInstanceTips3",
	},
	{
		Id:       250101,
		GlobalId: "ConfServerCode35",
	},
	{
		Id:       250103,
		GlobalId: "ConfServerCode52",
	},
	{
		Id:       270101,
		GlobalId: "ConfMonthCardTips1",
	},
	{
		Id:       280102,
		GlobalId: "ConfServerCode7",
	},
	{
		Id:       280104,
		GlobalId: "ConfCommon4",
	},
	{
		Id:       280202,
		GlobalId: "ConfServerCode7",
	},
	{
		Id:       280203,
		GlobalId: "ConfSummonExchangeTips",
	},
	{
		Id:       290101,
		GlobalId: "ConfPurchaseLimit",
	},
	{
		Id:       290103,
		GlobalId: "ConfServerCode41",
	},
	{
		Id:       300502,
		GlobalId: "ConfServerCode48",
	},
	{
		Id:       300503,
		GlobalId: "ConfServerCode56",
	},
	{
		Id:       300204,
		GlobalId: "ConfSettingTips3",
	},
	{
		Id:       320101,
		GlobalId: "ConfTopUpTip2",
	},
	{
		Id:       320102,
		GlobalId: "ConfShopFirstRecharge9",
	},
	{
		Id:       160110,
		GlobalId: "ConfBattleDateErr",
	},
	{
		Id:       220306,
		GlobalId: "ConfBattleDateErr",
	},
	{
		Id:       220405,
		GlobalId: "ConfBattleDateErr",
	},
	{
		Id:       330101,
		GlobalId: "ConfShopFirstRecharge9",
	},
	{
		Id:       330102,
		GlobalId: "ConfServerCode51",
	},
	{
		Id:       330103,
		GlobalId: "ConfTaskTips1",
	},
}

var serverCodeKeys = []int32{
	110101, 110102, 110103, 110104, 120101, 140201, 140202, 140203, 140301, 140302,
	140303, 140101, 140102, 140103, 140104, 130302, 130303, 130304, 130201, 130202,
	130204, 130101, 130102, 130501, 130503, 160102, 160103, 160104, 160105, 160106,
	160107, 160108, 160109, 170101, 170102, 170103, 170104, 170105, 170106, 170107,
	170202, 170203, 170304, 170305, 150103, 150104, 150105, 180105, 180302, 180303,
	180501, 180502, 180503, 190201, 190301, 190302, 190401, 190402, 190403, 200101,
	200102, 200201, 200202, 200301, 200302, 200303, 200401, 200402, 200601, 200602,
	210103, 210203, 220201, 220202, 220203, 220301, 220302, 220303, 250101, 250103,
	270101, 280102, 280104, 280202, 280203, 290101, 290103, 300502, 300503, 300204,
	320101, 320102, 160110, 220306, 220405, 330101, 330102, 330103,
}

func init() {
	ServerCodeData.data = make(map[int32]*ServerCodeCfg)
	for i := 0; i < len(serverCodeKeys); i++ {
		ServerCodeData.data[serverCodeKeys[i]] = serverCodeValues[i]
	}
}
