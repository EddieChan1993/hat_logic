package bi

import (
	"encoding/json"
	"fmt"
	gatePbgo "git.dhgames.cn/svr_comm/gclient/gate/pbgo"
	"git.dhgames.cn/svr_comm/gclient/utlog"
	"git.dhgames.cn/svr_comm/gclient/utlog/pbgo"
	"git.dhgames.cn/svr_comm/kite"
	"git.dhgames.cn/svr_comm/kite/utils/cast"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"hat_logic/tool/consul/static"
	"hat_logic/util"
)

type event struct {
	eventCode string
	eventName EventName
	eventType string
}

type EventValJson = map[string]interface{}
type EventName = string

// bi事件
const (
	Register   EventName = "register"
	Login                = "login"
	Logout               = "logout"
	Pay                  = "pay"
	LevelUp              = "level_up"
	ChangeName           = "change_name"
	Product              = "product"
	Player               = "player"
	Bag                  = "bag"
	Hero                 = "hero"
	Comm                 = "comm"
	Instance             = "instance"
	Mission              = "mission"
	Online               = "online"
)

var eventMap = map[EventName]event{
	Register:   {eventCode: "1100110001", eventType: "base", eventName: "register"},
	Login:      {eventCode: "1100110002", eventType: "base", eventName: "login"},
	Logout:     {eventCode: "1100110003", eventType: "base", eventName: "logout"},
	Pay:        {eventCode: "1100110004", eventType: "base", eventName: "pay"},
	LevelUp:    {eventCode: "1100110005", eventType: "base", eventName: "level_up"},
	ChangeName: {eventCode: "1100110007", eventType: "base", eventName: "change_name"},
	Product:    {eventCode: "1100210001", eventType: "product", eventName: "product"},
	Player:     {eventCode: "1100310001", eventType: "snapshot", eventName: "player"},
	Bag:        {eventCode: "1100310002", eventType: "snapshot", eventName: "bag"},
	Hero:       {eventCode: "1100310003", eventType: "snapshot", eventName: "hero"},
	Comm:       {eventCode: "1100310004", eventType: "snapshot", eventName: "comm"},
	Instance:   {eventCode: "1100410016", eventType: "game", eventName: "instance"},
	Mission:    {eventCode: "1100410025", eventType: "game", eventName: "mission"},
	Online:     {eventCode: "1100510002", eventType: "server", eventName: "online"},
}

// Appsflyer事件
const (
	AfPurchase EventName = "af_purchase"
)

func newEvent(eventName EventName, eventValue EventValJson) (*pbgo.PushLogReq, error) {
	if event, has := eventMap[eventName]; has {
		EventValueStr, err := json.Marshal(eventValue)
		if err != nil {
			return nil, err
		}
		return &pbgo.PushLogReq{
			EventCode:  event.eventCode,
			EventType:  event.eventType,
			EventName:  event.eventName,
			EventValue: string(EventValueStr),
			GameCd:     util.GameCd,
			CreateTs:   CreateTs(),
		}, nil
	}
	return nil, fmt.Errorf("unknown eventName: %s", eventName)
}

// PushLogBatch 批量推送
func PushLogBatch(eventName EventName, eventValue EventValJson) {
	if isReport, _, _ := static.StaticBi(); isReport {
		log, err := newEvent(eventName, eventValue)
		if err != nil {
			klog.Error("PushLogBatch err ", eventName, err)
			return
		}
		worker.batchRecv <- log
	}
}

// PushLogSingle 单条推送
func PushLogSingle(eventName EventName, eventValue EventValJson) {
	if isReport, _, _ := static.StaticBi(); isReport {
		log, err := newEvent(eventName, eventValue)
		if err != nil {
			klog.Error("PushLogSingle err ", eventName, err)
			return
		}
		worker.singleRecv <- log
	}
}

// ProxyAppsflyer AF推送
func ProxyAppsflyer(request *pbgo.ProxyAppsflyerReq) {
	if isReport, dc, cluster := static.StaticBi(); isReport {
		rsp, err := utlog.ProxyAppsflyer(dc, cluster, request, kite.Cast)
		if err != nil {
			klog.Error("utlog.ProxyAppsflyer err code:", err, rsp.ErrorCode)
		}
	}
}

// CreateTs 19位时间戳
func CreateTs() string {
	return cast.ToString(util.TimeNowUnixNano())
}

// =======================================

// RegisterInfo 创角日志信息
func RegisterInfo(req *gatePbgo.CreateRoleReq, rsp *gatePbgo.CreateRoleRsp) EventValJson {
	deviceInfo := EventValJson{
		"adid":         req.Client.Adid,
		"idfv":         req.Client.Idfv,
		"sm_id":        req.Client.Dhid,
		"imei":         req.Client.Imei,
		"android_id":   req.Client.AndroidId,
		"appsflyer_id": req.Client.AppsflyerId,
		"device_token": req.Client.DeviceToken,
		"mac_address":  req.Client.MacAddress,
		"device_model": req.Client.DeviceModel,
		"device_name":  req.Client.DviceName,
		"os_version":   req.Client.OsVersion,
		"language":     req.Client.Language,
		"network_type": req.Client.NetworkType,
		"app_version":  req.Client.AppVersion,
		"ip":           req.Client.Ip,
		"oaid":         req.Client.Reserved_2,
	}
	userInfo := EventValJson{
		"bundle_id":   req.Client.BundleId,
		"sub_package": req.Client.SubPackage,
		"server_id":   req.Sid,
		"user_id":     rsp.RoleId,
		"session_id":  req.Client.GetSessionId(),
		"account":     req.Account,
		"platform":    req.Client.Platform,
	}
	eventInfo := EventValJson{
		"device_info": deviceInfo,
		"user_info":   userInfo,
	}
	return eventInfo
}
