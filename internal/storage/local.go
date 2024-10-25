package storage

import (
	"fmt"
	"sync"
)

type Campaign struct {
}

type LocalStorage struct {
	data map[string]*Campaign
	mu   sync.Mutex
}

// NewLocalStorage creates a new storage instance
func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		data: make(map[string]*Campaign),
	}
}

func (s *LocalStorage) Set(key string, campaign *Campaign) (err error) {
	if campaign == nil {
		return fmt.Errorf("not valid or missed campaigh")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = campaign

	return err
}

func (s *LocalStorage) Get(key string) (*Campaign, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	campaign, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("item not found: %s", key)
	}
	return campaign, nil
}
