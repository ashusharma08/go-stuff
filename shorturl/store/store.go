package store

import "time"

type URL struct {
	Key       string
	Value     string
	CreatedAt time.Time
}

type Store interface {
	RedirectStore() Redirect
	HashStore() Hash
}

type MapStore struct {
	redirectStore Redirect
	hashStore     Hash
}

func NewMapStore() Store {
	return &MapStore{}
}

func (m *MapStore) RedirectStore() Redirect {
	if m.redirectStore == nil {
		m.redirectStore = NewRedirectStore()
	}
	return m.redirectStore
}
func (m *MapStore) HashStore() Hash {
	if m.hashStore == nil {
		m.hashStore = NewHashStore()
	}
	return m.hashStore
}

func NewStore() Store {
	return &MapStore{}
}
