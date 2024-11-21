package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type App struct {
	Logger *slog.Logger
	Server *http.Server
	Store  *sql.DB
}

func NewApp(port, logfile, database string) *App {
	fh, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Error creating/opening logfile: %v", err.Error()))
	}

	opts := slog.HandlerOptions{AddSource: true}
	applog := slog.New(slog.NewJSONHandler(fh, &opts))


	return &App{
		Logger: applog,
		// Store
	}
}

func (app *App) Run() error {
	s := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: time.Second * 30,
	}
	Server: &s,

	if err := app.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
