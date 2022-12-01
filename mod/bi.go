package mod

import (
	"encoding/json"
	"git.dhgames.cn/svr_comm/gclient/utlog/pbgo"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/pbgo/pbreq"
	"hat_logic/tool/bi"
	"hat_logic/util"
	"time"
)

type Bi struct {
	*pbreq.GameCtx
	biProduct *bi.BiProduct //玩家自身产销
}

//==================== func ====================//

type IModBiFun interface {
	Login()
	Upload(eventName bi.EventName, eventInfo bi.EventValJson)
	InitProduct() bi.IBiProduct
	AddProductLog(logs []*bi.ItemLog)
	AppsflyerUpload(eventName bi.EventName, eventValue bi.EventValJson)
}

func BiMod(gameCtx *pbreq.GameCtx) IModBiFun {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&Bi{}))
	if had {
		return ins.(IModBiFun)
	}
	return nil
}

func (b *Bi) Login() {
	role := RoleMod(b.GameCtx).RoleDigestInfo()
	client := b.Client()
	eventValue := bi.EventValJson{
		"user_info": bi.EventValJson{
			"bundle_id":   client.BundleId,
			"sub_package": client.SubPackage,
			"server_id":   role.Sid,
			"user_id":     role.RoleId,
			"user_name":   role.Name,
			"session_id":  client.GetSessionId(),
			"account":     role.Account,
			"platform":    client.Platform,
			"lv":          role.Lv,
			"vip":         0,
		},
		"device_info": bi.EventValJson{
			"adid":         client.Adid,
			"idfv":         client.Idfv,
			"sm_id":        client.Dhid,
			"imei":         client.Imei,
			"android_id":   client.AndroidId,
			"appsflyer_id": client.AppsflyerId,
			"device_token": client.DeviceToken,
			"mac_address":  client.MacAddress,
			"device_model": client.DeviceModel,
			"device_name":  client.DviceName,
			"os_version":   client.OsVersion,
			"language":     client.Language,
			"network_type": client.NetworkType,
			"app_version":  client.AppVersion,
			"ip":           client.Ip,
			"oaid":         client.Adid,
		},
		"event_info": "",
	}
	bi.PushLogBatch(bi.Login, eventValue)
}

// Upload 上报日志
func (b *Bi) Upload(eventName string, eventInfo bi.EventValJson) {
	eventValue := bi.EventValJson{
		"user_info":  b.biProduct.UserInfo,
		"event_info": eventInfo,
	}
	bi.PushLogBatch(eventName, eventValue)
}

// InitProduct 初始化产销日志记录
func (b *Bi) InitProduct() bi.IBiProduct {
	if b.biProduct.Logs != nil {
		klog.Warnf("%s 当前请求存在重复初始化产销日志", b.Log())
	}
	b.biProduct.Logs = make([]*bi.ItemLog, 0)
	return b.biProduct
}

// AddProductLog 添加产销记录
func (b *Bi) AddProductLog(logs []*bi.ItemLog) {
	if b.biProduct.Logs == nil {
		//没有被初始化，则不处理
		return
	}
	b.biProduct.Logs = append(b.biProduct.Logs, logs...)
}

// AppsflyerUpload AF日志上报
func (b *Bi) AppsflyerUpload(eventName string, eventValue bi.EventValJson) {
	client := b.Client()
	eventCurrency := "USD"
	eventValueByte, _ := json.Marshal(eventValue)
	body := bi.EventValJson{
		"appsflyer_id":   client.AppsflyerId,
		"advertising_id": client.Adid,
		"bundle_id":      client.BundleId,
		"eventName":      eventName,
		"eventValue":     string(eventValueByte),
		"eventCurrency":  eventCurrency,
		"ip":             client.Ip,
		"eventTime":      time.Now().Format("2006-01-02 15:04:05.00"),
		"af_events_api":  "true",
		"os_version":     client.OsVersion,
		"idfv":           client.Idfv,
		"att":            client.Att,
		"oaid":           client.Reserved_2,
		"android_id":     client.AndroidId,
		"imei":           client.Imei,
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		klog.Error("AppsflyerUpload bodyJson err ", err)
		return
	}
	request := &pbgo.ProxyAppsflyerReq{
		Bundle: client.BundleId,
		Body:   string(bodyJson),
	}
	bi.ProxyAppsflyer(request)
}

//========================mod base========================//

func (b *Bi) InitModAfter() {
}

func (b *Bi) InitMod(gameCtx *pbreq.GameCtx) {
	b.GameCtx = gameCtx
	client := b.Client()
	b.biProduct = &bi.BiProduct{
		RoleId: b.GameCtx.GetRoleId(),
		UserInfo: bi.EventValJson{
			"bundle_id":   client.BundleId,
			"sub_package": client.SubPackage,
			"server_id":   gameCtx.GetSid(),
			"user_id":     gameCtx.GetRoleId(),
			"session_id":  client.GetSessionId(),
		},
		BaseSerialNo: util.TimeNowUnixMilli(),
	}
}

func (b *Bi) RspSync(rsp *pbreq.RspSync) {
	return
}

func NewBi() *Bi {
	return &Bi{}
}
