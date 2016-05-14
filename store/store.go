// Package store provides a simple key-value store.
package store

import (
	"log"
	"os"
	"sync"
)

// Store is a simple key-value store.
type Store struct {
	mu sync.RWMutex
	m  map[string]string // The key-value store for the system.

	logger *log.Logger
}

// New returns a new Store.
func New() *Store {
	return &Store{
		m:      make(map[string]string),
		logger: log.New(os.Stderr, "[store] ", log.LstdFlags),
	}
}

// Open opens the store.
func (s *Store) Open() error {
	s.logger.Println("store opened")
	return nil
}

// Close closes the store.
func (s *Store) Close() error {
	return nil
}

// Get returns the value for the given key.
func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key], nil
}

// Set sets the value for the given key.
func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
	return nil
}

// Delete deletes the given key.
func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
	return nil
}
