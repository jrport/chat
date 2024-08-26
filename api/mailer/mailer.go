package mailer

import (
	"fmt"
	"log"
	"net/smtp"
)

type EmailConfiguration struct {
	Email    string
	Password string
	Host     string
	Port     string
}

type EmailService struct {
    Conf *EmailConfiguration
}

func NewEmailConfiguration(email string, password string, host string, port string) *EmailConfiguration {
	return &EmailConfiguration{
		Email:    email,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

func NewEmailService(conf *EmailConfiguration) *EmailService{
    return &EmailService{
        Conf: conf,
    }
}

func (e *EmailService)Run(DestinationChanel chan string){
    defer close(DestinationChanel)

    for {
        select {
        case email := <- DestinationChanel:
            log.Print("Sending email to ", email)
            e.sendEmail(email)
            log.Print("Sent email to ", email)
        }
    }

}

func (e *EmailService)sendEmail(email string) {
    message := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    message += "From: " + e.Conf.Email + "\r\n"
    message += "To: " + email + ";\r\n"
    message += "Subject: " + "Confirmação de cadastro;" + "\r\n"
    message += "\r\n Clique <a href='http://localhost:8080/verify?token='>aqui</a> para confirmar seu email\r\n"

	auth := smtp.PlainAuth("", e.Conf.Email, e.Conf.Password, e.Conf.Host)

	err := smtp.SendMail(e.Conf.Host+":"+e.Conf.Port, auth, e.Conf.Email, []string{email}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
