package server

import (
	"context"
	interface_modular "notification_service/internal/interface"
	"sync"
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

func (s *Worker) RunWorkers(ctx context.Context, wg *sync.WaitGroup) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.interfaceModule.WorkerModule.ScheduledNotificationWorker.RunScheduledNotifications(ctx)
	}()

	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.interfaceModule.WorkerModule.AppNotificationWorker.Start(ctx)
		}()
	}

	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.interfaceModule.WorkerModule.EmailNotificationWorker.Start(ctx)
		}()
	}
}
