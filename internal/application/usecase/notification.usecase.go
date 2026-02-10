package usecase

import (
	"context"
	"encoding/json"
	"notification_service/internal/application/dto"
	"notification_service/internal/application/helper"
	"notification_service/internal/application/mapper"
	app_model "notification_service/internal/application/model"
	"notification_service/internal/core/entity"
	"notification_service/internal/infrastructure/repos"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"
	"strconv"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
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
	userRepo           repos.UserNotificationRepo
	emailHelper        helper.EmailHelper
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
		return rabbitmq.NackDiscard
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

func (n *notificationUseCase) ProcessDeleteOldNotifications(ctx context.Context) error {
	before := time.Now().AddDate(0, 0, -30) // Delete notifications older than 30 days
	err := n.notificationRepo.DeleteOldNotifications(context.Background(), before)
	if err != nil {
		n.logger.Error("Failed to delete old notifications", "", zap.Error(err))
		return err
	}
	n.logger.Info("Old notifications (30 day ago) deleted successfully", "")
	return nil
}

func (n *notificationUseCase) ConsumeWorkGeneration(ctx context.Context, d rabbitmq.Delivery) rabbitmq.Action {
	notificationMessage, err := n.DecodeWorkMessage(d.Body)
	if err != nil {
		n.logger.Error("Failed to decode notification generation message", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	// build notification entities
	now := time.Now().UTC()
	notificationEntity := &entity.Notification{
		ID:              bson.NewObjectID(),
		Title:           notificationMessage.Title,
		Message:         notificationMessage.Message,
		Link:            notificationMessage.Link,
		SenderId:        notificationMessage.SenderID,
		ReceiverIds:     notificationMessage.ReceiverIDs,
		IsRead:          false,
		TriggerAt:       &now,
		CreatedAt:       now,
		UpdatedAt:       now,
		ImgUrl:          notificationMessage.ImageURL,
		CorrelationId:   notificationMessage.CorrelationID,
		CorrelationType: int32(notificationMessage.CorrelationType),
		IsActive:        true,
		IsSendMail:      true,
		IsPublished:     false,
	}

	// save before send noti
	err = n.notificationRepo.UpsertNotification(ctx, notificationEntity)
	if err != nil {
		n.logger.Error("Failed to save work generation notification", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	// send mail & notifications
	if len(notificationMessage.ReceiverIDs) == 0 {
		n.logger.Warn("No receiver IDs found for work generation notification", "")
		return rabbitmq.Ack
	}
	println("user-id", notificationMessage.ReceiverIDs[0])
	user, err := n.userRepo.GetUsersByID(ctx, notificationMessage.ReceiverIDs[0])
	if err != nil {
		n.logger.Error("Failed to get user for work generation notification", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	if user == nil {
		n.logger.Warn("User not found for work generation notification", "")
		return rabbitmq.Ack
	}

	if err := n.SendAIGenerationMail(ctx, notificationEntity, user); err != nil {
		n.logger.Error("Error sending AI generation email", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	message := &messaging.Message{
		Token: user.FCMToken,
		Data: map[string]string{
			"title":      notificationEntity.Title,
			"body":       notificationEntity.Message,
			"url":        utils.SafeStringWithDefault(notificationEntity.Link, "https://www.schedulr.site/images/ai-icon.webp"),
			"src":        utils.SafeStringWithDefault(notificationEntity.ImgUrl, "https://www.schedulr.site/schedule/daily"),
			"trigger_at": strconv.FormatInt(notificationEntity.TriggerAt.UnixMilli(), 10),
		},
	}

	firebaseClient, err := n.firebaseApp.Messaging(ctx)
	if err != nil {
		n.logger.Error("Failed to init firebase messaging", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	if _, err := firebaseClient.Send(ctx, message); err != nil {
		n.logger.Error("Error sending work generation notification", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	// update notification as sent
	err = n.notificationRepo.MarkIsPublished(ctx, []bson.ObjectID{notificationEntity.ID})
	if err != nil {
		n.logger.Error("Failed to mark work generation notification as published", "", zap.Error(err))
		return rabbitmq.NackDiscard
	}

	n.logger.Info("Work generation notification sent successfully", "")

	return rabbitmq.Ack
}

func (n *notificationUseCase) DecodeWorkMessage(body []byte) (*dto.WorkGenerationNotificationMessage, error) {
	var notifications dto.WorkGenerationNotificationMessage

	if err := json.Unmarshal(body, &notifications); err != nil {
		return nil, err
	}

	return &notifications, nil
}

func (n *notificationUseCase) SendAIGenerationMail(ctx context.Context, notification *entity.Notification, user *entity.User) error {
	err := n.emailHelper.SendAIWorkGenerationEmail(
		user.Email,
		app_model.EmailData{
			Title:      notification.Title,
			Message:    notification.Message,
			Link:       *notification.Link,
			ButtonText: "Click vào đây để xem công việc đã sinh.",
		},
	)
	return err
}
