package bi

/**
玩家自身产销bi
*/
import "git.dhgames.cn/svr_comm/kite/utils/cast"

type BiChannelId = int32 //渠道id

type IBiProduct interface {
	UploadProduct(channel BiChannelId, extend EventValJson)
}

//BiProduct 玩家自身产销
type BiProduct struct {
	Logs         []*ItemLog
	RoleId       int64
	UserInfo     EventValJson
	BaseSerialNo int64
}

type ItemLog struct {
	ProductNum int64
	ProductId  string
	LogType    int
	CreateTs   string
	Balance    int64
}

func (b *BiProduct) UploadProduct(channel BiChannelId, extend EventValJson) {
	logs := make([]EventValJson, 0, len(b.Logs))
	for _, log := range b.Logs {
		single := EventValJson{
			"product_num": log.ProductNum,
			"product_id":  log.ProductId,
			"log_type":    log.LogType,
			"extend":      extend,
			"create_ts":   log.CreateTs,
			"channel":     channel,
			"balance":     log.Balance,
		}
		logs = append(logs, single)
	}
	if len(logs) > 0 {
		eventInfo := EventValJson{
			"serial_no": cast.ToString(b.RoleId),
			"logs":      logs,
		}
		eventValue := EventValJson{
			"user_info":  b.UserInfo,
			"event_info": eventInfo,
		}
		PushLogBatch(Product, eventValue)
	}
	b.Logs = nil
	b.BaseSerialNo++
}
