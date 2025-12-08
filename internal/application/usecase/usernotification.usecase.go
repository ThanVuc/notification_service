package usecase

import (
	"context"
	"notification_service/internal/core/entity"
	"notification_service/internal/infrastructure/repos"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"
	"time"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
)

type userNotificationUseCase struct {
	mongodbConnector *mongolib.MongoConnector
	logger           log.Logger
	userRepo         repos.UserNotificationRepo
}

func (n *userNotificationUseCase) UpsertUserFCMToken(ctx context.Context, req *notification_service.UpsertUserFCMTokenRequest) (*common.EmptyResponse, error) {
	user := &entity.User{
		UserID:    req.GetUserId(),
		DeviceID:  req.GetDeviceId(),
		FCMToken:  req.GetFcmToken(),
		Email:     req.GetEmail(),
		UpdatedAt: time.Now().UTC(),
	}

	err := n.userRepo.UpsertUserNotification(ctx, user)

	if err != nil {
		return &common.EmptyResponse{
			Success: utils.ToBoolPointer(false),
			Message: utils.ToStringPointer("Failed to upsert user FCM token"),
			Error:   utils.DatabaseError(ctx, n.logger, err),
		}, err
	}

	resp := &common.EmptyResponse{
		Success: utils.ToBoolPointer(true),
		Message: utils.ToStringPointer("Successfully upserted user FCM token"),
		Error:   nil,
	}

	return resp, nil
}
