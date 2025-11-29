package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/pkg/utils"
	"notification_service/proto/common"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type notificationMapper struct{}

func (n *notificationMapper) FromProtoToEntity(protoNotification *common.Notification) *entity.Notification {
	objectId, err := bson.ObjectIDFromHex(*protoNotification.Id)

	if err != nil {
		objectId = bson.NewObjectID()
	}

	return &entity.Notification{
		ID:              objectId,
		Title:           protoNotification.Title,
		Message:         protoNotification.Message,
		SenderId:        protoNotification.SenderId,
		ReceiverIds:     protoNotification.ReceiverIds,
		IsRead:          protoNotification.IsRead,
		CreatedAt:       utils.FromTimeStampToTime(protoNotification.CreatedAt),
		UpdatedAt:       utils.FromTimeStampToTime(protoNotification.UpdateAt),
		Link:            protoNotification.Link,
		TriggerAt:       utils.FromTimeStampToTimePtr(protoNotification.TriggerAt),
		ImgUrl:          protoNotification.ImageUrl,
		CorrelationId:   protoNotification.CorrelationId,
		CorrelationType: protoNotification.CorrelationType,
	}
}
