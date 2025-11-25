package mapper

import (
	"notification_service/internal/core/entity"
	"notification_service/proto/common"
)

type (
	NotificationMapper interface {
		FromProtoToEntity(protoNotification *common.Notification) *entity.Notification
	}
)

func NewNotificationMapper() NotificationMapper {
	return &notificationMapper{}
}
