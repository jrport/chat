package handlers

import (
	"database/sql"
	"fmt"
	"jport/chat/backend/internal/errors"
	"jport/chat/backend/internal/app"
	"net/http"
)

type HandleWithDatabase  func(db *sql.DB, w http.ResponseWriter, r *http.Request) error

func SetupRoutes(app *app.App) {
	app.Muxer.HandleFunc("/login", HandleWithLoggingAndError(app, LoginHandler))
}

func HandleWithLoggingAndError(app *app.App, f HandleWithDatabase) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			http.Error(w, "Invalid Content-Type.", http.StatusBadRequest)
			return
		}

		app.Logger.Info(fmt.Sprintf("Incoming request from %s for %s", r.RemoteAddr, r.URL.Path))
		if err := f(app.Store, w, r); err != nil {
			switch e := err.(type) {
			case *errors.HttpError:
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
