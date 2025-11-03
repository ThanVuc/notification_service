package repos

import "notification_service/internal/infrastructure/base"

type RepoModule struct {
	NotificationRepo NotificationRepo
}

func NewRepoModule(
	baseModule *base.BaseModule,
) *RepoModule {
	notificationRepo := NewNotificationRepo(
		baseModule.MongoConnector,
		baseModule.Logger,
	)

	return &RepoModule{
		NotificationRepo: notificationRepo,
	}
}
