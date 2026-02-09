package store

import (
	"sync"
)

type Hash interface {
	GetHash() string
	SaveHash(hash string)
	GetSize() int
}

type HashStore struct {
	data []string
	mu   sync.RWMutex
}

func NewHashStore() Hash {
	return &HashStore{
		data: make([]string, 0),
	}
}

func (h *HashStore) GetHash() string {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.data) == 0 {
		return ""
	}
	val := h.data[0]
	h.data = h.data[1:]
	return val
}

func (h *HashStore) SaveHash(hash string) {
	if len(h.data) == 100 {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data = append(h.data, hash)
}

func (h *HashStore) GetSize() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.data)
}
