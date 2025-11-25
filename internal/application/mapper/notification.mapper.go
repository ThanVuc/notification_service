package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
)

type notificationMapper struct{}

func (n *notificationMapper) FromProtoToEntity(protoNotification *common.Notification) *entity.Notification {
	return &entity.Notification{
		ID:              *protoNotification.Id,
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
