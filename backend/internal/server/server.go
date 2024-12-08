package server

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	Logger *slog.Logger
	Server *http.Server
	Muxer  *http.ServeMux
	Store  *sql.DB
}

func NewApp(port string, logger *slog.Logger, store *sql.DB) *App{
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
	}

	app := &App{
		Logger: logger,
		Muxer:  mux,
		Server: server,
		Store: store, 
	}
	return app
}

func (app *App) Run() error {
	print("Server listening on - " + app.Server.Addr + "\n")
	app.Logger.Info("Server listening on - " + app.Server.Addr)
	if err := app.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
