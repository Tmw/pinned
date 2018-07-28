package pinned

import (
	"math/rand"
	"time"

	"github.com/tmw/pinned/backend/datafetcher"
	"github.com/tmw/pinned/backend/model"
)

// Pinned is our main quiz logic.
type Pinned struct {
	fetcher datafetcher.DataFetcher

	pins        []*model.Pin
	authorCache *AuthorCache

	NumChallanges int
	NumOptions    int
}

// Load is used to fetch Users and Pins
func (p *Pinned) Load() error {
	rand.Seed(time.Now().Unix())

	users, err := p.fetcher.FetchUsers()
	if err != nil {
		return err
	}

	p.authorCache.Add(users)

	pins, err := p.fetcher.FetchPins()
	if err != nil {
		return err
	}

	p.pins = pins
	return nil
}

// GetChallanges returns the amount of requested challanges.
func (p *Pinned) GetChallanges() []*model.Challange {
	cappedNumChallanges := cap(p.NumChallanges, len(p.pins))

	var challanges []*model.Challange
	indexes := rand.Perm(len(p.pins))

	for _, idx := range indexes[:cappedNumChallanges] {
		pickedPin := p.pins[idx]

		challange := &model.Challange{
			ID:      pickedPin.ID,
			Text:    pickedPin.Text,
			Options: p.getUniqueRandomAuthors(4, pickedPin.AuthorID),
			Author:  p.authorCache.Get(pickedPin.AuthorID),
		}

		challanges = append(challanges, challange)
	}

	return challanges
}

func (p *Pinned) getUniqueRandomAuthors(number int, primer string) []*model.Author {
	// first pluck unique keys
	authorKeys := []string{primer}
	keys := p.authorCache.Keys()

	for _, idx := range rand.Perm(len(keys)) {
		candidateKey := keys[idx]

		if !contains(authorKeys, candidateKey) {
			authorKeys = append(authorKeys, candidateKey)
		}

		// if desired number of authors is reached - break loop
		if len(authorKeys) == number {
			break
		}
	}

	// Ok; now look up actual author for every ID,
	// but do it in a randomized order.
	authors := []*model.Author{}
	for _, idx := range rand.Perm(number) {
		authors = append(authors, p.authorCache.Get(authorKeys[idx]))
	}

	return authors
}

// New returns a new initialized SuperSlack
func New(fetcher datafetcher.DataFetcher, numChallanges, numOptions int) *Pinned {
	return &Pinned{
		fetcher:     fetcher,
		pins:        []*model.Pin{},
		authorCache: NewAuthorCache(),

		NumChallanges: numChallanges,
		NumOptions:    numOptions,
	}
}
