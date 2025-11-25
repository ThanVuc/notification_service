package usecase

import (
	"notification_service/internal/application/mapper"
	"notification_service/internal/infrastructure"
)

type UsecaseModule struct {
	NotificationUseCase     NotificationUseCase
	UserNotificationUseCase UserNotificationUseCase
	ScheduledWorkerUseCase  ScheduledWorkerUseCase
}

func NewUsecaseModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *UsecaseModule {
	mapperModule := mapper.NewMapperModule()

	notificationUseCase := NewNotificationUseCase(
		infrastructureModule.BaseModule.MongoConnector,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.RepoModule.NotificationRepo,
		infrastructureModule.BaseModule.FirebaseApp,
		mapperModule.NotificationMapper,
	)

	userNotificationUseCase := NewUserNotificationUseCase(
		infrastructureModule.BaseModule.MongoConnector,
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.RepoModule.UserNotificationRepo,
	)

	scheduledWorkerUseCase := NewScheduledWorkerUseCase(
		infrastructureModule.BaseModule.Logger,
		infrastructureModule.RepoModule.NotificationRepo,
		infrastructureModule.BaseModule.FirebaseApp,
		infrastructureModule.RepoModule.UserNotificationRepo,
	)

	return &UsecaseModule{
		NotificationUseCase:     notificationUseCase,
		UserNotificationUseCase: userNotificationUseCase,
		ScheduledWorkerUseCase:  scheduledWorkerUseCase,
	}
}
