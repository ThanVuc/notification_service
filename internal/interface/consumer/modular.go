package consumer

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
)

type ConsumerModule struct {
	NotificationConsumer *ScheduledNotificationConsumer
	CronJobConsumer      *CronJobConsumer
}

func NewConsumerModule(
	applicationModuel *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *ConsumerModule {
	notificationConsumer := NewScheduledNotificationConsumer(
		applicationModuel.UsecaseModular.NotificationUseCase,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.BaseModule.EventBusConnector,
	)

	cronJobConsumer := NewCronJobConsumer(
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.BaseModule.EventBusConnector,
		applicationModuel.UsecaseModular.CronJobUseCase,
	)

	return &ConsumerModule{
		NotificationConsumer: notificationConsumer,
		CronJobConsumer:      cronJobConsumer,
	}
}
