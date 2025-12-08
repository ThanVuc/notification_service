package base

import (
	"notification_service/pkg/settings"

	firebase "firebase.google.com/go/v4"
	"github.com/thanvuc/go-core-lib/cache"
	"github.com/thanvuc/go-core-lib/eventbus"
	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/mongolib"
	"gopkg.in/gomail.v2"
)

type BaseModule struct {
	Logger            log.Logger
	EventBusConnector *eventbus.RabbitMQConnector
	CacheConnector    *cache.RedisCache
	MongoConnector    *mongolib.MongoConnector
	FirebaseApp       *firebase.App
	EmailDialer       *gomail.Dialer
}

func NewBaseModule(
	configuration *settings.Configuration,
) *BaseModule {
	logger := NewLogger(configuration)
	eventBusConnector := NewEventBus(configuration, logger)
	cacheConnector := NewRedis(configuration, logger)
	mongodbConnector := NewMongoDB(configuration, logger)
	firebaseApp := NewFirebaseApp(configuration, logger)
	emailDialer := NewEmailDialer(configuration, logger)

	return &BaseModule{
		Logger:            logger,
		EventBusConnector: eventBusConnector,
		CacheConnector:    cacheConnector,
		MongoConnector:    mongodbConnector,
		FirebaseApp:       firebaseApp,
		EmailDialer:       emailDialer,
	}
}
