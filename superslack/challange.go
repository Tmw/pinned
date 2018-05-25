package superslack

// Challange describes a single challange
type Challange struct {
	ID      string
	Text    string
	Options []*Author
}
