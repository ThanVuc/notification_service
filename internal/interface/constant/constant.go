package interface_constant

import (
	"fmt"

	"github.com/thanvuc/go-core-lib/eventbus"
)

const (
	// Service
	NOTIFICATION_SERVICE = "notification"

	// Feature
	SCHEDULED_NOTIFICATION = "scheduled_notification"
	WORK_GENERATION        = "generate_work"
	TEAM                   = "team"

	// Common
	EXCHANGE    = "exchange"
	QUEUE       = "queue"
	ROUTING_KEY = "routing_key"

	// Instance
	SCHEDULED_CONSUMER_NUMBER                  = 2
	NOTIFICATION_GENERATE_WORK_CONSUMER_NUMBER = 2
)

// Exchange
var (
	NOTIFICATION_EXCHANGE eventbus.ExchangeName = eventbus.ExchangeName(
		fmt.Sprintf(
			"%s_%s_%s",
			NOTIFICATION_SERVICE,
			SCHEDULED_NOTIFICATION,
			EXCHANGE,
		),
	)

	NOTIFICATION_GENERATE_WORK_EXCHANGE eventbus.ExchangeName = eventbus.ExchangeName(fmt.Sprintf(
		"%s_%s_%s",
		NOTIFICATION_SERVICE,
		WORK_GENERATION,
		EXCHANGE,
	))

	TEAM_EXCHANGE eventbus.ExchangeName = eventbus.ExchangeName(fmt.Sprintf(
		"%s_%s_%s",
		NOTIFICATION_SERVICE,
		TEAM,
		EXCHANGE,
	))
)

// Queue
var (
	NOTIFICATION_QUEUE = fmt.Sprintf(
		"%s_%s_%s",
		NOTIFICATION_SERVICE,
		SCHEDULED_NOTIFICATION,
		QUEUE,
	)

	NOTIFICATION_GENERATE_WORK_QUEUE = fmt.Sprintf(
		"%s_%s_%s",
		NOTIFICATION_SERVICE,
		WORK_GENERATION,
		QUEUE,
	)

	TEAM_QUEUE = fmt.Sprintf(
		"%s_%s_%s",
		NOTIFICATION_SERVICE,
		TEAM,
		QUEUE,
	)
)

// Routing Key
var (
	NOTIFICATION_ROUTING_KEY = fmt.Sprintf(
		"%s_%s_%s",
		NOTIFICATION_SERVICE,
		SCHEDULED_NOTIFICATION,
		ROUTING_KEY,
	)

	NOTIFICATION_GENERATE_WORK_ROUTING_KEY = fmt.Sprintf(
		"%s_%s",
		NOTIFICATION_SERVICE,
		WORK_GENERATION,
	)

	TEAM_ROUTING_KEY = fmt.Sprintf(
		"%s_%s",
		NOTIFICATION_SERVICE,
		TEAM,
	)
)

// cronjob name
const (
	DELETE_OLD_NOTIFICATIONS_CRONJOB = SCHEDULED_NOTIFICATION + "_delete_old_notifications_cronjob"
)
