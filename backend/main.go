package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tmw/pinned/backend/datafetcher"
	"github.com/tmw/pinned/backend/server"

	"github.com/joho/godotenv"
	"github.com/tmw/pinned/backend/superslack"
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

	fmt.Println("Superslack server starting at: http://localhost:4000")

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
