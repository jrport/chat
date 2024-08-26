package mailer

import (
	"chat/api/utils"
	"sync"
)

type EmailService struct {
	credentials *utils.MailCredentials
	MailingList chan string
	mu          sync.Mutex
}

func NewEmailService(c *utils.MailCredentials) *EmailService {
	return &EmailService{
		credentials: c,
        MailingList: make(chan string, 1024),
	}
}

func (e *EmailService) Run() {
    for {
        select {
        case email := <- e.MailingList:
            go e.SendEmail(email)
        }
    }
}

func (e *EmailService) SendEmail(email string) {
    e.mu.Lock()

    token := 1234
    

}
