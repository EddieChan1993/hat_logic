package rpc

import (
	"encoding/json"
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gdb/pbgo"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/rpc/rcst"
)

type MailContent struct {
	RoleId  int64            //角色id
	MailId  int32            //邮件id
	Rewards []*config.Reward //奖品id
	Args    string           //额外参数 json
}

type MailRewards struct {
	Type  int32 `json:"type"`
	Id    int32 `json:"id"`
	Count int64 `json:"num"`
}

//SendMail 发送邮件
func SendMail(mails ...*MailContent) error {
	if len(mails) == 0 {
		return nil
	}
	mailReq := make([]*pbgo.SendMailReq_Mail, len(mails))
	for i, obj := range mails {
		if obj == nil {
			continue
		}
		affix, err := cfg2mailReward(obj.Rewards)
		if err != nil {
			return err
		}
		mailReq[i] = &pbgo.SendMailReq_Mail{
			Roleid: int32(obj.RoleId),
			CfgId:  obj.MailId,
			Args:   obj.Args,
			Affix:  affix,
		}
	}

	resp, err := pbgo.Mail.SendMail(*rcst.RpcWhoAmIdGDB(), &pbgo.SendMailReq{Mails: mailReq})
	if err != nil {
		return err
	}
	if resp.ErrNo != 0 {
		return fmt.Errorf(resp.ErrMsg)
	}
	return nil
}

func cfg2mailReward(rewards []*config.Reward) (res []byte, err error) {
	mails := make([]*MailRewards, len(rewards))
	for i, reward := range rewards {
		mails[i] = &MailRewards{
			Type:  reward.Type,
			Id:    reward.Id,
			Count: reward.Count,
		}
	}
	bt, err := json.Marshal(mails)
	return bt, err
}

func Mail2cfgReward(Affix []byte) (res []*config.Reward, err error) {
	mails := make([]*MailRewards, 0)
	err = json.Unmarshal(Affix, &mails)
	if err != nil {
		return nil, err
	}
	res = make([]*config.Reward, len(mails))
	for i := range mails {
		tmp := (*config.Reward)(mails[i])
		res[i] = tmp
	}
	return res, err
}

//RoleMails 获取角色所有邮件
func RoleMails(roleId int64) ([]*pbgo.MailInfo, error) {
	resp, err := pbgo.Mail.RoleMails(*rcst.RpcWhoAmIdGDB(), &pbgo.RoleMailsReq{
		RoleId: int32(roleId),
	})
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf(resp.ErrMsg)
	}
	return resp.Mails, nil
}

//SyncMail 获取角色所有邮件
func SyncMail(roleId int64, mailTotal int32) ([]*pbgo.MailInfo, []*pbgo.MailAttach, error) {
	resp, err := pbgo.Mail.Sync(*rcst.RpcWhoAmIdGDB(), &pbgo.SyncMailReq{
		RoleId: int32(roleId),
		Opt: &pbgo.SyncMailReq_Option{
			MailTotal: mailTotal,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	if resp.ErrNo != 0 {
		return nil, nil, fmt.Errorf(resp.ErrMsg)
	}
	return resp.Mails, resp.Attach, nil
}

//ReadMail 阅读邮件
func ReadMail(roleId int64, mailId int32) error {
	resp, err := pbgo.Mail.Read(*rcst.RpcWhoAmIdGDB(), &pbgo.ReadMailReq{
		RoleId: int32(roleId),
		Mails:  []int32{mailId},
	})
	if err != nil {
		return err
	}
	if resp.ErrNo != 0 {
		return fmt.Errorf(resp.ErrMsg)
	}
	return nil
}

//DeleteMail 删除邮件
func DeleteMail(roleId int64, mailId []int32) error {
	resp, err := pbgo.Mail.DeleteMail(*rcst.RpcWhoAmIdGDB(), &pbgo.DeleteMailReq{
		RoleId: int32(roleId),
		Mails:  mailId,
	})
	if err != nil {
		return err
	}
	if resp.ErrNo != 0 {
		return fmt.Errorf(resp.ErrMsg)
	}
	return nil
}

//MailReward 领取邮件奖励
func MailReward(roleId int64, mailId []int32) (rsp []*pbgo.MailAttach, err error) {
	resp, err := pbgo.Mail.Reward(*rcst.RpcWhoAmIdGDB(), &pbgo.RewardMailReq{
		RoleId: int32(roleId),
		Mails:  mailId,
	})
	if err != nil {
		return nil, err
	}
	if resp.ErrNo != 0 {
		return nil, fmt.Errorf(resp.ErrMsg)
	}
	return resp.Attach, nil
}
