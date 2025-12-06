package usecase

import (
	"context"
	"notification_service/internal/core/entity"
	"notification_service/internal/infrastructure/repos"
	"strconv"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

const (
	dbFetchInterval           = 30 * time.Second
	notificationCheckInterval = 10 * time.Second
)

type scheduledWorkerUsecase struct {
	logger           log.Logger
	notificationRepo repos.NotificationRepo
	userRepo         repos.UserNotificationRepo
	firebaseApp      *firebase.App
}

func (s *scheduledWorkerUsecase) ProcessScheduledNotifications(ctx context.Context) error {
	dbTicker := time.NewTicker(dbFetchInterval)
	notificationTicker := time.NewTicker(notificationCheckInterval)
	firebaseClient, err := s.firebaseApp.Messaging(ctx)
	if err != nil {
		s.logger.Error("Failed to get Firebase Messaging client", "", zap.Error(err))
		return err
	}
	defer dbTicker.Stop()
	defer notificationTicker.Stop()

	notificationMap := make(map[string][]*entity.Notification)

	// Fetch immediately at startup
	if err := s.fetchScheduledNotifications(ctx, notificationMap); err != nil {
		s.logger.Error("Initial fetch failed", "", zap.Error(err))
	}

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("Scheduled notification worker stopped", "")
			return nil

		case <-dbTicker.C:
			// Clear previous entries to prevent duplicates
			notificationMap = make(map[string][]*entity.Notification)

			if err := s.fetchScheduledNotifications(ctx, notificationMap); err != nil {
				s.logger.Error("Failed to fetch scheduled notifications", "", zap.Error(err))
			}

		case <-notificationTicker.C:
			currentTime := time.Now().UTC().Format("15:04")

			notificationsToSend, exists := notificationMap[currentTime]
			if !exists {
				continue
			}
			tokenAndNotificationsMap := make(map[string][]*entity.Notification)
			receiverIDs := make([]string, 0)
			for _, notification := range notificationsToSend {
				receiverIDs = append(receiverIDs, notification.ReceiverIds...)
			}

			users, err := s.userRepo.GetUsersByIDs(ctx, receiverIDs)
			if err != nil {
				s.logger.Error("Failed to fetch users for scheduled notifications", "", zap.Error(err))
				continue
			}

			if len(users) == 0 {
				s.logger.Info("No users found for scheduled notifications", "")
				continue
			}

			// Map notifications to user tokens
			// One notification only has one receiver in this case
			for _, notification := range notificationsToSend {
				for _, user := range users {
					if user.FCMToken == "" || len(notification.ReceiverIds) == 0 {
						continue
					}

					if user.UserID == notification.ReceiverIds[0] {
						tokenAndNotificationsMap[user.FCMToken] = append(tokenAndNotificationsMap[user.FCMToken], notification)
					}
				}
			}

			for key, notifications := range tokenAndNotificationsMap {
				for _, notification := range notifications {
					s.logger.Info("Notification to send", "", zap.String("trigger_at", notification.TriggerAt.String()))
					imgUrl := ""
					if notification.ImgUrl != nil {
						imgUrl = *notification.ImgUrl
					}
					link := ""
					if notification.Link != nil {
						link = *notification.Link
					}
					triggerAt := ""
					if notification.TriggerAt != nil {
						triggerAt = strconv.FormatInt(notification.TriggerAt.UnixMilli(), 10)
					}

					message := &messaging.Message{
						Token: key,
						Data: map[string]string{
							"title":      notification.Title,
							"body":       notification.Message,
							"url":        link,
							"src":        imgUrl,
							"trigger_at": triggerAt,
						},
					}

					resp, err := firebaseClient.Send(ctx, message)
					if err != nil {
						s.logger.Error("Error sending scheduled notification", "", zap.Error(err))
					}
					s.logger.Info("Successfully sent scheduled notification", "", zap.String("response", resp))
					s.logger.Info("Notification: ", "", zap.String("Token", message.Token), zap.String("Title", message.Data["title"]))
				}
			}

			// Clean after sending
			delete(notificationMap, currentTime)
		}
	}
}

func (s *scheduledWorkerUsecase) fetchScheduledNotifications(
	ctx context.Context,
	notificationMap map[string][]*entity.Notification,
) error {
	now := time.Now().UTC()

	notifications, err := s.notificationRepo.GetNotificationsWithinTimeRange(
		ctx,
		now.Add(-1*time.Minute),
		now.Add(dbFetchInterval),
	)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		timeStr := notification.TriggerAt.UTC().Format("15:04")
		notificationMap[timeStr] = append(notificationMap[timeStr], notification)
	}

	return nil
}
