package consumer

import "notification_service/internal/domain"

type ConsumerModule struct {
}

func NewConsumerModule(
	domainModule *domain.DomainModule,
) *ConsumerModule {
	return &ConsumerModule{}
}
