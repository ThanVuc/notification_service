package usecase

import (
	"context"
	"notification_service/internal/infrastructure/repos"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	firebase "firebase.google.com/go/v4"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"github.com/wagslane/go-rabbitmq"
)

type (
	NotificationUseCase interface {
		GetNotifications(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsResponse, error)
		ConsumeScheduledNotification(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action
	}
)

func NewNotificationUseCase(
	mongodbConnector *mongolib.MongoConnector,
	logger log.Logger,
	notificationRepo repos.NotificationRepo,
	firebaseApp *firebase.App,
) NotificationUseCase {
	return &notificationUseCase{
		mongodbConnector: mongodbConnector,
		logger:           logger,
		notificationRepo: notificationRepo,
		firebaseApp:      firebaseApp,
	}
}
