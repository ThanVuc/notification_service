package worker

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
)

type WorkerModule struct {
	ScheduledNotificationWorker *ScheduledNotificationWorker
	AppNotificationWorker       *AppNotificationWorker
	EmailNotificationWorker     *EmailNotificationWorker
}

func NewWorkerModule(
	applicationModule *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *WorkerModule {
	ScheduledNotificationWorker := NewScheduledNotificationWorker(
		applicationModule.UsecaseModular.ScheduledWorkerUseCase,
		infrastructureModule.BaseModule.Logger,
	)

	AppNotificationWorker := NewAppNotificationWorker(
		applicationModule.UsecaseModular.NotificationUseCase,
		infrastructureModule.BaseModule.Logger,
	)

	EmailNotificationWorker := NewEmailNotificationWorker(
		applicationModule.UsecaseModular.NotificationUseCase,
		infrastructureModule.BaseModule.Logger,
	)

	return &WorkerModule{
		ScheduledNotificationWorker: ScheduledNotificationWorker,
		AppNotificationWorker:       AppNotificationWorker,
		EmailNotificationWorker:     EmailNotificationWorker,
	}
}
