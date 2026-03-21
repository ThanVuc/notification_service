package worker

import (
	"context"
	"notification_service/internal/application/usecase"
	"notification_service/pkg/utils"

	"github.com/thanvuc/go-core-lib/log"
)

type EmailNotificationWorker struct {
	useCase usecase.NotificationUseCase
	logger  log.Logger
}

func NewEmailNotificationWorker(
	useCase usecase.NotificationUseCase,
	logger log.Logger,
) *EmailNotificationWorker {
	return &EmailNotificationWorker{
		useCase: useCase,
		logger:  logger,
	}
}

func (w *EmailNotificationWorker) Start(ctx context.Context) {
	w.logger.Info("Email Notification Worker is starting...", "")
	utils.WithSafePanicSimple(ctx, w.logger, w.useCase.SendEmailNotifications)
}
