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
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func (r *notificationRepo) UpsertNotifications(ctx context.Context, notifications []*entity.Notification) error {
	collection := r.mongoConnector.GetCollection(constant.CollectionNotification)

	models := make([]mongo.WriteModel, 0, len(notifications))

	for _, n := range notifications {
		if !n.ID.IsZero() {
			n.ID = bson.NewObjectID()
		}

		filter := bson.M{"_id": n.ID}

		update := bson.M{
			"$set": bson.M{
				"title":            n.Title,
				"message":          n.Message,
				"link":             n.Link,
				"sender_id":        n.SenderId,
				"receiver_ids":     n.ReceiverIds,
				"is_read":          n.IsRead,
				"trigger_at":       n.TriggerAt,
				"img_url":          n.ImgUrl,
				"is_email_sent":    n.IsEmailSent,
				"is_active":        n.IsActive,
				"created_at":       n.CreatedAt,
				"updated_at":       n.UpdatedAt,
				"correlation_id":   n.CorrelationId,
				"correlation_type": n.CorrelationType,
				"is_published":     n.IsPublished,
			},
		}

		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)

		models = append(models, model)
	}

	if len(models) == 0 {
		return nil
	}

	_, err := collection.BulkWrite(ctx, models, options.BulkWrite().SetOrdered(false))
	return err
}

func (r *notificationRepo) GetNotificationsWithinTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*entity.Notification, error) {
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
	notificationIDs := make([]bson.ObjectID, len(notifications))
	for i, notification := range notifications {
		notificationIDs[i] = notification.ID
	}

	if err := r.InvalidateNotifications(ctx, notificationIDs); err != nil {
		return nil, fmt.Errorf("failed to invalidate notifications: %w", err)
	}

	return notifications, nil
}

func (r *notificationRepo) InvalidateNotifications(ctx context.Context, notificationIDs []bson.ObjectID) error {
	collection := r.mongoConnector.GetCollection(constant.CollectionNotification)
	_, err := collection.UpdateMany(
		ctx,
		bson.M{"_id": bson.M{"$in": notificationIDs}},
		bson.M{"$set": bson.M{"is_active": false}},
	)

	return err
}
