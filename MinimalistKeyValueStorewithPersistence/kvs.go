package main

import (
	"encoding/json"
	"os"
	"sync"
)

type KeyValueStore[K comparable, V any] struct {
	data   map[K]V
	mu     sync.RWMutex
	logger *os.File
}
type Command string

const (
	DELETE Command = "del"
	SET    Command = "set"
)

type LogEntry[K comparable, V any] struct {
	Op    Command
	Key   K
	Value V
}

func NewStore[K comparable, V any](filepath string) (*KeyValueStore[K, V], error) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	kvs := &KeyValueStore[K, V]{
		data:   make(map[K]V),
		logger: f,
	}
	if err := kvs.recover(); err != nil {
		return nil, err
	}
	return kvs, nil
}

func (s *KeyValueStore[K, V]) recover() error {
	return nil
}

func (s *KeyValueStore[K, V]) Get(key K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res, ok := s.data[key]
	return res, ok
}

func (s *KeyValueStore[K, V]) Set(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value

	entry := &LogEntry[K, V]{
		Op:    SET,
		Key:   key,
		Value: value,
	}
	if err := s.writeToLog(entry); err != nil {
		return err
	}
	return nil
}

func (s *KeyValueStore[K, V]) writeToLog(l *LogEntry[K, V]) error {
	bts, err := json.Marshal(l)
	if err != nil {
		return err
	}

	if _, err := s.logger.Write(append(bts, '\n')); err != nil {
		return err
	}
	return s.logger.Sync()
}

func (s *KeyValueStore[K, V]) Save() error {

	return nil
}

func (s *KeyValueStore[K, V]) Load() error {
	return nil
}
