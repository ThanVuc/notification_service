package consumer

import (
	"notification_service/internal/application"
)

type ConsumerModule struct {
}

func NewConsumerModule(
	applicationModuel *application.ApplicationModule,
) *ConsumerModule {
	return &ConsumerModule{}
}
