package mailer

import (
	// "chat/api/utils"
	// "context"
	"fmt"
	"net/smtp"
	// "time"
	// "google.golang.org/api/gmail/v1"
)

func Run() {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	// defer cancel()
	// Sender data
	from := "jr.net.solucoes@gmail.com"
	password := "Mui9bdwd!Qm7YVu"

	// Receiver email address
	to := []string{
		"joao6roberto@gmail.com",
	}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte("Subject: Test Subject\r\n" +
		"\r\n" +
		"This is the body of the email")

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
	fmt.Println("Email Sent!")
}
