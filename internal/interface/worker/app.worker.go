package worker

import (
	"context"
	"notification_service/internal/application/usecase"
	"notification_service/pkg/utils"

	"github.com/thanvuc/go-core-lib/log"
)

type AppNotificationWorker struct {
	useCase usecase.NotificationUseCase // Giả sử bạn có usecase xử lý logic gửi app
	logger  log.Logger
}

func NewAppNotificationWorker(
	useCase usecase.NotificationUseCase,
	logger log.Logger,
) *AppNotificationWorker {
	return &AppNotificationWorker{
		useCase: useCase,
		logger:  logger,
	}
}

func (w *AppNotificationWorker) Start(ctx context.Context) {
	w.logger.Info("App Notification Worker is starting...", "")
	utils.WithSafePanicSimple(ctx, w.logger, w.useCase.SendAppNotifications)
}
