package worker

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
)

type WorkerModule struct {
	ScheduledNotificationWorker *ScheduledNotificationWorker
}

func NewWorkerModule(
	applicationModule *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *WorkerModule {
	ScheduledNotificationWorker := NewScheduledNotificationWorker(
		applicationModule.UsecaseModular.ScheduledWorkerUseCase,
		infrastructureModule.BaseModule.Logger,
	)
	return &WorkerModule{
		ScheduledNotificationWorker: ScheduledNotificationWorker,
	}
}
