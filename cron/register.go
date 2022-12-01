package cron

import (
	"hat_logic/cron/task"
)

func registerCronTask() {
	cron.registerCronTask(&task.BiTask{})
}
