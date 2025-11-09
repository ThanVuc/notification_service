package controller

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
)

type ControllerModule struct {
	NotificationController     *NotificationController
	UserNotificationController *UserNotificationController
}

func NewControllerModule(
	applicationModule *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *ControllerModule {
	logger := infrastructureModule.BaseModule.Logger
	NotificationController := NewNotificationController(
		applicationModule.UsecaseModular.NotificationUseCase,
		logger,
	)

	UserNotificationController := NewUserNotificationController(
		applicationModule.UsecaseModular.UserNotificationUseCase,
		logger,
	)

	return &ControllerModule{
		NotificationController:     NotificationController,
		UserNotificationController: UserNotificationController,
	}
}
