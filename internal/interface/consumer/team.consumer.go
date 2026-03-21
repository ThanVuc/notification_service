package consumer

import (
	"context"
	"notification_service/internal/application/usecase"
	interface_constant "notification_service/internal/interface/constant"

	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/wagslane/go-rabbitmq"
)

type TeamNotificationConsumer struct {
	notificationUsecase usecase.NotificationUseCase
	logger              log.Logger
	connector           *eventbus.RabbitMQConnector
}

func NewTeamNotificationConsumer(
	notificationUsecase usecase.NotificationUseCase,
	logger log.Logger,
	connector *eventbus.RabbitMQConnector,
) *TeamNotificationConsumer {
	return &TeamNotificationConsumer{
		notificationUsecase: notificationUsecase,
		logger:              logger,
		connector:           connector,
	}
}

func (s *TeamNotificationConsumer) TeamConsume(ctx context.Context) {
	consumer := eventbus.NewConsumer(
		s.connector,
		interface_constant.TEAM_EXCHANGE,
		eventbus.ExchangeTypeDirect,
		interface_constant.TEAM_ROUTING_KEY,
		interface_constant.TEAM_QUEUE,
		interface_constant.NOTIFICATION_GENERATE_WORK_CONSUMER_NUMBER,
	)

	s.logger.Info("Starting team notification consumer", "")
	err := consumer.Consume(ctx, func(d rabbitmq.Delivery) rabbitmq.Action {
		return s.notificationUsecase.ConsumeTeamNotification(ctx, d)
	})

	if err != nil {
		s.logger.Error("Failed to start team notification consumer", "")
		return
	}
}
