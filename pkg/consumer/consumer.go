package consumers

import (
	"notification_service/global"
	"notification_service/pkg/loggers"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Consumer interface {
	RegisterConsumer(brokerMode brokerModeType, exchangeName exchangeNameType, routingKey routingKeyType, consumerName ConsumerNameType) (<-chan amqp.Delivery, error)
	createChannel() (*amqp.Channel, error)
	declareExchange(channel *amqp.Channel, brokerMode brokerModeType, exchangeName exchangeNameType) error
	declareQueue(channel *amqp.Channel) (amqp.Queue, error)
	bindQueue(channel *amqp.Channel, queueName string, exchangeName exchangeNameType, routingKey routingKeyType) error
	consumeMessages(channel *amqp.Channel, queueName string, consumerName ConsumerNameType) (<-chan amqp.Delivery, error)
	handleError(err error, message string, consumerName ConsumerNameType) bool
}

type consumer struct {
	logger     *loggers.LoggerZap
	connection *amqp.Connection
}

func NewConsumer() Consumer {
	return &consumer{
		logger:     global.Logger,
		connection: global.RabbitMQConnection,
	}
}

func (c *consumer) RegisterConsumer(brokerMode brokerModeType, exchangeName exchangeNameType, routingKey routingKeyType, consumerName ConsumerNameType) (<-chan amqp.Delivery, error) {
	channel, err := c.createChannel()
	if c.handleError(err, "Failed to create channel", consumerName) {
		return nil, err
	}

	err = c.declareExchange(channel, brokerMode, exchangeName)
	if c.handleError(err, "Failed to declare exchange", consumerName) {
		return nil, err
	}

	q, err := c.declareQueue(channel)
	if c.handleError(err, "Failed to declare queue", consumerName) {
		return nil, err
	}

	err = c.bindQueue(channel, q.Name, exchangeName, routingKey)
	if c.handleError(err, "Failed to bind queue", consumerName) {
		return nil, err
	}

	msgs, err := c.consumeMessages(channel, q.Name, consumerName)
	if c.handleError(err, "Failed to consume messages", consumerName) {
		return nil, err
	}

	return msgs, nil
}

func (c *consumer) createChannel() (*amqp.Channel, error) {
	channel, err := c.connection.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (c *consumer) declareExchange(channel *amqp.Channel, brokerMode brokerModeType, exchangeName exchangeNameType) error {
	err := channel.ExchangeDeclare(
		string(exchangeName),
		string(brokerMode),
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}
	return nil
}

func (c *consumer) declareQueue(channel *amqp.Channel) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		"",    // name
		true,  // durable
		false, // auto-deleted
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return amqp.Queue{}, err
	}

	return q, nil
}

func (c *consumer) bindQueue(channel *amqp.Channel, queueName string, exchangeName exchangeNameType, routingKey routingKeyType) error {
	err := channel.QueueBind(
		queueName,            // queue name
		string(routingKey),   // routing key
		string(exchangeName), // exchange
		false,                // no-wait
		nil,                  // arguments
	)

	if err != nil {
		return err
	}
	return nil
}

func (c *consumer) consumeMessages(channel *amqp.Channel, queueName string, consumerName ConsumerNameType) (<-chan amqp.Delivery, error) {
	msgs, err := channel.Consume(
		queueName,            // queue
		string(consumerName), // consumer
		false,                // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // arguments
	)

	if c.handleError(err, "Failed to consume messages", consumerName) {
		return nil, err
	}

	return msgs, nil
}

func (c *consumer) handleError(err error, message string, consumerName ConsumerNameType) bool {
	if err != nil {
		c.logger.Error(message, zap.Error(err), zap.String("consumer", string(consumerName)))
		return true
	}
	return false
}
