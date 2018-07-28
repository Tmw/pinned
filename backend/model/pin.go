package model

// Pin describes a single pin
type Pin struct {
	ID       string `json:"id"`
	AuthorID string `json:"author_id"`
	Text     string `json:"text"`
}
