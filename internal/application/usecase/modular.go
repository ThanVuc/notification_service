package usecase

import "notification_service/internal/infrastructure"

type UsecaseModule struct {
	NotificationUseCase     NotificationUseCase
	UserNotificationUseCase UserNotificationUseCase
}

func NewUsecaseModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *UsecaseModule {
	notificationUseCase := NewNotificationUseCase(
		infrastructureModule.BaseModule.MongoConnector,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.RepoModule.NotificationRepo,
		infrastructureModule.BaseModule.FirebaseApp,
	)

	userNotificationUseCase := NewUserNotificationUseCase(
		infrastructureModule.BaseModule.MongoConnector,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.RepoModule.UserNotificationRepo,
	)

	return &UsecaseModule{
		NotificationUseCase:     notificationUseCase,
		UserNotificationUseCase: userNotificationUseCase,
	}
}
