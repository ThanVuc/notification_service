package worker

import (
	"context"
	"notification_service/internal/application/usecase"
	"notification_service/pkg/utils"

	"github.com/thanvuc/go-core-lib/log"
)

type ScheduledNotificationWorker struct {
	scheduledWorkerUseCase usecase.ScheduledWorkerUseCase
	logger                 log.Logger
}

func NewScheduledNotificationWorker(
	scheduledWorkerUseCase usecase.ScheduledWorkerUseCase,
	logger log.Logger,
) *ScheduledNotificationWorker {
	return &ScheduledNotificationWorker{
		scheduledWorkerUseCase: scheduledWorkerUseCase,
		logger:                 logger,
	}
}

func (w *ScheduledNotificationWorker) RunScheduledNotifications(ctx context.Context) {
	utils.WithSafePanicSimple(ctx, w.logger, w.scheduledWorkerUseCase.ProcessScheduledNotifications)
}
