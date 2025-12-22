package interface_constant

import "github.com/thanvuc/go-core-lib/eventbus"

// base
const (
	SERVICE                = "notification"
	SCHEDULED_NOTIFICATION = "scheduled_notification"
	VIET_NAM_JOB           = "viet_nam_job"
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
	VIET_NAM_JOB_EXCHANGE eventbus.ExchangeName = VIET_NAM_JOB + EXCHANGE
)

// queues full names
const (
	NOTIFICATION_QUEUE string = SCHEDULED_NOTIFICATION + QUEUE
	TEST_JOB_QUEUE     string = VIET_NAM_JOB + SERVICE + QUEUE
)

// routing keys full names
const (
	NOTIFICATION_ROUTING_KEY  string = SCHEDULED_NOTIFICATION + ROUTING_KEY
	ONE_DAY_JOB_ROUTING_KEY   string = VIET_NAM_JOB + ".one_day"
	THREE_DAY_JOB_ROUTING_KEY string = VIET_NAM_JOB + ".three_day"
	TEST_JOB_ROUTING_KEY      string = VIET_NAM_JOB + ".test"
)

const (
	SCHEDULED_CONSUMER_NUMBER = 2
	CRONJOB_CONSUMER_NUMBER   = 1
)
