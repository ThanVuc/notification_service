package base

import "github.com/thanvuc/go-core-lib/cronjob"

func NewSchedulerManager() *cronjob.CronManager {
	return cronjob.NewCronManager()
}
