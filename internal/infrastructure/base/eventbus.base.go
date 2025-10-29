package base

import (
	"fmt"
	"notification_service/pkg/settings"

	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
)

func NewEventBus(
	configuration *settings.Configuration,
	logger log.Logger,
) *eventbus.RabbitMQConnector {
	uri := fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		configuration.RabbitMQ.User,
		configuration.RabbitMQ.Password,
		configuration.RabbitMQ.Host,
		configuration.RabbitMQ.Port,
	)

	connector, err := eventbus.NewConnector(
		uri,
		logger,
	)

	if err != nil {
		panic(err)
	}

	return connector
}
