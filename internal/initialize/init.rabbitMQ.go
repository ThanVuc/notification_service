package initialize

import (
	"fmt"
	"notification_service/global"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func InitRabbitMQ() error {
	var logger = global.Logger
	var config = global.Config.RabbitMQ

	var err error
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.User, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Error("Failed to connect to RabbitMQ", zap.Error(err), zap.String("url", url))
		return err
	}

	global.RabbitMQConnection = conn

	return nil
}

func CloseConnection() error {
	var logger = global.Logger
	if global.RabbitMQConnection != nil {
		err := global.RabbitMQConnection.Close()
		if err != nil {
			logger.Error("Failed to close RabbitMQ connection", zap.Error(err))
			return err
		}
		return nil
	}

	return nil
}
