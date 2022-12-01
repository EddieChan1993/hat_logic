package mod

/**
邮件
*/
import (
	"git.dhgames.cn/svr_comm/gclient/gdb/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
)

type Mail struct {
	pbreq.TModApi13
}

func (m *Mail) ReqMailRead(req *pbreq.ReqMailRead) *pbreq.RspMailRead {
	res := &pbreq.RspMailRead{
		MailId: req.MailId,
	}
	err := rpc.ReadMail(m.GetRoleId(), req.MailId)
	if err != nil {
		res.Status = 1
		return res
	}
	return res
}

func (m *Mail) ReqMailReward(req *pbreq.ReqMailReward) *pbreq.RspMailReward {
	res := &pbreq.RspMailReward{
		MailIds: req.MailIds,
	}
	mailAttach, err := rpc.MailReward(m.GetRoleId(), req.MailIds)
	if err != nil {
		res.Status = 1
		return res
	}
	rewardsCfgList := make([]*config.Reward, 0)
	for _, attach := range mailAttach {
		rewardsCfg, err := rpc.Mail2cfgReward(attach.Affix)
		if err != nil {
			res.Status = 2
			klog.Errorf("%s mailId %d mailCfg %d err %v", m.Log(), attach.MailId, attach.CfgId, err)
			return nil
		}
		rewardsCfgList = append(rewardsCfgList, rewardsCfg...)
	}
	rewards, err := GoodsMod(m.GameCtx).Rewards(rewardsCfgList)
	if err != nil {
		res.Status = 3
		klog.Error(m.Log(), err)
		return nil
	}
	res.Rewards = rewards
	return res
}

func (m *Mail) ReqMailDel(req *pbreq.ReqMailDel) *pbreq.RspMailDel {
	res := &pbreq.RspMailDel{
		MailIds: req.MailIds,
	}
	err := rpc.DeleteMail(m.GetRoleId(), req.MailIds)
	if err != nil {
		res.Status = 1
		return res
	}
	klog.Infof("%s ReqMailDel MailIds %v", m.Log(), req.MailIds)
	return res
}

//mailList 获取所有邮件
func (m *Mail) mailList() []*pbreq.MailMgr {
	rsp, err := rpc.RoleMails(m.GetRoleId())
	if err != nil {
		klog.Warn(m.Log(), err)
		return nil
	}
	return m.rpc2reqMail(rsp)
}

//syncLimitMails 自动删除多余邮件，并领取奖励
func (m *Mail) syncLimitMails() (mailList []*pbreq.MailMgr, rewards *pbreq.Rewards) {
	mails, attach, err := rpc.SyncMail(m.GetRoleId(), mcst.MaxMailLimit)
	if err != nil {
		klog.Warn(m.Log(), err)
		return nil, nil
	}
	mailList = m.rpc2reqMail(mails)
	//方法删除邮件的奖励
	rewards, err = m.drawMailRewards(attach)
	if err != nil {
		klog.Error(m.Log(), err)
		return nil, nil
	}
	return mailList, rewards
}

func (m *Mail) drawMailRewards(mailAttach []*pbgo.MailAttach) (rewards *pbreq.Rewards, err error) {
	if len(mailAttach) == 0 {
		return nil, nil
	}
	rewardsCfgList := make([]*config.Reward, 0)
	for _, attach := range mailAttach {
		rewardsCfg, err := rpc.Mail2cfgReward(attach.Affix)
		if err != nil {
			klog.Error(m.Log(), err)
			return nil, err
		}
		rewardsCfgList = append(rewardsCfgList, rewardsCfg...)
	}
	rewards, err = GoodsMod(m.GameCtx).Rewards(rewardsCfgList)
	if err != nil {
		klog.Error(m.Log(), err)
		return nil, err
	}
	return rewards, err
}

func (m *Mail) rpc2reqMail(mails []*pbgo.MailInfo) []*pbreq.MailMgr {
	mailList := make([]*pbreq.MailMgr, len(mails))
	for i, mail := range mails {
		mailList[i] = &pbreq.MailMgr{
			XId:             mail.MailId,
			Type:            mail.Type,
			CfgId:           mail.CfgId,
			Args:            mail.Args,
			Affix:           string(mail.Affix),
			ReadFlag:        mail.ReadFlag,
			RewardFlag:      mail.RewardFlag,
			FavoriteFlag:    mail.FavoriteFlag,
			SendTime:        mail.SendTime,
			ExpireTime:      mail.ExpireTime,
			AffixExpireTime: mail.AffixExpireTime,
			ReadTime:        mail.ReadTime,
			AffixTime:       mail.AffixTime,
			Title:           mail.Title,
			Content:         mail.Content,
		}
	}
	return mailList
}

func (m *Mail) SyncBefore() {
	m.syncLimitMails()
}

func (m *Mail) RspSync(rsp *pbreq.RspSync) {
	rsp.MailMgr = m.mailList()
}

func NewMail() *Mail {
	res := &Mail{}
	res.InitModApi(res)
	return res
}
