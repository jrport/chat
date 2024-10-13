package server

import "net/http"

func TestFunc(w http.ResponseWriter, r *http.Request) *ApiError{
	a := NewApiError(500, "caralho")
	return a
}
