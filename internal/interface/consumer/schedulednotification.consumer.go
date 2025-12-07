package consumer

import (
	"context"
	"notification_service/internal/application/usecase"
	interface_constant "notification_service/internal/interface/constant"

	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/wagslane/go-rabbitmq"
)

type ScheduledNotificationConsumer struct {
	notificationUsecase usecase.NotificationUseCase
	logger              log.Logger
	connector           *eventbus.RabbitMQConnector
}

func NewScheduledNotificationConsumer(
	notificationUsecase usecase.NotificationUseCase,
	logger log.Logger,
	connector *eventbus.RabbitMQConnector,
) *ScheduledNotificationConsumer {
	return &ScheduledNotificationConsumer{
		notificationUsecase: notificationUsecase,
		logger:              logger,
		connector:           connector,
	}
}

func (s *ScheduledNotificationConsumer) ScheduledNotificationConsume(ctx context.Context) {
	consumer := eventbus.NewConsumer(
		s.connector,
		interface_constant.NOTIFICATION_EXCHANGE,
		eventbus.ExchangeTypeTopic,
		interface_constant.NOTIFICATION_ROUTING_KEY,
		interface_constant.NOTIFICATION_QUEUE,
		interface_constant.SCHEDULED_CONSUMER_NUMBER,
	)
	s.logger.Info("Starting Scheduled Notification consumer", "")
	err := consumer.Consume(ctx, func(d rabbitmq.Delivery) rabbitmq.Action {
		return s.notificationUsecase.ConsumeScheduledNotification(ctx, d)
	})

	if err != nil {
		s.logger.Error("Failed to start ScheduledNotification consumer", "")
		return
	}
}
