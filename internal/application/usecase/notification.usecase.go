package usecase

import (
	"context"
	"notification_service/internal/infrastructure/repos"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type notificationUseCase struct {
	mongodbConnector *mongolib.MongoConnector
	logger           log.Logger
	notificationRepo repos.NotificationRepo
	firebaseApp      *firebase.App
}

func (n *notificationUseCase) GetNotifications(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsResponse, error) {
	// Implementation goes here
	// n.notificationRepo.GetNotificationsByRecipientID(req)
	return &notification_service.GetNotificationsResponse{
		Notifications: []*notification_service.Notification{
			&notification_service.Notification{
				Id:         "1234",
				Title:      "Hello",
				Message:    "Hello",
				Timestamp:  1234567789,
				SenderId:   "12345",
				ReceiverId: "12345",
				IsRead:     true,
			},
		},
	}, nil
}

func (n *notificationUseCase) ConsumeScheduledNotification(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action {
	n.logger.Info("Consumed scheduled notification", "", zap.ByteString("body", d.Body))
	messagingClient, err := n.firebaseApp.Messaging(ctx)
	if err != nil {
		return rabbitmq.NackRequeue
	}

	messagingClient.Send(ctx, &messaging.Message{})

	return rabbitmq.Ack
}
