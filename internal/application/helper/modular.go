package helper

import "notification_service/internal/infrastructure"

type HelperModule struct {
	EmailHelper *EmailHelper
}

func NewMapperModule(
	infrastructureModule *infrastructure.InfrastructureModule,
) *HelperModule {
	emailHelper := NewEmailHelper(
		infrastructureModule.BaseModule.EmailDialer,
		infrastructureModule.BaseModule.Logger,
	)
	return &HelperModule{
		EmailHelper: emailHelper,
	}
}
