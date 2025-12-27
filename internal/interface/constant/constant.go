package interface_constant

import (
	"github.com/thanvuc/go-core-lib/eventbus"
)

// base
const (
	SERVICE                = "notification"
	SCHEDULED_NOTIFICATION = "scheduled_notification"
)

// exchange, queue, routing key
const (
	EXCHANGE    = ".exchange"
	QUEUE       = ".queue"
	ROUTING_KEY = ".routing_key"
)

// exchanges full names
const (
	NOTIFICATION_EXCHANGE eventbus.ExchangeName = SCHEDULED_NOTIFICATION + EXCHANGE
)

// queues full names
const (
	NOTIFICATION_QUEUE string = SCHEDULED_NOTIFICATION + QUEUE
)

// routing keys full names
const (
	NOTIFICATION_ROUTING_KEY string = SCHEDULED_NOTIFICATION + ROUTING_KEY
)

const (
	SCHEDULED_CONSUMER_NUMBER = 2
)

// cronjob name
const (
	DELETE_OLD_NOTIFICATIONS_CRONJOB = SCHEDULED_NOTIFICATION + "_delete_old_notifications_cronjob"
)
