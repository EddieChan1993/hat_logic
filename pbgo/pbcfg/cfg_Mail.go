package config

type MailCfg struct {
	Id      int32
	Title   string
	Content string
	From    string
}

type MailTable struct {
	data map[int32]*MailCfg
}

var MailData = &MailTable{
	data: map[int32]*MailCfg{},
}

func (table *MailTable) Get(id int32) *MailCfg {
	return table.data[id]
}

func (table *MailTable) GetAll() []int32 {
	return mailKeys
}

var mailValues = []*MailCfg{
	{
		Id:      1,
		Title:   "福利派送",
		Content: "亲爱的冒险家，欢迎来到<color=#f12b2b>Click Deity</color>！感谢与您相遇，请收下我们的礼物：",
		From:    "GM",
	},
	{
		Id:      2,
		Title:   "更新补偿",
		Content: "亲爱的冒险家，本次更新已完成，希望能给您带来更好的游戏体验，请收下我们的礼物：",
		From:    "GM",
	},
	{
		Id:      3,
		Title:   "官方社区",
		Content: "亲爱的冒险家，关注我们的facebook官方账号可以第一时间获取活动更新哦！Facebook官方账号链接：<u><color=#2b8bf1><link>ww.baidu.com</link></color></u>",
		From:    "GM",
	},
	{
		Id:      4,
		Title:   "充值补单",
		Content: "亲爱的冒险家，您的充值未到账问题已经解决，附件是补发给您的物品，给您带来的不便敬请谅解",
		From:    "GM",
	},
	{
		Id:      5,
		Title:   "排行榜个人奖励",
		Content: "亲爱的冒险家，您在昨日的世界首领排行中，获得了以下个人奖励：",
		From:    "GM",
	},
}

var mailKeys = []int32{
	1, 2, 3, 4, 5,
}

func init() {
	MailData.data = make(map[int32]*MailCfg)
	for i := 0; i < len(mailKeys); i++ {
		MailData.data[mailKeys[i]] = mailValues[i]
	}
}
