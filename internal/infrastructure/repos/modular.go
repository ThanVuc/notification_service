package repos

import "notification_service/internal/infrastructure/base"

type RepoModule struct {
}

func NewRepoModule(
	baseModule *base.BaseModule,
) *RepoModule {
	return &RepoModule{}
}
