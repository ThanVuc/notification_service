package interface_modular

import (
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
	"notification_service/internal/interface/consumer"
	"notification_service/internal/interface/controller"
	"notification_service/internal/interface/cronjob"
	"notification_service/internal/interface/worker"
)

type InterfaceModule struct {
	ControllerModule *controller.ControllerModule
	ConsumerModule   *consumer.ConsumerModule
	WorkerModule     *worker.WorkerModule
	CronJobModule    *cronjob.CronJobModule
}

func NewInterfaceModule(
	applicationModule *application.ApplicationModule,
	infrastructureModule *infrastructure.InfrastructureModule,
) *InterfaceModule {
	controllerModule := controller.NewControllerModule(applicationModule, infrastructureModule)
	consumerModule := consumer.NewConsumerModule(applicationModule, infrastructureModule)
	workerModule := worker.NewWorkerModule(applicationModule, infrastructureModule)
	cronJobModule := cronjob.NewCronJobModule(applicationModule, infrastructureModule)

	return &InterfaceModule{
		ControllerModule: controllerModule,
		ConsumerModule:   consumerModule,
		WorkerModule:     workerModule,
		CronJobModule:    cronJobModule,
	}
}
