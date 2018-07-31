package main

import (
	"log"

	"github.com/tmw/pinned/backend/datafetcher"
	"github.com/tmw/pinned/backend/pinned"
	"github.com/tmw/pinned/backend/server"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type config struct {
	NumChallanges int    `env:"NUM_CHALLANGES" envDefault:"5"`
	NumAuthors    int    `env:"NUM_AUTHORS" envDefault:"4"`
	SlackToken    string `env:"SLACK_TOKEN,required"`
	SlackChannel  string `env:"SLACK_CHANNEL" envDefault:"general"`
}

var (
	p   *pinned.Pinned
	srv *server.Server
)

func init() {
	godotenv.Load()

	cfg := new(config)
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	p = pinned.New(
		datafetcher.New(cfg.SlackToken, cfg.SlackChannel),
		cfg.NumChallanges,
		cfg.NumAuthors,
	)
	srv = server.New(p)
}

func main() {
	// calling load will fetch all required objects
	p.Load()

	// start HTTP server
	srv.Start(4000)
}
