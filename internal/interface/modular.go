package interface_modular

import (
	"notification_service/internal/domain"
	"notification_service/internal/interface/consumer"
	"notification_service/internal/interface/controller"
)

type InterfaceModule struct {
	ControllerModule *controller.ControllerModule
	ConsumerModule   *consumer.ConsumerModule
}

func NewInterfaceModule(
	domainModule *domain.DomainModule,
) *InterfaceModule {
	controllerModule := controller.NewControllerModule(domainModule)
	consumerModule := consumer.NewConsumerModule(domainModule)

	return &InterfaceModule{
		ControllerModule: controllerModule,
		ConsumerModule:   consumerModule,
	}
}
