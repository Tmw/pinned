package model

// User describes a single user
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	IsBot  bool   `json:"-"`
}

// IndexKey returns the key used to index this object in our Index.
// Satisfies the Indexable interface.
func (u *User) IndexKey() string {
	return u.ID
}
