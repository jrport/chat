package handles

import (
	"log/slog"
	"net/http"
)


type HandlerError struct {
	Error string
	Code  int
}

type HandleFunc func(http.ResponseWriter, *http.Request)
type HandleFuncWithError func(http.ResponseWriter, *http.Request)*HandlerError

func WithError(h HandleFuncWithError)HandleFunc{
    return func(w http.ResponseWriter, r *http.Request){
        if err := h(w, r); err != nil {
            slog.Error(err.Error)
            http.Error(w, err.Error, err.Code)
        }
    }
}

