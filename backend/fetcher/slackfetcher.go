package fetcher

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
func (sf *SlackFetcher) FetchUsers() ([]*model.User, error) {
	// fetch from Slack API
	slackUsers, err := sf.client.GetUsers()
	if err != nil {
		return nil, err
	}

	// Transform into model.User structs
	users := make([]*model.User, len(slackUsers))
	for i, user := range slackUsers {
		users[i] = &model.User{
			ID:     user.ID,
			Name:   user.Name,
			Avatar: user.Profile.Image192,
			IsBot:  user.IsBot,
		}
	}

	return users, nil
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
			ID:       uuid.NewV4(),
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
