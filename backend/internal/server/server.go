package server

import (
	"net/http"
	"chat/backend/internal/store"
)

type Server struct {
	port  string
	Mu    *http.ServeMux
	Users *store.UserStore
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
		Mu:   http.NewServeMux(),
	}
}

func (s *Server) Run() error {
	// s.users = store.NewUserStore()
	s.registerRoutes()

	return http.ListenAndServe(s.port, s.Mu)
}
