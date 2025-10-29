package server

import (
	"notification_service/internal/interface/controller"
	"notification_service/pkg/settings"

	"github.com/thanvuc/go-core-lib/log"
)

type ServerModule struct {
	NotificationServer *NotificationServer
}

func NewServer(
	configuration *settings.Configuration,
	logger log.Logger,
	controllerModule *controller.ControllerModule,
) *ServerModule {
	notificationServer := NewNotificationServer(configuration, logger, controllerModule)

	return &ServerModule{
		NotificationServer: notificationServer,
	}
}
