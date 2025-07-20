package consumers

import (
	"notification_service/global"
	"notification_service/pkg/loggers"
)

type ConsumerFactory interface {
}

type consumerFactory struct {
	consumer Consumer
	logger   *loggers.LoggerZap
}

func NewConsumerFactory() ConsumerFactory {
	return &consumerFactory{
		consumer: NewConsumer(),
		logger:   global.Logger,
	}
}
