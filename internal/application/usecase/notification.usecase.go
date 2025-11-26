package usecase

import (
	"context"
	"notification_service/internal/application/mapper"
	"notification_service/internal/infrastructure/repos"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	firebase "firebase.google.com/go/v4"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type notificationUseCase struct {
	mongodbConnector   *mongolib.MongoConnector
	logger             log.Logger
	notificationRepo   repos.NotificationRepo
	firebaseApp        *firebase.App
	notificationMapper mapper.NotificationMapper
}

func (n *notificationUseCase) GetNotifications(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsResponse, error) {
	// Implementation goes here
	// n.notificationRepo.GetNotificationsByRecipientID(req)
	return &notification_service.GetNotificationsResponse{
		Notifications: []*notification_service.Notification{
			&notification_service.Notification{
				Id:          "1234",
				Title:       "Hello",
				Message:     "Hello",
				ReceiverIds: []string{"12345", "45678"},
				SenderId:    "12345",
				IsRead:      true,
			},
		},
	}, nil
}

func (n *notificationUseCase) ConsumeScheduledNotification(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action {
	notificationBody := new(common.Notification)
	requestId := d.Headers["request_id"].(string)
	err := proto.Unmarshal(d.Body, notificationBody)
	if err != nil {
		n.logger.Error("Failed to unmarshal scheduled notification", requestId, zap.Error(err))
		return rabbitmq.NackDiscard
	}

	// Process the scheduled notification
	notificationEntity := n.notificationMapper.FromProtoToEntity(notificationBody)
	err = n.notificationRepo.SaveNotification(ctx, notificationEntity)
	if err != nil {
		n.logger.Error("Failed to save scheduled notification", requestId, zap.Error(err))
		return rabbitmq.NackRequeue
	}

	return rabbitmq.Ack
}
