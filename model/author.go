package model

// Author describes a single author
type Author struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// AuthorStore can be used to map IDs to Authors
type AuthorStore map[string]Author
