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

	m := BuildMessageBody(app_constant.APP_EMAIL, to, data.Title, textBody, htmlBody)

	// Send using the helper's dialer
	return h.emailDialer.DialAndSend(m)
}

func (h *EmailHelper) SendAIWorkGenerationEmail(to string, data app_model.EmailData) error {
	htmlBody, err := RenderEmailTemplate(app_template.AIWorkGenerationEmailHTML, data)
	if err != nil {
		return err
	}
	textBody, err := RenderEmailTemplate(app_template.AIWorkGenerationEmailPlain, data)
	if err != nil {
		return err
	}
	m := BuildMessageBody(app_constant.APP_EMAIL, to, data.Title, textBody, htmlBody)

	// Send using the helper's dialer
	return h.emailDialer.DialAndSend(m)
}

func BuildMessageBody(from, to, title, textBody, htmlBody string) *gomail.Message {
	m := gomail.NewMessage()

	// From + To + Subject
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)

	// Add text fallback
	m.AddAlternative("text/plain", textBody)

	// Add HTML version
	m.AddAlternative("text/html", htmlBody)
	return m
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

func RenderEmailTemplate(tpl string, data app_model.EmailData) (string, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
