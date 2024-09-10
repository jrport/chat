package handles

import (
	"chat/api/handles/auth/registrations"
	"chat/api/handles/auth/sessions"
	"chat/api/handles/utils"
	"chat/api/services/mailer"
	"fmt"
	"log/slog"
	"net/http"
)

type Router struct {
	Mailer *mailer.Mailer
}

func NewRouter(mailer *mailer.Mailer) *Router{
	return &Router{
		Mailer: mailer,
	}
}

func (r *Router) SetupRoutes() {
	slog.Info("Setting up routes")
	
	registrationHandle := HandleWithMailer{
		mailerService: r.Mailer,
		handleFunc: registration.RegistrationHandleFunc,
	}

	http.Handle("/register", registrationHandle)
	http.HandleFunc("/verify", HandlerFuncWithError(registration.ValidationHandle))
	http.HandleFunc("/login", HandlerFuncWithError(sessions.NewSessionHandle))
	http.HandleFunc("/logout", HandlerFuncWithError(sessions.TerminateSessionHandle))
	

	slog.Info("Routes ready")
}

func HandlerFuncWithError(h HandlerWithError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Info(fmt.Sprintf("Error answering route %v | %v", r.URL.Path, err.Error()))
			if err.Code == http.StatusInternalServerError {
				http.Error(w, "Error: ", err.Code)
				return
			}
			http.Error(w, err.Message, err.Code)
			return
		}
	}
}

type HandlerWithError func(http.ResponseWriter, *http.Request) *utils.ResponseError

type HandlerWithErrorAndMailer func(http.ResponseWriter, *http.Request, *mailer.Mailer) *utils.ResponseError

type HandleWithMailer struct {
	mailerService *mailer.Mailer
	handleFunc HandlerWithErrorAndMailer
}

func (h HandleWithMailer)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handleFunc(w, r, h.mailerService); err != nil {
		slog.Info(fmt.Sprintf("Error answering route %v | %v", r.URL.Path, err.Error()))
		if err.Code == http.StatusInternalServerError {
			http.Error(w, "Unexpected error on response", err.Code)
		}
		http.Error(w, err.Message, err.Code)
	}
}
