package application

import (
	"notification_service/internal/application/helper"
	"notification_service/internal/application/mapper"
	"notification_service/internal/application/usecase"
	"notification_service/internal/infrastructure"
)

type ApplicationModule struct {
	MapperModular  *mapper.MapperModule
	UsecaseModular *usecase.UsecaseModule
	HelperModule   *helper.HelperModule
}

func NewApplicationModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *ApplicationModule {
	mapperModule := mapper.NewMapperModule()
	usecaseModule := usecase.NewUsecaseModule(infrastructureModule)
	helperModule := helper.NewMapperModule(infrastructureModule)

	return &ApplicationModule{
		MapperModular:  mapperModule,
		UsecaseModular: usecaseModule,
		HelperModule:   helperModule,
	}
}
