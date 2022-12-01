package core

import (
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	gdbpb "git.dhgames.cn/svr_comm/gclient/gdb/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbdb"
	"hat_logic/rpc"
	"hat_logic/util"
	"sync/atomic"
)

//Login 登陆
func Login(req *pbgo.LaunchReq) error {
	oldPlayer, had := playerMgr.getPlayer(req.RoleId)
	if had {
		if req.Session != oldPlayer.sessionId {
			//踢下线
			err := oldPlayer.kickPlayer(cst.KickReasonLogin)
			if err != nil {
				klog.Error(oldPlayer.gameCtx.Log(), err)
				return err
			}
			oldPlayer.changeSessionId(req.Session)
		}
		oldPlayer.online()
		return nil
	}

	player := NewPlayer(playerMgr.ctx, req)
	defer func() {
		if panicErr := recover(); panicErr != nil {
			//捕获初始化mod产生的panic，删掉失败的player，防止再次登陆进入未被成功初始化的player
			playerMgr.delForcePlayer(req.RoleId)
			klog.Error(player.gameCtx.Log(), panicErr)
			util.PanicStack()
			rpc.KickBySession(player.sessionId, cst.KickInitPlayer)
			return
		}
	}()
	err := player.loadMod()
	if err != nil {
		return err
	}
	player.online()
	go player.run()
	playerMgr.addPlayer(req.RoleId, player)
	klog.Infof("%s 登陆成功", player.gameCtx.Log())
	return nil
}

//HotUpdateLogin 热更登陆
func HotUpdateLogin(req *pbgo.LaunchReq, status cst.PlayerStatus, DbBin []byte, tmpData map[string][]byte) error {
	player := NewPlayer(playerMgr.ctx, req)
	defer func() {
		if panicErr := recover(); panicErr != nil {
			//捕获初始化mod产生的panic，删掉失败的player，防止再次登陆进入未被成功初始化的player
			playerMgr.delForcePlayer(req.RoleId)
			klog.Error(player.gameCtx.Log(), panicErr)
			util.PanicStack()
			rpc.KickBySession(player.sessionId, cst.KickInitPlayer)
			return
		}
	}()
	err := player.hotLoadMod(DbBin)
	if err != nil {
		klog.Error(player.gameCtx.Log(), err)
		return err
	}
	player.restoreStatus(status)      //恢复玩家状态
	player.reloadTempDataMod(tmpData) //加载玩家临时数据
	go player.run()
	playerMgr.addPlayer(req.RoleId, player)
	klog.Infof("%s 登陆成功", player.gameCtx.Log())
	return nil
}

//WaitReconnect 断线
func WaitReconnect(req *pbgo.WaitReconnectReq) error {
	player, has := playerMgr.getPlayer(req.RoleId)
	if !has {
		return nil
	}
	if player.sessionId != req.Session {
		//只操作，当前在线的玩家离线落地，那些被顶号的玩家不处理落地情况
		return nil
	}
	player.offline()
	return nil
}

//Push 推送消息
func Push(req *pbgo.PushReq) {
	player, has := playerMgr.getPlayer(req.RoleId)
	if !has {
		klog.Warnf("角色id %d 未登陆  sessionId %d", req.RoleId, req.Session)
		return
	}
	player.sendMgr(req)
	return
}

//Reconnect 消息重连
func Reconnect(req *pbgo.ReconnectReq) error {
	oldPlayer, has := playerMgr.getPlayer(req.RoleId)
	if !has {
		return nil
	}
	if req.Session != oldPlayer.sessionId {
		err := oldPlayer.kickPlayer(cst.KickReconnect)
		if err != nil {
			klog.Error(err)
			return err
		}
		oldPlayer.changeSessionId(req.Session)
	}
	oldPlayer.online()
	return nil
}

func KickByRoleId(roleId cst.RoleId) error {
	oldPlayer, had := playerMgr.getPlayer(roleId)
	if !had {
		return nil
	}
	oldPlayer.kickPlayer(cst.KickGMKick)
	oldPlayer.closePlayer(cst.SignalOnlyExit)
	oldPlayer.gameCtx.WaitWg()
	<-oldPlayer.ctx.Done() //ctx关闭，证明 信号已经完全处理完成
	klog.Warnf("===== %s 被 GM 踢下线 =====", oldPlayer.gameCtx.Log())
	return nil
}

func KickBySid(sid cst.DeInt) {
	playerMgr.Lock()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			klog.Error(panicErr)
			util.PanicStack()
		}
		playerMgr.Unlock()
	}()
	klog.Info("======GM KickBySid start=====")
	playerMgr.players.Range(func(key, value interface{}) bool {
		player, isOk := value.(*Player)
		if isOk && player != nil {
			if player.gameCtx.GetSid() != sid {
				return true
			}
			klog.Infof("%s batchClosePlayer start", player.gameCtx.Log())
			player.kickPlayer(cst.KickGMStopSvr)
			player.closePlayer(cst.SignalExitWithSaveDb)
			player.gameCtx.WaitWg()
			<-player.ctx.Done() //ctx关闭，证明 信号已经完全处理完成
			klog.Infof("%s batchClosePlayer end", player.gameCtx.Log())
		}
		return true
	})
	klog.Info("======GM KickBySid end=====")
}

//KickBeforeHotUpdate 热更前踢人
func KickBeforeHotUpdate() {
	playerMgr.Lock()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			klog.Error(panicErr)
			util.PanicStack()
		}
		playerMgr.Unlock()
	}()
	klog.Info("======GM KickToHotUpdate start=====")
	playerMgr.players.Range(func(key, value interface{}) bool {
		player, isOk := value.(*Player)
		if isOk && player != nil {
			klog.Infof("%s batchClosePlayer start", player.gameCtx.Log())
			player.kickPlayer(cst.KickToHotUpdate)
			player.closePlayer(cst.SignalExitWithSaveDb)
			player.gameCtx.WaitWg()
			<-player.ctx.Done() //ctx关闭，证明 信号已经完全处理完成
			klog.Infof("%s batchClosePlayer end", player.gameCtx.Log())
		}
		return true
	})
	klog.Info("======GM KickToHotUpdate end=====")
}

func GetRoleData(roleId cst.RoleId) ([]byte, error) {
	//获取db数据
	gdbRes, err := rpc.LoadLogic(roleId)
	if err != nil {
		klog.Errorf("roleId %d error %v", roleId, err)
		return nil, err
	}
	roleData := &pbdb.RoleData{}
	err = proto.Unmarshal(gdbRes.Bin, roleData)
	if err != nil {
		klog.Errorf("roleId %d error %v", roleId, err)
		return nil, err
	}
	btData, err := jsoniter.Marshal(roleData)
	if err != nil {
		klog.Errorf("roleId %d error %v", roleId, err)
		return nil, err
	}
	return btData, nil
}

func SetRoleData(roleId cst.RoleId, data string) error {
	roleData := &pbdb.RoleData{}
	err := jsoniter.Unmarshal([]byte(data), roleData)
	if err != nil {
		klog.Errorf("roleId %d error %v", roleId, err)
		return err
	}
	btRes, err := proto.Marshal(roleData)
	if err != nil {
		klog.Errorf("roleId %d error %v", roleId, err)
		return err
	}
	ret := &gdbpb.SaveLogicReq{
		RoleId: rpc.TransRoleId(roleId),
		Bin:    btRes,
	}
	//rpc保存数据
	_, err = rpc.SaveLogic(ret)
	if err != nil {
		klog.Errorf("roleId %d error %v", roleId, err)
		return err
	}
	klog.Infof("==== roleId %d GM 修改玩家数据成功 ====", roleId)
	return nil
}

//OnlinePeople 当前所有在线人数
func OnlinePeople() cst.DeInt {
	return atomic.LoadInt32(&playerMgr.onlinePlayers)
}
