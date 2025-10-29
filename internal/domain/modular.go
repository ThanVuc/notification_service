package domain

import (
	"notification_service/internal/domain/mapper"
	"notification_service/internal/domain/usecase"
	"notification_service/internal/infrastructure"
)

type DomainModule struct {
	MapperModular *mapper.MapperModule
	UserModular   *usecase.UsecaseModule
}

func NewDomainModular(
	infrastructureModule *infrastructure.InfrastructureModule,
) *DomainModule {
	mapperModule := mapper.NewMapperModule()
	usecaseModule := usecase.NewUsecaseModule(infrastructureModule)

	return &DomainModule{
		MapperModular: mapperModule,
		UserModular:   usecaseModule,
	}
}
