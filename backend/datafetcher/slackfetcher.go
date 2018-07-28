package datafetcher

import (
	"fmt"

	"github.com/nlopes/slack"
	uuid "github.com/satori/go.uuid"
	"github.com/tmw/pinned/backend/model"
)

// SlackFetcher fetches users and pins from the Slack API
type SlackFetcher struct {
	client      *slack.Client
	channelName string
}

// FetchUsers fetches users from Slack
func (sf *SlackFetcher) FetchUsers() ([]*model.Author, error) {
	// fetch from Slack API
	users, err := sf.client.GetUsers()
	if err != nil {
		return nil, err
	}

	// iterate over Slack results, inflating into Author models
	authors := []*model.Author{}
	for _, u := range users {
		// inflate to Author object
		a := &model.Author{
			ID:     u.ID,
			Name:   u.Name,
			Avatar: u.Profile.Image192,
		}

		// and persist the author in the map
		authors = append(authors, a)
	}

	return authors, nil
}

// FetchPins fetches pins from Slack
func (sf *SlackFetcher) FetchPins() ([]*model.Pin, error) {
	channelID, err := sf.channelID()
	if err != nil {
		return nil, err
	}

	items, _, err := sf.client.ListPins(channelID)
	if err != nil {
		return nil, err
	}

	var pins []*model.Pin
	for _, item := range items {
		// early return if item is not a message
		if item.Message == nil {
			continue
		}

		p := &model.Pin{
			ID:       uuid.Must(uuid.NewV4()).String(),
			AuthorID: item.Message.User,
			Text:     item.Message.Text,
		}

		// persist the pin
		pins = append(pins, p)
	}

	return pins, nil
}

func (sf *SlackFetcher) channelID() (string, error) {
	channels, err := sf.client.GetChannels(true)
	if err != nil {
		return "", err
	}

	for _, c := range channels {
		if c.Name == sf.channelName {
			return c.ID, nil
		}
	}

	return "", fmt.Errorf("Could not find channel with name %s", sf.channelName)
}

// New sets up and returns a new SlackFetcher
func New(slackToken, channelName string) *SlackFetcher {
	return &SlackFetcher{
		client:      slack.New(slackToken),
		channelName: channelName,
	}
}
