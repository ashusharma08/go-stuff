package store

import (
	"sync"

	serr "github.com/esoptra/go-prac/shorturl/error"
)

type Redirect interface {
	Get(key string) (interface{}, error)
	Create(key string, value interface{})
	Delete(key string)
	Update(key string, value interface{})
}

type RedirectStore struct {
	data sync.Map
}

func NewRedirectStore() Redirect {
	return &RedirectStore{
		data: sync.Map{},
	}
}

func (m *RedirectStore) Get(key string) (interface{}, error) {
	if v, ok := m.data.Load(key); !ok {
		return nil, serr.NewError(serr.ERR_KEY_NOT_FOUND, "Invalid key")
	} else {
		return v, nil
	}
}

func (m *RedirectStore) Create(key string, value interface{}) {
	m.data.Store(key, value)
}

func (m *RedirectStore) Delete(key string) {
	m.data.Delete(key)
}

func (m *RedirectStore) Update(key string, value interface{}) {
	m.data.Swap(key, value)
}
