package core

import (
	"context"
	"git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	gdbpb "git.dhgames.cn/svr_comm/gclient/gdb/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"google.golang.org/protobuf/proto"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbdb"
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
	"hat_logic/tool/consul/static"
	"hat_logic/util"
	"sync/atomic"
	"time"
)

const maxReceiveCh = 3

type Player struct {
	receive      chan *pbreq.ReqMsg          //消息收发通道
	playerSignal chan cst.PlayerSignalTyp    //系统内部踢出进程信号
	ctx          context.Context             //ctx进程管理
	cancel       context.CancelFunc          //ctx进程取消函数
	roleId       cst.RoleId                  //玩家id
	sessionId    cst.SessionId               //玩家sessionid
	gameCtx      *pbreq.GameCtx              //模块之间公用的数据
	modBase      []IModBase                  //基础模块集
	modReqApi    map[cst.GrpIdTyp]IModReqApi //模块之间相互调用的api方法集
	modTmpData   map[ModNameTyp]ITempDataMod //临时数据模块
	modRole      IRoleMod
	modDb        map[ModNameTyp]IModDb //DB模块集
	ticker       *time.Ticker          //定时存储数据
	status       cst.PlayerStatus      //玩家在线状态
}

func NewPlayer(ctx context.Context, req *pbgo.LaunchReq) *Player {
	childCtx, childCancel := context.WithCancel(ctx)
	playerIns := &Player{
		receive:      make(chan *pbreq.ReqMsg, maxReceiveCh),
		ctx:          childCtx,
		cancel:       childCancel,
		roleId:       req.RoleId,
		sessionId:    req.Session,
		modReqApi:    map[cst.GrpIdTyp]IModReqApi{},
		gameCtx:      pbreq.NewGameCtx(childCtx),
		playerSignal: make(chan cst.PlayerSignalTyp, 1),
	}
	playerIns.ticker = time.NewTicker(playerIns.randSaveDbTickerTime())
	playerIns.gameCtx.SetPlayerInfo(&pbreq.PlayerInfo{
		RoleId:    req.RoleId,
		SessionId: req.Session,
		Sid:       req.Sid,
		TaskFn:    make(chan pbreq.TaskFn, 3),
		Client:    req.Client,
	})
	return playerIns
}

//run 运行玩家worker进程
func (this_ *Player) run() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			klog.Error(this_.gameCtx.Log(), panicErr)
			util.PanicStack()
			_, err := rpc.KickBySession(this_.sessionId, cst.KickModPanic)
			if err != nil {
				klog.Warn(this_.gameCtx.Log(), err)
			}
			this_.loginOutWithoutSaveDb()
		}
		this_.ticker.Stop()
		this_.cancel() //执行完成，取消所有玩家ctx
		klog.Infof("%s worker退出", this_.gameCtx.Log())
	}()
	for {
		select {
		case signalId := <-this_.playerSignal:
			klog.Infof("%s 玩家信号", this_.gameCtx.Log())
			this_.playerSignalEvent(signalId)
			return
		case <-this_.ticker.C:
			//定时处理
			klog.Infof("%s 定时信号", this_.gameCtx.Log())
			this_.timeEvent()
		case msg, ok := <-this_.receive:
			if !ok {
				return
			}
			this_.handleMsg(msg)
		case fn, ok := <-this_.gameCtx.TaskFnCh():
			if !ok {
				return
			}
			fn()
		}
	}
}

//timeEvent 定时任务事件处理
func (this_ *Player) timeEvent() {
	switch atomic.LoadUint32(&this_.status) {
	case cst.PlayerStatusLoginOut:
		klog.Infof("%s 开始退出流程", this_.gameCtx.Log())
		if !this_.isOnline() {
			this_.closePlayer(cst.SignalExitWithSaveDb)
		}
	case cst.PlayerStatusLogin:
		klog.Infof("%s 定时saveDB开始", this_.gameCtx.Log())
		this_.saveDb()
	}
}

func (this_ *Player) closePlayer(sig cst.PlayerSignalTyp) {
	if this_.status == cst.PlayerStatusClose {
		return
	}
	atomic.StoreUint32(&this_.status, cst.PlayerStatusClose) //标记关闭状态
	playerMgr.delPlayer(this_.roleId)                        //删除玩家管理
	this_.playerSignal <- sig                                //发送停止信号
}

func (this_ *Player) playerSignalEvent(signalId cst.PlayerSignalTyp) {
	switch signalId {
	case cst.SignalExitWithSaveDb:
		this_.modRole.OfflineHandle()        //执行玩家离线事件
		err := rpc.RoleOffline(this_.roleId) //标记离线状态
		if err != nil {
			//不能返回
			klog.Errorf("%s RoleOffline %v", this_.gameCtx.Log(), err)
		}
		//处理残留消息，提前close ch，range缓冲通道才能不阻塞
		close(this_.receive)
		this_.handleLeftMsg()
		err = this_.saveDb()
		if err != nil {
			//不能返回
			klog.Errorf("%s 落地失败 %v", this_.gameCtx.Log(), err)
		}
	case cst.SignalOnlyExit:
		close(this_.receive)
	}
}

//handleMsg 消息统一处理
func (this_ *Player) handleMsg(m *pbreq.ReqMsg) {
	res := this_.dispatch(m)
	if res == nil {
		klog.Errorf("%s 返回消息为空", this_.gameCtx.Log())
		return
	}
	this_.callGate(m.GetGrpId(), m.GetCmdId(), res)
	//klog.Info(this_.gameCtx.Log(), "消息处理完成")
}

//handleLeftMsg 处理通道内未残留消息
func (this_ *Player) handleLeftMsg() {
	for msg := range this_.receive {
		if msg == nil {
			klog.Infof("%s 消息为空", this_.gameCtx.Log())
			continue
		}
		this_.dispatch(msg)
	}
}

//dispatch 消息分发
func (this_ *Player) dispatch(msg *pbreq.ReqMsg) proto.Message {
	if msg.GetCmdId() == pbreq.CmdReqSyncId &&
		msg.GetGrpId() == pbreq.GrpId10 {
		//sync请求执行
		return this_.InitModSync()
	} else {
		//获取对应模块
		modIns := this_.modReqApi[msg.GetGrpId()]
		if modIns == nil {
			klog.Panic(this_.gameCtx.Log(), "该GrpId对应的模块没有注册", "GrpId=", msg.GetGrpId())
			return nil
		}
		//获取到对应pbid的请求结构体实例
		ReqPb := pbreq.GetPbReq(msg.GetGrpId(), msg.GetCmdId())
		if ReqPb == nil {
			klog.Panic(this_.gameCtx.Log(), "该协议未注册", msg.GetGrpId(), msg.GetCmdId())
			return nil
		}
		err := proto.Unmarshal(msg.GetData(), ReqPb)
		if err != nil {
			klog.Panic(this_.gameCtx.Log(), "协议解析出错", err)
			return nil
		}
		msg.ProtoMsg = ReqPb
		modIns.ReqBefore()
		return modIns.Dispatch(msg)
	}
}

//InitModSync 初始化注册模块的sync方法
func (this_ *Player) InitModSync() proto.Message {
	resp := &pbreq.RspSync{}
	for _, modIns := range this_.modBase {
		modIns.SyncBefore()
	}
	for _, modIns := range this_.modBase {
		modIns.RspSync(resp)
	}
	klog.Infof("%s sync完成", this_.gameCtx.Log())
	return resp
}

//callGate 发消息给gogate
func (this_ *Player) callGate(grp, cmd cst.GrpIdTyp, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		klog.Error(err)
		return
	}
	req := &pbgo.SendBySessionReq{Session: this_.sessionId, Grp: grp, Cmd: cmd, Bin: data}
	_, err = rpc.SendBySession(req)
	if err != nil {
		klog.Error(err)
		return
	}
}

// sendMgr 发送消息
func (this_ *Player) sendMgr(req *pbgo.PushReq) error {
	msg := pbreq.NewReqMsg(req.Grp, req.Cmd, req.Bin)
	if this_.status == cst.PlayerStatusClose {
		//TODO 应该把响应proto.Message返回，而不是nil
		return nil
	}
	timeoutC := time.NewTimer(cst.ReqMsgTimeOut)
	defer timeoutC.Stop()
	select {
	case this_.receive <- msg:
		//klog.Info(this_.gameCtx.Log(), "消息发送完成")
	case <-timeoutC.C:
		//请求超时检测，上个请求还没还行完成触发
		klog.Infof("%s 请求超时", this_.gameCtx.Log())
	}
	return nil
}

func (this_ *Player) loginOutWithoutSaveDb() {
	playerMgr.delPlayer(this_.roleId)
	//更新db玩家状态
	err := rpc.RoleOffline(this_.roleId)
	if err != nil {
		klog.Warn(this_.gameCtx.Log(), err)
	}
}

//kickPlayer 踢玩家下线通知
func (this_ *Player) kickPlayer(reason string) error {
	_, err := rpc.KickBySession(this_.sessionId, reason, kite.Cast)
	if err != nil {
		klog.Error(err)
		return err
	}
	klog.Infof("%s 踢玩家下线完成 %s", this_.gameCtx.Log(), reason)
	return nil
}

//changeSessionId 改变sessionId
func (this_ *Player) changeSessionId(sessionId cst.SessionId) {
	if sessionId != this_.sessionId {
		klog.Infof("%s 顶号 newSessionId %d oldSessionId %d", this_.gameCtx.Log(), sessionId, this_.sessionId)
		this_.sessionId = sessionId
		this_.gameCtx.SetSessionId(sessionId)
	}
}

//saveDb 数据落地
func (this_ *Player) saveDb() error {
	klog.Infof("%s 开始落地", this_.gameCtx.Log())
	//数据落地
	res, err := this_.memoryToDb(this_.roleId)
	if err != nil {
		klog.Error(this_.gameCtx.Log(), err)
		return err
	}
	//rpc保存数据
	_, err = rpc.SaveLogic(res)
	if err != nil {
		klog.Error(this_.gameCtx.Log(), err)
		return err
	}
	klog.Infof("%s 落地完成", this_.gameCtx.Log())
	return nil
}

//loadMod 加载数据到Mod内存
func (this_ *Player) loadMod() error {
	var err error
	//获取db数据
	gdbRes, err := rpc.LoadLogic(this_.roleId)
	if err != nil {
		klog.Error(err)
		return err
	}
	err = this_.dbToMemory(gdbRes.Bin)
	if err != nil {
		klog.Error(err)
		return err
	}
	this_.runInitModAfter()
	this_.modRole.OnlineHandle()
	return nil
}

//hotLoadMod 热更数据加载
func (this_ *Player) hotLoadMod(bin []byte) error {
	var err error
	//获取db数据
	err = this_.dbToMemory(bin)
	if err != nil {
		klog.Error(err)
		return err
	}
	this_.runInitModAfter()
	this_.modRole.OnlineHandle()
	return nil
}

func (this_ *Player) dbToMemory(bin []byte) error {
	roleData := &pbdb.RoleData{}
	modMgr := newModMgr()
	for _, modIns := range modMgr.modDb {
		modIns.InitDbRoleData(roleData)
	}
	//proto.UnmarshalOptions{Merge: true}
	//proto默认会忽略掉内嵌引用类型值为空的字段，将不被反解析生成初始化值
	//通过配置merge，解析中为空值，则使用被初始化的值
	err := proto.UnmarshalOptions{Merge: true}.Unmarshal(bin, roleData)
	if err != nil {
		klog.Error(err)
		return err
	}
	apiFunc := make(map[ModNameTyp]cst.IApiFunc)
	modReqApi := make(map[cst.GrpIdTyp]IModReqApi)
	modDb := make(map[ModNameTyp]IModDb)
	modBase := make([]IModBase, 0, len(modMgr.allMod()))
	modTmp := make(map[ModNameTyp]ITempDataMod)
	for _, modName := range modMgr.allMod() {
		modIns := modMgr.getModIns(modName)
		modIns.InitMod(this_.gameCtx)

		if modRoleIns, isOk := modMgr.isRoleMod(modIns); isOk {
			this_.modRole = modRoleIns
		}
		if mod, isOk := modMgr.isTempDataMod(modIns); isOk {
			modTmp[modName] = mod
		}
		//请求模块
		grpId := modMgr.getGrpIdByName(modName)
		if grpId != 0 {
			modReqApi[grpId] = modIns.(IModReqApi)
		}
		//db模块
		_, isModDb := modMgr.modDb[modName]
		if isModDb {
			modDb[modName] = modIns.(IModDb)
			modDb[modName].InitLoadDb(roleData)
		}
		apiFunc[modName] = modIns
		modBase = append(modBase, modIns)
	}
	this_.gameCtx.RegApiFunc(apiFunc)
	this_.modBase = modBase
	this_.modReqApi = modReqApi
	this_.modDb = modDb
	this_.modTmpData = modTmp
	return nil
}

//modDataToDbBin mod数据转为db bin
func (this_ *Player) modDataToDbBin() []byte {
	allData := &pbdb.RoleData{}
	for _, mod := range this_.modDb {
		mod.Save(allData)
	}
	bin, err := proto.Marshal(allData)
	if err != nil {
		klog.Error(this_.gameCtx.Log(), err)
		return nil
	}
	return bin
}

//执行所有mod的InitModAfter方法
func (this_ *Player) runInitModAfter() {
	for _, mod := range this_.modBase {
		mod.InitModAfter()
	}
}
func (this_ *Player) memoryToDb(roleId int64) (rsp *gdbpb.SaveLogicReq, err error) {
	ret := &gdbpb.SaveLogicReq{
		RoleId: rpc.TransRoleId(roleId),
		Bin:    nil,
	}
	ret.Bin = this_.modDataToDbBin()
	return ret, nil
}

func (this_ *Player) online() {
	if !atomic.CompareAndSwapUint32(&this_.status, cst.PlayerStatusLoginOut, cst.PlayerStatusLogin) {
		return
	}
	this_.ticker.Reset(this_.randSaveDbTickerTime())
	gdbRsp, err := rpc.RoleOnline(this_.roleId)
	if err != nil {
		klog.Error(this_.gameCtx.Log(), "err", err)
	} else if gdbRsp.ErrNo != 0 {
		klog.Error(this_.gameCtx.Log(), "gdbErr", gdbRsp.ErrMsg)
	}
	klog.Infof("%s 玩家上线", this_.gameCtx.Log())
}

func (this_ *Player) offline() {
	if !atomic.CompareAndSwapUint32(&this_.status, cst.PlayerStatusLogin, cst.PlayerStatusLoginOut) {
		return
	}
	this_.ticker.Reset(time.Second*time.Duration(static.StaticWaitReconnectWaitTime()) + cst.PlayerOfflineDbSaveTick)
	klog.Infof("%s 玩家离线", this_.gameCtx.Log())
}

func (this_ *Player) isOnline() bool {
	return atomic.LoadUint32(&this_.status) == cst.PlayerStatusLogin
}

//reloadTempDataMod 重载临时数据到模块中
func (this_ *Player) reloadTempDataMod(tmpData map[ModNameTyp][]byte) {
	if tmpData == nil {
		return
	}
	for typ, mod := range this_.modTmpData {
		mod.TmpDataUnmarshal(tmpData[typ])
	}
}

//getTempDataMod 保存临时数据
func (this_ *Player) getTempDataMod() map[ModNameTyp][]byte {
	res := make(map[ModNameTyp][]byte)
	for typ, mod := range this_.modTmpData {
		res[typ] = mod.TmpDataMarshal()
	}
	return res
}

//restoreStatus 恢复状态
func (this_ *Player) restoreStatus(status cst.PlayerStatus) {
	if atomic.CompareAndSwapUint32(&this_.status, cst.PlayerStatusLoginOut, cst.PlayerStatusLogin) {
		klog.Infof("%s 热更上线 状态 %d", this_.gameCtx.Log(), status)
	}
	if status == cst.PlayerStatusLoginOut {
		this_.offline()
		return
	}
}

//randSaveDbTickerTime 存储间隔db间隔时间
func (this_ *Player) randSaveDbTickerTime() time.Duration {
	//离散数据存储间隔，防止热更后，玩家存储间隔时间过于集中
	return time.Duration(cst.DeInt(cst.PlayerOnlineDbSaveTick)+util.RandInt32(60)) * time.Second
}
