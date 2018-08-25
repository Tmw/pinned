package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getChallengesHandler(w http.ResponseWriter, r *http.Request) {
	challenges := s.pinned.GetChallenges()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(challenges)
}
