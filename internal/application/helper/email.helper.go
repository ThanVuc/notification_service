package helper

import (
	"bytes"
	app_constant "notification_service/internal/application/constant"
	app_model "notification_service/internal/application/model"
	app_template "notification_service/internal/application/template"
	"text/template"

	"gopkg.in/gomail.v2"
)

type EmailHelper struct {
	emailDialer *gomail.Dialer
}

func NewEmailHelper(
	emailDialer *gomail.Dialer,
) *EmailHelper {
	return &EmailHelper{
		emailDialer: emailDialer,
	}
}

func (h *EmailHelper) SendScheduledWorkEmail(to string, data app_model.EmailData) error {
	htmlBody, textBody, err := RenderEmail(data)
	if err != nil {
		return err
	}

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
