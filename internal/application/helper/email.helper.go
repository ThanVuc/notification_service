package helper

import (
	"bytes"
	"fmt"
	app_constant "notification_service/internal/application/constant"
	app_model "notification_service/internal/application/model"
	app_template "notification_service/internal/application/template"
	"text/template"

	"github.com/thanvuc/go-core-lib/log"
	"gopkg.in/gomail.v2"
)

type EmailHelper struct {
	emailDialer *gomail.Dialer
	logger      log.Logger
}

func NewEmailHelper(
	emailDialer *gomail.Dialer,
	logger log.Logger,
) *EmailHelper {
	return &EmailHelper{
		emailDialer: emailDialer,
		logger:      logger,
	}
}

func (h *EmailHelper) SendScheduledWorkEmail(to string, data app_model.EmailData) error {
	htmlBody, textBody, err := RenderEmail(data)
	if err != nil {
		return err
	}

	emailContent := fmt.Sprintf("Title: %s, message: %s, Link: %s")
	h.logger.Info("Sending email to: "+to+", With content: "+emailContent, "")

	m := gomail.NewMessage()

	// From + To + Subject
	m.SetHeader("From", app_constant.APP_EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Title)

	// Add text fallback
	m.AddAlternative("text/plain", textBody)

	// Add HTML version
	m.AddAlternative("text/html", htmlBody)

	// Send using the helper's dialer
	return h.emailDialer.DialAndSend(m)
}

func RenderEmail(data app_model.EmailData) (htmlBody string, textBody string, err error) {
	// Parse HTML template
	tHTML, err := template.New("email-html").Parse(app_template.EmailHTML)
	if err != nil {
		return "", "", err
	}

	var htmlBuf bytes.Buffer
	if err := tHTML.Execute(&htmlBuf, data); err != nil {
		return "", "", err
	}

	// Parse text fallback
	tText, err := template.New("email-text").Parse(app_template.EmailPlain)
	if err != nil {
		return "", "", err
	}

	var textBuf bytes.Buffer
	if err := tText.Execute(&textBuf, data); err != nil {
		return "", "", err
	}

	return htmlBuf.String(), textBuf.String(), nil
}
