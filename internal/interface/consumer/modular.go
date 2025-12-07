package consumer

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
)

type ConsumerModule struct {
	NotificationConsumer *ScheduledNotificationConsumer
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

	return &ConsumerModule{
		NotificationConsumer: notificationConsumer,
	}
}
