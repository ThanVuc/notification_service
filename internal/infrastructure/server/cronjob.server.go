package server

import (
	"context"
	interface_modular "notification_service/internal/interface"
)

type CronJob struct {
	interfaceModule *interface_modular.InterfaceModule
}

func NewCronJob(
	interfaceModule *interface_modular.InterfaceModule,
) *CronJob {
	return &CronJob{
		interfaceModule: interfaceModule,
	}
}

func (s *CronJob) RunCronjob(ctx context.Context) {
	cronJobModule := s.interfaceModule.CronJobModule
	cronJobModule.NotificationCronJob.DeleteOldNotificationsCronJob(ctx)
}
