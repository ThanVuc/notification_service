package base

import (
	"context"
	"fmt"
	"time"

	"notification_service/internal/core/entity"
	"notification_service/pkg/settings"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
)

func NewMongoDB(
	configuration *settings.Configuration,
	logger log.Logger,
) *mongolib.MongoConnector {
	cfg := createMongoConfiguration(configuration)

	const maxRetries = 10
	const retryInterval = 6 * time.Second

	var err error
	for i := 1; i <= maxRetries; i++ {
		mongoConnector, err := mongolib.NewMongoConnector(context.Background(), cfg)
		if err == nil {
			logger.Info("MongoDB connected successfully", "")
			// create collections, validators, indexes
			err := createCollections(mongoConnector)
			if err != nil {
				logger.Error("Failed to create collections", "")
			}

			return mongoConnector
		}

		logger.Warn(fmt.Sprintf("Failed to connect to MongoDB (attempt %d/%d): %v", i, maxRetries, err), "")
		time.Sleep(retryInterval * time.Duration(i)) // Exponential backoff
	}

	logger.Error("Could not connect to MongoDB after maximum retries", "")
	panic(fmt.Sprintf("Could not connect to MongoDB after %d attempts: %v", maxRetries, err))
}

func createMongoConfiguration(configuration *settings.Configuration) mongolib.MongoConnectorConfig {
	return mongolib.MongoConnectorConfig{
		URI:      configuration.Mongo.URI,
		Database: configuration.Mongo.Database,
		Username: configuration.Mongo.Username,
		Password: configuration.Mongo.Password,
	}
}

func createCollections(
	connector *mongolib.MongoConnector,
) error {
	var errs []error

	if err := entity.CreateUserCollection(connector); err != nil {
		errs = append(errs, fmt.Errorf("user collection: %w", err))
	}
	if err := entity.CreateNotificationCollection(connector); err != nil {
		errs = append(errs, fmt.Errorf("notification collection: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to create collections: %v", errs)
	}
	return nil
}
