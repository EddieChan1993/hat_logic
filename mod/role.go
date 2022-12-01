package mod

/**
玩家
*/
import (
	pbgo2 "git.dhgames.cn/svr_comm/gclient/gchat/pbgo"
	"git.dhgames.cn/svr_comm/gclient/gdb/constant"
	"git.dhgames.cn/svr_comm/gclient/gdb/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/cast"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/core/cst"
	"hat_logic/mod/mcst"
	config "hat_logic/pbgo/pbcfg"
	"hat_logic/pbgo/pbdb"
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
	"hat_logic/tool/bi"
	"hat_logic/util"
	"time"
)

const maxNameLen = 18 //名字最大字符个数

type Role struct {
	pbreq.TModApi17
	*pbgo.RoleDigest
	isTodayFirstLogin bool //今日是否首次登陆
	dbData            *pbdb.RoleMgr
	temporaryData     map[cst.GrpIdTyp][]byte //临时数据存储，热更需要加载进去
}

//==================== api ====================//

type IModRoleFun interface {
	RoleDigestInfo() *pbgo.RoleDigest
	RoleAccount() mcst.DefaultInt
	RoleState() mcst.DefaultInt
	RoleLastOfflineTime() mcst.Int64Typ
	RoleCreateAt() mcst.Int64Typ
	IsUnlockMod(lockId mcst.LockTyp) bool //是否解锁某个功能
}

func RoleMod(gameCtx *pbreq.GameCtx) IModRoleFun {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&Role{}))
	if had {
		return ins.(IModRoleFun)
	}
	return nil
}

//==================== api ====================//

func (this_ *Role) ReqChangeAvatar(req *pbreq.ReqChangeAvatar) *pbreq.RspChangeAvatar {
	res := &pbreq.RspChangeAvatar{AvatarId: req.AvatarId}
	err := this_.UpdateAvatar(req.AvatarId)
	if err != nil {
		klog.Warn(this_.Log(), err)
		return res
	}
	if err != nil {
		res.Status = 2
		return res
	}
	klog.Infof("%s ReqChangeAvatar avatarId %d ", this_.Log(), req.AvatarId)
	return res
}

func (this_ *Role) ReqChangeName(req *pbreq.ReqChangeName) *pbreq.RspChangeName {
	res := &pbreq.RspChangeName{Name: req.Name}
	klog.Info(this_.Log(), req.Name)
	if len([]rune(req.Name)) > maxNameLen {
		//名字不能超过最大字符个数
		res.Status = 4
		return res
	}
	costItem := make(mcst.ItemIdMapNum)
	if this_.dbData.Setting.ChangedName {
		//非首次改名，需要消耗物品
		CstChangeNameCost := config.ConstData.Get(mcst.CstChangeNameCost)
		costItem = mcst.ItemIdMapNum{
			mcst.ItemIdDIAMOND: mcst.BagNum(CstChangeNameCost.Value),
		}
		err := BagMod(this_.GameCtx).SubItem(costItem)
		if err != nil {
			res.Status = 2
			return res
		}
	}
	err := rpc.RoleDigestSave(&pbgo.SaveDigestReq{
		Digest: &pbgo.RoleDigest{
			RoleId: this_.RoleDigest.RoleId,
			Name:   req.Name,
		},
		HasDigest: &pbgo.HasRoleDigest{Name: true},
	})
	if err != nil {
		BagMod(this_.GameCtx).AddItem(costItem)
		if err.Error() == constant.ErrMsgDupName {
			res.Status = 3
			return res
		}
		res.Status = 1
		return res
	}
	OldName := this_.RoleDigest.Name
	this_.RoleDigest.Name = req.Name
	BiMod(this_.GameCtx).Upload(bi.ChangeName, bi.EventValJson{
		"before_name": OldName,
		"after_name":  req.Name,
	})
	err = rpc.ReUpdateChatDigest(&pbgo2.UpdateDigestReq{Digest: this_.toChatDigest()})
	if err != nil {
		klog.Warn(this_.Log(), err)
	}
	this_.dbData.Setting.ChangedName = true
	klog.Infof("%s ReqChangeName name %s", this_.Log(), req.Name)
	return res
}

//========================tool========================//

//UpdateAvatar 更新头像
func (this_ *Role) UpdateAvatar(avatarId mcst.DefaultInt) error {
	err := rpc.RoleDigestSave(&pbgo.SaveDigestReq{
		Digest: &pbgo.RoleDigest{
			RoleId: this_.RoleDigest.RoleId,
			Logo:   avatarId,
		},
		HasDigest: &pbgo.HasRoleDigest{Logo: true},
	})
	if err != nil {
		return err
	}
	this_.RoleDigest.Logo = avatarId
	err = rpc.ReUpdateChatDigest(&pbgo2.UpdateDigestReq{Digest: this_.toChatDigest()})
	if err != nil {
		return err
	}
	return nil
}

func (this_ *Role) toChatDigest() *pbgo2.RoleDigest {
	return &pbgo2.RoleDigest{
		Id:    int64(this_.RoleDigest.RoleId),
		Name:  this_.RoleDigest.Name,
		Level: int64(this_.RoleDigest.Lv),
		Icon:  cast.ToString(this_.RoleDigest.Logo),
		Job:   pbgo2.RoleDigest_None,
		Frame: cast.ToString(this_.RoleDigest.HeadFrame),
		Title: cast.ToString(this_.RoleDigest.HeadTitle),
	}
}

//RoleDigestInfo 玩家基础信息
func (this_ *Role) RoleDigestInfo() *pbgo.RoleDigest {
	return this_.RoleDigest
}

//RoleAccount 账号
func (this_ *Role) RoleAccount() mcst.DefaultInt {
	return this_.RoleDigest.Account
}

//RoleLastOfflineTime 玩家最近离线时间(时间戳毫秒）
func (this_ *Role) RoleLastOfflineTime() mcst.Int64Typ {
	return this_.RoleDigest.LastOfflineTime * util.Milli
}

//RoleState 玩家状态
func (this_ *Role) RoleState() mcst.DefaultInt {
	return this_.RoleDigest.Status
}

func (this_ *Role) BiLogout() {
	BiMod(this_.GameCtx).Upload(bi.Logout, make(bi.EventValJson))
}

// BiSnapshot 角色数据快照
func (this_ *Role) BiSnapshot() {
	bag := BagMod(this_.GameCtx)
	eventInfo := bi.EventValJson{
		"uid":        this_.RoleId,
		"name":       this_.Name,
		"lv":         this_.Lv,
		"reg":        this_.Ctime,
		"login":      this_.LastOnlineTime,
		"power":      this_.Power,
		"gold":       bag.GetItem(mcst.ItemIdCOIN),
		"diamond":    bag.GetItem(mcst.ItemIdDIAMOND),
		"friend_num": 0,
	}
	BiMod(this_.GameCtx).Upload(bi.Player, eventInfo)
}

//RoleCreateAt 玩家创建时间
func (this_ *Role) RoleCreateAt() mcst.Int64Typ {
	return this_.RoleDigest.Ctime
}

//IsUnlockMod 是否解锁某个功能
func (this_ *Role) IsUnlockMod(lockId mcst.LockTyp) bool {
	return true
}

//========================mod base========================//

func (this_ *Role) InitDbRoleData(db *pbdb.RoleData) {
	db.RoleMgr = &pbdb.RoleMgr{
		Setting:   &pbdb.Setting{},
		ExtStatus: &pbdb.ExtStatus{},
	}
}

func (this_ *Role) InitLoadDb(db *pbdb.RoleData) {
	this_.dbData = db.RoleMgr
}

func (this_ *Role) Save(db *pbdb.RoleData) {
	db.RoleMgr = this_.dbData
}

//OnlineHandle 玩家上线事件
func (this_ *Role) OnlineHandle() {
	CronTasksMod(this_.GameCtx).RunTasks()
	BiMod(this_.GameCtx).Login()
}

//OfflineHandle 玩家离线事件
func (this_ *Role) OfflineHandle() {
	this_.BiLogout()
	this_.BiSnapshot()
}

func (this_ *Role) InitMod(gameCtx *pbreq.GameCtx) {
	this_.GameCtx = gameCtx
	gdbRole, err := rpc.RoleDigest(this_.GameCtx.GetRoleId())
	if err != nil {
		klog.Warn(this_.Log(), err)
	}
	this_.RoleDigest = gdbRole.Digest
	if this_.RoleDigest.Logo == 0 {
		//默认第一个头像
		allHead := config.HeadIconData.GetAll()
		err = this_.UpdateAvatar(allHead[0])
		if err != nil {
			klog.Warn(this_.Log(), err)
		}
	}
}

func (this_ *Role) SyncBefore() {
	if this_.dbData.ExtStatus.LastLoginZeroAt != util.ZeroToday() {
		this_.isTodayFirstLogin = true
	} else {
		this_.isTodayFirstLogin = false
	}
	this_.dbData.ExtStatus.LastLoginZeroAt = util.ZeroToday()
}

func (this_ *Role) RspSync(resp *pbreq.RspSync) {
	resp.SvrTime = time.Now().Unix()
	gdbRole := this_.RoleDigest
	dbData := this_.dbData
	resp.RoleMgr = &pbreq.RoleMgr{
		Digest: &pbreq.Role{
			Id:              gdbRole.RoleId,
			Name:            gdbRole.Name,
			Logo:            gdbRole.Logo,
			Exp:             mcst.DefaultInt(gdbRole.Exp),
			Vexp:            mcst.DefaultInt(gdbRole.Vexp),
			Sid:             gdbRole.Sid,
			Status:          gdbRole.Status,
			LastOfflineTime: gdbRole.LastOfflineTime,
		},
		Settings: &pbreq.Settings{
			IsFirstChangeName: dbData.Setting.ChangedName == false,
		},
		ExtStatus: &pbreq.ExtStatus{
			IsTodayFirstLogin: this_.isTodayFirstLogin,
			NextDayCd:         util.NextDayCd(),
		},
	}
}

func NewRole() *Role {
	res := &Role{}
	res.InitModApi(res)
	return res
}
