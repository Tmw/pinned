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
