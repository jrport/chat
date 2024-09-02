package handles

import (
	"chat/api/handles/registration"
	"chat/api/handles/utils"
	"log"
	"net/http"
)

func SetupRoutes() {
	log.Print("Setting up routes")

	http.HandleFunc("/register", WithError(registration.RegistrationHandle))

	log.Print("Routes ready")
}



func WithError(h HandleWithError) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			log.Printf("Error answering route %v\nOf type %v", r.URL.Path, err.Error())
			if err.Code == http.StatusInternalServerError {
				http.Error(w, "Unexpected error on response", err.Code)	
			}
			http.Error(w, err.Message, err.Code)
		}
	}
}

type HandleWithError func(http.ResponseWriter, *http.Request) *utils.ResponseError
