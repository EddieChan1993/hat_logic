package core

import (
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"google.golang.org/protobuf/proto"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbdb"
	"hat_logic/pbgo/pbreq"
	"hat_logic/util"
)

type ModNameTyp = string

//==================== Mod 接口 ====================//

//IRoleMod 角色模块事件
type IRoleMod interface {
	OnlineHandle()  //上线事件
	OfflineHandle() //离线事件
}

//ITempDataMod 临时数据模块,需要热更时初始化的数据
type ITempDataMod interface {
	TmpDataUnmarshal([]byte) //反解析
	TmpDataMarshal() []byte  //解析
}

type IMod interface {
	IModDb
	IModReqApi
}

// IModBase 基础mod模块，所有定义的mod都需要实现该接口
type IModBase interface {
	InitMod(gameCtx *pbreq.GameCtx) //初始化mod结构体里的参数
	ApiModBefore()                  //调用apiMod前执行
	InitModAfter()                  //所有mod初始化完成之后执行，可用于玩家登陆结算
	SyncBefore()                    //在执行所有mod的Sync请求前执行（不要写任何逻辑初始化代码，因为热更不会调用）
	RspSync(rsp *pbreq.RspSync)     //sync仅作为数据同步，切勿在此写任何发奖或复杂逻辑
}

//IModDb 数据存储模块，mod需要存储则实现该接口
type IModDb interface {
	IModBase
	InitDbRoleData(db *pbdb.RoleData) //初始化db参数，用于数据解析（注意：该方法下不能调用其他模块）
	InitLoadDb(db *pbdb.RoleData)     //将db的数据load mod模块变量中
	Save(db *pbdb.RoleData)           //mod变量的更新，最终数据落地，切勿写任何逻辑代码
}

//IModReqApi api请求模块，mod需要给客户端通讯则实现该接口
type IModReqApi interface {
	IModBase
	GrpId() int32                                //获取请求的grpid
	ReqBefore()                                  //请求前执行
	Dispatch(reqMsg *pbreq.ReqMsg) proto.Message //请求协议根据cmdid进行分发到对应接口实现
}

//==================== TModMgr ====================//

type TModMgr struct {
	modList         []ModNameTyp
	modAll          map[ModNameTyp]IModBase     //基础模块
	modReqApi       map[cst.GrpIdTyp]IModReqApi //api请求模块
	modNameMpReqApi map[ModNameTyp]cst.GrpIdTyp //api请求模块 名和grpid关系映射
	modDb           map[ModNameTyp]IModDb       //数据存储模块
}

func initModMgr() *TModMgr {
	return &TModMgr{
		modAll:          map[ModNameTyp]IModBase{},
		modReqApi:       map[cst.GrpIdTyp]IModReqApi{},
		modNameMpReqApi: map[ModNameTyp]cst.GrpIdTyp{},
		modDb:           map[ModNameTyp]IModDb{},
	}
}

func (this_ *TModMgr) getGrpIdByName(modName ModNameTyp) cst.GrpIdTyp {
	return this_.modNameMpReqApi[modName]
}

func (this_ *TModMgr) getModIns(modName ModNameTyp) IModBase {
	return this_.modAll[modName]
}

//allMod 获取所有模块实例
func (this_ *TModMgr) allMod() []ModNameTyp {
	return this_.modList
}

//RegModBase 注册基础模块实例
func (this_ *TModMgr) RegModBase(mod IModBase) {
	modName := util.GetModName(mod)
	this_.checkRepeat(modName)
	this_.modAll[modName] = mod
	this_.modList = append(this_.modList, modName)
}

//RegCompleteMod 注册所有类型模块
func (this_ *TModMgr) RegCompleteMod(mod IMod) {
	modName := util.GetModName(mod)
	this_.checkRepeat(modName)
	this_.modDb[modName] = mod
	this_.modAll[modName] = mod
	this_.modReqApi[mod.GrpId()] = mod
	this_.modNameMpReqApi[modName] = mod.GrpId() //api请求模块 名和grpid关系映射
	this_.modList = append(this_.modList, modName)
}

//RegReqApiMod 注册请求模块实例
func (this_ *TModMgr) RegReqApiMod(mod IModReqApi) {
	modName := util.GetModName(mod)
	this_.checkRepeat(modName)
	this_.modReqApi[mod.GrpId()] = mod
	this_.modNameMpReqApi[modName] = mod.GrpId() //api请求模块 名和grpid关系映射
	this_.modAll[modName] = mod
	this_.modList = append(this_.modList, modName)
}

//RegDbMod 注册db模块实例
func (this_ *TModMgr) RegDbMod(mod IModDb) {
	modName := util.GetModName(mod)
	this_.checkRepeat(modName)
	this_.modDb[modName] = mod
	this_.modAll[modName] = mod
	this_.modList = append(this_.modList, modName)
}

func (this_ *TModMgr) checkRepeat(modName ModNameTyp) {
	_, had := this_.modAll[modName]
	if had {
		klog.Panic(modName, "重复注册")
	}
}

//isRoleMod 是否是IRoleModHandle模块
func (this_ *TModMgr) isRoleMod(modIns IModBase) (roleModIns IRoleMod, isOk bool) {
	roleModIns, isOk = modIns.(IRoleMod)
	return
}

//isTempDataMod 是否是零时数据模块
func (this_ *TModMgr) isTempDataMod(modIns IModBase) (roleModIns ITempDataMod, isOk bool) {
	roleModIns, isOk = modIns.(ITempDataMod)
	return
}
