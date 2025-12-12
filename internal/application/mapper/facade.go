package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"
)

type (
	NotificationMapper interface {
		FromProtoListToEntities(notif []*common.Notification, isPublished bool) []*entity.Notification
		FromEntitiesToProtoList(notifications []*entity.Notification) []*notification_service.Notification
	}
)

func NewNotificationMapper() NotificationMapper {
	return &notificationMapper{}
}
