package fetcher

import (
	"github.com/tmw/pinned/backend/model"
)

// Fetcher describes an interface that is responsible for
// actually fetching the required data from an external API.
// An example implementation is the slackfetcher.go
type Fetcher interface {
	FetchUsers() ([]*model.User, error)
	FetchPins() ([]*model.Pin, error)
}
