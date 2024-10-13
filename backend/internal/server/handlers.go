package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)
type ApiHandleFunc func(http.ResponseWriter, *http.Request) *ApiError
type ApiError struct {
	StatusCode int   `json:"status"`
	Msg        error `json:"msg"`
}

func (a ApiError) Error() string {
	return a.Msg.Error()
}

func NewApiError(status int, msg string) *ApiError {
	return &ApiError{
		StatusCode: status,
		Msg:        fmt.Errorf(msg),
	}
}

func MakeHandleWithError(f ApiHandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := f(w, r); err != nil {
			logger.Error(err.Error())
			if err.StatusCode == http.StatusInternalServerError {
				logger.Error("pqp")
				json.NewEncoder(w).Encode(
					NewApiError(
						http.StatusInternalServerError,
						"Pqp",
						// http.StatusText(http.StatusInternalServerError),
					),
				)
			}
		}
	}
}
