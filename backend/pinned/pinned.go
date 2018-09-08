package pinned

import (
	"fmt"
	"math/rand"
	"regexp"
	"sync"
	"time"

	"github.com/tmw/pinned/backend/fetcher"
	"github.com/tmw/pinned/backend/model"
)

// Pinned is our main quiz logic.
type Pinned struct {
	fetcher fetcher.Fetcher

	pins      []*model.Pin
	userIndex *Index

	NumChallenges int
	NumChoices    int
}

// Load fetches Users and Pins using the provided fetcher
func (p *Pinned) Load() error {
	rand.Seed(time.Now().Unix())

	users, err := p.fetcher.FetchUsers()
	if err != nil {
		return err
	}

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

	// unfurl all usernames
	for _, pin := range p.pins {
		pin.Text = p.unfurlMention(pin.Text)
	}

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

// unfurlMention will take mention sequences of Slack and convert them
// into their actual user- or channel names. It'll look upt the user IDs
// using the user index and replace the <@39cm9pd> with the users name.
// For the <@here|here> and <#CJDF989D|general> it'll use the second part
// after the pipe symbol.
func (p *Pinned) unfurlMention(message string) string {
	r := userMentionRegex()

	// replace occurrences of <@U9390238> into their actual usernames
	message = r.ReplaceAllStringFunc(message, func(m string) string {
		userID := r.FindStringSubmatch(m)[1]
		if user := p.userIndex.Get(userID).(*model.User); user != nil {
			return fmt.Sprintf("<@%s>", user.Name)
		}

		// When the user is not found in the index for some reason,
		// there's not much else to do then return the ol' john doe :)
		return "<@unknown>"
	})

	// replace channel references (<#C989DF|general>) or channel mentions
	// (<@here|here>) into plane <@here> and <#general>. Front-end will make sure
	// these sequences are highlighted.
	r = channelMentionRegex()
	message = r.ReplaceAllStringFunc(message, func(m string) string {
		result := r.FindStringSubmatch(m)
		switch result[1] {
		case "#": // channel reference (eg. #general)
			return fmt.Sprintf("<#%s>", result[2])

		case "!": //channel mention (eg. @here)
			return fmt.Sprintf("<%s>", result[2])

		default:
			return m
		}
	})

	return message
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

var userRegexInstance *regexp.Regexp
var userRegexInstanceOnce sync.Once

func userMentionRegex() *regexp.Regexp {
	// making sure we're only compiling the regex once
	userRegexInstanceOnce.Do(func() {
		userRegexInstance = regexp.MustCompile("<@([a-zA-Z0-9]+)>")
	})

	return userRegexInstance
}

var channelRegexInstance *regexp.Regexp
var channelRegexInstanceOnce sync.Once

func channelMentionRegex() *regexp.Regexp {
	// making sure we're only compiling the regex once
	channelRegexInstanceOnce.Do(func() {
		channelRegexInstance = regexp.MustCompile("<(!|#).+\\|(@?[a-zA-Z0-9]+)>")
	})

	return channelRegexInstance
}

// New returns a new initialized Pinned
func New(fetcher fetcher.Fetcher, numChallenges, numChoices int) *Pinned {

	// Initialize the rest of the pinned instance
	return &Pinned{
		fetcher:   fetcher,
		pins:      []*model.Pin{},
		userIndex: NewIndex(),

		NumChallenges: numChallenges,
		NumChoices:    numChoices,
	}
}
