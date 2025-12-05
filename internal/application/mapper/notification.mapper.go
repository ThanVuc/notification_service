package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
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
		println(notif.GetTitle()+": ", *notif.TriggerAt)
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
			IsEmailSent:     notif.IsEmailSent,
			IsActive:        notif.IsActive,
		})
	}
	return notifications
}
