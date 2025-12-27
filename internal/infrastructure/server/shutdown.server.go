package server

import (
	"context"
	"notification_service/internal/infrastructure/base"
	interface_modular "notification_service/internal/interface"
	"sync"

	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

func GracefulShutdown(
	wg *sync.WaitGroup,
	baseModule *base.BaseModule,
	interfaceModule *interface_modular.InterfaceModule,
) {
	wg.Add(1)
	interfaceModule.CronJobModule.CronManager.Shutdown(wg)

	wg.Add(1)
	logger := baseModule.Logger
	err := baseModule.MongoConnector.GracefulClose(context.Background(), wg)
	handleError(logger, err, "MongoDB connection closed successfully")

	wg.Add(1)
	err = baseModule.CacheConnector.Close(wg)
	handleError(logger, err, "Redis connection closed successfully")

	wg.Add(1)
	baseModule.EventBusConnector.Close(wg)

	wg.Add(1)
	err = logger.Sync(wg)
	handleError(logger, err, "Logger synced successfully")

	wg.Wait()
}

func handleError(logger log.Logger, err error, successMessage string) {
	if err != nil {
		logger.Error("An error occurred", "", zap.Error(err))
	} else {
		logger.Info(successMessage, "")
	}
}
