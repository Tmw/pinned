package superslack

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nlopes/slack"
	uuid "github.com/satori/go.uuid"
	"github.com/tmw/slack-service/model"
)

// SuperSlack is our own magic on top of the Slack SDK
type SuperSlack struct {
	client *slack.Client

	channelName string
	pins        []*model.Pin
	authorCache *AuthorCache
}

// Load is used to fetch Users and Pins
func (s *SuperSlack) Load() error {
	rand.Seed(time.Now().Unix())

	if err := s.loadUsers(); err != nil {
		return err
	}

	return s.loadPins()
}

func (s *SuperSlack) loadUsers() error {
	users, err := s.client.GetUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		// inflate to Author object
		a := &model.Author{
			ID:     u.ID,
			Name:   u.Name,
			Avatar: u.Profile.Image192,
		}

		// and persist the author in the map
		s.authorCache.Set(u.ID, a)
	}

	return nil
}

func (s *SuperSlack) loadPins() error {
	channelID, err := s.channelID()
	if err != nil {
		return err
	}

	pins, _, err := s.client.ListPins(channelID)

	if err != nil {
		return err
	}

	for _, item := range pins {
		// early return if pinned item is not a message
		if item.Message == nil {
			continue
		}

		// Pull original author object from cache
		originalAuthor := s.authorCache.Get(item.Message.User)

		p := &model.Pin{
			ID:     uuid.Must(uuid.NewV4()).String(),
			Author: originalAuthor,
			Text:   item.Message.Text,
		}

		// persist the pin
		s.pins = append(s.pins, p)
	}

	return nil
}

func (s *SuperSlack) channelID() (string, error) {
	channels, err := s.client.GetChannels(true)
	if err != nil {
		return "", err
	}

	for _, c := range channels {
		if c.Name == s.channelName {
			return c.ID, nil
		}
	}

	return "", fmt.Errorf("Could not find channel with name %s", s.channelName)
}

// GetChallanges returns the amount of requested challanges.
func (s *SuperSlack) GetChallanges(numChallanges int) []*model.Challange {

	var challanges []*model.Challange
	indexes := rand.Perm(len(s.pins))

	for _, idx := range indexes[:numChallanges] {
		pickedPin := s.pins[idx]
		originalAuthor := pickedPin.Author

		challange := &model.Challange{
			ID:      pickedPin.ID,
			Text:    pickedPin.Text,
			Options: s.getUniqueRandomAuthors(4, originalAuthor.ID),
			Author:  originalAuthor,
		}

		challanges = append(challanges, challange)
	}

	// TODO: There's actually loads of improvements here:
	// * The suggested authors ARE unique, but is done very sloppy..

	return challanges
}

// TODO: Rewrite this so it'll create a copy of the author array,
// drop the original author from it, pick a random one, drop the picked one from the list and repeat.
// that way we can get rid of the if-already-in-list, try forever loop..

// We can also reuse the same technique to pick random challanges.
// maybe let it be an generic function.
func (s *SuperSlack) getUniqueRandomAuthors(number int, primer string) []*model.Author {
	// first pluck unique keys
	authorKeys := []string{primer}

	keys := s.authorCache.Keys()

	for {
		candidateKey := keys[rand.Intn(len(keys))]
		if !contains(authorKeys, candidateKey) {
			authorKeys = append(authorKeys, candidateKey)
		}

		// if desired number of authors is reached - break loop
		if len(authorKeys) == number {
			break
		}
	}

	// Ok; now look up actual author for every ID
	authors := []*model.Author{}
	for _, id := range authorKeys {
		authors = append(authors, s.authorCache.Get(id))
	}

	// and return!
	return authors
}

func contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func (s *SuperSlack) findPinByID(pinID string) (*model.Pin, error) {
	for _, p := range s.pins {
		if p.ID == pinID {
			return p, nil
		}
	}

	return nil, fmt.Errorf("Unable to find pin with id %s", pinID)
}

// CheckAnswer checks answer against the truth
func (s *SuperSlack) CheckAnswer(challangeID, answeredUserID string) bool {
	pin, err := s.findPinByID(challangeID)
	if err != nil {
		return false
	}

	return pin.Author.ID == answeredUserID
}

// New returns a new initialized SuperSlack
func New(slackToken, channelName string) SuperSlack {
	return SuperSlack{
		client:      slack.New(slackToken),
		channelName: channelName,
		pins:        []*model.Pin{},
		authorCache: NewAuthorCache(),
	}
}
