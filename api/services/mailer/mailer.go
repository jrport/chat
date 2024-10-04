package mailer

import (
	"log/slog"
	"net/smtp"
)

type MailOptions struct {
	Email string
	Token string
	Host  string
	Port  string
}

type Mailer struct {
	Feed   chan MailOrder
	Config *MailOptions
}

type MailOrder struct {
	Destination string
	MailType    MailAction
	Params 		map[string]string
}

func NewMailer(config *MailOptions) *Mailer {
	return &Mailer{
        Feed: make(chan MailOrder, 1024),
        Config: config,
    }
}

func NewMailerOrder(email string, action MailAction, params *map[string]string) *MailOrder{
	return &MailOrder{
		Destination: email,
		MailType: action,
		Params: *params,
	}	
}

func (m *Mailer) GetAuth() smtp.Auth{
	return smtp.PlainAuth("", m.Config.Email, m.Config.Token, m.Config.Host)
}

func (m *Mailer)IssueMail(order MailOrder) {
	m.Feed <- order
	return
}

func (m *Mailer)Run() {
	slog.Info("Running mail service...")

	for {
		job, ok := <- m.Feed
		if !ok {
			slog.Error("Error on ordering mail action")
			continue
		}
		switch job.MailType {
		case RecoverPasswordMail:
			token := job.Params["resetToken"]
			go m.SendForgotPasswordMail(job.Destination, token)	
		case ValidationTokenMail:
			token := job.Params["verificationToken"]
			go m.SendValidationMail(job.Destination, token)
		}		
	}
}
