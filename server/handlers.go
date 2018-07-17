package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getChallangesHandler(w http.ResponseWriter, r *http.Request) {
	challanges := s.superslack.GetChallanges()
	json, err := json.Marshal(challanges)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong! :("))
	}

	w.WriteHeader(200)
	w.Write(json)
}

func (s *Server) checkAnswerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Hello, world! =) "))
}
