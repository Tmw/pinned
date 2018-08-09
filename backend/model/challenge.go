package model

import (
	"github.com/satori/go.uuid"
)

// Challenge describes a single challenge
type Challenge struct {
	ID      uuid.UUID `json:"id"`
	Text    string    `json:"text"`
	Options []*Author `json:"options"`
	Author  *Author   `json:"author"`
}
