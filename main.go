package main

import (
	"log"
	"os"

	"github.com/tmw/slack-service/datafetcher"
	"github.com/tmw/slack-service/server"

	"github.com/joho/godotenv"
	"github.com/tmw/slack-service/superslack"
)

const (
	// NumChallanges indicates the number of challanges returned
	NumChallanges = 5

	// NumAuthors indicates the number of authors suggested with each challange
	NumAuthors = 4
)

var (
	ss  *superslack.SuperSlack
	srv *server.Server
)

func init() {
	godotenv.Load()

	slackToken := envOrPanic("SLACK_TOKEN")
	slackChannel := envOrPanic("SLACK_CHANNEL")

	ss = superslack.New(
		datafetcher.New(slackToken, slackChannel),
		NumChallanges,
		NumAuthors,
	)
	srv = server.New(ss)
}

func main() {
	// calling load will fetch all required objects
	ss.Load()

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
