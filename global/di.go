package global

import (
	"context"
	"notification_service/internal/application"
	"notification_service/internal/infrastructure"
	"notification_service/internal/infrastructure/base"
	"notification_service/internal/infrastructure/server"
	interface_modular "notification_service/internal/interface"
	"notification_service/pkg/settings"
	"sync"
)

type DIContainer struct {
	configuration        *settings.Configuration
	applicationModule    *application.ApplicationModule
	infrastructureModule *infrastructure.InfrastructureModule
	interfaceModule      *interface_modular.InterfaceModule
}

func NewDIContainer() *DIContainer {
	configuration := base.LoadConfiguration()
	infrastructureModule := infrastructure.NewInfrastructure(configuration)
	applicationModule := application.NewApplicationModule(infrastructureModule)
	interfaceModule := interface_modular.NewInterfaceModule(applicationModule, infrastructureModule)

	return &DIContainer{
		applicationModule:    applicationModule,
		infrastructureModule: infrastructureModule,
		interfaceModule:      interfaceModule,
		configuration:        configuration,
	}
}

func (c *DIContainer) GetApplicationModule() *application.ApplicationModule {
	return c.applicationModule
}

func (c *DIContainer) GetInfrastructureModule() *infrastructure.InfrastructureModule {
	return c.infrastructureModule
}

func (c *DIContainer) GetInterfaceModule() *interface_modular.InterfaceModule {
	return c.interfaceModule
}

func (c *DIContainer) StartGrpcServer(ctx context.Context, wg *sync.WaitGroup) {
	notificationServer := server.NewNotificationServer(
		c.configuration,
		c.infrastructureModule.BaseModule.Logger,
		c.interfaceModule.ControllerModule,
	)
	c.infrastructureModule.BaseModule.Logger.Info("gRPC server started", "")
	notificationServer.RunServers(ctx, wg)
}

func (c *DIContainer) StartComsumerWorkers(ctx context.Context, wg *sync.WaitGroup) {
	consumerWorker := server.NewConsumerWorker(c.interfaceModule)
	consumerWorker.RunConsumers(ctx)
}

func (c *DIContainer) StartWorker(ctx context.Context, wg *sync.WaitGroup) {
	workerServer := server.NewWorker(c.interfaceModule)
	c.infrastructureModule.BaseModule.Logger.Info("Worker started", "")
	workerServer.RunWorkers()
}

func (c *DIContainer) StartCronJobs(ctx context.Context, wg *sync.WaitGroup) {
	cronJobServer := server.NewCronJob(c.interfaceModule)
	c.infrastructureModule.BaseModule.Logger.Info("Cronjob started", "")
	cronJobServer.RunCronjob(ctx)
}

func (c *DIContainer) GracefulShutdown(wg *sync.WaitGroup) {
	server.GracefulShutdown(
		wg,
		c.infrastructureModule.BaseModule,
		c.interfaceModule,
	)
}
