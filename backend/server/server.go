package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	pinned "github.com/tmw/pinned/backend"
)

// Server sets up a basic HTTP server
type Server struct {
	pinned   *pinned.Pinned
	instance *http.Server
}

// Start starts the HTTP server on the given port
func (s *Server) Start(port int) {
	handler := s.setupRoutes()
	log.Printf("Starting server at http://127.0.0.1:%d", port)

	// setup a http.Server with default timeouts in place
	s.instance = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start listening or error out
	if err := s.instance.ListenAndServe(); err != nil {
		log.Fatalf("error while starting server: %v\n", err)
	}
}

// Stop does a graceful stop of the server
func (s *Server) Stop() {
	s.instance.Close()
}

// New configures and returns a Server instance
func New(pinned *pinned.Pinned) *Server {
	return &Server{
		pinned: pinned,
	}
}
