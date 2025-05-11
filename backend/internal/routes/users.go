package routes

import (
	"chat/backend/internal/server"
	"net/http"
)

func registrationHandle(s *server.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Users
	}
}
