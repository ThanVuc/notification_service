package entity

import (
	"context"
	"notification_service/internal/core/constant"
	"time"

	"github.com/thanvuc/go-core-lib/mongolib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string        `bson:"user_id" json:"user_id"`
	FCMToken  string        `bson:"fcm_token" json:"fcm_token"`
	DeviceID  string        `bson:"device_id" json:"device_id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

func (u *User) CollectionName() string {
	return constant.CollectionUser
}

func CreateUserCollection(
	connector *mongolib.MongoConnector,
) error {
	ctx := context.Background()

	userValidator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"_id", "user_id", "fcm_token", "device_id", "created_at", "updated_at"},
			"properties": bson.M{
				"_id": bson.M{
					"bsonType":    []string{"string", "objectId"},
					"description": "ID",
				},
				"user_id": bson.M{
					"bsonType":    "string",
					"description": "user ID",
				},
				"fcm_token": bson.M{
					"bsonType":    "string",
					"description": "FCM token",
				},
				"device_id": bson.M{
					"bsonType":    "string",
					"description": "Device ID",
				},
				"created_at": bson.M{
					"bsonType":    "date",
					"description": "Creation timestamp",
				},
				"updated_at": bson.M{
					"bsonType":    "date",
					"description": "Last update timestamp",
				},
			},
		},
	}

	userIdxs := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "device_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	return connector.CreateCollection(ctx, constant.CollectionUser, userValidator, userIdxs)
}
