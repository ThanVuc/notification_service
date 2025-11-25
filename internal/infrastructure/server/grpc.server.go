package server

import (
	"context"
	"fmt"
	"net"
	"notification_service/internal/interface/controller"
	"notification_service/pkg/settings"
	"notification_service/proto/notification_service"
	"sync"

	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type NotificationServer struct {
	logger           log.Logger
	configuration    *settings.Configuration
	controllerModule *controller.ControllerModule
}

func NewNotificationServer(
	configuration *settings.Configuration,
	logger log.Logger,
	controllerModule *controller.ControllerModule,
) *NotificationServer {
	return &NotificationServer{
		logger:           logger,
		configuration:    configuration,
		controllerModule: controllerModule,
	}
}

func (ns *NotificationServer) RunServers(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	ns.runServiceServer(ctx, wg)
}

// create server factory
func (ns *NotificationServer) createServer() *grpc.Server {
	server := grpc.NewServer()

	// register services here
	notification_service.RegisterNotificationServiceServer(server, ns.controllerModule.NotificationController)
	notification_service.RegisterUserNotificationServiceServer(server, ns.controllerModule.UserNotificationController)

	return server
}

func (ns *NotificationServer) runServiceServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := ns.createListener()
	if err != nil {
		ns.logger.Error("Failed to create listener",
			"", zap.Error(err),
		)
		return
	}

	// Create a new gRPC server instance
	server := ns.createServer()

	// Gracefully handle server shutdown
	go ns.gracefullyShutdownServer(ctx, server)

	// Server listening on the specified port
	ns.serverListening(server, lis)
}

func (ns *NotificationServer) gracefullyShutdownServer(ctx context.Context, server *grpc.Server) {
	<-ctx.Done()
	ns.logger.Info("gRPC server is shutting down...", "")
	server.GracefulStop()
	ns.logger.Info("gRPC server stopped gracefully!", "")
}

func (ns *NotificationServer) serverListening(server *grpc.Server, lis net.Listener) {
	ns.logger.Info(fmt.Sprintf("gRPC server listening on %s:%d", ns.configuration.Server.Host, lis.Addr().(*net.TCPAddr).Port), "")
	if err := server.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			ns.logger.Info("gRPC server exited normally", "")
		} else {
			ns.logger.Error("Failed to serve gRPC server",
				"", zap.Error(err),
			)
		}
	}
}

func (ns *NotificationServer) createListener() (net.Listener, error) {
	err := error(nil)
	lis := net.Listener(nil)
	lis, err = net.Listen("tcp", fmt.Sprintf("%s:%d", ns.configuration.Server.Host, ns.configuration.Server.NotificationPort))
	if err != nil {
		ns.logger.Error("Failed to listen: %v", "", zap.Error(err))
		return nil, err
	}

	return lis, nil
}
