package superslack

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tmw/slack-service/datafetcher"
	"github.com/tmw/slack-service/model"
)

// SuperSlack is our main quiz logic.
type SuperSlack struct {
	fetcher datafetcher.DataFetcher

	pins        []*model.Pin
	authorCache *AuthorCache

	NumChallanges int
	NumOptions    int
}

// Load is used to fetch Users and Pins
func (s *SuperSlack) Load() error {
	rand.Seed(time.Now().Unix())

	users, err := s.fetcher.FetchUsers()
	if err != nil {
		return err
	}

	s.authorCache.Add(users)

	pins, err := s.fetcher.FetchPins()
	if err != nil {
		return err
	}

	s.pins = pins
	return nil
}

// GetChallanges returns the amount of requested challanges.
func (s *SuperSlack) GetChallanges() []*model.Challange {
	cappedNumChallanges := cap(s.NumChallanges, len(s.pins))

	var challanges []*model.Challange
	indexes := rand.Perm(len(s.pins))

	for _, idx := range indexes[:cappedNumChallanges] {
		pickedPin := s.pins[idx]

		challange := &model.Challange{
			ID:      pickedPin.ID,
			Text:    pickedPin.Text,
			Options: s.getUniqueRandomAuthors(4, pickedPin.AuthorID),
		}

		challanges = append(challanges, challange)
	}

	return challanges
}

func (s *SuperSlack) getUniqueRandomAuthors(number int, primer string) []*model.Author {
	// first pluck unique keys
	authorKeys := []string{primer}
	keys := s.authorCache.Keys()

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
		authors = append(authors, s.authorCache.Get(authorKeys[idx]))
	}

	return authors
}

func (s *SuperSlack) findPinByID(pinID string) (*model.Pin, error) {
	for _, p := range s.pins {
		if p.ID == pinID {
			return p, nil
		}
	}

	return nil, fmt.Errorf("Unable to find pin with id %s", pinID)
}

// CheckAnswers checks answer against the truth and returns either
// the number of correct answers or an error
func (s *SuperSlack) CheckAnswers(answers []*model.Answer) (int, error) {
	if len(answers) < s.NumChallanges {
		return 0, fmt.Errorf("Not enough answers provided")
	}

	if len(answers) > s.NumChallanges {
		return 0, fmt.Errorf("Too much answers provided")
	}

	// happy path
	score := 0

	for _, answer := range answers {
		pin, err := s.findPinByID(answer.ChallangeID)
		if err != nil {
			return 0, fmt.Errorf("challange_id %v invalid", answer.ChallangeID)
		}

		if pin.AuthorID == answer.AuthorID {
			score++
		}
	}

	return score, nil
}

// TODO: Refactor New method so it'll take an option struct

// New returns a new initialized SuperSlack
func New(fetcher datafetcher.DataFetcher, numChallanges, numOptions int) *SuperSlack {
	return &SuperSlack{
		fetcher:     fetcher,
		pins:        []*model.Pin{},
		authorCache: NewAuthorCache(),

		NumChallanges: numChallanges,
		NumOptions:    numOptions,
	}
}
