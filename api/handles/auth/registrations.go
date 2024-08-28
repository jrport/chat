package auth

import (
	"chat/api/handles"
	"chat/api/models"
	"chat/api/services/mailer"
	"chat/api/utils"
	"fmt"
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegistrationHandle struct {
	mailService *mailer.EmailService
	resultChan  chan error
}

func (h *UserRegistrationHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.ServeHTTPWithError(w, r); err != nil {
		slog.Info(err.Error)
		http.Error(w, err.Error, err.Code)
	}
}

func (h *UserRegistrationHandle) ServeHTTPWithError(w http.ResponseWriter, r *http.Request) *handles.HandlerError {
	if r.Method != http.MethodPost {
		return &handles.HandlerError{Error: "Method not allowed", Code: http.StatusMethodNotAllowed}
	}
	if r.Header.Get("Content-Type") != "application/json" {
		return &handles.HandlerError{Error: "Invalid MIMEType on registration", Code: http.StatusUnsupportedMediaType}
	}

	credentialsDecoded, err := utils.JsonToCredentials(&r.Body)
	if err != nil {
		return &handles.HandlerError{Error: err.Error(), Code: http.StatusBadRequest}
	}

	id, err := models.CreateUser(credentialsDecoded)
	credentialsDecoded.ID = id.InsertedID.(primitive.ObjectID)
	if err != nil {
		return &handles.HandlerError{Error: err.Error(), Code: http.StatusBadRequest}
	}

	done := h.mailService.Subscribe(mailer.Job{
		Account: credentialsDecoded,
		Action:  mailer.Confirmation,
		Result:  h.resultChan,
	})

	if err := <-done; err != nil {
		return &handles.HandlerError{
			Code:  http.StatusInternalServerError,
			Error: "Error on email confirmation, please try again later!",
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Cadastro concluído com sucesso, %v!", credentialsDecoded.Email)))
	return nil
}

func verificationHandle(w http.ResponseWriter, r *http.Request) *handles.HandlerError {
	slog.Info("Validating email")
	if r.Method != http.MethodGet {
		http.Error(w, "Unallowed method", http.StatusMethodNotAllowed)
	}

	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "No validation token", http.StatusBadRequest)
	}

	err := models.ValidateUser(token)
	switch err.(type) {
	case models.InvalidTokenErr:
		return &handles.HandlerError{Error: "Invalid token", Code: http.StatusBadRequest}
	case models.AlreadyValidatedError:
		return &handles.HandlerError{Error: "Token already validated", Code: http.StatusOK}
	case error:
		return &handles.HandlerError{Error: "Error on valdition, please try again later", Code: http.StatusBadRequest}
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "User verified successfully!")
	slog.Info("Email validated")
	return nil
}

func init() {
	mailCredentials, err := utils.GetMailerCredentials()
	if err != nil {
		slog.Error(fmt.Sprintf("Error on fetching env variables: %v", err.Error()))
	}

	MailService := mailer.NewEmailService(mailCredentials)
	registrationHandle := &UserRegistrationHandle{
		mailService: MailService,
	}

	go MailService.Run()
	defer MailService.Close()

	http.Handle("/register", registrationHandle)
	http.HandleFunc("/verify", handles.WithError(verificationHandle))
}
