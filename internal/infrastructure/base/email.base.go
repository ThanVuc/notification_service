package base

import (
	"notification_service/pkg/settings"

	"github.com/thanvuc/go-core-lib/log"
	"gopkg.in/gomail.v2"
)

func NewEmailDialer(
	configuration *settings.Configuration,
	logger log.Logger,
) *gomail.Dialer {
	emailSettings := configuration.Email
	dialer := gomail.NewDialer(
		emailSettings.Host,
		emailSettings.Port,
		emailSettings.Username,
		emailSettings.Password,
	)
	return dialer
}
