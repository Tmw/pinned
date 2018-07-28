package model

import (
	"github.com/satori/go.uuid"
)

// Challange describes a single challange
type Challange struct {
	ID      uuid.UUID `json:"id"`
	Text    string    `json:"text"`
	Options []*Author `json:"options"`
	Author  *Author   `json:"author"`
}
