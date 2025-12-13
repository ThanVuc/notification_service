package usecase

import (
	"context"
	"notification_service/internal/application/helper"
	"notification_service/internal/application/mapper"
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
		GetNotificationsByRecipientId(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsByRecipientIdResponse, error)
		ConsumeScheduledNotification(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action
		MarkNotificationsAsRead(ctx context.Context, req *common.IDsRequest) (*common.EmptyResponse, error)
		DeleteNotificationById(ctx context.Context, req *common.IDRequest) (*common.EmptyResponse, error)
		GetNotificationByWorkId(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsByWorkIdResponse, error)
	}

	UserNotificationUseCase interface {
		UpsertUserFCMToken(ctx context.Context, req *notification_service.UpsertUserFCMTokenRequest) (*common.EmptyResponse, error)
	}

	ScheduledWorkerUseCase interface {
		ProcessScheduledNotifications(ctx context.Context) error
	}
)

func NewNotificationUseCase(
	mongodbConnector *mongolib.MongoConnector,
	logger log.Logger,
	notificationRepo repos.NotificationRepo,
	firebaseApp *firebase.App,
	notificationMapper mapper.NotificationMapper,
) NotificationUseCase {
	return &notificationUseCase{
		mongodbConnector:   mongodbConnector,
		logger:             logger,
		notificationRepo:   notificationRepo,
		firebaseApp:        firebaseApp,
		notificationMapper: notificationMapper,
	}
}

func NewUserNotificationUseCase(
	mongodbConnector *mongolib.MongoConnector,
	logger log.Logger,
	userRepo repos.UserNotificationRepo,
) UserNotificationUseCase {
	return &userNotificationUseCase{
		mongodbConnector: mongodbConnector,
		logger:           logger,
		userRepo:         userRepo,
	}
}

func NewScheduledWorkerUseCase(
	logger log.Logger,
	notificationRepo repos.NotificationRepo,
	firebaseApp *firebase.App,
	userRepo repos.UserNotificationRepo,
	emailHelper helper.EmailHelper,
) ScheduledWorkerUseCase {
	return &scheduledWorkerUsecase{
		logger:           logger,
		notificationRepo: notificationRepo,
		firebaseApp:      firebaseApp,
		userRepo:         userRepo,
		emailHelper:      emailHelper,
	}
}
