package base

import (
	"notification_service/pkg/settings"

	"github.com/thanvuc/go-core-lib/config"
)

func LoadConfiguration() *settings.Configuration {
	configuration := &settings.Configuration{}
	err := config.LoadConfig(configuration, "./")

	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	return configuration
}
