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

type Notification struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Message     string    `bson:"message" json:"message"`
	Link        string    `bson:"link" json:"link"`
	SenderId    string    `bson:"sender_id" json:"sender_id"`
	ReceiverIds []string  `bson:"receiver_ids" json:"receiver_ids"`
	IsRead      bool      `bson:"is_read" json:"is_read"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

func (n *Notification) CollectionName() string {
	return constant.CollectionNotification
}

func CreateNotificationCollection(
	connector *mongolib.MongoConnector,
) error {
	ctx := context.Background()

	notificationValidator := bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"_id", "message", "sender_id", "receiver_ids", "created_at", "updated_at"},
			"properties": bson.M{
				"_id": bson.M{
					"bsonType":    "string",
					"description": "ID",
				},
				"message": bson.M{
					"bsonType":    "string",
					"description": "Notification message",
				},
				"link": bson.M{
					"bsonType":    "string",
					"description": "Notification link",
				},
				"sender_id": bson.M{
					"bsonType":    "string",
					"description": "Sender ID",
				},
				"receiver_ids": bson.M{
					"bsonType":    "array",
					"description": "Receiver IDs",
					"items": bson.M{
						"bsonType": "string",
					},
				},
				"is_read": bson.M{
					"bsonType":    "bool",
					"description": "Read status",
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

	notificationIdxs := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "receiver_ids", Value: 1}},
			Options: options.Index().SetName("idx_receiver_ids"),
		},
		{
			Keys:    bson.D{{Key: "is_read", Value: 1}},
			Options: options.Index().SetName("idx_is_read"),
		},
		{
			Keys:    bson.D{{Key: "created_at", Value: -1}},
			Options: options.Index().SetName("idx_created_at"),
		},
	}

	return connector.CreateCollection(ctx, constant.CollectionNotification, notificationValidator, notificationIdxs)
}
