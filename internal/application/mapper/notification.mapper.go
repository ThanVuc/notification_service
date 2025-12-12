package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type notificationMapper struct{}

func (n *notificationMapper) FromProtoListToEntities(notif []*common.Notification, isPublished bool) []*entity.Notification {
	notifications := make([]*entity.Notification, 0, len(notif))
	for _, notif := range notif {
		objectId, err := bson.ObjectIDFromHex(*notif.Id)

		if err != nil {
			objectId = bson.NewObjectID()
		}

		nowTimestampt := time.Now().UnixMilli()
		notifications = append(notifications, &entity.Notification{
			ID:              objectId,
			Title:           notif.Title,
			Message:         notif.Message,
			SenderId:        notif.SenderId,
			ReceiverIds:     notif.ReceiverIds,
			IsRead:          notif.IsRead,
			CreatedAt:       utils.FromTimeStampToTime(nowTimestampt),
			UpdatedAt:       utils.FromTimeStampToTime(nowTimestampt),
			Link:            notif.Link,
			TriggerAt:       utils.FromTimeStampToTimePtr(notif.TriggerAt),
			ImgUrl:          notif.ImageUrl,
			CorrelationId:   notif.CorrelationId,
			CorrelationType: notif.CorrelationType,
			IsPublished:     isPublished,
			IsSendMail:      notif.IsSendMail,
			IsActive:        notif.IsActive,
		})
	}
	return notifications
}

func (n *notificationMapper) FromEntitiesToProtoList(notifications []*entity.Notification) []*notification_service.Notification {
	protoNotifications := make([]*notification_service.Notification, 0, len(notifications))
	for _, notification := range notifications {
		protoNotifications = append(protoNotifications, n.FromEntityToProto(notification))
	}
	return protoNotifications
}

func (n *notificationMapper) FromEntityToProto(notification *entity.Notification) *notification_service.Notification {
	return &notification_service.Notification{
		Id:              notification.ID.Hex(),
		Title:           notification.Title,
		Message:         notification.Message,
		SenderId:        notification.SenderId,
		ReceiverIds:     notification.ReceiverIds,
		IsRead:          notification.IsRead,
		IsActive:        notification.IsActive,
		IsSendMail:      notification.IsSendMail,
		Link:            notification.Link,
		TriggerAt:       utils.FromTimePtrToTimeStamp(notification.TriggerAt),
		ImageUrl:        notification.ImgUrl,
		CorrelationId:   notification.CorrelationId,
		CorrelationType: notification.CorrelationType,
	}
}
