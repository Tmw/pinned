package main

import (
	"encoding/json"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tmw/slack-service/superslack"
)

var ss superslack.SuperSlack

const (
	// NumChallanges indicates the number of challanges returned
	NumChallanges = 5

	// NumAuthors indicates the number of authors suggested with each challange
	NumAuthors = 4
)

func init() {
	godotenv.Load()
}

func main() {
	// setup superslack
	ss = superslack.New(envOrPanic("SLACK_TOKEN"), envOrPanic("SUPERSLACK_CHANNEL"))
	ss.Load()

	// start HTTP server
	startServer()
}

func envOrPanic(key string) string {
	v := os.Getenv(key)

	if v == "" {
		log.Fatalf("No %s environment variable found", v)
	}

	return v
}

func getChallangeHandler(w http.ResponseWriter, r *http.Request) {

	challanges := ss.GetChallanges(NumChallanges)
	json, err := json.Marshal(challanges)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong! :("))
	}

	w.WriteHeader(200)
	w.Write(json)
}

func checkAnswerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Hello, world! =) "))
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/challanges", getChallangeHandler)
	mux.HandleFunc("/check", checkAnswerHandler)
	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":4000", handler)
}
