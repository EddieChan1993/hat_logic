package api

import (
	"git.dhgames.cn/svr_comm/gclient/gchat/pbgo"
	"hat_logic/rpc"
	"strconv"
)

const (
	tagWorld1 = "world-1"
)

type GChat struct {
}

func (G *GChat) LoginChat(req *pbgo.LoginChatReq) (*pbgo.LoginChatRsp, error) {
	res := &pbgo.LoginChatRsp{}
	roleDigest, err := rpc.RoleDigest(req.RoleID)
	if err != nil {
		return res, err
	}
	digest := roleDigest.Digest
	res.Digest = &pbgo.RoleDigest{
		Id:    int64(digest.RoleId),
		Name:  digest.Name,
		Level: int64(digest.Lv),
		Icon:  strconv.Itoa(int(digest.Logo)),
		Job:   0,
		Frame: strconv.Itoa(int(digest.HeadFrame)),
		Title: strconv.Itoa(int(digest.HeadTitle)),
	}
	res.ChannelInfo = []*pbgo.ChannelInfo{
		{
			ChannelType: pbgo.ChannelType_World,
			Tag:         tagWorld1,
			Name:        "世界服",
			Icon:        "",
		},
	}
	return res, nil
}

func (G *GChat) CanPrivateTalk(req *pbgo.CanPrivateTalkReq) (*pbgo.CanPrivateTalkRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GChat) ChannelList(req *pbgo.ChannelListReq) (*pbgo.ChannelListRsp, error) {
	//TODO implement me
	panic("implement me")
}
