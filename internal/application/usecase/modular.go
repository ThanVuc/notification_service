package usecase

import "notification_service/internal/infrastructure"

type UsecaseModule struct {
	NotificationUseCase NotificationUseCase
}

func NewUsecaseModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *UsecaseModule {
	notificationUseCase := NewNotificationUseCase(
		infrastructureModule.BaseModule.MongoConnector,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.RepoModule.NotificationRepo,
	)
	return &UsecaseModule{
		NotificationUseCase: notificationUseCase,
	}
}
