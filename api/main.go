package main

import (
	_ "chat/api/handles/auth"
	"log/slog"
	"net/http"
	"os"
)


func main() {
    slog.Info("Booting the server...")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        slog.Info(err.Error())
        os.Exit(0)
    }
}
