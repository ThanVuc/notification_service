package base

import (
	"context"
	"fmt"
	"notification_service/pkg/settings"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/thanvuc/go-core-lib/log"
	"google.golang.org/api/option"
)

func NewFirebaseApp(
	config *settings.Configuration,
	logger log.Logger,
) *firebase.App {
	opt := option.WithCredentialsFile(config.Firebase.Path)
	ctx := context.Background()
	const maxRetries = 10
	var app *firebase.App
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		app, err = firebase.NewApp(ctx, nil, opt)
		if err == nil {
			// ✅ Successfully initialized Firebase
			logger.Info("Firebase connected successfully", "")
			return app
		}

		// ❌ Log the failure
		wait := time.Duration(attempt*2) * time.Second
		logger.Warn(
			fmt.Sprintf("Failed to initialize Firebase App (attempt %d/%d): %v. Retrying in %v...",
				attempt, maxRetries, err, wait),
			"",
		)

		// Wait before next attempt or exit if context canceled
		select {
		case <-ctx.Done():
			logger.Error("Firebase initialization canceled by context", "")
			return nil
		case <-time.After(wait):
			continue
		}
	}

	// 🔥 All retries failed
	logger.Error("Failed to initialize Firebase App after multiple retries: "+err.Error(), "")
	return nil
}
