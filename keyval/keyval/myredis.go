package keyval

import (
	"sync"
	"time"
)

type Data struct {
	Expiry *time.Time
	Value  any
}
type Store struct {
	mu   sync.RWMutex
	data map[any]*Data
}

func NewStore() *Store {
	s := &Store{
		data: make(map[any]*Data),
	}
	go s.cleanup()
	return s
}
func (s *Store) Get(key any) any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.data[key]; ok {
		return v.Value
	}
	return "nodata"
}

func (s *Store) cleanup() {
	for {
		s.mu.Lock()
		for k, v := range s.data {
			if v == nil || v.Expiry == nil {
				continue
			}
			if v.Expiry.Before(time.Now()) {
				delete(s.data, k)
			}
		}
		s.mu.Unlock()
		time.Sleep(2 * time.Second)
	}
}

type Options func(*Data)

func WithExpiry(dur time.Duration) Options {
	return func(d *Data) {
		dur := time.Now().Add(dur)
		d.Expiry = &dur
	}
}

func (s *Store) Set(key, value any, opts ...Options) {
	s.mu.Lock()
	defer s.mu.Unlock()
	dt := &Data{
		Value: value,
	}
	for _, o := range opts {
		o(dt)
	}
	s.data[key] = dt
}

func (s *Store) Del(key any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
