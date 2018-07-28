package model

// Challange describes a single challange
type Challange struct {
	ID      string    `json:"id"`
	Text    string    `json:"text"`
	Options []*Author `json:"options"`
	Author  *Author   `json:"author"`
}
