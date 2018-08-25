package datafetcher

import (
	"github.com/tmw/pinned/backend/model"
)

// DataFetcher describes an interface that is responsible for
// actually fetching the required data
type DataFetcher interface {
	FetchUsers() ([]*model.User, error)
	FetchPins() ([]*model.Pin, error)
}
