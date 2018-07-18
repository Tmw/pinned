package model

// Answer described a single answer given by the user
type Answer struct {
	ChallangeID string `json:"challange_id"`
	AuthorID    string `json:"author_id"`
}
