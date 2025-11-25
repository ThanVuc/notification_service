package interface_constant

import "github.com/thanvuc/go-core-lib/eventbus"

const (
	SCHEDULED_NOTIFICATION = "scheduled_notification"
)

const (
	EXCHANGE    = ".exchange"
	QUEUE       = ".queue"
	ROUTING_KEY = ".routing_key"
)

const (
	NOTIFICATION_EXCHANGE eventbus.ExchangeName = SCHEDULED_NOTIFICATION + EXCHANGE
)

const (
	SCHEDULED_CONSUMER_NUMBER = 2
)
