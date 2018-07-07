package model

// Pin describes a single pin
type Pin struct {
	ID     string `json:"id"`
	Author Author `json:"author"`
	Text   string `json:"text"`
}
