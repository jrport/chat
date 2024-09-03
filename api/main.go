package main

import (
	"chat/api/handles"
	"chat/api/services/mailer"
	"log"
	"net/http"

	"github.com/goloop/env"
)

func init() {
	log.SetFlags(log.Ltime)
	log.SetFlags(log.Ldate)

	if err := env.Load(".env"); err != nil {
		log.Fatalf("%v", err.Error())
	}
}

func main() {
	log.Print("Booting server")

	MailConfig := mailer.MailOptions{
		Email: env.Get("EMAIL"),
		Token: env.Get("TOKEN"),
		Host:  env.Get("HOST"),
		Port:  env.Get("PORT"),
	}

	mailer := mailer.NewMailer(&MailConfig)
	router := handles.NewRouter(mailer)
	router.SetupRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Shutting down server...")
	}
}
