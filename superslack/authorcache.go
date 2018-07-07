package superslack

import "github.com/tmw/slack-service/model"

// AuthorCache acts as a key-value cache for author IDs and their Author models
type AuthorCache map[string]*model.Author

// Get will return either a pointer to model.Author or nil given its ID
func (a AuthorCache) Get(id string) *model.Author {
	return a[id]
}

// Set will insert or override a author in the authorCache
func (a AuthorCache) Set(id string, author *model.Author) {
	a[id] = author
}

// Keys returns the available keys within the cache
func (a AuthorCache) Keys() []string {
	keys := []string{}
	for k := range a {
		keys = append(keys, k)
	}
	return keys
}

// NewAuthorCache returns an initialized AuthorCache
func NewAuthorCache() *AuthorCache {
	return &AuthorCache{}
}
