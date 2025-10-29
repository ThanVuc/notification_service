package controller

import "notification_service/internal/domain"

type ControllerModule struct {
}

func NewControllerModule(
	domainModule *domain.DomainModule,
) *ControllerModule {
	return &ControllerModule{}
}
