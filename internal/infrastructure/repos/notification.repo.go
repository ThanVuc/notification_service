package repos

import (
	"notification_service/proto/common"
	"notification_service/proto/notification_service"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
)

type notificationRepo struct {
	mongoConnector *mongolib.MongoConnector
	logger         log.Logger
}

func (r *notificationRepo) GetNotificationsByRecipientID(request *common.IDRequest) (*notification_service.GetNotificationsResponse, error) {
	// Implementation goes here
	return nil, nil
}
