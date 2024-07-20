package mailer

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
	"github.com/tuanchill/lofola-api/pkg/logger"
	"gopkg.in/gomail.v2"
)

// EmailData holds the data to be passed to the template
type EmailData struct {
	From    string
	To      string
	Subject string
	Body    string
}

func SendMail(data models.EmailData) error {
	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s@gmail.com", global.Config.Server.AppName))
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	// Set up the SMTP dialer
	d := gomail.NewDialer(global.Config.Mail.Host, global.Config.Mail.Port, global.Config.Mail.UserName, global.Config.Mail.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		logger.LogError(fmt.Sprintf("Failed to send email: %v", err))
		return err
	}

	logger.LogInfo("Email sent successfully!")
	return nil
}

func SendOptMail(data models.DataOtpMail) error {
	// Parse the template file
	tmpl, err := template.ParseFiles("./templates/otp.html")
	if err != nil {
		logger.LogError(fmt.Sprintf("Failed to parse template: %v", err))
		return err
	}

	// Execute the template with data
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		logger.LogError(fmt.Sprintf("Failed to execute template: %v", err))
		return err
	}

	// send email
	emailData := models.EmailData{
		From:    global.Config.Mail.UserName,
		To:      data.To,
		Subject: data.Title,
		Body:    body.String(),
	}

	if err := SendMail(emailData); err != nil {
		return err
	}
	return nil
}
