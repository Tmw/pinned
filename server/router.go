package server

import (
	"net/http"

	"github.com/rs/cors"
)

func (s *Server) setupRoutes() http.Handler {
	mux := http.NewServeMux()

	// configure routes
	mux.HandleFunc("/challanges", s.getChallangesHandler)
	mux.HandleFunc("/check", s.checkAnswerHandler)

	// return CORS enabled http.Handler
	return cors.Default().Handler(mux)
}
