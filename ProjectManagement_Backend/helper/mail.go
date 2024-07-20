package helper

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendEmail sends an email to the specified recipient.
// It uses SMTP server configuration from environment variables.
// Returns an error if the email couldn't be sent.
func SendEmail(recipient string) error {
    
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	// Convert SMTP_PORT to int
	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal("Invalid SMTP port")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", emailUser)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Password Reset")
	m.SetBody("text/html", `<p>Click <a href="https://example.com/reset-password?token=your-reset-token">here</a> to reset your password.</p>`)

	// Set up SMTP server details
	d := gomail.NewDialer(smtpServer, smtpPortInt, emailUser, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
