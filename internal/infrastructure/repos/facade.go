package repos

import (
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
)

type (
	NotificationRepo interface {
		GetNotificationsByRecipientID(request *common.IDRequest) (*notification_service.GetNotificationsResponse, error)
	}
)

func NewNotificationRepo(
	mongoConnector *mongolib.MongoConnector,
	logger log.Logger,
) NotificationRepo {
	return &notificationRepo{
		mongoConnector: mongoConnector,
		logger:         logger,
	}
}
