package controller

import (
	"context"
	"notification_service/internal/application/usecase"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	"github.com/thanvuc/go-core-lib/log"
)

type UserNotificationController struct {
	notification_service.UnimplementedUserNotificationServiceServer
	userNotificationUseCase usecase.UserNotificationUseCase
	logger                  log.Logger
}

func NewUserNotificationController(
	userNotificationUseCase usecase.UserNotificationUseCase,
	logger log.Logger,
) *UserNotificationController {
	return &UserNotificationController{
		userNotificationUseCase: userNotificationUseCase,
		logger:                  logger,
	}
}

func (c *UserNotificationController) UpsertUserFCMToken(ctx context.Context, req *notification_service.UpsertUserFCMTokenRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, c.logger, req, c.userNotificationUseCase.UpsertUserFCMToken)
}
