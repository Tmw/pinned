package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getChallangesHandler(w http.ResponseWriter, r *http.Request) {
	challanges := s.superslack.GetChallanges()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(challanges)
}

func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	errorResponse(w, "This route could not be found")
}

func errorResponse(w http.ResponseWriter, message string) {
	err := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{false, message}

	json.NewEncoder(w).Encode(err)
}
