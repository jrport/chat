package main

import (
	"chat/api/internal/server"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil)).With("package", "main")

func main() {
	logger.Info("Starting server...")
	server.SetupRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error(fmt.Sprintf("Terminating server with error: %s", err.Error()))
		os.Exit(1)
	}
	logger.Info("Terminating peacefully server...")
}
