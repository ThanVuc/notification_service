package application

import (
	"notification_service/internal/application/mapper"
	"notification_service/internal/application/usecase"
	"notification_service/internal/infrastructure"
)

type ApplicationModule struct {
	MapperModular  *mapper.MapperModule
	UsecaseModular *usecase.UsecaseModule
}

func NewApplicationModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *ApplicationModule {
	mapperModule := mapper.NewMapperModule()
	usecaseModule := usecase.NewUsecaseModule(infrastructureModule)

	return &ApplicationModule{
		MapperModular:  mapperModule,
		UsecaseModular: usecaseModule,
	}
}
