package main

import (
	"chat/api/handles"
	"chat/api/services/mailer"
	"log/slog"
	"net/http"
	"os"

	"github.com/goloop/env"
)

func init() {
	if err := env.Load(".env"); err != nil {
		slog.Error("%v" + err.Error())
		os.Exit(1)
	}
}

func main() {
	slog.Info("Booting server")

	MailConfig := mailer.MailOptions{
		Email: env.Get("EMAIL"),
		Token: env.Get("TOKEN"),
		Host:  env.Get("HOST"),
		Port:  env.Get("PORT"),
	}

	mailer := mailer.NewMailer(&MailConfig)
	router := handles.NewRouter(mailer)
	go mailer.Run()
	router.SetupRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Shutting down server...")
		os.Exit(1)
	}
}
