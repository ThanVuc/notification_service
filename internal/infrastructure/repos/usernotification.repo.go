package repos

import (
	"context"
	"notification_service/internal/core/constant"
	"notification_service/internal/core/entity"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type userNotificationRepo struct {
	mongoConnector *mongolib.MongoConnector
	logger         log.Logger
}

func (r *userNotificationRepo) UpsertUserNotification(ctx context.Context, user *entity.User) error {
	collection := r.mongoConnector.GetCollection(constant.CollectionUser)
	filter := bson.M{
		"user_id":   user.UserID,
		"device_id": user.DeviceID,
	}
	update := bson.M{
		"$set": bson.M{
			"fcm_token":  user.FCMToken,
			"updated_at": user.UpdatedAt,
		},
		"$setOnInsert": bson.M{
			"created_at": user.CreatedAt,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update, options.UpdateOne().SetUpsert(true))
	return err
}
