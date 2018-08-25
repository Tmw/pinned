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

	pins      []*model.Pin
	userIndex *Index

	NumChallenges int
	NumChoices    int
}

// Load fetches Users and Pins using the provided DataFetcher
func (p *Pinned) Load() error {
	rand.Seed(time.Now().Unix())

	users, err := p.fetcher.FetchUsers()
	if err != nil {
		return err
	}

	// NOTE: Not sure why we need to iterate over the slice
	// again and append it to an []Indexable slice to make it
	// be accepted by our variadic function that ingests Indexables..

	// but hey - such is life for now.. :)

	indexables := make([]Indexable, len(users))
	for i, user := range users {
		indexables[i] = user
	}

	p.userIndex.Store(indexables...)

	pins, err := p.fetcher.FetchPins()
	if err != nil {
		return err
	}

	p.pins = pins
	return nil
}

// GetChallenges returns the requested amount of challenges
func (p *Pinned) GetChallenges() []*model.Challenge {
	numChallenges := max(p.NumChallenges, len(p.pins))
	challenges := make([]*model.Challenge, numChallenges)
	indexes := rand.Perm(len(p.pins))

	for i, idx := range indexes[:numChallenges] {
		pin := p.pins[idx]

		challenges[i] = &model.Challenge{
			ID:      pin.ID,
			Text:    pin.Text,
			Choices: p.generateChoices(4, pin.AuthorID),
			Author:  p.userIndex.Get(pin.AuthorID).(*model.User),
		}
	}

	return challenges
}

func (p *Pinned) generateChoices(number int, authorID string) []*model.User {
	// first pluck author IDs but already provide actual author ID
	userIDs := []string{authorID}
	keys := p.userIndex.Keys()

	for _, idx := range rand.Perm(len(keys)) {
		k := keys[idx]

		// Skip bot users for available choices
		if p.userIndex.Get(k).(*model.User).IsBot {
			continue
		}

		// Skip if the user is already in the list
		if contains(userIDs, k) {
			continue
		}

		// Add author id to the list of ids
		userIDs = append(userIDs, k)

		// Break the loop once we've reached desired number of users
		if len(userIDs) == number {
			break
		}
	}

	// Last step is to look up actual users by ID in randomized order
	users := make([]*model.User, number)
	for i, idx := range rand.Perm(number) {
		users[i] = p.userIndex.Get(userIDs[idx]).(*model.User)
	}

	return users
}

// New returns a new initialized Pinned
func New(fetcher datafetcher.DataFetcher, numChallenges, numChoices int) *Pinned {
	return &Pinned{
		fetcher:   fetcher,
		pins:      []*model.Pin{},
		userIndex: NewIndex(),

		NumChallenges: numChallenges,
		NumChoices:    numChoices,
	}
}
