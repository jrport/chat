package auth

import (
	"chat/api/models"
	"chat/api/services/mailer"
	"chat/api/utils"
	"fmt"
	"log/slog"
	"net/http"
)

type UserRegistrationHandle struct {
    mailService *mailer.EmailService
}

func (h *UserRegistrationHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        slog.Info("Call on registration with invalid method")
        http.Error(w, "Must be post with registration info", http.StatusMethodNotAllowed)
        return
    }
    if r.Header.Get("Content-Type") != "application/json" {
        slog.Info("Invalid MIMEType on registration")
        http.Error(w, "Must be json", http.StatusUnsupportedMediaType)
        return
    }    

    credentialsDecoded, err := utils.JsonToCredentials(&r.Body)
    if err != nil {
        slog.Error(err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = models.CreateUser(credentialsDecoded)
    if err != nil {
        slog.Error(err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Write([]byte(fmt.Sprintf("Cadastro concluído com sucesso, %v!", credentialsDecoded.Email)))
    h.mailService.MailingList <- credentialsDecoded.Email
}

func init() {
    mailCredentials, err := utils.GetMailerCredentials()
    if err != nil {
        slog.Error(fmt.Sprintf("Error on fetching env variables: %v", err.Error()))
    }

    mailService := mailer.NewEmailService(mailCredentials)
    registrationHandle := &UserRegistrationHandle{
        mailService: mailService,
    }

    go mailService.Run()

    http.Handle("/register", registrationHandle)
}
