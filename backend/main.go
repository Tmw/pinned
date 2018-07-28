package main

import (
	"log"
	"os"

	"github.com/tmw/pinned/backend/datafetcher"
	"github.com/tmw/pinned/backend/pinned"
	"github.com/tmw/pinned/backend/server"

	"github.com/joho/godotenv"
)

const (
	// NumChallanges indicates the number of challanges returned
	NumChallanges = 5

	// NumAuthors indicates the number of authors suggested with each challange
	NumAuthors = 4
)

var (
	p   *pinned.Pinned
	srv *server.Server
)

func init() {
	godotenv.Load()

	slackToken := envOrPanic("SLACK_TOKEN")
	slackChannel := envOrPanic("SLACK_CHANNEL")

	p = pinned.New(
		datafetcher.New(slackToken, slackChannel),
		NumChallanges,
		NumAuthors,
	)
	srv = server.New(p)
}

func main() {
	// calling load will fetch all required objects
	p.Load()

	// start HTTP server
	srv.Start(4000)
}

func envOrPanic(key string) string {
	v := os.Getenv(key)

	if v == "" {
		log.Fatalf("No %s environment variable found\n", key)
	}

	return v
}
