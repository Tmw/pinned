package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmw/slack-service/superslack"
)

func main() {
	godotenv.Load()

	s := superslack.New(envOrPanic("SLACK_TOKEN"), envOrPanic("SUPERSLACK_CHANNEL"))

	s.Load()

	challange := s.GetChallange()

}

func envOrPanic(key string) string {
	v := os.Getenv(key)

	if v == "" {
		log.Fatalf("No %s environment variable found", v)
	}

	return v
}
