package core

import (
	"context"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/core/cst"
	"hat_logic/util"
	"sync"
	"sync/atomic"
)

//==================== playerMgr ====================//

//TPlayerMgr 管理player
type TPlayerMgr struct {
	sync.RWMutex
	onlinePlayers cst.DeInt          //在线人数
	players       sync.Map           //存储player集
	ctx           context.Context    //根ctx进程管理器
	cancel        context.CancelFunc //根ctx进程取消函数
}

var playerMgr *TPlayerMgr

func InitPlayerMgr() {
	ctx, cancel := context.WithCancel(context.Background())
	playerMgr = &TPlayerMgr{
		players: sync.Map{},
		ctx:     ctx,
		cancel:  cancel,
	}
}

//stopServer 关闭playerMgr
func (this_ *TPlayerMgr) stopServer() {
	this_.batchClosePlayer()
	klog.Info("player 管理器关闭")
}

//batchClosePlayer 批量踢出玩家进程
func (this_ *TPlayerMgr) batchClosePlayer() {
	this_.Lock()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			klog.Error(panicErr)
			util.PanicStack()
		}
		this_.Unlock()
	}()
	klog.Info("======batchClosePlayer all start=====")
	this_.players.Range(func(key, value interface{}) bool {
		player, isOk := value.(*Player)
		if isOk && player != nil {
			klog.Infof("%s batchClosePlayer start", player.gameCtx.Log())
			player.kickPlayer(cst.KickStopSvr)
			player.closePlayer(cst.SignalExitWithSaveDb)
			<-player.ctx.Done() //ctx关闭，证明 信号已经完全处理完成
			klog.Infof("%s batchClosePlayer end", player.gameCtx.Log())
		}
		return true
	})
	klog.Info("======batchClosePlayer all end=====")
}

//addPlayer 添加player
func (this_ *TPlayerMgr) addPlayer(roleId int64, p *Player) (player *Player, loaded bool) {
	old, has := this_.players.LoadOrStore(roleId, p)
	if has {
		return old.(*Player), true
	}
	this_.countOnlinePlayers(1)
	return old.(*Player), false
}

//delPlayer 删除player
func (this_ *TPlayerMgr) delPlayer(roleId int64) {
	old, had := this_.players.Load(roleId)
	if !had {
		return
	}
	//关闭player进程
	oldPlayer := old.(*Player)
	this_.players.Delete(roleId)
	this_.countOnlinePlayers(-1)
	klog.Infof("%s 删除注册", oldPlayer.gameCtx.Log())
}

//delForcePlayer 强制删除player
func (this_ *TPlayerMgr) delForcePlayer(roleId int64) {
	old, had := this_.players.LoadAndDelete(roleId)
	if !had {
		return
	}
	//关闭player进程
	oldPlayer := old.(*Player)
	oldPlayer.cancel()
	this_.countOnlinePlayers(-1)
	klog.Infof("%s 删除注册", oldPlayer.gameCtx.Log())
}

//getPlayer 获取player
func (this_ *TPlayerMgr) getPlayer(roleId int64) (*Player, bool) {
	this_.RLock()
	defer this_.RUnlock()
	w, has := this_.players.Load(roleId)
	if has {
		return w.(*Player), true
	}
	return nil, false
}

//countOnlinePlayers 记录系统当前人数
func (this_ *TPlayerMgr) countOnlinePlayers(delets cst.DeInt) {
	now := atomic.AddInt32(&this_.onlinePlayers, delets)
	if now < 0 {
		atomic.StoreInt32(&this_.onlinePlayers, 0)
	}
	klog.Infof("当前人数 %d", now)
}
