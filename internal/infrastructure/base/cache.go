package base

import (
	"fmt"
	"notification_service/pkg/settings"
	"strconv"

	"github.com/thanvuc/go-core-lib/cache"
	"github.com/thanvuc/go-core-lib/log"
)

func NewRedis(
	configuration *settings.Configuration,
	logger log.Logger,
) *cache.RedisCache {
	println("HOST: " + fmt.Sprintf("%s:%s", configuration.Redis.Host, strconv.Itoa(configuration.Redis.Port)))
	redisClient := cache.NewRedisCache(cache.Config{
		Addr:     fmt.Sprintf("%s:%s", configuration.Redis.Host, strconv.Itoa(configuration.Redis.Port)),
		DB:       configuration.Redis.DB,
		Password: configuration.Redis.Password,
		PoolSize: configuration.Redis.PoolSize,
		MinIdle:  configuration.Redis.MinIdle,
	})

	if err := redisClient.Ping(); err != nil {
		logger.Error("Failed to connect to Redis", "")
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	} else {
		logger.Info("Redis connection established successfully", "")
	}

	return redisClient
}
