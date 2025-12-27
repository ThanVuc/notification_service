package cronjob

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"

	"github.com/thanvuc/go-core-lib/cronjob"
)

type CronJobModule struct {
	CronManager         *cronjob.CronManager
	NotificationCronJob *NotificationCronJob
}

func NewCronJobModule(
	applicationModule *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *CronJobModule {
	notificationCronJob := NewNotificationCronJob(
		applicationModule.UsecaseModular.NotificationUseCase,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.BaseModule.CacheConnector,
		infrastructureModule.BaseModule.CronManager,
	)

	return &CronJobModule{
		CronManager:         infrastructureModule.BaseModule.CronManager,
		NotificationCronJob: notificationCronJob,
	}
}
