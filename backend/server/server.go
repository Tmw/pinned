package server

import (
	"net/http"

	"github.com/tmw/pinned/backend/superslack"
)

// Server sets up a basic HTTP server
type Server struct {
	superslack *superslack.SuperSlack
}

// Start starts the HTTP server on the given port
func (s *Server) Start(port int) {
	handler := s.setupRoutes()
	http.ListenAndServe(":4000", handler)
}

// New configures and returns a Server instance
func New(superslack *superslack.SuperSlack) *Server {
	return &Server{
		superslack: superslack,
	}
}
