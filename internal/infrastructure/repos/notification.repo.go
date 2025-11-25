package repos

import (
	"context"
	"fmt"
	"notification_service/internal/core/constant"
	"notification_service/internal/core/entity"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"
	"time"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type notificationRepo struct {
	mongoConnector *mongolib.MongoConnector
	logger         log.Logger
}

func (r *notificationRepo) GetNotificationsByRecipientID(request *common.IDRequest) (*notification_service.GetNotificationsResponse, error) {
	// Implementation goes here
	return nil, nil
}

func (r *notificationRepo) SaveNotification(ctx context.Context, notification *entity.Notification) error {
	collection := r.mongoConnector.GetCollection(constant.CollectionNotification)
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": notification.ID},
		bson.M{"$set": notification},
		options.UpdateOne().SetUpsert(true),
	)

	return err
}

func (r *notificationRepo) GetNotificationsWithinTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*entity.Notification, error) {
	println(startTime.String())
	println(endTime.String())
	collection := r.mongoConnector.GetCollection(constant.CollectionNotification)
	filter := bson.M{
		"trigger_at": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
		"is_active": true,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []*entity.Notification
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}

	// Mark notifications as inactive after fetching
	notificationIDs := make([]string, len(notifications))
	for i, notification := range notifications {
		notificationIDs[i] = notification.ID
	}

	if err := r.InvalidateNotifications(ctx, notificationIDs); err != nil {
		return nil, fmt.Errorf("failed to invalidate notifications: %w", err)
	}

	return notifications, nil
}

func (r *notificationRepo) InvalidateNotifications(ctx context.Context, notificationIDs []string) error {
	collection := r.mongoConnector.GetCollection(constant.CollectionNotification)
	_, err := collection.UpdateMany(
		ctx,
		bson.M{"_id": bson.M{"$in": notificationIDs}},
		bson.M{"$set": bson.M{"is_active": false}},
	)

	return err
}
