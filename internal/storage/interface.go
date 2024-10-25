package storage

import "pet-bidder/internal/types"

type Storage interface {
	BulkSet(campaigns []types.AdCampaign) error
	Set(campaign *types.AdCampaign) error
}
