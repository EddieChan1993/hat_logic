package core

import (
	"bytes"
	"encoding/gob"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	klog "git.dhgames.cn/svr_comm/gcore/glog"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"hat_logic/tool/ants"
	"hat_logic/util"
	"sync"
)

//=====热更=====

type AllTransData struct {
	PlayersData []*TransData //所有玩家内存数据
}
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
	count := playerCount()
	dataAll := make([]*TransData, count)
	var err error
	var wg sync.WaitGroup
	var index cst.DeInt
	playerMgr.players.Range(func(key, value interface{}) bool {
		player, isOk := value.(*Player)
		if !isOk || player == nil {
			return true
		}
		i := index
		index++
		wg.Add(1)
		err = ants.AntsGo.Submit(func() {
			defer func() {
				if panicErr := recover(); panicErr != nil {
					klog.Error(panicErr)
					util.PanicStack()
				}
				wg.Done()
				klog.Infof("%s send 完成", player.gameCtx.Log())
			}()
			klog.Infof("%s send 开始", player.gameCtx.Log())
			oldStatus := player.status
			player.closePlayer(cst.SignalOnlyExit)
			klog.Infof("%s send Done 1", player.gameCtx.Log())
			<-player.ctx.Done() //ctx关闭，证明 信号已经完全处理完成
			klog.Infof("%s send Done 2", player.gameCtx.Log())
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
			dataAll[i] = ins
		})
		if err != nil {
			wg.Done()
			klog.Error(err)
			return true
		}
		return true
	})
	wg.Wait()
	klog.Info("======HotSendData wait end=====")
	//过滤0值
	dataAllNew := filterNil(dataAll)
	if len(dataAllNew) != 0 {
		byRes, err := encode(&AllTransData{PlayersData: dataAllNew})
		if err != nil {
			klog.Error(err)
			return
		}
		send(byRes)
	}
	klog.Info("======HotSendData all end=====")
}

func HotReceiveData(data []byte) {
	allData := &AllTransData{}
	err := decode(data, allData)
	if err != nil {
		klog.Error(err)
		return
	}
	var wg sync.WaitGroup
	klog.Infof("ReceiveData Players Count %d", len(allData.PlayersData))
	for _, reqTmp := range allData.PlayersData {
		if reqTmp == nil {
			klog.Warnf("req nil")
			continue
		}
		req := reqTmp
		wg.Add(1)
		err = ants.AntsGo.Submit(func() {
			defer func() {
				if panicErr := recover(); panicErr != nil {
					klog.Error(panicErr)
					util.PanicStack()
				}
				wg.Done()
			}()
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
		})
		if err != nil {
			klog.Error(err)
			wg.Done()
			continue
		}
	}
	wg.Wait()
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

// playerCount 进程管理中玩家数量
func playerCount() cst.DeInt {
	var count cst.DeInt
	playerMgr.players.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// filterNil 过滤掉0值
func filterNil(data []*TransData) []*TransData {
	dataAllNew := make([]*TransData, 0, len(data))
	for i, per := range data {
		if per == nil {
			continue
		}
		dataAllNew = append(dataAllNew, data[i])
	}
	return dataAllNew
}
