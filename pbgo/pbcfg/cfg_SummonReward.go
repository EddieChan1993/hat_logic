package config

type SummonRewardCfg struct {
	Id        int32
	Type      int32
	Reward    []*Reward
	Weight    int32
	Guarantee int32
}

type SummonRewardTable struct {
	data                     map[int32]*SummonRewardCfg
	summonRewardTypeIndexMap map[int32][]int32
}

var SummonRewardData = &SummonRewardTable{
	data: map[int32]*SummonRewardCfg{},
	summonRewardTypeIndexMap: map[int32][]int32{
		101: {10101, 10102, 10103, 10104, 10105, 10106, 10107, 10108, 10109, 10110, 10111, 10112, 10113, 10114, 10115, 10116, 10117, 10118, 10119, 10120, 10121, 10122, 10123, 10124, 10125, 10126, 10127, 10128, 10129, 10130, 10131, 10132, 10133, 10134, 10135, 10136, 10137, 10138, 10139, 10140, 10141, 10142, 10143, 10144, 10145, 10146, 10147, 10148, 10149, 10150, 10151, 10152, 10153, 10154, 10155, 10156, 10157},
		102: {11101, 11102, 11103, 11104, 11105, 11106, 11107, 11108, 11109, 11110, 11111, 11112, 11113, 11114, 11115, 11116, 11117, 11118, 11119, 11120, 11121, 11122, 11123, 11124, 11125, 11126, 11127, 11128, 11129, 11130, 11131, 11132, 11133, 11134, 11135, 11136, 11137, 11138, 11139, 11140, 11141, 11142, 11143, 11144, 11145, 11146, 11147, 11148, 11149, 11150, 11151, 11152, 11153, 11154, 11155, 11156, 11157},
		103: {12101, 12102, 12103, 12104, 12105, 12106, 12107, 12108, 12109, 12110, 12111, 12112, 12113, 12114, 12115, 12116, 12117, 12118, 12119, 12120, 12121, 12122, 12123, 12124, 12125, 12126, 12127, 12128, 12129, 12130, 12131, 12132, 12133, 12134, 12135, 12136, 12137, 12138, 12139, 12140, 12141, 12142, 12143, 12144, 12145, 12146, 12147, 12148, 12149, 12150, 12151, 12152, 12153, 12154, 12155, 12156, 12157},
		104: {13101, 13102, 13103, 13104, 13105, 13106, 13107, 13108, 13109, 13110, 13111, 13112, 13113, 13114, 13115, 13116, 13117, 13118, 13119, 13120, 13121, 13122, 13123, 13124, 13125, 13126, 13127, 13128, 13129, 13130, 13131, 13132, 13133, 13134, 13135, 13136, 13137, 13138, 13139, 13140, 13141, 13142, 13143, 13144, 13145, 13146, 13147, 13148, 13149, 13150, 13151, 13152, 13153, 13154, 13155, 13156, 13157},
		105: {14101, 14102, 14103, 14104, 14105, 14106, 14107, 14108, 14109, 14110, 14111, 14112, 14113, 14114, 14115, 14116, 14117, 14118, 14119, 14120, 14121, 14122, 14123, 14124, 14125, 14126, 14127, 14128, 14129, 14130, 14131, 14132, 14133, 14134, 14135, 14136, 14137, 14138, 14139, 14140, 14141, 14142, 14143, 14144, 14145, 14146, 14147, 14148, 14149, 14150, 14151, 14152, 14153, 14154, 14155, 14156, 14157},
		106: {15101, 15102, 15103, 15104, 15105, 15106, 15107, 15108, 15109, 15110, 15111, 15112, 15113, 15114, 15115, 15116, 15117, 15118, 15119, 15120, 15121, 15122, 15123, 15124, 15125, 15126, 15127, 15128, 15129, 15130, 15131, 15132, 15133, 15134, 15135, 15136, 15137, 15138, 15139, 15140, 15141, 15142, 15143, 15144, 15145, 15146, 15147, 15148, 15149, 15150, 15151, 15152, 15153, 15154, 15155, 15156, 15157},
		107: {16101, 16102, 16103, 16104, 16105, 16106, 16107, 16108, 16109, 16110, 16111, 16112, 16113, 16114, 16115, 16116, 16117, 16118, 16119, 16120, 16121, 16122, 16123, 16124, 16125, 16126, 16127, 16128, 16129, 16130, 16131, 16132, 16133, 16134, 16135, 16136, 16137, 16138, 16139, 16140, 16141, 16142, 16143, 16144, 16145, 16146, 16147, 16148, 16149, 16150, 16151, 16152, 16153, 16154, 16155, 16156, 16157},
		108: {17101, 17102, 17103, 17104, 17105, 17106, 17107, 17108, 17109, 17110, 17111, 17112, 17113, 17114, 17115, 17116, 17117, 17118, 17119, 17120, 17121, 17122, 17123, 17124, 17125, 17126, 17127, 17128, 17129, 17130, 17131, 17132, 17133, 17134, 17135, 17136, 17137, 17138, 17139, 17140, 17141, 17142, 17143, 17144, 17145, 17146, 17147, 17148, 17149, 17150, 17151, 17152, 17153, 17154, 17155, 17156, 17157},
		109: {18101, 18102, 18103, 18104, 18105, 18106, 18107, 18108, 18109, 18110, 18111, 18112, 18113, 18114, 18115, 18116, 18117, 18118, 18119, 18120, 18121, 18122, 18123, 18124, 18125, 18126, 18127, 18128, 18129, 18130, 18131, 18132, 18133, 18134, 18135, 18136, 18137, 18138, 18139, 18140, 18141, 18142, 18143, 18144, 18145, 18146, 18147, 18148, 18149, 18150, 18151, 18152, 18153, 18154, 18155, 18156, 18157},
		110: {19101, 19102, 19103, 19104, 19105, 19106, 19107, 19108, 19109, 19110, 19111, 19112, 19113, 19114, 19115, 19116, 19117, 19118, 19119, 19120, 19121, 19122, 19123, 19124, 19125, 19126, 19127, 19128, 19129, 19130, 19131, 19132, 19133, 19134, 19135, 19136, 19137, 19138, 19139, 19140, 19141, 19142, 19143, 19144, 19145, 19146, 19147, 19148, 19149, 19150, 19151, 19152, 19153, 19154, 19155, 19156, 19157},
		201: {20001, 20002, 20003, 20004, 20005, 20006, 20007, 20008, 20009, 20010, 20011, 20012, 20013, 20014, 20015, 20016, 20017, 20018, 20019, 20020, 20021, 20022, 20023, 20024, 20025, 20026, 20027, 20028, 20029, 20030, 20031, 20032, 20033, 20034, 20035, 20036, 20037, 20038, 20039, 20040, 20041, 20042, 20043, 20044, 20045, 20046, 20047, 20048, 20049, 20050, 20051, 20052, 20053, 20054, 20055, 20056, 20057, 20058, 20059},
		202: {21001, 21002, 21003, 21004, 21005, 21006, 21007, 21008, 21009, 21010, 21011, 21012, 21013, 21014, 21015, 21016, 21017, 21018, 21019, 21020, 21021, 21022, 21023, 21024, 21025, 21026, 21027, 21028, 21029, 21030, 21031, 21032, 21033, 21034, 21035, 21036, 21037, 21038, 21039, 21040, 21041, 21042, 21043, 21044, 21045, 21046, 21047, 21048, 21049, 21050, 21051, 21052, 21053, 21054, 21055, 21056, 21057, 21058, 21059},
		203: {22001, 22002, 22003, 22004, 22005, 22006, 22007, 22008, 22009, 22010, 22011, 22012, 22013, 22014, 22015, 22016, 22017, 22018, 22019, 22020, 22021, 22022, 22023, 22024, 22025, 22026, 22027, 22028, 22029, 22030, 22031, 22032, 22033, 22034, 22035, 22036, 22037, 22038, 22039, 22040, 22041, 22042, 22043, 22044, 22045, 22046, 22047, 22048, 22049, 22050, 22051, 22052, 22053, 22054, 22055, 22056, 22057, 22058, 22059},
		204: {23001, 23002, 23003, 23004, 23005, 23006, 23007, 23008, 23009, 23010, 23011, 23012, 23013, 23014, 23015, 23016, 23017, 23018, 23019, 23020, 23021, 23022, 23023, 23024, 23025, 23026, 23027, 23028, 23029, 23030, 23031, 23032, 23033, 23034, 23035, 23036, 23037, 23038, 23039, 23040, 23041, 23042, 23043, 23044, 23045, 23046, 23047, 23048, 23049, 23050, 23051, 23052, 23053, 23054, 23055, 23056, 23057, 23058, 23059},
		205: {24001, 24002, 24003, 24004, 24005, 24006, 24007, 24008, 24009, 24010, 24011, 24012, 24013, 24014, 24015, 24016, 24017, 24018, 24019, 24020, 24021, 24022, 24023, 24024, 24025, 24026, 24027, 24028, 24029, 24030, 24031, 24032, 24033, 24034, 24035, 24036, 24037, 24038, 24039, 24040, 24041, 24042, 24043, 24044, 24045, 24046, 24047, 24048, 24049, 24050, 24051, 24052, 24053, 24054, 24055, 24056, 24057, 24058, 24059},
		206: {25001, 25002, 25003, 25004, 25005, 25006, 25007, 25008, 25009, 25010, 25011, 25012, 25013, 25014, 25015, 25016, 25017, 25018, 25019, 25020, 25021, 25022, 25023, 25024, 25025, 25026, 25027, 25028, 25029, 25030, 25031, 25032, 25033, 25034, 25035, 25036, 25037, 25038, 25039, 25040, 25041, 25042, 25043, 25044, 25045, 25046, 25047, 25048, 25049, 25050, 25051, 25052, 25053, 25054, 25055, 25056, 25057, 25058, 25059},
		207: {26001, 26002, 26003, 26004, 26005, 26006, 26007, 26008, 26009, 26010, 26011, 26012, 26013, 26014, 26015, 26016, 26017, 26018, 26019, 26020, 26021, 26022, 26023, 26024, 26025, 26026, 26027, 26028, 26029, 26030, 26031, 26032, 26033, 26034, 26035, 26036, 26037, 26038, 26039, 26040, 26041, 26042, 26043, 26044, 26045, 26046, 26047, 26048, 26049, 26050, 26051, 26052, 26053, 26054, 26055, 26056, 26057, 26058, 26059},
		208: {27001, 27002, 27003, 27004, 27005, 27006, 27007, 27008, 27009, 27010, 27011, 27012, 27013, 27014, 27015, 27016, 27017, 27018, 27019, 27020, 27021, 27022, 27023, 27024, 27025, 27026, 27027, 27028, 27029, 27030, 27031, 27032, 27033, 27034, 27035, 27036, 27037, 27038, 27039, 27040, 27041, 27042, 27043, 27044, 27045, 27046, 27047, 27048, 27049, 27050, 27051, 27052, 27053, 27054, 27055, 27056, 27057, 27058, 27059},
		209: {28001, 28002, 28003, 28004, 28005, 28006, 28007, 28008, 28009, 28010, 28011, 28012, 28013, 28014, 28015, 28016, 28017, 28018, 28019, 28020, 28021, 28022, 28023, 28024, 28025, 28026, 28027, 28028, 28029, 28030, 28031, 28032, 28033, 28034, 28035, 28036, 28037, 28038, 28039, 28040, 28041, 28042, 28043, 28044, 28045, 28046, 28047, 28048, 28049, 28050, 28051, 28052, 28053, 28054, 28055, 28056, 28057, 28058, 28059},
		210: {29001, 29002, 29003, 29004, 29005, 29006, 29007, 29008, 29009, 29010, 29011, 29012, 29013, 29014, 29015, 29016, 29017, 29018, 29019, 29020, 29021, 29022, 29023, 29024, 29025, 29026, 29027, 29028, 29029, 29030, 29031, 29032, 29033, 29034, 29035, 29036, 29037, 29038, 29039, 29040, 29041, 29042, 29043, 29044, 29045, 29046, 29047, 29048, 29049, 29050, 29051, 29052, 29053, 29054, 29055, 29056, 29057, 29058, 29059},
		301: {30001, 30002, 30003, 30004, 30005, 30006, 30007, 30008, 30009, 30010},
	},
}

func (table *SummonRewardTable) Get(id int32) *SummonRewardCfg {
	return table.data[id]
}

func (table *SummonRewardTable) GetAll() []int32 {
	return summonRewardKeys
}

var summonRewardValues = []*SummonRewardCfg{
	{
		Id:   10101,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    15001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   10102,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    25001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   10103,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    35001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   10104,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    45001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   10105,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    14001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   10106,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    24001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   10107,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    34001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   10108,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    44001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   10109,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    53001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   10110,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    63001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   10111,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    73001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   10112,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    83001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   10113,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    93001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   10114,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    103001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   10115,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   10116,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   10117,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   10118,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10119,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10120,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10121,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10122,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10123,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10124,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10125,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10126,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10127,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   10128,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10129,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10130,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10131,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10132,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10133,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10134,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10135,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10136,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10137,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   10138,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10139,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10140,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   10141,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10142,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10143,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10144,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10145,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10146,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10147,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10148,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10149,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10150,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10151,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10152,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10153,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10154,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10155,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10156,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   10157,
		Type: 101,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11101,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    115001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   11102,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    125001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   11103,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    135001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   11104,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    145001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   11105,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    114001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   11106,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    124001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   11107,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    134001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   11108,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    144001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   11109,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    153001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   11110,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    163001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   11111,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    173001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   11112,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    183001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   11113,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    193001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   11114,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    203001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   11115,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   11116,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   11117,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   11118,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11119,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11120,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11121,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11122,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11123,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11124,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11125,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11126,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11127,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   11128,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11129,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11130,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11131,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11132,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11133,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11134,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11135,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11136,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11137,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   11138,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11139,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11140,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   11141,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11142,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11143,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11144,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11145,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11146,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11147,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11148,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11149,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11150,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11151,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11152,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11153,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11154,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11155,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11156,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   11157,
		Type: 102,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12101,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    215001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   12102,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    225001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   12103,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    235001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   12104,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    245001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   12105,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    214001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   12106,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    224001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   12107,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    234001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   12108,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    244001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   12109,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    253001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   12110,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    263001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   12111,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    273001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   12112,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    283001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   12113,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    293001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   12114,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    303001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   12115,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   12116,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   12117,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   12118,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12119,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12120,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12121,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12122,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12123,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12124,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12125,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12126,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12127,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   12128,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12129,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12130,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12131,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12132,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12133,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12134,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12135,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12136,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12137,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   12138,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12139,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12140,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   12141,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12142,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12143,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12144,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12145,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12146,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12147,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12148,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12149,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12150,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12151,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12152,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12153,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12154,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12155,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12156,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   12157,
		Type: 103,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13101,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    315001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   13102,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    325001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   13103,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    335001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   13104,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    345001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   13105,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    314001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   13106,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    324001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   13107,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    334001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   13108,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    344001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   13109,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    353001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   13110,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    363001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   13111,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    373001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   13112,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    383001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   13113,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    393001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   13114,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    403001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   13115,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   13116,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   13117,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   13118,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13119,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13120,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13121,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13122,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13123,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13124,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13125,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13126,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13127,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   13128,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13129,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13130,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13131,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13132,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13133,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13134,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13135,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13136,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13137,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   13138,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13139,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13140,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   13141,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13142,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13143,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13144,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13145,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13146,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13147,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13148,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13149,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13150,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13151,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13152,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13153,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13154,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13155,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13156,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   13157,
		Type: 104,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14101,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    415001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   14102,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    425001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   14103,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    435001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   14104,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    445001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   14105,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    414001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   14106,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    424001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   14107,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    434001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   14108,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    444001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   14109,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    453001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   14110,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    463001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   14111,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    473001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   14112,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    483001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   14113,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    493001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   14114,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    503001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   14115,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   14116,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   14117,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   14118,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14119,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14120,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14121,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14122,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14123,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14124,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14125,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14126,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14127,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   14128,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14129,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14130,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14131,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14132,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14133,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14134,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14135,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14136,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14137,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   14138,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14139,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14140,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   14141,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14142,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14143,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14144,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14145,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14146,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14147,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14148,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14149,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14150,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14151,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14152,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14153,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14154,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14155,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14156,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   14157,
		Type: 105,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15101,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    515001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   15102,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    525001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   15103,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    535001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   15104,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    545001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   15105,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    514001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   15106,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    524001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   15107,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    534001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   15108,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    544001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   15109,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    553001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   15110,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    563001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   15111,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    573001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   15112,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    583001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   15113,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    593001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   15114,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    603001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   15115,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   15116,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   15117,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   15118,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15119,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15120,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15121,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15122,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15123,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15124,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15125,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15126,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15127,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   15128,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15129,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15130,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15131,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15132,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15133,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15134,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15135,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15136,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15137,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   15138,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15139,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15140,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   15141,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15142,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15143,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15144,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15145,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15146,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15147,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15148,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15149,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15150,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15151,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15152,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15153,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15154,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15155,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15156,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   15157,
		Type: 106,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16101,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    615001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   16102,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    625001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   16103,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    635001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   16104,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    645001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   16105,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    614001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   16106,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    624001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   16107,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    634001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   16108,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    644001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   16109,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    653001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   16110,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    663001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   16111,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    673001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   16112,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    683001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   16113,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    693001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   16114,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    703001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   16115,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   16116,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   16117,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   16118,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16119,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16120,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16121,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16122,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16123,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16124,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16125,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16126,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16127,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   16128,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16129,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16130,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16131,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16132,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16133,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16134,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16135,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16136,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16137,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   16138,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16139,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16140,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   16141,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16142,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16143,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16144,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16145,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16146,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16147,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16148,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16149,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16150,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16151,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16152,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16153,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16154,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16155,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16156,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   16157,
		Type: 107,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17101,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    715001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   17102,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    725001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   17103,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    735001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   17104,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    745001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   17105,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    714001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   17106,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    724001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   17107,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    734001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   17108,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    744001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   17109,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    753001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   17110,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    763001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   17111,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    773001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   17112,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    783001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   17113,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    793001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   17114,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    803001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   17115,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   17116,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   17117,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   17118,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17119,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17120,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17121,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17122,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17123,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17124,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17125,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17126,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17127,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   17128,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17129,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17130,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17131,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17132,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17133,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17134,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17135,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17136,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17137,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   17138,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17139,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17140,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   17141,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17142,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17143,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17144,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17145,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17146,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17147,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17148,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17149,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17150,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17151,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17152,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17153,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17154,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17155,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17156,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   17157,
		Type: 108,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18101,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    815001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   18102,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    825001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   18103,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    835001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   18104,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    845001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   18105,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    814001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   18106,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    824001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   18107,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    834001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   18108,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    844001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   18109,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    853001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   18110,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    863001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   18111,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    873001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   18112,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    883001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   18113,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    893001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   18114,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    903001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   18115,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   18116,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   18117,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   18118,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18119,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18120,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18121,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18122,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18123,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18124,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18125,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18126,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18127,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   18128,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18129,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18130,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18131,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18132,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18133,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18134,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18135,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18136,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18137,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   18138,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18139,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18140,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   18141,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18142,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18143,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18144,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18145,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18146,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18147,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18148,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18149,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18150,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18151,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18152,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18153,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18154,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18155,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18156,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   18157,
		Type: 109,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19101,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    915001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   19102,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    925001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   19103,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    935001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   19104,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    945001,
				Count: 1,
			},
		},
		Weight: 25,
	},
	{
		Id:   19105,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    914001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   19106,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    924001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   19107,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    934001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   19108,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    944001,
				Count: 1,
			},
		},
		Weight: 500,
	},
	{
		Id:   19109,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    953001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   19110,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    963001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   19111,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    973001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   19112,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    983001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   19113,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    993001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   19114,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    1003001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   19115,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 30,
			},
		},
		Weight: 25000,
	},
	{
		Id:   19116,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 50,
			},
		},
		Weight: 15000,
	},
	{
		Id:   19117,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 70,
			},
		},
		Weight: 10000,
	},
	{
		Id:   19118,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19119,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19120,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19121,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19122,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19123,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19124,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19125,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19126,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19127,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 1490,
	},
	{
		Id:   19128,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19129,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19130,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19131,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19132,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19133,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19134,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19135,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19136,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19137,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 100,
	},
	{
		Id:   19138,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19139,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19140,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 1500,
	},
	{
		Id:   19141,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19142,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19143,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19144,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19145,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19146,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19147,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19148,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19149,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19150,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19151,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19152,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19153,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19154,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19155,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19156,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   19157,
		Type: 110,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight:    1500,
		Guarantee: 1,
	},
	{
		Id:   20001,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    15001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20002,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    25001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20003,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    35001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20004,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    45001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20005,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    54001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   20006,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    64001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   20007,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    74001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   20008,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    84001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   20009,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    94001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   20010,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    104001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   20011,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    55001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20012,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    65001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20013,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    75001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20014,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    85001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20015,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    95001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20016,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    105001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   20017,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   20018,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   20019,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   20020,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20021,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20022,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20023,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20024,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20025,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20026,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20027,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20028,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20029,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   20030,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20031,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20032,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20033,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20034,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20035,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   20036,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20037,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20038,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   20039,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   20040,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20041,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20042,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20043,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20044,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20045,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20046,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20047,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20048,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20049,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20050,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20051,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20052,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20053,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20054,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20055,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20056,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20057,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20058,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   20059,
		Type: 201,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21001,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    115001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21002,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    125001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21003,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    135001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21004,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    145001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21005,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    154001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   21006,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    164001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   21007,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    174001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   21008,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    184001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   21009,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    194001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   21010,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    204001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   21011,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    155001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21012,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    165001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21013,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    175001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21014,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    185001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21015,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    195001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21016,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    205001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   21017,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   21018,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   21019,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   21020,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21021,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21022,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21023,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21024,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21025,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21026,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21027,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21028,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21029,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   21030,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21031,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21032,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21033,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21034,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21035,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   21036,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21037,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21038,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   21039,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   21040,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21041,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21042,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21043,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21044,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21045,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21046,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21047,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21048,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21049,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21050,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21051,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21052,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21053,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21054,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21055,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21056,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21057,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21058,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   21059,
		Type: 202,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22001,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    215001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22002,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    225001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22003,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    235001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22004,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    245001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22005,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    254001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   22006,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    264001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   22007,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    274001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   22008,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    284001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   22009,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    294001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   22010,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    304001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   22011,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    255001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22012,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    265001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22013,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    275001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22014,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    285001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22015,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    295001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22016,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    305001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   22017,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   22018,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   22019,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   22020,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22021,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22022,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22023,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22024,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22025,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22026,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22027,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22028,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22029,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   22030,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22031,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22032,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22033,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22034,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22035,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   22036,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22037,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22038,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   22039,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   22040,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22041,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22042,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22043,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22044,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22045,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22046,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22047,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22048,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22049,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22050,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22051,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22052,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22053,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22054,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22055,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22056,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22057,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22058,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   22059,
		Type: 203,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23001,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    315001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23002,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    325001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23003,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    335001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23004,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    345001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23005,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    354001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   23006,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    364001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   23007,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    374001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   23008,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    384001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   23009,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    394001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   23010,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    404001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   23011,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    355001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23012,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    365001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23013,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    375001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23014,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    385001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23015,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    395001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23016,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    405001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   23017,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   23018,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   23019,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   23020,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23021,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23022,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23023,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23024,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23025,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23026,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23027,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23028,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23029,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   23030,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23031,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23032,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23033,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23034,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23035,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   23036,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23037,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23038,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   23039,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   23040,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23041,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23042,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23043,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23044,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23045,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23046,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23047,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23048,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23049,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23050,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23051,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23052,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23053,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23054,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23055,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23056,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23057,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23058,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   23059,
		Type: 204,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24001,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    415001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24002,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    425001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24003,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    435001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24004,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    445001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24005,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    454001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   24006,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    464001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   24007,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    474001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   24008,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    484001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   24009,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    494001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   24010,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    504001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   24011,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    455001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24012,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    465001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24013,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    475001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24014,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    485001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24015,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    495001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24016,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    505001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   24017,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   24018,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   24019,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   24020,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24021,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24022,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24023,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24024,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24025,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24026,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24027,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24028,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24029,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   24030,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24031,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24032,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24033,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24034,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24035,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   24036,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24037,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24038,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   24039,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   24040,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24041,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24042,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24043,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24044,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24045,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24046,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24047,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24048,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24049,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24050,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24051,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24052,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24053,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24054,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24055,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24056,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24057,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24058,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   24059,
		Type: 205,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25001,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    515001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25002,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    525001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25003,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    535001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25004,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    545001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25005,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    554001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   25006,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    564001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   25007,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    574001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   25008,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    584001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   25009,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    594001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   25010,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    604001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   25011,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    555001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25012,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    565001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25013,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    575001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25014,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    585001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25015,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    595001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25016,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    605001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   25017,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   25018,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   25019,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   25020,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25021,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25022,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25023,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25024,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25025,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25026,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25027,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25028,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25029,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   25030,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25031,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25032,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25033,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25034,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25035,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   25036,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25037,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25038,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   25039,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   25040,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25041,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25042,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25043,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25044,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25045,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25046,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25047,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25048,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25049,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25050,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25051,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25052,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25053,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25054,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25055,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25056,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25057,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25058,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   25059,
		Type: 206,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26001,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    615001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26002,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    625001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26003,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    635001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26004,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    645001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26005,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    654001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   26006,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    664001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   26007,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    674001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   26008,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    684001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   26009,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    694001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   26010,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    704001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   26011,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    655001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26012,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    665001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26013,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    675001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26014,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    685001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26015,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    695001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26016,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    705001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   26017,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   26018,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   26019,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   26020,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26021,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26022,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26023,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26024,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26025,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26026,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26027,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26028,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26029,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   26030,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26031,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26032,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26033,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26034,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26035,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   26036,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26037,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26038,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   26039,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   26040,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26041,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26042,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26043,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26044,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26045,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26046,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26047,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26048,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26049,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26050,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26051,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26052,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26053,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26054,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26055,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26056,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26057,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26058,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   26059,
		Type: 207,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27001,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    715001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27002,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    725001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27003,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    735001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27004,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    745001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27005,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    754001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   27006,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    764001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   27007,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    774001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   27008,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    784001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   27009,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    794001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   27010,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    804001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   27011,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    755001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27012,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    765001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27013,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    775001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27014,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    785001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27015,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    795001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27016,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    805001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   27017,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   27018,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   27019,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   27020,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27021,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27022,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27023,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27024,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27025,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27026,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27027,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27028,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27029,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   27030,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27031,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27032,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27033,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27034,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27035,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   27036,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27037,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27038,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   27039,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   27040,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27041,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27042,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27043,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27044,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27045,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27046,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27047,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27048,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27049,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27050,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27051,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27052,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27053,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27054,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27055,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27056,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27057,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27058,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   27059,
		Type: 208,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28001,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    815001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28002,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    825001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28003,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    835001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28004,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    845001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28005,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    854001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   28006,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    864001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   28007,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    874001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   28008,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    884001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   28009,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    894001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   28010,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    904001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   28011,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    855001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28012,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    865001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28013,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    875001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28014,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    885001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28015,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    895001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28016,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    905001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   28017,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   28018,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   28019,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   28020,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28021,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28022,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28023,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28024,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28025,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28026,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28027,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28028,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28029,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   28030,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28031,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28032,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28033,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28034,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28035,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   28036,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28037,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28038,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   28039,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   28040,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28041,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28042,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28043,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28044,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28045,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28046,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28047,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28048,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28049,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28050,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28051,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28052,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28053,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28054,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28055,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28056,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28057,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28058,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   28059,
		Type: 209,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29001,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    915001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29002,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    925001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29003,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    935001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29004,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    945001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29005,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    954001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   29006,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    964001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   29007,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    974001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   29008,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    984001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   29009,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    994001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   29010,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    1004001,
				Count: 1,
			},
		},
		Weight: 333,
	},
	{
		Id:   29011,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    955001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29012,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    965001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29013,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    975001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29014,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    985001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29015,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    995001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29016,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  2,
				Id:    1005001,
				Count: 1,
			},
		},
		Weight: 200,
	},
	{
		Id:   29017,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 90,
			},
		},
		Weight: 4000,
	},
	{
		Id:   29018,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 150,
			},
		},
		Weight: 4000,
	},
	{
		Id:   29019,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  1,
				Id:    2,
				Count: 210,
			},
		},
		Weight: 5000,
	},
	{
		Id:   29020,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    1,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29021,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    2,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29022,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    3,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29023,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    4,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29024,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    5,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29025,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    6,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29026,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    7,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29027,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    8,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29028,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    9,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29029,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  4,
				Id:    10,
				Count: 1,
			},
		},
		Weight: 3000,
	},
	{
		Id:   29030,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29031,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29032,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29033,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29034,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29035,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   29036,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29037,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29038,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    1000,
		Guarantee: 1,
	},
	{
		Id:   29039,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight: 1000,
	},
	{
		Id:   29040,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10721,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29041,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10921,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29042,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11122,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29043,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11322,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29044,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11421,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29045,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29046,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11612,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29047,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11712,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29048,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11812,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29049,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29050,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12011,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29051,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12112,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29052,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12212,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29053,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12312,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29054,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12412,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29055,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12511,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29056,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12611,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29057,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12711,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29058,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12811,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   29059,
		Type: 210,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    12911,
				Count: 1,
			},
		},
		Weight: 2250,
	},
	{
		Id:   30001,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10021,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30002,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10121,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30003,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10222,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30004,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10321,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30005,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10421,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30006,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10521,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30007,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10622,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30008,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    10822,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30009,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11021,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
	{
		Id:   30010,
		Type: 301,
		Reward: []*Reward{
			{
				Type:  3,
				Id:    11222,
				Count: 1,
			},
		},
		Weight:    500,
		Guarantee: 1,
	},
}

var summonRewardKeys = []int32{
	10101, 10102, 10103, 10104, 10105, 10106, 10107, 10108, 10109, 10110,
	10111, 10112, 10113, 10114, 10115, 10116, 10117, 10118, 10119, 10120,
	10121, 10122, 10123, 10124, 10125, 10126, 10127, 10128, 10129, 10130,
	10131, 10132, 10133, 10134, 10135, 10136, 10137, 10138, 10139, 10140,
	10141, 10142, 10143, 10144, 10145, 10146, 10147, 10148, 10149, 10150,
	10151, 10152, 10153, 10154, 10155, 10156, 10157, 11101, 11102, 11103,
	11104, 11105, 11106, 11107, 11108, 11109, 11110, 11111, 11112, 11113,
	11114, 11115, 11116, 11117, 11118, 11119, 11120, 11121, 11122, 11123,
	11124, 11125, 11126, 11127, 11128, 11129, 11130, 11131, 11132, 11133,
	11134, 11135, 11136, 11137, 11138, 11139, 11140, 11141, 11142, 11143,
	11144, 11145, 11146, 11147, 11148, 11149, 11150, 11151, 11152, 11153,
	11154, 11155, 11156, 11157, 12101, 12102, 12103, 12104, 12105, 12106,
	12107, 12108, 12109, 12110, 12111, 12112, 12113, 12114, 12115, 12116,
	12117, 12118, 12119, 12120, 12121, 12122, 12123, 12124, 12125, 12126,
	12127, 12128, 12129, 12130, 12131, 12132, 12133, 12134, 12135, 12136,
	12137, 12138, 12139, 12140, 12141, 12142, 12143, 12144, 12145, 12146,
	12147, 12148, 12149, 12150, 12151, 12152, 12153, 12154, 12155, 12156,
	12157, 13101, 13102, 13103, 13104, 13105, 13106, 13107, 13108, 13109,
	13110, 13111, 13112, 13113, 13114, 13115, 13116, 13117, 13118, 13119,
	13120, 13121, 13122, 13123, 13124, 13125, 13126, 13127, 13128, 13129,
	13130, 13131, 13132, 13133, 13134, 13135, 13136, 13137, 13138, 13139,
	13140, 13141, 13142, 13143, 13144, 13145, 13146, 13147, 13148, 13149,
	13150, 13151, 13152, 13153, 13154, 13155, 13156, 13157, 14101, 14102,
	14103, 14104, 14105, 14106, 14107, 14108, 14109, 14110, 14111, 14112,
	14113, 14114, 14115, 14116, 14117, 14118, 14119, 14120, 14121, 14122,
	14123, 14124, 14125, 14126, 14127, 14128, 14129, 14130, 14131, 14132,
	14133, 14134, 14135, 14136, 14137, 14138, 14139, 14140, 14141, 14142,
	14143, 14144, 14145, 14146, 14147, 14148, 14149, 14150, 14151, 14152,
	14153, 14154, 14155, 14156, 14157, 15101, 15102, 15103, 15104, 15105,
	15106, 15107, 15108, 15109, 15110, 15111, 15112, 15113, 15114, 15115,
	15116, 15117, 15118, 15119, 15120, 15121, 15122, 15123, 15124, 15125,
	15126, 15127, 15128, 15129, 15130, 15131, 15132, 15133, 15134, 15135,
	15136, 15137, 15138, 15139, 15140, 15141, 15142, 15143, 15144, 15145,
	15146, 15147, 15148, 15149, 15150, 15151, 15152, 15153, 15154, 15155,
	15156, 15157, 16101, 16102, 16103, 16104, 16105, 16106, 16107, 16108,
	16109, 16110, 16111, 16112, 16113, 16114, 16115, 16116, 16117, 16118,
	16119, 16120, 16121, 16122, 16123, 16124, 16125, 16126, 16127, 16128,
	16129, 16130, 16131, 16132, 16133, 16134, 16135, 16136, 16137, 16138,
	16139, 16140, 16141, 16142, 16143, 16144, 16145, 16146, 16147, 16148,
	16149, 16150, 16151, 16152, 16153, 16154, 16155, 16156, 16157, 17101,
	17102, 17103, 17104, 17105, 17106, 17107, 17108, 17109, 17110, 17111,
	17112, 17113, 17114, 17115, 17116, 17117, 17118, 17119, 17120, 17121,
	17122, 17123, 17124, 17125, 17126, 17127, 17128, 17129, 17130, 17131,
	17132, 17133, 17134, 17135, 17136, 17137, 17138, 17139, 17140, 17141,
	17142, 17143, 17144, 17145, 17146, 17147, 17148, 17149, 17150, 17151,
	17152, 17153, 17154, 17155, 17156, 17157, 18101, 18102, 18103, 18104,
	18105, 18106, 18107, 18108, 18109, 18110, 18111, 18112, 18113, 18114,
	18115, 18116, 18117, 18118, 18119, 18120, 18121, 18122, 18123, 18124,
	18125, 18126, 18127, 18128, 18129, 18130, 18131, 18132, 18133, 18134,
	18135, 18136, 18137, 18138, 18139, 18140, 18141, 18142, 18143, 18144,
	18145, 18146, 18147, 18148, 18149, 18150, 18151, 18152, 18153, 18154,
	18155, 18156, 18157, 19101, 19102, 19103, 19104, 19105, 19106, 19107,
	19108, 19109, 19110, 19111, 19112, 19113, 19114, 19115, 19116, 19117,
	19118, 19119, 19120, 19121, 19122, 19123, 19124, 19125, 19126, 19127,
	19128, 19129, 19130, 19131, 19132, 19133, 19134, 19135, 19136, 19137,
	19138, 19139, 19140, 19141, 19142, 19143, 19144, 19145, 19146, 19147,
	19148, 19149, 19150, 19151, 19152, 19153, 19154, 19155, 19156, 19157,
	20001, 20002, 20003, 20004, 20005, 20006, 20007, 20008, 20009, 20010,
	20011, 20012, 20013, 20014, 20015, 20016, 20017, 20018, 20019, 20020,
	20021, 20022, 20023, 20024, 20025, 20026, 20027, 20028, 20029, 20030,
	20031, 20032, 20033, 20034, 20035, 20036, 20037, 20038, 20039, 20040,
	20041, 20042, 20043, 20044, 20045, 20046, 20047, 20048, 20049, 20050,
	20051, 20052, 20053, 20054, 20055, 20056, 20057, 20058, 20059, 21001,
	21002, 21003, 21004, 21005, 21006, 21007, 21008, 21009, 21010, 21011,
	21012, 21013, 21014, 21015, 21016, 21017, 21018, 21019, 21020, 21021,
	21022, 21023, 21024, 21025, 21026, 21027, 21028, 21029, 21030, 21031,
	21032, 21033, 21034, 21035, 21036, 21037, 21038, 21039, 21040, 21041,
	21042, 21043, 21044, 21045, 21046, 21047, 21048, 21049, 21050, 21051,
	21052, 21053, 21054, 21055, 21056, 21057, 21058, 21059, 22001, 22002,
	22003, 22004, 22005, 22006, 22007, 22008, 22009, 22010, 22011, 22012,
	22013, 22014, 22015, 22016, 22017, 22018, 22019, 22020, 22021, 22022,
	22023, 22024, 22025, 22026, 22027, 22028, 22029, 22030, 22031, 22032,
	22033, 22034, 22035, 22036, 22037, 22038, 22039, 22040, 22041, 22042,
	22043, 22044, 22045, 22046, 22047, 22048, 22049, 22050, 22051, 22052,
	22053, 22054, 22055, 22056, 22057, 22058, 22059, 23001, 23002, 23003,
	23004, 23005, 23006, 23007, 23008, 23009, 23010, 23011, 23012, 23013,
	23014, 23015, 23016, 23017, 23018, 23019, 23020, 23021, 23022, 23023,
	23024, 23025, 23026, 23027, 23028, 23029, 23030, 23031, 23032, 23033,
	23034, 23035, 23036, 23037, 23038, 23039, 23040, 23041, 23042, 23043,
	23044, 23045, 23046, 23047, 23048, 23049, 23050, 23051, 23052, 23053,
	23054, 23055, 23056, 23057, 23058, 23059, 24001, 24002, 24003, 24004,
	24005, 24006, 24007, 24008, 24009, 24010, 24011, 24012, 24013, 24014,
	24015, 24016, 24017, 24018, 24019, 24020, 24021, 24022, 24023, 24024,
	24025, 24026, 24027, 24028, 24029, 24030, 24031, 24032, 24033, 24034,
	24035, 24036, 24037, 24038, 24039, 24040, 24041, 24042, 24043, 24044,
	24045, 24046, 24047, 24048, 24049, 24050, 24051, 24052, 24053, 24054,
	24055, 24056, 24057, 24058, 24059, 25001, 25002, 25003, 25004, 25005,
	25006, 25007, 25008, 25009, 25010, 25011, 25012, 25013, 25014, 25015,
	25016, 25017, 25018, 25019, 25020, 25021, 25022, 25023, 25024, 25025,
	25026, 25027, 25028, 25029, 25030, 25031, 25032, 25033, 25034, 25035,
	25036, 25037, 25038, 25039, 25040, 25041, 25042, 25043, 25044, 25045,
	25046, 25047, 25048, 25049, 25050, 25051, 25052, 25053, 25054, 25055,
	25056, 25057, 25058, 25059, 26001, 26002, 26003, 26004, 26005, 26006,
	26007, 26008, 26009, 26010, 26011, 26012, 26013, 26014, 26015, 26016,
	26017, 26018, 26019, 26020, 26021, 26022, 26023, 26024, 26025, 26026,
	26027, 26028, 26029, 26030, 26031, 26032, 26033, 26034, 26035, 26036,
	26037, 26038, 26039, 26040, 26041, 26042, 26043, 26044, 26045, 26046,
	26047, 26048, 26049, 26050, 26051, 26052, 26053, 26054, 26055, 26056,
	26057, 26058, 26059, 27001, 27002, 27003, 27004, 27005, 27006, 27007,
	27008, 27009, 27010, 27011, 27012, 27013, 27014, 27015, 27016, 27017,
	27018, 27019, 27020, 27021, 27022, 27023, 27024, 27025, 27026, 27027,
	27028, 27029, 27030, 27031, 27032, 27033, 27034, 27035, 27036, 27037,
	27038, 27039, 27040, 27041, 27042, 27043, 27044, 27045, 27046, 27047,
	27048, 27049, 27050, 27051, 27052, 27053, 27054, 27055, 27056, 27057,
	27058, 27059, 28001, 28002, 28003, 28004, 28005, 28006, 28007, 28008,
	28009, 28010, 28011, 28012, 28013, 28014, 28015, 28016, 28017, 28018,
	28019, 28020, 28021, 28022, 28023, 28024, 28025, 28026, 28027, 28028,
	28029, 28030, 28031, 28032, 28033, 28034, 28035, 28036, 28037, 28038,
	28039, 28040, 28041, 28042, 28043, 28044, 28045, 28046, 28047, 28048,
	28049, 28050, 28051, 28052, 28053, 28054, 28055, 28056, 28057, 28058,
	28059, 29001, 29002, 29003, 29004, 29005, 29006, 29007, 29008, 29009,
	29010, 29011, 29012, 29013, 29014, 29015, 29016, 29017, 29018, 29019,
	29020, 29021, 29022, 29023, 29024, 29025, 29026, 29027, 29028, 29029,
	29030, 29031, 29032, 29033, 29034, 29035, 29036, 29037, 29038, 29039,
	29040, 29041, 29042, 29043, 29044, 29045, 29046, 29047, 29048, 29049,
	29050, 29051, 29052, 29053, 29054, 29055, 29056, 29057, 29058, 29059,
	30001, 30002, 30003, 30004, 30005, 30006, 30007, 30008, 30009, 30010,
}

func init() {
	SummonRewardData.data = make(map[int32]*SummonRewardCfg)
	for i := 0; i < len(summonRewardKeys); i++ {
		SummonRewardData.data[summonRewardKeys[i]] = summonRewardValues[i]
	}
}

func (table *SummonRewardTable) GetByType(Type int32) (res []*SummonRewardCfg) {
	for _, i := range table.summonRewardTypeIndexMap[Type] {
		res = append(res, table.Get(i))
	}
	return
}
