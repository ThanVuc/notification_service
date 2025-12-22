package consumer

import (
	"context"
	"notification_service/internal/application/usecase"
	interface_constant "notification_service/internal/interface/constant"

	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/wagslane/go-rabbitmq"
)

type CronJobConsumer struct {
	logger         log.Logger
	connector      *eventbus.RabbitMQConnector
	cronJobUseCase usecase.CronJobUseCase
}

func NewCronJobConsumer(
	logger log.Logger,
	connector *eventbus.RabbitMQConnector,
	cronJobUseCase usecase.CronJobUseCase,
) *CronJobConsumer {
	return &CronJobConsumer{
		logger:         logger,
		connector:      connector,
		cronJobUseCase: cronJobUseCase,
	}
}

func (c *CronJobConsumer) DeleteOldNotificationsConsume(ctx context.Context) {
	c.logger.Info("Starting job DeleteOldNotificationsConsume consumer", "")
	consumer := eventbus.NewConsumer(
		c.connector,
		interface_constant.VIET_NAM_JOB_EXCHANGE,
		eventbus.ExchangeTypeTopic,
		interface_constant.TEST_JOB_ROUTING_KEY,
		interface_constant.TEST_JOB_QUEUE,
		interface_constant.CRONJOB_CONSUMER_NUMBER,
	)
	err := consumer.Consume(ctx, func(d rabbitmq.Delivery) rabbitmq.Action {
		return c.cronJobUseCase.ProcessDeleteOldNotifications(ctx, d)
	})

	if err != nil {
		c.logger.Error("Failed to start CronJob consumer", "")
		return
	}
}
