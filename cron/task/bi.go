package task

import (
	"hat_logic/core"
	"hat_logic/tool/bi"
	"hat_logic/tool/consul"
	"time"
)

type BiTask struct {
}

func (b *BiTask) CronTime() time.Duration {
	return time.Second
}

func (b *BiTask) CronTask() {
	t := time.Now()
	if t.Minute()%2 == 0 && t.Second() == 0 {
		whoAmI := consul.WhoAmI()
		eventValue := bi.EventValJson{
			"event_info": bi.EventValJson{
				"server_id":  whoAmI.Index,
				"platform":   whoAmI.Cluster,
				"online_num": core.OnlinePeople(),
				"dump_time":  t.Format("2006-01-02 15:04:05"),
			},
		}
		bi.PushLogSingle(bi.Online, eventValue)
	}
}
