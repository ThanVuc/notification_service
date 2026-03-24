package worker

import (
	"context"
	"notification_service/internal/application/usecase"
	"notification_service/pkg/utils"

	"github.com/thanvuc/go-core-lib/log"
)

type DirectEmailWorker struct {
	useCase usecase.NotificationUseCase
	logger  log.Logger
}

func NewDirectEmailWorker(
	useCase usecase.NotificationUseCase,
	logger log.Logger,
) *DirectEmailWorker {
	return &DirectEmailWorker{
		useCase: useCase,
		logger:  logger,
	}
}

func (w *DirectEmailWorker) Start(ctx context.Context) {
	w.logger.Info("Direct Email Worker is starting...", "")
	utils.WithSafePanicSimple(ctx, w.logger, w.useCase.SendDirectEmailNotifications)
}
