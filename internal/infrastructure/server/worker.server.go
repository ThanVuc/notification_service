package server

import (
	"context"
	interface_modular "notification_service/internal/interface"
)

type Worker struct {
	interfaceModule *interface_modular.InterfaceModule
}

func NewWorker(
	interfaceModule *interface_modular.InterfaceModule,
) *Worker {
	return &Worker{
		interfaceModule: interfaceModule,
	}
}

func (s *Worker) RunWorkers() {
	ctx := context.Background()
	go s.interfaceModule.WorkerModule.ScheduledNotificationWorker.RunScheduledNotifications(ctx)
}
