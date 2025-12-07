package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/proto/common"
)

type (
	NotificationMapper interface {
		FromProtoListToEntities(notif []*common.Notification, isPublished bool) []*entity.Notification
	}
)

func NewNotificationMapper() NotificationMapper {
	return &notificationMapper{}
}
