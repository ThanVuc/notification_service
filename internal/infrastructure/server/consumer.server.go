package server

import (
	interface_modular "notification_service/internal/interface"
)

type ConsumerWorker struct {
	interfaceModule *interface_modular.InterfaceModule
}

func NewConsumerWorker(
	interfaceModule *interface_modular.InterfaceModule,
) *ConsumerWorker {
	return &ConsumerWorker{
		interfaceModule: interfaceModule,
	}
}

func (s *ConsumerWorker) RunConsumers() {

}
