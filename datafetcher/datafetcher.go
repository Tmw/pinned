package datafetcher

import (
	"github.com/tmw/slack-service/model"
)

// DataFetcher describes an interface that is responsible
// that is actually responsible for talking to slack.
type DataFetcher interface {
	FetchUsers() ([]*model.Author, error)
	FetchPins() ([]*model.Pin, error)
}
