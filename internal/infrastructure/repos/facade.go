package repos

import (
	"context"
	"notification_service/internal/core/entity"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"
	"time"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	NotificationRepo interface {
		GetNotificationsByRecipientID(request *common.IDRequest) (*notification_service.GetNotificationsResponse, error)
		SaveNotification(ctx context.Context, notification *entity.Notification) error
		GetNotificationsWithinTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*entity.Notification, error)
		InvalidateNotifications(ctx context.Context, notificationIDs []bson.ObjectID) error
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
