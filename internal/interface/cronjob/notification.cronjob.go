package cronjob

import (
	"context"
	"notification_service/internal/application/usecase"
	interface_constant "notification_service/internal/interface/constant"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/thanvuc/go-core-lib/cache"
	"github.com/thanvuc/go-core-lib/cronjob"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type NotificationCronJob struct {
	notificationUseCase usecase.NotificationUseCase
	logger              log.Logger
	redisClient         *cache.RedisCache
	cronJobManager      *cronjob.CronManager
}

func NewNotificationCronJob(
	notificationUseCase usecase.NotificationUseCase,
	logger log.Logger,
	redisClient *cache.RedisCache,
	cronjobManager *cronjob.CronManager,
) *NotificationCronJob {
	return &NotificationCronJob{
		notificationUseCase: notificationUseCase,
		logger:              logger,
		redisClient:         redisClient,
		cronJobManager:      cronjobManager,
	}
}

func (c *NotificationCronJob) DeleteOldNotificationsCronJob(ctx context.Context) {
	// 1 monthly job to delete old notifications
	loc, _ := time.LoadLocation(interface_constant.LOCATION_HCM)
	jobScheduler := cronjob.NewCronScheduler(c.redisClient, interface_constant.DELETE_OLD_NOTIFICATIONS_CRONJOB, cron.WithLocation(loc))

	c.cronJobManager.AddScheduler(jobScheduler)

	err := jobScheduler.ScheduleCronJob("0 0 1 * *", func() {
		c.notificationUseCase.ProcessDeleteOldNotifications(ctx)
	})
	if err != nil {
		c.logger.Error("Failed to schedule DeleteOldNotificationsCronJob", "", zap.Error(err))
	}
	jobScheduler.Start()
}
