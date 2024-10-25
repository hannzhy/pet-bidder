package storage

import (
	"fmt"
	"sync"

	"pet-bidder/internal/types"
)

type LocalStorage struct {
	campaigns []types.AdCampaign
	mu        sync.Mutex
}

// NewLocalStorage creates a new storage instance
func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		campaigns: make([]types.AdCampaign, 0),
	}
}

func (s *LocalStorage) Set(campaign *types.AdCampaign) error {
	if campaign == nil {
		return fmt.Errorf("not valid or missed campaigh")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.campaigns = append(s.campaigns, *campaign)

	return nil
}

func (s *LocalStorage) BulkSet(campaigns []types.AdCampaign) error {
	if len(campaigns) <= 0 {
		return fmt.Errorf("no data to set")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.campaigns = append(s.campaigns, campaigns...)

	return nil
}
