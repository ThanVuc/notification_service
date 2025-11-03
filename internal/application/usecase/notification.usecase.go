package usecase

import (
	"context"
	"notification_service/internal/infrastructure/repos"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
)

type notificationUseCase struct {
	mongodbConnector *mongolib.MongoConnector
	logger           log.Logger
	notificationRepo repos.NotificationRepo
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
