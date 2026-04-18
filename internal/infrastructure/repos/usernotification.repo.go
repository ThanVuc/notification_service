package repos

import (
	"context"
	"notification_service/internal/core/constant"
	"notification_service/internal/core/entity"
	"time"

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

	now := time.Now()

	// 1. Remove token from other users
	_, err := collection.DeleteMany(ctx, bson.M{
		"fcm_token": user.FCMToken,
		"user_id":   bson.M{"$ne": user.UserID},
	})
	if err != nil {
		return err
	}

	// 2. Upsert by fcm_token
	filter := bson.M{
		"fcm_token": user.FCMToken,
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":    user.UserID,
			"device_id":  user.DeviceID,
			"email":      user.Email,
			"updated_at": now,
		},
		"$setOnInsert": bson.M{
			"created_at": user.CreatedAt,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update, options.UpdateOne().SetUpsert(true))
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

func (r *userNotificationRepo) GetUsersByID(ctx context.Context, userID string) (*entity.User, error) {
	collection := r.mongoConnector.GetCollection(constant.CollectionUser)
	filter := bson.M{
		"user_id": userID,
	}
	var user entity.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
