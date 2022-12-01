package rpc

import (
	"fmt"
	"git.dhgames.cn/svr_comm/gclient/gchat"
	"git.dhgames.cn/svr_comm/gclient/gchat/pbgo"
)

//ReUpdateChatDigest 更新digest后同时更新聊天信息
func ReUpdateChatDigest(req *pbgo.UpdateDigestReq) (err error) {
	resp, err := gchat.UpdateDigest(req)
	if err != nil {
		return err
	}
	if resp.ErrCode != 0 {
		return fmt.Errorf("code %d errmsg %s", resp.ErrCode, resp.ErrContent)
	}
	return nil
}
