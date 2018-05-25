package superslack

// Author describes a single author
type Author struct {
	ID     string
	Name   string
	Avatar string
}

// AuthorStore can be used to map IDs to Authors
type AuthorStore map[string]Author
