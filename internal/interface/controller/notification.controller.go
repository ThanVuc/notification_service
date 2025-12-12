package controller

import (
	"context"
	"notification_service/internal/application/usecase"
	"notification_service/pkg/utils"
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	"github.com/thanvuc/go-core-lib/log"
)

type NotificationController struct {
	notification_service.UnimplementedNotificationServiceServer
	notificationUseCase usecase.NotificationUseCase
	logger              log.Logger
}

func NewNotificationController(
	notificationUseCase usecase.NotificationUseCase,
	logger log.Logger,
) *NotificationController {
	return &NotificationController{
		notificationUseCase: notificationUseCase,
		logger:              logger,
	}
}

func (c *NotificationController) GetNotificationsByRecipientId(ctx context.Context, req *common.IDRequest) (*notification_service.GetNotificationsByRecipientIdResponse, error) {
	return utils.WithSafePanic(ctx, c.logger, req, c.notificationUseCase.GetNotificationsByRecipientId)
}

func (c *NotificationController) MarkNotificationsAsRead(ctx context.Context, req *common.IDsRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, c.logger, req, c.notificationUseCase.MarkNotificationsAsRead)
}

func (c *NotificationController) DeleteNotificationById(ctx context.Context, req *common.IDRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, c.logger, req, c.notificationUseCase.DeleteNotificationById)
}
