package superslack

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nlopes/slack"
	uuid "github.com/satori/go.uuid"
)

// SuperSlack is our own magic on top of the Slack SDK
type SuperSlack struct {
	client *slack.Client

	channelName string
	pins        PinStore
	authors     AuthorStore
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
		a := Author{
			ID:     u.ID,
			Name:   u.Name,
			Avatar: u.Profile.Image192,
		}

		// and persist the author in the map
		s.authors[u.ID] = a
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
		// Currently only support pinned messages
		if item.Message != nil {

			// find original author object
			p := Pin{
				ID:     uuid.Must(uuid.NewV4()).String(),
				Author: s.authors[item.Message.User],
				Text:   item.Message.Text,
			}

			// persist the pin
			s.pins = append(s.pins, p)
		}
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

// GetChallange returns a single challange
func (s *SuperSlack) GetChallange() Challange {

	// grab random pin
	pickedPin := s.pins[rand.Intn(len(s.pins))]

	originalAuthor := &pickedPin.Author

	challange := Challange{
		ID:      pickedPin.ID,
		Text:    pickedPin.Text,
		Options: s.getUniqueRandomAuthors(4, originalAuthor.ID),
	}

	return challange
}

func (s *SuperSlack) getUniqueRandomAuthors(number int, primer string) []*Author {
	// first pluck unique keys
	authorKeys := []string{primer}

	// grab all keys (since Authors are stored in a map)
	keys := []string{}
	for k := range s.authors {
		keys = append(keys, k)
	}

	for {
		candidateKey := keys[rand.Intn(len(keys))]
		if !contains(authorKeys, candidateKey) {
			authorKeys = append(authorKeys, candidateKey)
		}

		if len(authorKeys) == number {
			break
		}
	}

	// Ok; now look up actual author for every ID
	authors := []*Author{}
	for _, id := range authorKeys {
		author := s.authors[id]
		authors = append(authors, &author)
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

// CheckAnswer checks answer against the truth
func (s *SuperSlack) CheckAnswer(challangeId, answeredUserId string) bool {
	return true
}

// New returns a new initialized SuperSlack
func New(slackToken, channelName string) SuperSlack {
	return SuperSlack{
		client:      slack.New(slackToken),
		channelName: channelName,
		pins:        PinStore{},
		authors:     AuthorStore{},
	}
}
