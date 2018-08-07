package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) setupRoutes() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/challanges", s.getChallangesHandler).Methods("GET")
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return mux
}
