package services

import (
	"net/smtp"
)

func SendSimpleMail(SMTPServer string, SMTPUsername string, SMTPPassword string, SenderEmail string, RecipientEmail string) error {
	message := []byte("To: " + RecipientEmail + "\r\n" +
		"Subject: Test email\r\n" +
		"\r\n" +
		"This is a test email message.\r\n")

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPServer)

	err := smtp.SendMail(SMTPServer, auth, SenderEmail, []string{RecipientEmail}, message)
	if err != nil {
		return err
	}
	return nil
}
