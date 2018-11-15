package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/tmw/pinned/backend/fetcher"
	"github.com/tmw/pinned/backend/pinned"
	"github.com/tmw/pinned/backend/server"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const demoSlackToken = "the_office"

type config struct {
	NumChallenges int    `env:"NUM_CHALLENGES" envDefault:"5"`
	NumChoices    int    `env:"NUM_CHOICES" envDefault:"4"`
	SlackToken    string `env:"SLACK_TOKEN,required"`
	SlackChannel  string `env:"SLACK_CHANNEL" envDefault:"general"`
	ServerPort    int    `env:"SERVER_PORT" envDefault:"4000"`
}

var (
	p   *pinned.Pinned
	srv *server.Server
	cfg *config
)

func init() {
	godotenv.Load()

	cfg = new(config)
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	p = pinned.New(
		getFetcher(),
		cfg.NumChallenges,
		cfg.NumChoices,
	)
	srv = server.New(p)
}

func getFetcher() fetcher.Fetcher {
	if cfg.SlackToken == demoSlackToken {
		return fetcher.NewOfficeFetcher()
	}

	return fetcher.New(cfg.SlackToken, cfg.SlackChannel)
}

func main() {
	// calling load will fetch all required objects
	if err := p.Load(); err != nil {
		log.Fatalf("Error while pre-fetching: %s", err)
	}

	// start HTTP server async
	go srv.Start(cfg.ServerPort)

	// wait for kill signal and gracefully exit
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)

	<-sc

	log.Println("Stopping server ...")
	srv.Stop()
}
