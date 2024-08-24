package main

import (
	"chat/api/db"
	"chat/api/handles"
	"chat/api/mailer"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting things up...")
    log.Printf("Starting mailer...")
    go mailer.Run()

	log.Println("Listening on port 8080")

	http.HandleFunc("/login", WithError(handles.LoginUserHandle))
	http.HandleFunc("/register", WithError(handles.RegisterUserHandle))
    http.HandleFunc("/logout", WithError(WithAuth(handles.LogOutHandle)))
	http.HandleFunc("/ping", WithError(WithAuth(handles.PingRoute)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Closing server: %v", err.Error())
	}
}

type HandleWithError func(http.ResponseWriter, *http.Request) error

func WithError(f HandleWithError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("Erro na resposta ao cliente: %v", err.Error())
		}
	}
}

func WithAuth(f HandleWithError) HandleWithError{
    return func(w http.ResponseWriter, r *http.Request) error {
        sessionId, err := r.Cookie("sessionId")
        if err != nil {
            http.Error(w, "Sessão não autenticada", http.StatusForbidden)
            return err
        }

        cookie, err := db.FindCookie(sessionId.Value)
        if err != nil {
            http.Error(w, "Sessão não autenticada", http.StatusForbidden)
            return err
        }
        if cookie.ExpireAt.Before(time.Now().UTC()) {
            err = db.DeleteCookie(cookie)
            http.Error(w, "Sessão não autenticada", http.StatusForbidden)
            return err
        }

        log.Printf("Sessão %v validada", cookie.Value)
        return f(w, r)
    }
}
