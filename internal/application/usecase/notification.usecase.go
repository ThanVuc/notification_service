package usecase

import (
	"context"
	"notification_service/internal/application/mapper"
	"notification_service/internal/infrastructure/repos"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	firebase "firebase.google.com/go/v4"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"github.com/wagslane/go-rabbitmq"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (n *notificationUseCase) GetNotificationsByRecipientId(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsByRecipientIdResponse, error) {
	notifications, err := n.notificationRepo.GetNotificationsByRecipientID(ctx, req)
	requestId := utils.GetRequestIDFromOutgoingContext(ctx)
	if err != nil {
		n.logger.Error("Failed to get notifications by recipient ID", requestId, zap.Error(err))
		return nil, err
	}
	notificationProtos := n.notificationMapper.FromEntitiesToProtoList(notifications)
	resp := &notification_service.GetNotificationsByRecipientIdResponse{
		Notifications: notificationProtos,
	}
	return resp, nil
}

func (n *notificationUseCase) ConsumeScheduledNotification(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action {
	notificationsBody := common.Notifications{}
	requestId := d.Headers["request_id"].(string)
	err := proto.Unmarshal(d.Body, &notificationsBody)
	if err != nil {
		n.logger.Error("Failed to unmarshal scheduled notification", requestId, zap.Error(err))
		return rabbitmq.NackDiscard
	}

	// Process the scheduled notification
	notificationEntities := n.notificationMapper.FromProtoListToEntities(notificationsBody.Notifications, false)
	err = n.notificationRepo.UpsertNotifications(ctx, notificationEntities)
	if err != nil {
		n.logger.Error("Failed to save scheduled notification", requestId, zap.Error(err))
		return rabbitmq.NackRequeue
	}

	return rabbitmq.Ack
}

func (n *notificationUseCase) MarkNotificationsAsRead(ctx context.Context, req *common.IDsRequest) (*common.EmptyResponse, error) {
	objectIds := make([]bson.ObjectID, 0, len(req.Ids))
	requestId := utils.GetRequestIDFromOutgoingContext(ctx)
	for _, idStr := range req.Ids {
		objectId, err := bson.ObjectIDFromHex(idStr)
		if err != nil {
			n.logger.Warn("Invalid notification ID", requestId, zap.Error(err))
			continue
		}
		objectIds = append(objectIds, objectId)
	}

	err := n.notificationRepo.MarkNotificationsAsRead(ctx, objectIds)
	if err != nil {
		return nil, err
	}
	return &common.EmptyResponse{}, nil
}

func (n *notificationUseCase) DeleteNotificationById(ctx context.Context, req *common.IDRequest) (*common.EmptyResponse, error) {
	objectId, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	err = n.notificationRepo.DeleteNotificationById(ctx, objectId)
	if err != nil {
		return nil, err
	}
	return &common.EmptyResponse{}, nil
}

func (n *notificationUseCase) GetNotificationByWorkId(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsByWorkIdResponse, error) {
	notifications, err := n.notificationRepo.GetNotificationByWorkId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	notificationProtos := n.notificationMapper.FromNotificationEntitiesToWorkNotificationsProto(notifications)
	return &notification_service.GetNotificationsByWorkIdResponse{
		Notifications: notificationProtos,
	}, nil
}
