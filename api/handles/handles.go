package handles

import (
	"chat/api/handles/registration"
	"chat/api/handles/utils"
	"chat/api/services/mailer"
	"log"
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
	log.Print("Setting up routes")

	http.HandleFunc("/register", WithError(registration.RegistrationHandle))
	http.HandleFunc("/verify", WithError(registration.ValidationHandle))

	log.Print("Routes ready")
}

func WithError(h HandleWithError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			log.Printf("Error answering route %v | %v", r.URL.Path, err.Error())
			if err.Code == http.StatusInternalServerError {
				http.Error(w, "Unexpected error on response", err.Code)
			}
			http.Error(w, err.Message, err.Code)
		}
	}
}

type HandleWithError func(http.ResponseWriter, *http.Request) *utils.ResponseError
