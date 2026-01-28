package consumer

import (
	"context"
	"notification_service/internal/application/usecase"
	interface_constant "notification_service/internal/interface/constant"

	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/wagslane/go-rabbitmq"
)

type WorkGenerationConsumer struct {
	notificationUsecase usecase.NotificationUseCase
	logger              log.Logger
	connector           *eventbus.RabbitMQConnector
}

func NewWorkGenerationConsumer(
	notificationUsecase usecase.NotificationUseCase,
	logger log.Logger,
	connector *eventbus.RabbitMQConnector,
) *WorkGenerationConsumer {
	return &WorkGenerationConsumer{
		notificationUsecase: notificationUsecase,
		logger:              logger,
		connector:           connector,
	}
}

func (s *WorkGenerationConsumer) WorkGenerationConsume(ctx context.Context) {
	consumer := eventbus.NewConsumer(
		s.connector,
		interface_constant.NOTIFICATION_GENERATE_WORK_EXCHANGE,
		eventbus.ExchangeTypeDirect,
		interface_constant.NOTIFICATION_GENERATE_WORK_ROUTING_KEY,
		interface_constant.NOTIFICATION_GENERATE_WORK_QUEUE,
		interface_constant.NOTIFICATION_GENERATE_WORK_CONSUMER_NUMBER,
	)

	s.logger.Info("Starting notification generate works consumer", "")
	err := consumer.Consume(ctx, func(d rabbitmq.Delivery) rabbitmq.Action {
		return s.notificationUsecase.ConsumeWorkGeneration(ctx, d)
	})

	if err != nil {
		s.logger.Error("Failed to start WorkGeneration consumer", "")
		return
	}
}
