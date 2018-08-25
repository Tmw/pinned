package server

import (
	"fmt"
	"net/http"

	"github.com/tmw/pinned/backend/pinned"
)

// Server sets up a basic HTTP server
type Server struct {
	pinned *pinned.Pinned
}

// Start starts the HTTP server on the given port
func (s *Server) Start(port int) {
	handler := s.setupRoutes()
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}

// New configures and returns a Server instance
func New(pinned *pinned.Pinned) *Server {
	return &Server{
		pinned: pinned,
	}
}
