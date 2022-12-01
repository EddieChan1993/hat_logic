package core

import (
	"bytes"
	"encoding/gob"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"hat_logic/util"
)

//=====热更=====

type TransData struct {
	Session  int64
	RoleId   int64
	Sid      int32
	Status   cst.PlayerStatus      //玩家在线状态
	ReadCh   [][]byte              //未处理玩的请求消息
	TempData map[ModNameTyp][]byte //临时数据
	DbBin    []byte                //游戏数据
	Client   *pbgo.LogicClientInfo //client数据
}

func HotSendData(send func(data []byte)) {
	playerMgr.Lock()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			klog.Error(panicErr)
			util.PanicStack()
		}
		playerMgr.Unlock()
	}()
	klog.Info("======HotSendData all start=====")
	playerMgr.players.Range(func(key, value interface{}) bool {
		player, isOk := value.(*Player)
		if isOk && player != nil {
			klog.Infof("%s send 开始", player.gameCtx.Log())
			oldStatus := player.status
			player.closePlayer(cst.SignalOnlyExit)
			<-player.ctx.Done() //ctx关闭，证明 信号已经完全处理完成

			ins := &TransData{
				Session:  player.sessionId,
				RoleId:   player.gameCtx.GetRoleId(),
				Sid:      player.gameCtx.GetSid(),
				Status:   oldStatus,
				TempData: player.getTempDataMod(),
				DbBin:    player.modDataToDbBin(),
				Client:   player.gameCtx.Client(),
			}
			//处理未完成消息，receive此时已经关闭，如果没有消息，receive应为空
			//我：为什么此处不直接把receive消息去dispatch处理了，这样就可以不转移这消息数据了?（马：减少热更的迁移时间）
			for msg := range player.receive {
				if msg == nil {
					klog.Infof("%s 消息为空", player.gameCtx.Log())
					continue
				}
				reqMsg := pbreq.NewReqMsg(msg.GetGrpId(), msg.GetCmdId(), msg.GetData())
				reaCh, err := encode(reqMsg)
				if err != nil {
					klog.Infof("%s 消息解析失败 %v", player.gameCtx.Log(), err)
					continue
				}
				ins.ReadCh = append(ins.ReadCh, reaCh)
				klog.Infof("%s 消息装填完成", player.gameCtx.Log())
			}
			byRes, err := encode(ins)
			if err != nil {
				klog.Error(player.gameCtx.Log(), err)
				return true
			}
			send(byRes)
			klog.Infof("%s send 完成", player.gameCtx.Log())
		}
		return true
	})
	klog.Info("======HotSendData all end=====")
}

func HotReceiveData(data []byte) {
	req := &TransData{}
	err := decode(data, req)
	if err != nil {
		klog.Error(err)
		return
	}
	err = HotUpdateLogin(&pbgo.LaunchReq{
		Session: req.Session,
		RoleId:  req.RoleId,
		Sid:     req.Sid,
		Client:  req.Client,
	}, req.Status, req.DbBin, req.TempData)
	if err != nil {
		klog.Errorf("%d receive %v", req.RoleId, err)
		return
	}
	//处理热更前未完成的消息
	for i, ch := range req.ReadCh {
		msg := &pbreq.ReqMsg{}
		err = decode(ch, msg)
		if err != nil {
			klog.Errorf("%d decode", req.RoleId, err)
			break
		}
		if msg == nil {
			continue
		}
		player, had := playerMgr.getPlayer(req.RoleId)
		if !had {
			klog.Warnf("%d getPlayer none", req.RoleId)
			break
		}
		player.receive <- pbreq.NewReqMsg(msg.GetGrpId(), msg.GetCmdId(), msg.GetData())
		klog.Infof("%d 消息补发 %d 完成", req.RoleId, i)
	}
	klog.Infof("%d receive 完成", req.RoleId)
	return
}

func encode(m interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte, m interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(m)
	if err != nil {
		return err
	}
	return nil
}
