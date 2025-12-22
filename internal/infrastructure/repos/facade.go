package repos

import (
	"context"
	"notification_service/internal/core/entity"
	"notification_service/proto/common"
	"time"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	NotificationRepo interface {
		GetNotificationsByRecipientID(ctx context.Context, request *common.IDRequest) ([]*entity.Notification, error)
		UpsertNotifications(ctx context.Context, notifications []*entity.Notification) error
		GetNotificationsWithinTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*entity.Notification, error)
		InvalidateNotifications(ctx context.Context, notificationIDs []bson.ObjectID) error
		MarkIsPublished(ctx context.Context, notificationID []bson.ObjectID) error
		MarkNotificationsAsRead(ctx context.Context, notificationID []bson.ObjectID) error
		DeleteNotificationById(ctx context.Context, notificationID bson.ObjectID) error
		GetNotificationByWorkId(ctx context.Context, workId string) ([]*entity.Notification, error)
		DeleteOldNotifications(ctx context.Context, before time.Time) error
	}

	UserNotificationRepo interface {
		UpsertUserNotification(ctx context.Context, user *entity.User) error
		GetUsersByIDs(ctx context.Context, userIDs []string) ([]*entity.User, error)
	}
)

func NewNotificationRepo(
	mongoConnector *mongolib.MongoConnector,
	logger log.Logger,
) NotificationRepo {
	return &notificationRepo{
		mongoConnector: mongoConnector,
		logger:         logger,
	}
}

func NewUserNotificationRepo(
	mongoConnector *mongolib.MongoConnector,
	logger log.Logger,
) UserNotificationRepo {
	return &userNotificationRepo{
		mongoConnector: mongoConnector,
		logger:         logger,
	}
}
