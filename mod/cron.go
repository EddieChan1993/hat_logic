package mod

/**
玩家定时任务相关
*/
import (
	"context"
	"git.dhgames.cn/svr_comm/kite/utils/klog"
	"google.golang.org/protobuf/proto"
	"hat_logic/core/cst"
	"hat_logic/pbgo/pbreq"
	"hat_logic/rpc"
	"hat_logic/util"
	"time"
)

const cronPerSec = 3 * time.Second //秒级间隔

type CronTasks struct {
	*pbreq.GameCtx
}
type IModCronTasksFunc interface {
	RunTasks()
}

func CronTasksMod(gameCtx *pbreq.GameCtx) IModCronTasksFunc {
	ins, had := gameCtx.GetApiFunc(util.GetModName(&CronTasks{}))
	if had {
		return ins.(IModCronTasksFunc)
	}
	return nil
}

//======================== cron ========================//

//RunTasks 执行所有任务
func (c *CronTasks) RunTasks() {
	c.NewTask(cronPerSec, func() {})
}

//======================== task ========================//

//========================tool========================//

//Push 发送推送消息
func (c *CronTasks) Push(grp, cmd cst.GrpIdTyp, msg proto.Message) {
	rpc.SendPush(c.GetSessionId(), grp, cmd, msg)
}

func (c *CronTasks) NewTask(cronTick time.Duration, fn pbreq.TaskFn) {
	c.GoRun(func(ctx context.Context) {
		ticker := time.NewTicker(cronTick)
		defer func() {
			if panicErr := recover(); panicErr != nil {
				klog.Error(c.Log(), panicErr)
				util.PanicStack()
			}
			ticker.Stop()
		}()
		for {
			select {
			case <-c.GetCtx().Done():
				return
			case <-ticker.C:
				c.GameCtx.SendTask(fn)
			}
		}
	})
}

//========================mod base========================//

func (c *CronTasks) InitMod(gameCtx *pbreq.GameCtx) {
	c.GameCtx = gameCtx
}

func (c *CronTasks) RspSync(rsp *pbreq.RspSync) {
	return
}

func NewCronTasks() *CronTasks {
	return &CronTasks{}
}
