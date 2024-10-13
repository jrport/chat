package server

import (
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil)).With("package", "server")

func SetupRoutes() {
	http.HandleFunc("GET /test", MakeHandleWithError(TestFunc))
}
