package server

import (
	"context"
	"sync"

	"github.com/thanvuc/go-core-lib/cache"
	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"go.uber.org/zap"
)

func GracefulShutdown(
	wg *sync.WaitGroup,
	logger log.Logger,
	mongoConector mongolib.MongoConnector,
	redisDb cache.RedisCache,
	eventBusConnector eventbus.RabbitMQConnector,
) {
	wg.Add(1)
	err := mongoConector.GracefulClose(context.Background(), wg)
	handleError(logger, err, "MongoDB connection closed successfully")

	wg.Add(1)
	err = redisDb.Close(wg)
	handleError(logger, err, "Redis connection closed successfully")

	wg.Add(1)
	eventBusConnector.Close(wg)

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
