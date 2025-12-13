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
		FromNotificationEntitiesToWorkNotificationsProto(notifs []*entity.Notification) []*notification_service.WorkNotification
	}
)

func NewNotificationMapper() NotificationMapper {
	return &notificationMapper{}
}
