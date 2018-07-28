package model

import (
	"github.com/satori/go.uuid"
)

// Pin describes a single pin
type Pin struct {
	ID       uuid.UUID `json:"id"`
	AuthorID string    `json:"author_id"`
	Text     string    `json:"text"`
}
