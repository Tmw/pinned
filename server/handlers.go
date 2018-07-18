package server

import (
	"encoding/json"
	"net/http"

	"github.com/tmw/slack-service/model"
)

func (s *Server) getChallangesHandler(w http.ResponseWriter, r *http.Request) {
	challanges := s.superslack.GetChallanges()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(challanges)
}

func (s *Server) checkAnswersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	var answers []*model.Answer
	err := json.NewDecoder(r.Body).Decode(&answers)
	if err != nil {
		errorResponse(w, "Error decoding JSON")
		return
	}

	numCorrect, err := s.superslack.CheckAnswers(answers)
	if err != nil {
		errorResponse(w, "Error processing answers: "+err.Error())
		return
	}

	resp := struct {
		NumCorrect int `json:"num_correct"`
	}{numCorrect}

	json.NewEncoder(w).Encode(resp)
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
