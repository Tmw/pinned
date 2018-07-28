package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) setupRoutes() http.Handler {
	mux := mux.NewRouter()

	// configure routes
	mux.HandleFunc("/api/challanges", s.getChallangesHandler).Methods("GET")
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// custom 404 response
	mux.NotFoundHandler = http.HandlerFunc(s.notFoundHandler)

	// return CORS enabled http.Handler
	return mux
}
