package usecase

import (
	"context"
	"notification_service/internal/infrastructure/repos"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
)

type (
	NotificationUseCase interface {
		GetNotifications(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsResponse, error)
	}
)

func NewNotificationUseCase(
	mongodbConnector *mongolib.MongoConnector,
	logger log.Logger,
	notificationRepo repos.NotificationRepo,
) NotificationUseCase {
	return &notificationUseCase{
		mongodbConnector: mongodbConnector,
		logger:           logger,
		notificationRepo: notificationRepo,
	}
}
