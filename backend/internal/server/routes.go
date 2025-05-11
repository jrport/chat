package server

import (
	"chat/backend/internal/store"
	"net/http"
)

type registrationHandle struct {
	userStore *store.UserStore
}

func (h registrationHandle)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Server) registerRoutes() {
	s.Mu.Handle("POST /auth/register", registrationHandle{userStore: s.users})
}
