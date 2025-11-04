package interface_modular

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
	"notification_service/internal/interface/consumer"
	"notification_service/internal/interface/controller"
)

type InterfaceModule struct {
	ControllerModule *controller.ControllerModule
	ConsumerModule   *consumer.ConsumerModule
}

func NewInterfaceModule(
	applicationModule *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *InterfaceModule {
	controllerModule := controller.NewControllerModule(applicationModule, infrastructureModule)
	consumerModule := consumer.NewConsumerModule(applicationModule, infrastructureModule)

	return &InterfaceModule{
		ControllerModule: controllerModule,
		ConsumerModule:   consumerModule,
	}
}
