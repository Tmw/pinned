package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func (s *Server) setupRoutes() http.Handler {
	mux := mux.NewRouter()

	// configure routes
	mux.HandleFunc("/api/challanges", s.getChallangesHandler).Methods("GET")

	// custom 404 response
	mux.NotFoundHandler = http.HandlerFunc(s.notFoundHandler)

	// return CORS enabled http.Handler
	return cors.Default().Handler(mux)
}
