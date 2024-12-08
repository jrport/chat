package handlers

import (
	"fmt"
	"jport/chat/backend/internal/errors"
	"jport/chat/backend/internal/server"
	"net/http"
)

type HandleFuncWithError func(w http.ResponseWriter, r *http.Request) error

func SetupRoutes(app *server.App) {
	app.Muxer.HandleFunc("/register", HandleWithLoggingAndError(app, loginHandler))
}

func HandleWithLoggingAndError(app *server.App, f HandleFuncWithError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info(fmt.Sprintf("Incoming request from %s for %s", r.RemoteAddr, r.URL.Path))
		if err := f(w, r); err != nil {
			switch e := err.(type) {
			case errors.HttpError:
				app.Logger.Warn(fmt.Sprintf("From: %v | %v", r.RemoteAddr, e.Error()))
				http.Error(w, e.Msg, e.Status)
				return
			default:
				app.Logger.Error(e.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
	}
}
