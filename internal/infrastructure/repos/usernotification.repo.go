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
			"email":      user.Email,
		},
		"$setOnInsert": bson.M{
			"created_at": user.CreatedAt,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update, options.UpdateOne().SetUpsert(true))
	return err
}

func (r *userNotificationRepo) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*entity.User, error) {
	collection := r.mongoConnector.GetCollection(constant.CollectionUser)
	filter := bson.M{
		"user_id": bson.M{
			"$in": userIDs,
		},
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*entity.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
