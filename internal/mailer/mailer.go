package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"os"
	"strconv"

	"path/filepath"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

//go:embed templates
var templateFS embed.FS

type Mailer struct {
	dialer *gomail.Dialer
	sender string //"no-reply@gmail.com"
	logger echo.Logger
}

type EmailData struct {
	AppName string
	Subject string
	Meta    interface{}
}

func NewMailer(logger echo.Logger) Mailer {
	mailPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		logger.Fatal("Invalid MAIL_PORT value", err)
	}
	mailHost := os.Getenv("MAIL_HOST")
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailSender := os.Getenv("MAIL_SENDER")

	dialer := gomail.NewDialer(mailHost, mailPort, mailUsername, mailPassword)

	return Mailer{
		dialer: dialer,
		sender: mailSender,
		logger: logger,
	}
}

func (mailer Mailer) Send(recipient string, templateFile string, data EmailData) error {
	absolutePath := filepath.Join("templates", templateFile)
	tmpl, err := template.ParseFS(templateFS, absolutePath)
	if err != nil {
		mailer.logger.Error(err)
		return err
	}

	data.AppName = os.Getenv("APP_NAME")

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		mailer.logger.Error(err)
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		mailer.logger.Error(err)
		return err
	}

	message := gomail.NewMessage()
	message.SetHeader("To", recipient)
	message.SetHeader("From", mailer.sender)
	message.SetHeader("Subject", subject.String())
	message.SetBody("text/html", htmlBody.String())

	if err := mailer.dialer.DialAndSend(message); err != nil {
		mailer.logger.Error(err)
	}

	return nil
}
