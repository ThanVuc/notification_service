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
	ID              bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title           string        `bson:"title" json:"title"`
	Message         string        `bson:"message" json:"message"`
	Link            *string       `bson:"link" json:"link"`
	SenderId        string        `bson:"sender_id" json:"sender_id"`
	ReceiverIds     []string      `bson:"receiver_ids" json:"receiver_ids"`
	IsRead          bool          `bson:"is_read" json:"is_read"`
	TriggerAt       *time.Time    `bson:"trigger_at" json:"trigger_at"`
	ImgUrl          *string       `bson:"img_url" json:"img_url"`
	IsEmailSent     bool          `bson:"is_email_sent" json:"is_email_sent"`
	IsActive        bool          `bson:"is_active" json:"is_active"`
	CreatedAt       time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time     `bson:"updated_at" json:"updated_at"`
	CorrelationId   string        `bson:"correlation_id,omitempty" json:"correlation_id,omitempty"`
	CorrelationType string        `bson:"correlation_type,omitempty" json:"correlation_type,omitempty"`
	IsPublished     bool          `bson:"is_published" json:"is_published"`
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
					"bsonType":    []string{"objectId"},
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
				"is_published": bson.M{
					"bsonType":    "bool",
					"description": "Published status",
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
