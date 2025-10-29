package main

import (
	"context"
	"notification_service/cmd"
	"notification_service/global"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	diContainer := global.NewDIContainer()
	logger := diContainer.GetInfrastructureModule().BaseModule.Logger
	// run consumer in another goroutine
	cmd.RunConsumer(ctx, &wg, diContainer)

	// run grpc server in main goroutine
	cmd.RunGrpcServer(ctx, &wg, diContainer)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("shutting down notification service...", "")
	cancel()
	diContainer.GracefulShutdown(&wg)
}
