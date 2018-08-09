package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) setupRoutes() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/challenges", s.getChallengesHandler).Methods("GET")
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return mux
}
