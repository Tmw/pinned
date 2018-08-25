package pinned

// Indexable described the interface required to make objects indexable
type Indexable interface {
	IndexKey() string
}

// Index is an in-memory key-value based index that indexes `Indexable` objects
type Index map[string]Indexable

// Get will return either a pointer to the indexed model or nil
// based on the given ID
func (idx Index) Get(id string) Indexable {
	return idx[id]
}

// Store is a variadic function that accepts Indexables to store.
// It will override objects in the index where the IndexKey() matches
func (idx Index) Store(indexables ...Indexable) {
	for _, indexable := range indexables {
		idx[indexable.IndexKey()] = indexable
	}
}

// Keys returns the available keys within the cache
func (idx Index) Keys() []string {
	keys := []string{}
	for key := range idx {
		keys = append(keys, key)
	}

	return keys
}

// NewIndex returns an initialized Index
func NewIndex() *Index {
	return &Index{}
}
