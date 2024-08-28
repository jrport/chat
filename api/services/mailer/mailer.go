package mailer

import (
	"chat/api/utils"
	"fmt"
	"log/slog"
	"net/smtp"
	"sync"
)

type MailAction int

const (
	Confirmation MailAction = iota
	Reset
	Change
)

type Job struct {
	Account *utils.Account
	Action  MailAction
	Result  chan error
}

type EmailService struct {
	credentials *utils.MailCredentials
	MailQueue   chan Job
	mu          sync.Mutex
}

func NewEmailService(c *utils.MailCredentials) *EmailService {
	return &EmailService{
		credentials: c,
		MailQueue:   make(chan Job, 1024),
	}
}

func (e *EmailService) Subscribe(job Job) chan error{
    e.MailQueue <- job
    return job.Result
}

func (e *EmailService) Run() {
	slog.Info("Running emailing service")
	for {
		job := <-e.MailQueue
		slog.Info(fmt.Sprintf("Sending %v email to %v", job.Action, job.Account))
		switch job.Action {
		case Confirmation:
			go e.VerifyAccount(&job)
		default:
			slog.Error("Invalid mailing job type!")
		}
	}
}

func (e *EmailService)Close() {
    slog.Info("Closing channels")
    close(e.MailQueue)
    slog.Info("All channels closed")
} 

func (e *EmailService) SendEmail(destination string, message *[]byte) {
	err := smtp.SendMail(
		e.credentials.Host+":"+e.credentials.Port,
		e.credentials.Auth(),
		e.credentials.Email,
		[]string{destination},
		*message,
	)

	if err != nil {
		slog.Error(fmt.Sprintf("Error on sending confirmation email: %v", err.Error()))
		e.mu.Unlock()
		return
	}

	slog.Info(fmt.Sprintf("Email sent to: %v", destination))
}
