package global

import (
	"context"
	"notification_service/internal/domain"
	"notification_service/internal/infrastructure"
	"notification_service/internal/infrastructure/base"
	"notification_service/internal/infrastructure/server"
	interface_modular "notification_service/internal/interface"
	"notification_service/pkg/settings"
	"sync"
)

type DIContainer struct {
	configuration        *settings.Configuration
	domainModule         *domain.DomainModule
	infrastructureModule *infrastructure.InfrastructureModule
	interfaceModule      *interface_modular.InterfaceModule
}

func NewDIContainer() *DIContainer {
	configuration := base.LoadConfiguration()
	infrastructureModule := infrastructure.NewInfrastructure(configuration)
	domainModule := domain.NewDomainModular(infrastructureModule)
	interfaceModule := interface_modular.NewInterfaceModule(domainModule)

	return &DIContainer{
		domainModule:         domainModule,
		infrastructureModule: infrastructureModule,
		interfaceModule:      interfaceModule,
		configuration:        configuration,
	}
}

func (c *DIContainer) GetDomainModule() *domain.DomainModule {
	return c.domainModule
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

	notificationServer.RunServers(ctx, wg)
}

func (c *DIContainer) StartComsumerWorkers(ctx context.Context, wg *sync.WaitGroup) {
	consumerWorker := server.NewConsumerWorker(c.interfaceModule)
	consumerWorker.RunConsumers()
}

func (c *DIContainer) GracefulShutdown(wg *sync.WaitGroup) {
	server.GracefulShutdown(
		wg,
		c.infrastructureModule.BaseModule.Logger,
		*c.infrastructureModule.BaseModule.MongoConnector,
		*c.infrastructureModule.BaseModule.CacheConnector,
		*c.infrastructureModule.BaseModule.EventBusConnector,
	)
}
