package global

import (
	"notification_service/pkg/loggers"
	"notification_service/pkg/settings"

	"github.com/streadway/amqp"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: This package defines global variables that are used throughout the application.
*/
var (
	Config             settings.Config
	Logger             *loggers.LoggerZap
	RabbitMQConnection *amqp.Connection
)
