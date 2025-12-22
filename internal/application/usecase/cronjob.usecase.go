package usecase

import (
	"context"
	"notification_service/internal/infrastructure/repos"
	"time"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type cronJobUseCase struct {
	logger           log.Logger
	notificationRepo repos.NotificationRepo
}

func (u *cronJobUseCase) ProcessDeleteOldNotifications(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action {
	before := time.Now().AddDate(0, 0, -30) // Delete notifications older than 30 days
	requestId := d.Headers["request_id"].(string)
	err := u.notificationRepo.DeleteOldNotifications(context.Background(), before)
	if err != nil {
		u.logger.Error("Failed to delete old notifications", requestId, zap.Error(err))
		return rabbitmq.NackRequeue
	}
	u.logger.Info("Old notifications deleted successfully", requestId)
	return rabbitmq.Ack
}
