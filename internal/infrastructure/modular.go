package infrastructure

import (
	"notification_service/internal/infrastructure/base"
	"notification_service/internal/infrastructure/repos"
	"notification_service/pkg/settings"
)

type InfrastructureModule struct {
	BaseModule *base.BaseModule
	RepoModule *repos.RepoModule
}

func NewInfrastructure(
	configuration *settings.Configuration,
) *InfrastructureModule {
	baseModule := base.NewBaseModule(configuration)
	repoModule := repos.NewRepoModule(baseModule)

	return &InfrastructureModule{
		BaseModule: baseModule,
		RepoModule: repoModule,
	}
}
