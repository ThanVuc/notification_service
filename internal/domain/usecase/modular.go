package usecase

import "notification_service/internal/infrastructure"

type UsecaseModule struct {
}

func NewUsecaseModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *UsecaseModule {
	return &UsecaseModule{}
}
