package base

import (
	"notification_service/pkg/settings"
	"os"

	"github.com/thanvuc/go-core-lib/log"
)

func NewLogger(
	configuration *settings.Configuration,
) log.Logger {
	env := os.Getenv("GO_ENV")
	logger := log.NewLogger(log.Config{
		Env:   env,
		Level: configuration.Log.Level,
	})

	return logger
}
